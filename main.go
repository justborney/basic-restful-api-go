package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

type Todos []Todo

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/todos", todoIndex)
	router.HandleFunc("/todos/{todoId}", todoShow)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func todoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Learn go programming language"},
		Todo{Name: "Get some sleep"},
	}

	json.NewEncoder(w).Encode(todos)
}

func todoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
