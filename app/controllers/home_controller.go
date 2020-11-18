package controllers

import (
	"net/http"
	"text/template"

	"github.com/202lp1/colms/models"
)

var tmpl = template.Must(template.ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/home/index.html"))

func Home(w http.ResponseWriter, req *http.Request) {

	d := models.Item{Title: "Ping", Notes: "Pong"}

	err := tmpl.ExecuteTemplate(w, "home/indexPage", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
