package handlers

import (
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-03/030-Building-a-simple-template-cache/pkg/render"
	"net/http"
)

// Home : handler function for "/"
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About : handler function for "/about"
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
