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
	// TODO: simplify routing using mux
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateChat(w, r)
			return
		}
	})

	http.HandleFunc("/chat/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			handler.GetChat(w, r)
			return

		case http.MethodPost:
			if strings.HasSuffix(r.URL.Path, "/message") {
				handler.CreateMessage(w, r)
				return
			}
			http.NotFound(w, r)
			return
		case http.MethodDelete:
			handler.DeleteChat(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
