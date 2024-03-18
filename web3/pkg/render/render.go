package render

import (
	"fmt"
	"html/template"
	"net/http"
	"web3/pkg/helpers"
	"web3/models"
)

var tmplCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string, pd *models.PageData){
	var tmpl *template.Template
	var err error
	_, inMap := tmplCache[t]
	if !inMap {
		err = makeTemplaceCache(t)
		helpers.ErrorCheck(err)
		fmt.Println("Template in cache")

		tmpl = tmplCache[t]
		err = tmpl.Execute(w, pd)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func makeTemplaceCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base_layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil{
		return err
	}
	tmplCache[t] = tmpl
	return nil
}