package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func parseChatID(path string) (uint, error) {
	parts := strings.Split(path, "/")
	if len(parts) != 3 {
		return 0, errors.New("invalid path")
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil || id <= 0 {
		return 0, errors.New("invalid id")
	}

	return uint(id), nil
}

func parseLimit(r *http.Request) int {
	l := r.URL.Query().Get("limit")
	if l == "" {
		return 20
	}

	val, err := strconv.Atoi(l)
	if err != nil || val <= 0 || val > 100 {
		return 20
	}

	return val
}
