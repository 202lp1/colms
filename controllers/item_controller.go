package controllers

import (
	"fmt"
	"net/http"

	"text/template"

	"github.com/202lp1/colms/models"
	//"github.com/gorilla/mux"
)

var tmpl = template.Must(template.ParseGlob("web/*"))

func Tablaget(w http.ResponseWriter, req *http.Request) {

	//t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	//t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")

	//vars := mux.Vars(req)
	//fmt.Println("id=", vars["id"])
	//fmt.Fprintf(w, "tablaget page ", vars["id"])
	// you access the cached templates with the defined name, not the filename
	d := models.Item{Title: "Sean", Notes: "nnn"}

	err := tmpl.ExecuteTemplate(w, "home/indexPage", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Tablalist(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "tablalist page ")
}
