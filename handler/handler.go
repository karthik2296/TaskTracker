package handler

import (
	"TaskTracker/storage"
	"encoding/json"
	"net/http"
)

type CreateTaskRequest struct {
	Title     string `json:"title"`
	New       bool   `json:"new"`
	Completed bool   `json:"completed"`
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	json.NewEncoder(w).Encode(storage.GetAll())
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	task := storage.NewAdd(req.Title)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalud request", http.StatusBadRequest)
		return
	}

	task := storage.TaskCompleted(req.Title)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
