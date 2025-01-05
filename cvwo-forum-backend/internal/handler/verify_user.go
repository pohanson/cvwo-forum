package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pohanson/cvwo-forum/internal/usersession"
)

func VerifyUserHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := usersession.GetUserFromCtx(r.Context())
	if !ok {
		http.Error(w, "User Session Error", http.StatusUnauthorized)
		return
	}
	jsonResult, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	w.Write(jsonResult)
}
