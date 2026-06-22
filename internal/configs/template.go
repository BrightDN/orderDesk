package configs

import (
	"html/template"
	"io"
	"io/fs"
	"maps"
	"os"
	"path/filepath"
	"strings"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/services/companies"
	"github.com/brightDN/orderDesk/internal/services/companies/suppliers"
	"github.com/labstack/echo/v4"
)

type Template struct {
	Identity  IdentityConfig
	Suppliers *suppliers.SupplierService
}

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	page := templatePath(name)
	executeName := templateName(name)
	isPartial := strings.HasPrefix(name, "partials/")
	isComponent := strings.HasPrefix(name, "components/")
	base, err := templateBase(page)
	if err != nil {
		return err
	}

	files, err := templateFiles(page)
	if err != nil {
		return err
	}

	tmpl := template.New("").Funcs(t.templateFuncMap(c))
	tmpl, err = tmpl.ParseFiles(files...)
	if err != nil {
		return err
	}

	templateData, err := t.templateData(data, c, isPartial, isComponent, base == "businessBase")
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, executeName, templateData)
}

func (t *Template) templateFuncMap(c echo.Context) template.FuncMap {
	return template.FuncMap{
		"route": func(name string, params ...interface{}) string {
			if c == nil || c.Echo() == nil {
				return ""
			}
			return c.Echo().Reverse(name, params...)
		},
	}
}

func templatePath(name string) string {
	name = strings.TrimPrefix(name, "/")
	if strings.HasPrefix(name, "pages/") ||
		strings.HasPrefix(name, "partials/") ||
		strings.HasPrefix(name, "components/") {
		return filepath.Join("templates", name+".html")
	}

	return filepath.Join("templates", "pages", name+".html")
}

func templateName(name string) string {
	return strings.TrimSuffix(filepath.Base(name), filepath.Ext(name))
}

func (t *Template) templateData(data any, c echo.Context, isPartial, isComponent bool, needsSuppliers bool) (any, error) {
	csrf, _ := c.Get("csrf").(string)
	feedback, _ := flash.Pop(c)
	employee := c.Get("employee")

	switch values := data.(type) {
	case nil:
		withGlobals := map[string]any{
			"csrf":        csrf,
			"isPartial":   isPartial,
			"isComponent": isComponent,
			"identity":    t.Identity,
		}
		if employee != nil {
			withGlobals["employee"] = employee
		}
		if feedback != nil {
			withGlobals["feedback"] = feedback
		}
		if err := t.addSuppliers(withGlobals, employee, c, needsSuppliers); err != nil {
			return nil, err
		}
		return withGlobals, nil
	case map[string]any:
		withGlobals := make(map[string]any, len(values)+4)
		maps.Copy(withGlobals, values)
		if _, ok := withGlobals["csrf"]; !ok {
			withGlobals["csrf"] = csrf
		}
		if _, ok := withGlobals["isPartial"]; !ok {
			withGlobals["isPartial"] = isPartial
		}
		if _, ok := withGlobals["isComponent"]; !ok {
			withGlobals["isComponent"] = isComponent
		}
		if _, ok := withGlobals["identity"]; !ok {
			withGlobals["identity"] = t.Identity
		}
		if _, ok := withGlobals["employee"]; !ok && employee != nil {
			withGlobals["employee"] = employee
		}
		if _, ok := withGlobals["feedback"]; !ok && feedback != nil {
			withGlobals["feedback"] = feedback
		}
		if err := t.addSuppliers(withGlobals, withGlobals["employee"], c, needsSuppliers); err != nil {
			return nil, err
		}
		return withGlobals, nil
	case map[string]string:
		withGlobals := make(map[string]any, len(values)+4)
		for key, value := range values {
			withGlobals[key] = value
		}
		if _, ok := withGlobals["csrf"]; !ok {
			withGlobals["csrf"] = csrf
		}
		if _, ok := withGlobals["isPartial"]; !ok {
			withGlobals["isPartial"] = isPartial
		}
		if _, ok := withGlobals["isComponent"]; !ok {
			withGlobals["isComponent"] = isComponent
		}
		if _, ok := withGlobals["identity"]; !ok {
			withGlobals["identity"] = t.Identity
		}
		if _, ok := withGlobals["employee"]; !ok && employee != nil {
			withGlobals["employee"] = employee
		}
		if _, ok := withGlobals["feedback"]; !ok && feedback != nil {
			withGlobals["feedback"] = feedback
		}
		if err := t.addSuppliers(withGlobals, withGlobals["employee"], c, needsSuppliers); err != nil {
			return nil, err
		}
		return withGlobals, nil
	default:
		return data, nil
	}
}

func (t *Template) addSuppliers(data map[string]any, employee any, c echo.Context, needsSuppliers bool) error {
	if !needsSuppliers || t.Suppliers == nil {
		return nil
	}

	if _, ok := data["Suppliers"]; ok {
		return nil
	}

	if existing, ok := data["suppliers"]; ok {
		data["Suppliers"] = existing
		return nil
	}

	empl, ok := employee.(companies.Employee)
	if !ok {
		return nil
	}

	suppliers, err := t.Suppliers.GetAllByCompany(c, int32(empl.CompanyId))
	if err != nil {
		return err
	}

	data["Suppliers"] = suppliers
	return nil
}

func templateFiles(page string) ([]string, error) {
	page = filepath.Clean(page)

	rel, err := filepath.Rel("templates", page)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(rel, string(filepath.Separator))

	// HTMX partials:
	// templates/partials/*
	// skip base layouts entirely
	if len(parts) > 0 && parts[0] == "partials" {
		files := []string{}

		err = filepath.WalkDir("templates/components", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() || filepath.Ext(path) != ".html" {
				return nil
			}

			files = append(files, path)
			return nil
		})
		if err != nil {
			return nil, err
		}

		err = filepath.WalkDir("templates/partials", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() || filepath.Ext(path) != ".html" {
				return nil
			}

			if path != page {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}

		files = append(files, page)
		return files, nil
	}

	// Component templates:
	// templates/components/*
	// skip base layouts entirely
	if len(parts) > 0 && parts[0] == "components" {
		files := []string{}

		err = filepath.WalkDir("templates/components", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() || filepath.Ext(path) != ".html" {
				return nil
			}

			files = append(files, path)
			return nil
		})
		if err != nil {
			return nil, err
		}

		err = filepath.WalkDir("templates/partials", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() || filepath.Ext(path) != ".html" {
				return nil
			}

			files = append(files, path)
			return nil
		})
		if err != nil {
			return nil, err
		}

		files = append(files, page)
		return files, nil
	}

	base, err := templateBase(page)
	if err != nil {
		return nil, err
	}

	baseFile := filepath.Join("templates", base+".html")
	files := []string{baseFile}

	err = filepath.WalkDir("templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || filepath.Ext(path) != ".html" {
			return nil
		}

		path = filepath.Clean(path)

		if path == filepath.Clean(page) || path == filepath.Clean(baseFile) {
			return nil
		}

		rel, err := filepath.Rel("templates", path)
		if err != nil {
			return err
		}

		dir := strings.Split(rel, string(filepath.Separator))[0]

		if dir == "pages" || dir == rel {
			return nil
		}

		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	files = append(files, page)
	return files, nil
}

func templateBase(page string) (string, error) {
	contents, err := os.ReadFile(page)
	if err != nil {
		return "", err
	}

	for _, base := range []string{"adminBase", "businessBase", "centerBase"} {
		if strings.Contains(string(contents), `template "`+base+`"`) {
			return base, nil
		}
	}

	return "base", nil
}
