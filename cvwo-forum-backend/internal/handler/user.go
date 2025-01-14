package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pohanson/cvwo-forum/internal/model"
	"github.com/pohanson/cvwo-forum/internal/repository"
	"github.com/pohanson/cvwo-forum/internal/usersession"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := model.User{}
	if err := user.DecodeFromJson(r.Body); err != nil {
		log.Println(err)
		writeDecodingError(err, w)
		return
	}
	if err := user.ValidateAll(); err != nil {
		log.Println(err)
		writeValidationError(err, w)
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
		return
	}
	usersession.PutSesUser(r, result)
	w.Write(jsonResult)
	w.WriteHeader(200)
}
