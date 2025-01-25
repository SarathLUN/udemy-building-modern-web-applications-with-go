package main

import (
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/075-writing-tests-for-our-POST-handlers/internal/config"
	"github.com/go-chi/chi/v5"
	"testing"
)

func TestSetupRoutes(t *testing.T) {
	var app config.AppConfig
	mux := setupRoutes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
	// do nothing;
	default:
		t.Errorf("want chi.Mux, got %T", v)
	}
}
