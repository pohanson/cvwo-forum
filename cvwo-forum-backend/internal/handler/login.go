package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/pohanson/cvwo-forum/internal/repository"
)

type loginFormData struct {
	Username string `json:"username"`
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	formData := loginFormData{}
	if err := json.NewDecoder(r.Body).Decode(&formData); err != nil {
		log.Println(err)
		writeDecodingError(err, w)
	}

	repo := repository.GetPgRepo()
	result, err := repo.User.GetByUsername(ctx, formData.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			log.Println(err)
		}
		return
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Println("Error converting json result:", err)
		http.Error(w, "Unknown Error", http.StatusBadRequest)
	}
	w.Write(jsonResult)
}
