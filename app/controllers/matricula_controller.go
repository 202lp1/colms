package controllers

import (
	"log"
	"net/http"
	"strconv"
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
		n, err := strconv.Atoi(r.FormValue("alumno_id"))
		//if err == nil {
		//	fmt.Printf("%d of type %T", n, n)
		//}
		if err != nil {
			log.Printf("Invalid ID: %v - %v\n", n, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			//n = 1
			return
		}

		d.AlumnoId = n
		if id != "" {
			if err := cfig.DB.Save(&d).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return //err
			}

		} else { //https://gorm.io/fr_FR/docs/transactions.html
			if err := cfig.DB.Create(&d).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return //err
			}
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

	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		//log.Printf("No save  %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}

	http.Redirect(w, r, "/matricula/index", 301)
}
