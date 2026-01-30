package handler

import (
	database "chatAPI/internal/db"
	"chatAPI/internal/model"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type createMessageRequest struct {
	Text string `json:"text"`
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 4 || parts[3] != "message" {
		http.NotFound(w, r)
		return
	}

	chatID, err := strconv.Atoi(parts[2]) // <-- правильный индекс
	if err != nil || chatID <= 0 {
		http.Error(w, "Invalid chat ID", http.StatusBadRequest)
		return
	}

	var chat model.Chat
	if err := database.GORM_DB.First(&chat, chatID).Error; err != nil {
		http.NotFound(w, r)
		return
	}

	var req createMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	text := strings.TrimSpace(req.Text)
	if len(text) == 0 || len(text) > 5000 {
		http.Error(w, "invalid text", http.StatusBadRequest)
		return
	}

	message := model.Message{
		ChatID: uint(chatID),
		Text:   text,
	}

	if err := database.GORM_DB.Create(&message).Error; err != nil {
		http.Error(w, "failed to create message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}
