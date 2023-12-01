package handlers

import (
	"td-project/models"
	"encoding/json"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetTasks()
	if err != nil {
		log.Printf("Ocorreu um erro ao obter tasks: %v", err)
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}