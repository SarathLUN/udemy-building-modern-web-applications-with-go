package handlers

import (
	"fmt"
	"net/http"

	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/062-creating-a-handler-that-return-json/pkg/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/062-creating-a-handler-that-return-json/pkg/models"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/062-creating-a-handler-that-return-json/pkg/renderers"
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
	m.App.Session.Put(r.Context(), "remote_ip", r.RemoteAddr)
	renderers.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About : handler function for "/about"
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"
	// query client IP address and pass to template
	stringMap["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")
	// send data to template
	renderers.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability search for availability
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")
	w.Write([]byte(fmt.Sprintf("start date is %s, and end date is %s", start, end)))
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
