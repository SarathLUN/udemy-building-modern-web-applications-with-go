package main

import (
	"net/http"

	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-04/040-developing-our-own-middleware/pkg/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-04/040-developing-our-own-middleware/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func setup_routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
