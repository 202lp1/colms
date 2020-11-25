package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewMatricula struct {
	Name    string
	IsEdit  bool
	Data    models.Matricula
	Widgets []models.Matricula
}

var tmplm = template.Must(template.New("foo").Funcs(cfig.FuncMap).ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/matricula/index.html", "web/matricula/form.html"))

func MatriculaList(w http.ResponseWriter, req *http.Request) {
	// Create
	//cfig.DB.Create(&models.Matricula{Name: "Juan", City: "Juliaca"})
	lis := []models.Matricula{}
	if err := cfig.DB.Find(&lis).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//log.Printf("lis: %v", lis)
	data := ViewMatricula{
		Name:    "Matricula",
		Widgets: lis,
	}

	err := tmplm.ExecuteTemplate(w, "matricula/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func MatriculaForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	var d models.Matricula
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
		d.Semestre = r.FormValue("semestre")
		ida := r.FormValue("alumno_id")
		log.Printf("ida=: %v", ida)

		d.AlumnoId = 3
		if id != "" {
			cfig.DB.Save(&d)
		} else {
			cfig.DB.Create(&d)
		}
		http.Redirect(w, r, "/matricula/index", 301)
	}

	data := ViewMatricula{
		Name:   "Matricula",
		Data:   d,
		IsEdit: IsEdit,
	}

	err := tmplm.ExecuteTemplate(w, "matricula/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func MatriculaDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]//log.Printf("del id=: %v", id)
	var d models.Matricula
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cfig.DB.Unscoped().Delete(&d)
	http.Redirect(w, r, "/matricula/index", 301)
}
