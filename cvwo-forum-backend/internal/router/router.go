package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Setup() chi.Router {

	r := chi.NewRouter()
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	return r
}
