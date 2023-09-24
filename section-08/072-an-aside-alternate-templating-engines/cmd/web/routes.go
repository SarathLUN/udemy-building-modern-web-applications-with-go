package main

import (
	"github.com/bmizerany/pat"
	"net/http"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-08/072-an-aside-alternate-templating-engines/internal/handlers"
)

// routes handles application routes
func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/send", http.HandlerFunc(handlers.SendData))

	return mux
}
