package main

import (
	"github.com/go-chi/chi/v5"
	"web3/pkg/config"
	"web3/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler{
	// Mux http multiplexer
	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	return mux
}