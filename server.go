package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name  string
	Email string
	Id    string
}

type Users []User

func allUsers(w http.ResponseWriter, r *http.Request) {
	
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
