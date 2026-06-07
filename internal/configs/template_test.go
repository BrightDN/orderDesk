package configs

import (
	"html/template"
	"os"
	"path/filepath"
	"runtime"
	"testing"
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
