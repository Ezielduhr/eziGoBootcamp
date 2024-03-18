package handlers

import (
	"net/http"
	render "web3/pkg/render"
	"web3/models"
	"web3/pkg/config"
)

type Repository struct{
	// used to easily swap out components in the webapp
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(ac *config.AppConfig) *Repository{
	return &Repository{
		App: ac,
	}
}

func NewHandlers(r *Repository){
	Repo = r
}

func (m *Repository) HomeHandler(w http.ResponseWriter, request *http.Request){
	render.RenderTemplate(w, "home_page.tmpl",
	&models.PageData{})
}

func (m *Repository) AboutHandler(w http.ResponseWriter, request *http.Request){

	strMap := make(map[string]string)
	strMap["title"] = "About Us"
	strMap["intro"] = "This page is where we talk about ourselves"
	render.RenderTemplate(w, "about_page.tmpl",
	&models.PageData{StrMap: strMap})
}
