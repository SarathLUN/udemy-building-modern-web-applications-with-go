package main

import (
	"net/http"

	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-04/038-Using-package-pat-for-routing/pkg/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-04/038-Using-package-pat-for-routing/pkg/handlers"
	"github.com/bmizerany/pat"
)

func setup_routes(app *config.AppConfig) http.Handler{
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	
	return mux
}