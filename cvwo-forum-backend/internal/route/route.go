package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pohanson/cvwo-forum/internal/handler"
	"github.com/pohanson/cvwo-forum/internal/usersession"
)

func All(r chi.Router) {
	r.Post("/user", handler.CreateUserHandler)
	r.Post("/login", handler.LoginUserHandler)
	r.Get("/logout", func(w http.ResponseWriter, r *http.Request) { usersession.RemoveSesUser(r); return })
	r.Get("/verifyUser", handler.VerifyUserHandler)
}
