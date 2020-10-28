package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	//http.HandleFunc("/tabla", tablalist) /{id}
	http.HandleFunc("/tabla", tablaget)

	http.ListenAndServe(":8090", nil)
}
func tablaget(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query() //req.Param("id")
	fmt.Println("id=", id)
	fmt.Fprintf(w, "tablaget page ", id)
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
