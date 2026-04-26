package configs

import (
	"html/template"
	"io"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type Template struct{}

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	files := []string{"templates/base.html"}

	components, err := filepath.Glob("templates/components/*.html")
	if err != nil {
		return err
	}
	files = append(files, components...)
	files = append(files, "templates/pages/"+name+".html")

	tmpl := template.Must(template.ParseFiles(files...))
	return tmpl.ExecuteTemplate(w, name, data)
}
