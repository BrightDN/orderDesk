package configs

import (
	"html/template"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

type Template struct{}

// func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
// 	files := []string{"templates/*.html"}

// 	components, err := filepath.Glob("templates/components/*.html")
// 	if err != nil {
// 		return err
// 	}
// 	files = append(files, components...)
// 	files = append(files, "templates/pages/"+name+".html")

// 	tmpl := template.Must(template.ParseFiles(files...))
// 	return tmpl.ExecuteTemplate(w, name, data)
// }

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	page := filepath.Join("templates", "pages", name+".html")

	files, err := templateFiles(page)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, name, templateData(data, c))
}

func templateData(data any, c echo.Context) any {
	csrf, _ := c.Get("csrf").(string)

	switch values := data.(type) {
	case nil:
		return map[string]any{
			"csrf": csrf,
		}
	case map[string]any:
		withGlobals := make(map[string]any, len(values)+1)
		for key, value := range values {
			withGlobals[key] = value
		}
		if _, ok := withGlobals["csrf"]; !ok {
			withGlobals["csrf"] = csrf
		}
		return withGlobals
	case map[string]string:
		withGlobals := make(map[string]any, len(values)+1)
		for key, value := range values {
			withGlobals[key] = value
		}
		if _, ok := withGlobals["csrf"]; !ok {
			withGlobals["csrf"] = csrf
		}
		return withGlobals
	default:
		return data
	}
}

func templateFiles(page string) ([]string, error) {
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
