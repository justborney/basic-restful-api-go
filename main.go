package main

import (
	"encoding/json"
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

type UserService struct {
	users map[int]*User
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[int]*User),
	}
}

func main() {
	userService := NewUserService()
	user1 := &User{ID: 1, Token: "token1", Name: "John Doe", Age: 25}
	user2 := &User{ID: 2, Token: "token2", Name: "Jane Smith", Age: 30}
	userService.AddUser(user1)
	userService.AddUser(user2)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", index)
	router.HandleFunc("/user/{id:[0-9]+}", userService.GetUser).Methods("GET")
	router.HandleFunc("/user/{id:[0-9]+}", userService.UpdateUser).Methods("POST")

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Available endpoints\nGET /user/\nPOST /user/")
}

// GetUser обрабатывает GET запрос
func (us *UserService) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		log.Printf("Invalid user ID", err.Error())
		return
	}

	user, ok := us.users[id]
	if !ok {
		http.NotFound(w, r)
		log.Print("NotFound")
		return
	}

	// Кодируем User в JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser обрабатывает POST запрос
func (us *UserService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		log.Print("Invalid user ID", err.Error())
		return
	}

	user, ok := us.users[id]
	if !ok {
		http.NotFound(w, r)
		log.Print("NotFound")
		return
	}

	// Декодируем тело запроса
	var newUser User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		log.Print("Invalid JSON", err.Error())
		return
	}

	// Проверка на то, что ID и токен не изменены
	if newUser.ID != id || newUser.Token != user.Token {
		http.Error(w, "Cannot change ID or Token", http.StatusBadRequest)
		log.Print("Cannot change ID or Token")
		return
	}

	// Обновляем информацию
	user.Name = newUser.Name
	user.Age = newUser.Age

	fmt.Fprint(w, "User info updated succesfully")
	log.Printf("User info updated succesfully. ID=%d", id)
}

func (us *UserService) AddUser(user *User) {
	us.users[user.ID] = user
}
