package main

import (
	"net/http"

	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-04/039-using-chi-for-routing/pkg/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-04/039-using-chi-for-routing/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func setup_routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
  mux.Use(middleware.Logger)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
