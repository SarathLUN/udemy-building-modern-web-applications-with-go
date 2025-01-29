package renderers

import (
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/077-writing-tests-for-our-Render-package/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	session.Put(r.Context(), "flash", "123") // this will test will pass
	//session.Put(r.Context(), "flash", "1234") // this will fail the test
	result := AddDefaultData(&td, r)
	if result.FlashMessage != "123" {
		t.Error("flash value of 123 not found in the session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates/"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = tc
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	var ww myWriter
	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error rendering template", err)
	}
	err = RenderTemplate(&ww, r, "non-existent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("RenderTemplate should have returned an error")
	}
}

func getSession() (*http.Request, error) {

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates/"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
