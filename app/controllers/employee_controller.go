package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ViewData struct {
	Name string
	//Widgets []models.Empleado
}

var tmple = template.Must(template.ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/employee/index.html"))

func EmployeeList(w http.ResponseWriter, req *http.Request) {
	lis := []models.Empleado{}

	/////////

	dsn := "docker:docker@tcp(mysql-db:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:paswd@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect database %v", err)
	}
	///////

	log.Printf("db is connected: %v", db)

	// Migrate the schema
	db.AutoMigrate(&models.Empleado{})

	// Create
	//db.Create(&models.Empleado{Name: "Juan", City: "Juliaca"})

	db.Find(&lis)

	log.Printf("lis: %v", lis)
	//d := ViewData{Name: "Angel"}

	//for lis {

	//}
	fmt.Fprintf(w, "%+v", lis)
	/*d := models.Empleado{Name: "Angel", City: "Juliaca"}
	err := tmple.ExecuteTemplate(w, "employee/indexPage", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/
}

func EmployeeGet(w http.ResponseWriter, req *http.Request) {

	d := models.Empleado{Name: "Angel", City: "Juliaca"}

	err := tmple.ExecuteTemplate(w, "employee/indexPage", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
