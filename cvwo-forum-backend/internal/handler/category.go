package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pohanson/cvwo-forum/internal/repository"
)

func GetAllCategoryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	repo := repository.GetPgRepo()
	result, err := repo.Category.GetAll(ctx)

	if err != nil {
		log.Println("Error reading category", err)
		http.Error(w, "Unknown Error", http.StatusBadRequest)
		return
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Println("Error converting json result: ", err)
		http.Error(w, "Unknown Error", http.StatusBadRequest)
		return
	}
	w.Write(jsonResult)
}
