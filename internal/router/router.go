package router

import (
	"chatAPI/internal/handler"
	"net/http"
	"strings"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.CreateChat(w, r)
			return
		}
		http.NotFound(w, r)
	})

	mux.HandleFunc("/chat/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			handler.GetChat(w, r)
			return

		case http.MethodDelete:
			handler.DeleteChat(w, r)
			return

		case http.MethodPost:
			if strings.HasSuffix(r.URL.Path, "/message") {
				handler.CreateMessage(w, r)
				return
			}
		}

		http.NotFound(w, r)
	})

	return mux
}
