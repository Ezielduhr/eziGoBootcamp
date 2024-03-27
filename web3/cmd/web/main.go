package main

import (
	"fmt"
	scs "github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"encoding/gob"
	"time"
	config "web3/pkg/config"
	handlers "web3/pkg/handlers"
	"web3/models"
)

var sessionManager *scs.SessionManager
var app config.AppConfig

func main() {

	gob.Register(models.Article{})

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = false
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	app.Session = sessionManager

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	fmt.Println("Starting server, Listening on http://localhost:8080")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
