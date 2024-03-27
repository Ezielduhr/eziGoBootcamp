package main

import (
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"web3/pkg/config"
	"web3/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// Mux http multiplexer
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(LogRequestInfo)

	mux.Use(NoSurf)
	mux.Use(SetupSession)

	mux.Get("/", handlers.Repo.HomeHandler)

	mux.Get("/about", handlers.Repo.AboutHandler)

	mux.Get("/login", handlers.Repo.LoginHandler)

	mux.Get("/page", handlers.Repo.PageHandler)

	mux.Get("/makepost", handlers.Repo.MakePostHandler)
	mux.Post("/makepost", handlers.Repo.PostMakePostHandler)

	mux.Get("/article-received", handlers.Repo.ArticleReceived)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))


	return mux
}
