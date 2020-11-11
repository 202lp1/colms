package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
	"github.com/gorilla/mux"
)

type ViewData struct { //opcional
	Name    string
	Widgets []models.Empleado
}

var tmple = template.Must(template.ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/employee/index.html", "web/employee/form.html"))

func EmployeeList(w http.ResponseWriter, req *http.Request) {

	// Create
	//db.Create(&models.Empleado{Name: "Juan", City: "Juliaca"})

	lis := []models.Empleado{}
	cfig.DB.Find(&lis)
	log.Printf("lis: %v", lis)
	//data := ViewData{
	//	Name:    "John Smith",
	//	Widgets: lis,
	//}
	err := tmple.ExecuteTemplate(w, "employee/indexPage", lis)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func EmployeeGet(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	log.Printf("id=: %v", vars["id"])

	var d models.Empleado
	cfig.DB.First(&d, vars["id"]) // find product with integer primary key

	//d := models.Empleado{Name: "Angel", City: "Juliaca"}

	err := tmple.ExecuteTemplate(w, "employee/formPage", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
