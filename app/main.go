package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/controllers"
	"github.com/202lp1/colms/models"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {

	dsn := "docker:docker@tcp(mysql-db:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	cfig.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("failed to connect database %v", err)
	}

	log.Printf("db is connected: %v", cfig.DB)
	// Migrate the schema
	cfig.DB.AutoMigrate(&models.Empleado{})

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
