package main

import (
	"net/http"

	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-06/044-setting-up-our-project/pkg/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-06/044-setting-up-our-project/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func setupRoutes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
