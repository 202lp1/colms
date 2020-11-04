package main

import (
	"fmt"
	"net/http"

	//"text/template"

	//"github.com/202lp1/colms/models"
	"github.com/202lp1/colms/controllers"
	"github.com/gorilla/mux"
)

//var tmpl = template.Must(template.ParseGlob("web/*"))

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/login", login)
	r.HandleFunc("/tabla", controllers.Tablalist)
	r.HandleFunc("/tabla/{id}", controllers.Tablaget)

	http.ListenAndServe(":8090", r)
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Home page ")
}

func login(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Login page ")
}
