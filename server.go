package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User object
type User struct {
	Name  string
	Email string
	ID    int
}

// Users array
type Users []User

func allUsers(w http.ResponseWriter, r *http.Request) {

	users := Users{
		User{Name: "Marc", Email: "777marc@gmail.com", ID: 1},
	}

	fmt.Println("all users endpoint")
	json.NewEncoder(w).Encode(users)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", allUsers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
