package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/users", getUsers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{Name: "HieuAM", Email: "hieuam@example.com"},
		{Name: "DevOps", Email: "devops@example.com"},
	}

	json.NewEncoder(w).Encode(users)
}
