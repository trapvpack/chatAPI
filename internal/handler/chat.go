package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"chatAPI/internal/db"
	"chatAPI/internal/model"

	"gorm.io/gorm"
)

type createChatRequest struct {
	Title string `json:"title"`
}

func CreateChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req createChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	chat := model.Chat{
		Title: req.Title,
	}

	if err := database.GORM_DB.Create(&chat).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(chat)
}

func GetChat(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.NotFound(w, r)
		return
	}

	chatID, err := strconv.Atoi(parts[2])
	if err != nil || chatID <= 0 {
		http.Error(w, "invalid chat ID", http.StatusBadRequest)
		return
	}

	limit := 20
	if l := r.URL.Query().Get("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil && val > 0 && val <= 100 {
			limit = val
		}
	}

	var chat model.Chat
	if err := database.GORM_DB.Preload("Messages", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at desc").Limit(limit)
	}).First(&chat, chatID).Error; err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chat)
}
