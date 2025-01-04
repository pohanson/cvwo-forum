package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/pohanson/cvwo-forum/internal/model"
	"github.com/pohanson/cvwo-forum/internal/repository"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := model.User{}
	err := user.DecodeFromJson(r.Body)
	if err != nil {
		log.Println(err)
		writeDecodingError(err, w)
		return
	}
	user.Role = 1
	repo := repository.GetPgRepo()
	result, err := repo.User.Create(ctx, &user)
	if err != nil {
		log.Println("Error creating user:", err)
		writeDbInsertError(err, w)
		return
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Println("Error converting json result: ", err)
		http.Error(w, "Unknown Error", http.StatusBadRequest)
	}
	w.Write(jsonResult)
}

func writeDecodingError(err error, w http.ResponseWriter) {
	if err != nil {
		var maxLengthError *model.MaxLengthViolationErr
		var missingFieldError *model.MissingFieldErr
		var syntaxError *json.SyntaxError
		switch {
		case errors.As(err, &maxLengthError):
			fallthrough
		case errors.As(err, &missingFieldError):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.As(err, &syntaxError):
			http.Error(w, fmt.Sprint("Error decoding json: ", err.Error()), http.StatusBadRequest)
		default:
			http.Error(w, "Error decoding json", http.StatusBadRequest)
		}
	}
}

func writeDbInsertError(err error, w http.ResponseWriter) {
	var uniqueViolationErr *repository.UniqueViolationErr
	switch {
	case errors.As(err, &uniqueViolationErr):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, "Unknown Error", http.StatusBadRequest)
	}
}
