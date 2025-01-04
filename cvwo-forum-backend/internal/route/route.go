package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/pohanson/cvwo-forum/internal/handler"
)

func All(r chi.Router) {
	r.Post("/user", handler.CreateUserHandler)
	r.Post("/login", handler.LoginUserHandler)
}
