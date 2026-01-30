package main

import (
	database "chatAPI/internal/db"
	"chatAPI/internal/handler"
	_ "chatAPI/migrations"
	"net/http"
	"strings"

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

	http.HandleFunc("/chat", handler.CreateChat)

	http.HandleFunc("/chat/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/message") && r.Method == http.MethodPost {
			handler.CreateMessage(w, r)
			return
		}

		if r.Method == http.MethodGet {
			handler.GetChat(w, r)
			return
		}

		http.NotFound(w, r)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
