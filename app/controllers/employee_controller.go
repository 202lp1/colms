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
	//cfig.DB.Create(&models.Empleado{Name: "Juan", City: "Juliaca"})

	lis := []models.Empleado{}
	if err := cfig.DB.Find(&lis).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

func EmployeeGet(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	//db.First(&product, "code = ?", "D42") // find product with code D42
	var d models.Empleado
	if id != "" {
		if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == "POST" {
		d.Name = r.FormValue("name")
		d.City = r.FormValue("city")
		if id != "" {
			cfig.DB.Save(&d)
		} else {
			cfig.DB.Create(&d)
		}
		http.Redirect(w, r, "/employee/index", 301)
	}

	err := tmple.ExecuteTemplate(w, "employee/formPage", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
