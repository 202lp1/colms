package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewAlumno struct {
	Name    string
	IsEdit  bool
	Data    models.Alumno
	Widgets []models.Alumno
}

var tmpla = template.Must(template.New("foo").Funcs(cfig.FuncMap).ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/alumno/index.html", "web/alumno/form.html"))

func AlumnoList(w http.ResponseWriter, req *http.Request) {
	// Create
	//cfig.DB.Create(&models.Alumno{Name: "Juan", City: "Juliaca"})
	lis := []models.Alumno{}
	if err := cfig.DB.Find(&lis).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//log.Printf("lis: %v", lis)
	data := ViewAlumno{
		Name:    "Alumno",
		Widgets: lis,
	}

	err := tmpla.ExecuteTemplate(w, "alumno/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AlumnoForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	var d models.Alumno
	IsEdit := false
	if id != "" {
		IsEdit = true
		if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == "POST" {
		log.Printf("POST id=: %v", id)
		d.Nombres = r.FormValue("nombres")
		d.Codigo = r.FormValue("codigo")
		if id != "" {
			cfig.DB.Save(&d)
		} else {
			cfig.DB.Create(&d)
		}
		http.Redirect(w, r, "/alumno/index", 301)
	}

	data := ViewAlumno{
		Name:   "Alumno",
		Data:   d,
		IsEdit: IsEdit,
	}

	err := tmpla.ExecuteTemplate(w, "alumno/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AlumnoDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]//log.Printf("del id=: %v", id)
	var d models.Alumno
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cfig.DB.Unscoped().Delete(&d)
	http.Redirect(w, r, "/alumno/index", 301)
}
