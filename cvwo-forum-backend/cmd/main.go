package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/pohanson/cvwo-forum/internal/database"
	"github.com/pohanson/cvwo-forum/internal/route"
	"github.com/pohanson/cvwo-forum/internal/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	database.MakeMigration("cvwo-forum")
	defer database.GetDb().Close()

	r := router.Setup()
	route.All(r)
	log.Println("Started server on port 5000 at http://localhost:5000")
	log.Fatalln(http.ListenAndServe("localhost:5000", r))
}
