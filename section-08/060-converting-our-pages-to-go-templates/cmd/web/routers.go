package main

import (
	"net/http"

	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/060-converting-our-pages-to-go-templates/pkg/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/060-converting-our-pages-to-go-templates/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func setupRoutes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation", handlers.Repo.Reservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}