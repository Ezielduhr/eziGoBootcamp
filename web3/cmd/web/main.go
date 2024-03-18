package main

import (
	"fmt"
	"log"
	"net/http"
	"web3/pkg/config"
	handlers "web3/pkg/handlers"

)

func main(){
	var app config.AppConfig
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