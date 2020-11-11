package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/models"
)

var tmple = template.Must(template.ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/employee/index.html"))

func EmployeeGet(w http.ResponseWriter, req *http.Request) {

	d := models.Empleado{Name: "Angel", City: "Juliaca"}

	err := tmple.ExecuteTemplate(w, "employee/indexPage", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func EmployeeList(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "employe list page ooooo ")
}
