package router

import (
	database "chatAPI/internal/db"
	"chatAPI/internal/handler"
	"chatAPI/internal/repository"
	"chatAPI/internal/usecase"
	"net/http"
	"strings"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	repo := repository.NewChatRepository(database.GORM_DB)
	uc := usecase.NewChatUsecase(repo)
	chatHandler := handler.NewChatHandler(uc)

	mux.HandleFunc("/chats", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			chatHandler.CreateChat(w, r)
			return
		}
		http.NotFound(w, r)
	})

	mux.HandleFunc("/chats/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			chatHandler.GetChat(w, r)
			return

		case http.MethodDelete:
			chatHandler.DeleteChat(w, r)
			return

		case http.MethodPost:
			if strings.HasSuffix(r.URL.Path, "/messages") {
				handler.CreateMessage(w, r)
				return
			}
		}

		http.NotFound(w, r)
	})

	return mux
}
