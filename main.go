package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/login", login)
	//http.HandleFunc("/tabla", tablalist) /{id}
	r.HandleFunc("/tabla/{id}", tablaget)

	http.ListenAndServe(":8090", r)
}
func tablaget(w http.ResponseWriter, req *http.Request) {
	//params := req.URL.Query() //gin.Param("id") ["param1"]
	vars := mux.Vars(req)

	//fmt.Println("id=", params.Get("id"))
	//fmt.Fprintf(w, "tablaget page ", params.Get("id"))
	fmt.Println("id=", vars["id"])
	fmt.Fprintf(w, "tablaget page ", vars["id"])
}

func tablalist(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "tablalist page ")
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Home page ")
}

func login(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Login page ")
}
