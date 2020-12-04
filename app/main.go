package main

import (
	//"fmt"
	"log"
	"net/http"
	"os"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/controllers"
	"github.com/202lp1/colms/models"
	"github.com/gorilla/mux" //gin
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Hello World!")
	//})
	cfig.DB, err = connectDBmysql()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	log.Printf("db is connected: %v", cfig.DB)

	// Migrate the schema
	cfig.DB.AutoMigrate(&models.Empleado{})
	cfig.DB.AutoMigrate(&models.Alumno{})
	cfig.DB.AutoMigrate(&models.Matricula{})
	//cfig.DB.Create(&models.Empleado{Name: "Juan", City: "Juliaca"})

	r := NewRouter()

	//fs := http.FileServer(http.Dir("assets/"))
    //http.Handle("/static/", http.StripPrefix("/static/", fs))
    //r := mux.NewRouter().StrictSlash(true)
 	//r.Handle("/", http.FileServer(http.Dir("assets/")))




	r.HandleFunc("/", controllers.Home).Methods("GET")

	r.HandleFunc("/item/index", controllers.ItemList).Methods("GET")

	r.HandleFunc("/employee/index", controllers.EmployeeList).Methods("GET")
	r.HandleFunc("/employee/form", controllers.EmployeeForm).Methods("GET", "POST")
	r.HandleFunc("/employee/delete", controllers.EmployeeDel).Methods("GET")

	r.HandleFunc("/alumno/index", controllers.AlumnoList).Methods("GET")
	r.HandleFunc("/alumno/form", controllers.AlumnoForm).Methods("GET", "POST")
	r.HandleFunc("/alumno/delete", controllers.AlumnoDel).Methods("GET")

	r.HandleFunc("/matricula/index", controllers.MatriculaList).Methods("GET")
	r.HandleFunc("/matricula/form", controllers.MatriculaForm).Methods("GET", "POST")
	r.HandleFunc("/matricula/delete", controllers.MatriculaDel).Methods("GET")

	//http.ListenAndServe(":80", r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("port: %v", port)
	http.ListenAndServe(":"+port, r)

}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Choose the folder to serve
	staticDir := "/assets/"

	// Create the route
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	return router
}


func connectDBmysql() (c *gorm.DB, err error) {
	dsn := "docker:docker@tcp(mysql-db:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "docker:docker@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return conn, err
}

func connectDB() (c *gorm.DB, err error) {
	////dsn := "docker:docker@tcp(mysql-db:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "docker:docker@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := "user=rdcrofkuaysjfd password=6b7d82e87153eaa84691f4fc26edc8fe38776c193fd8155d19cbcc6c031ac4b3 host=ec2-54-158-190-214.compute-1.amazonaws.com dbname=d83unplc9i729a port=5432 sslmode=require TimeZone=Asia/Shanghai"
	//dsn := "user=postgres password=postgres2 dbname=users_test host=localhost port=5435 sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return conn, err
}
