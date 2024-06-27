package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kar5960/url_shortener/database"
	"github.com/kar5960/url_shortener/helpers"
)

type ShortenRequest struct {
	LongURL string `json:"long_url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if req.LongURL == "" {
		http.Error(w, "empty long_url", http.StatusBadRequest)
		return
	}
	shortURL := helpers.GenerateShortURL(req.LongURL)
	err := database.StoreURL(shortURL, req.LongURL)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	resp := ShortenResponse{ShortURL: shortURL}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
}

func Resolve(w http.ResponseWriter, r *http.Request) {
	shortURL := strings.TrimPrefix(r.URL.Path, "/")

	longURL, err := database.RetrieveURL(shortURL)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
