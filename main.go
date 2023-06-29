package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
	fmt.Fprintln(w, "Todo index!")
}

func todoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
