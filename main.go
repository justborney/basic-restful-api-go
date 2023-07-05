package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// User структура пользователя
type User struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/user/{id}", getUserById).Methods("GET")
	router.HandleFunc("/user/{id}", updateUserById).Methods("POST")

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Available endpoints\nGET /user/\nPOST /user/")
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.ParseInt(vars["id"], 10, 64)
	fmt.Fprintf(w, "id: %d", userId)
}
