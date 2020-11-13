package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/controllers"
	"github.com/gorilla/mux"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {

	cfig.DB, err = connectDB()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	log.Printf("db is connected: %v", cfig.DB)

	// Migrate the schema
	//cfig.DB.AutoMigrate(&models.Empleado{})

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

func connectDB() (c *gorm.DB, err error) {
	dsn := "docker:docker@tcp(mysql-db:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "docker:docker@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return conn, err
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Home page x2222")
}

func login(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Login page ")
}
