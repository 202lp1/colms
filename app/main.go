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
//var tmpl = template.Must(template.ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl"))

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/login", login)
	r.HandleFunc("/tabla", controllers.Tablalist)
	r.HandleFunc("/tabla/{id}", controllers.Tablaget)
	r.HandleFunc("/item", controllers.Itemlist)
	r.HandleFunc("/item/{id}", controllers.Itemget)

	r.HandleFunc("/employee", controllers.EmployeeList)
	r.HandleFunc("/employee/{id}", controllers.EmployeeGet)

	http.ListenAndServe(":8080", r)
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Home page x2222")
}

func login(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Login page ")
}
