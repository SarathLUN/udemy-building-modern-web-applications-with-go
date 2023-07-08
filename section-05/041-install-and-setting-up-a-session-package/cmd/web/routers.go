package main

import (
	"net/http"

	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-05/041-install-and-setting-up-a-session-package/pkg/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-05/041-install-and-setting-up-a-session-package/pkg/handlers"
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
