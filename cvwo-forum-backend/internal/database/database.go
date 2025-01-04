package database

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var dbInstance *sql.DB
var once sync.Once

func GetDb() *sql.DB {
	if dbInstance == nil {
		once.Do(func() {
			connStr := os.Getenv("POSTGRESDB_URL")
			db, err := sql.Open("postgres", connStr)
			if err != nil {
				log.Fatalln(err)
			} else {
				dbInstance = db
			}
		})
	}
	return dbInstance
}
func MakeMigration(dbName string) {
	if dbInstance == nil {
		GetDb()
	}
	driver, err := postgres.WithInstance(dbInstance, &postgres.Config{})
	if err != nil {
		log.Fatalln("Error getting postgres instance: ", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://internal/database/migration", dbName, driver)
	if err != nil {
		log.Fatalln("Creating new migration instance error: ", err)
	}
	switch err := m.Up(); {
	case err == nil:
		log.Println("Applied Migrations")
	case err == migrate.ErrNoChange:
		log.Println("No Migrations Made")
	default:
		log.Fatalln("Error applying migration: ", err)
	}
}
