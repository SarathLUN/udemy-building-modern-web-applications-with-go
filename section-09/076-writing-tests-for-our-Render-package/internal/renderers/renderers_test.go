package renderers

import (
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/076-writing-tests-for-our-Render-package/internal/models"
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
