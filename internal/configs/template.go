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
	"github.com/labstack/echo/v4"
)

type Template struct{}

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	page := templatePath(name)
	executeName := templateName(name)
	isPartial := strings.HasPrefix(name, "partials/")

	files, err := templateFiles(page)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, executeName, templateData(data, c, isPartial))
}

func templatePath(name string) string {
	if strings.HasPrefix(name, "pages/") ||
		strings.HasPrefix(name, "partials/") {
		return filepath.Join("templates", name+".html")
	}

	return filepath.Join("templates", "pages", name+".html")
}

func templateName(name string) string {
	return strings.TrimSuffix(filepath.Base(name), filepath.Ext(name))
}

func templateData(data any, c echo.Context, isPartial bool) any {
	csrf, _ := c.Get("csrf").(string)
	feedback, _ := flash.Pop(c)

	switch values := data.(type) {
	case nil:
		withGlobals := map[string]any{
			"csrf":      csrf,
			"isPartial": isPartial,
		}
		if feedback != nil {
			withGlobals["feedback"] = feedback
		}
		return withGlobals
	case map[string]any:
		withGlobals := make(map[string]any, len(values)+2)
		maps.Copy(withGlobals, values)
		if _, ok := withGlobals["csrf"]; !ok {
			withGlobals["csrf"] = csrf
		}
		if _, ok := withGlobals["isPartial"]; !ok {
			withGlobals["isPartial"] = isPartial
		}
		if _, ok := withGlobals["feedback"]; !ok && feedback != nil {
			withGlobals["feedback"] = feedback
		}
		return withGlobals
	case map[string]string:
		withGlobals := make(map[string]any, len(values)+2)
		for key, value := range values {
			withGlobals[key] = value
		}
		if _, ok := withGlobals["csrf"]; !ok {
			withGlobals["csrf"] = csrf
		}
		if _, ok := withGlobals["isPartial"]; !ok {
			withGlobals["isPartial"] = isPartial
		}
		if _, ok := withGlobals["feedback"]; !ok && feedback != nil {
			withGlobals["feedback"] = feedback
		}
		return withGlobals
	default:
		return data
	}
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
