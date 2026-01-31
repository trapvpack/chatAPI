package main

import (
	database "chatAPI/internal/db"
	"chatAPI/internal/router"
	_ "chatAPI/migrations"
	"net/http"

	"github.com/pressly/goose/v3"
)

func main() {
	if err := database.ConnectToDatabase(); err != nil {
		panic(err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.Up(database.SQL_DB, "migrations"); err != nil {
		panic(err)
	}

	mux := router.New()

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
