package configs

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestRoutedTemplatesHaveMatchingDefinitions(t *testing.T) {
	chdirRepoRoot(t)

	names := []string{
		"adminCompanyDetails",
		"adminCompanyInvites",
		"adminCompanyOverview",
		"app/companySettings",
		"app/newOrder",
		"app/orderHistory",
		"app/suppliers",
		"app/userSettings",
		"auth/forgot-password",
		"auth/login",
		"auth/select-company",
		"auth/signup",
		"error",
		"partials/companyList",
		"partials/inviteList",
		"support/contact",
	}

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			page := templatePath(name)
			if _, err := os.Stat(page); err != nil {
				t.Fatalf("template page %q: %v", page, err)
			}

			files, err := templateFiles(page)
			if err != nil {
				t.Fatalf("template files: %v", err)
			}

			tmpl, err := template.ParseFiles(files...)
			if err != nil {
				t.Fatalf("parse templates: %v", err)
			}

			executeName := templateName(name)
			if tmpl.Lookup(executeName) == nil {
				t.Fatalf("missing template definition %q in %s", executeName, page)
			}
		})
	}
}

func TestTemplateRouteFunc(t *testing.T) {
	e := echo.New()
	e.GET("/test/:id", func(c echo.Context) error { return nil }).Name = "test.route"

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	tmpl := &Template{}
	routeFunc, ok := tmpl.templateFuncMap(c)["route"].(func(string, ...interface{}) string)
	if !ok {
		t.Fatal("expected route func in template func map")
	}

	got := routeFunc("test.route", 123)
	want := "/test/123"
	if got != want {
		t.Fatalf("expected %q got %q", want, got)
	}
}

func chdirRepoRoot(t *testing.T) {
	t.Helper()

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("locate test file")
	}

	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
	if err := os.Chdir(root); err != nil {
		t.Fatalf("chdir repo root: %v", err)
	}
}
