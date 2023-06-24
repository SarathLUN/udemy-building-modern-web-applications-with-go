package handlers

import (
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/031-Building-a-more-complex-template-cache/pkg/renderers"
	"net/http"
)

// Home : handler function for "/"
func Home(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "home.page.tmpl")
}

// About : handler function for "/about"
func About(w http.ResponseWriter, r *http.Request) {
	renderers.RenderTemplate(w, "about.page.tmpl")
}
