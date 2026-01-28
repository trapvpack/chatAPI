package main

import (
	database "chatAPI/internal/db"
	_ "chatAPI/migrations"

	"github.com/pressly/goose/v3"
)

func main() {
	err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.Up(database.SQL_DB, "migrations"); err != nil {
		panic(err)
	}
}
