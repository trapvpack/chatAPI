package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"chatAPI/internal/usecase"
)

type ChatHandler struct {
	uc *usecase.ChatUsecase
}

func NewChatHandler(uc *usecase.ChatUsecase) *ChatHandler {
	return &ChatHandler{uc: uc}
}

type createChatRequest struct {
	Title string `json:"title"`
}

func (h *ChatHandler) CreateChat(w http.ResponseWriter, r *http.Request) {
	var req createChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	chat, err := h.uc.CreateChat(r.Context(), req.Title)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidTitle) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(chat)
}

func (h *ChatHandler) GetChat(w http.ResponseWriter, r *http.Request) {
	id, err := parseChatID(r.URL.Path)
	if err != nil {
		http.Error(w, "invalid chat id", http.StatusBadRequest)
		return
	}

	limit := parseLimit(r)

	chat, err := h.uc.GetChat(r.Context(), id, limit)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chat)
}

func (h *ChatHandler) DeleteChat(w http.ResponseWriter, r *http.Request) {
	id, err := parseChatID(r.URL.Path)
	if err != nil {
		http.Error(w, "invalid chat id", http.StatusBadRequest)
		return
	}

	if err := h.uc.DeleteChat(r.Context(), id); err != nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
