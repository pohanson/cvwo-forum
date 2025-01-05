package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pohanson/cvwo-forum/internal/usersession"
)

func Setup() chi.Router {

	r := chi.NewRouter()
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Logger)
	r.Use(usersession.SessionUserMiddleware)
	r.Use(middleware.Recoverer)
	return r
}
