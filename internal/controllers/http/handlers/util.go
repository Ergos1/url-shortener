package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func ParseShortUrl(r *http.Request) string {
	url := chi.URLParam(r, "shortUrl")
	return url
}
