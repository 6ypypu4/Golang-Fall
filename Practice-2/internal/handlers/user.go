package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		handleGetUser(w, r)
	case http.MethodPost:
		handlePostUser(w, r)
	default:
		http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

// GET /user
func handleGetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid id"})
		return
	}

	w.WriteHeader(http.StatusOK)
	//тут нужно добавить подтягивание из /data/users.json
	json.NewEncoder(w).Encode(map[string]int{"user_id": int(id)})
}

// POST /user
func handlePostUser(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil || data.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid name"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	//тут нужно добавить подтягивание в /data/users.json через try

	json.NewEncoder(w).Encode(map[string]string{"created": data.Name})
}

// работа с данными

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var users []User
var path string = "Practice-2/data/users.json"

func loadData(dataPath string) {
	f, err := os.Open("data/users.json")
	if err != nil {
		log.Println(err)
		return
	}

	json.NewDecoder(f).Decode(&users)

	f.Close()
}

func saveData(dataPath string) {
	f, err := os.Create("data/users.json")
	if err != nil {
		log.Println(err)
		return
	}
	json.NewEncoder(f).Encode(users)

	f.Close()
}
func Init() {
	loadData(os.Args[1])
}
