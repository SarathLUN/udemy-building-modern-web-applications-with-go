package handlers

import (
	"net/http"

	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/035-Optimizing-our-template-cache-by-using-an-application-config/pkg/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/035-Optimizing-our-template-cache-by-using-an-application-config/pkg/renderers"
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
	renderers.RenderTemplate(w, "home.page.html")
}

// About : handler function for "/about"
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "about.page.html")
}
