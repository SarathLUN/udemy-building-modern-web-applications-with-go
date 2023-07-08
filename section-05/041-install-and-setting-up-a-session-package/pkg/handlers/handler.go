package handlers

import (
	"net/http"

	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-05/041-install-and-setting-up-a-session-package/pkg/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-05/041-install-and-setting-up-a-session-package/pkg/models"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-05/041-install-and-setting-up-a-session-package/pkg/renderers"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers set the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home : handler function for "/"
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About : handler function for "/about"
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"
	// send data to template
	renderers.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
