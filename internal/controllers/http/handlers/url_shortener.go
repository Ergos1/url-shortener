package handlers

import (
	"context"
	"net/http"

	"github.com/Ergos1/url-shortener.git/internal/app/core"
	"github.com/Ergos1/url-shortener.git/internal/controllers/http/dtos"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
)

type Service interface {
	CreateOrGetShortUrl(ctx context.Context, url core.CreateOrGetShortUrlRequest) (string, error)
	GetOriginalUrl(ctx context.Context, shortUrl string) (string, error)
}

type UrlShortenerHandler struct {
	service Service
}

func NewUrlShortenerHandler(service Service) *UrlShortenerHandler {
	return &UrlShortenerHandler{
		service: service,
	}
}

func (h *UrlShortenerHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/shorten", h.CreateOrGetShortUrl)
	r.Get("/{shortUrl}", h.GetOriginalUrl)

	return r
}

func (h *UrlShortenerHandler) CreateOrGetShortUrl(w http.ResponseWriter, r *http.Request) {
	var createOrGetShortUrlReq dtos.CreateOrGetShortUrlRequest
	if err := render.DecodeJSON(r.Body, &createOrGetShortUrlReq); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, dtos.Response{Err: ErrInvalidRequestBody.Error()})
		return
	}

	err := validator.New().Struct(createOrGetShortUrlReq)
	if err != nil {
		render.Status(r, http.StatusUnprocessableEntity)
		render.JSON(w, r, dtos.Response{Err: err.Error()})
		return
	}

	shortUrl, err := h.service.CreateOrGetShortUrl(r.Context(), createOrGetShortUrlReq.ToServiceRequest())
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, dtos.Response{Err: err.Error()})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, dtos.CreateOrGetShortUrlResponse{ShortUrl: shortUrl})
}

func (h *UrlShortenerHandler) GetOriginalUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := ParseShortUrl(r)
	originalUrl, err := h.service.GetOriginalUrl(r.Context(), shortUrl)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, dtos.Response{Err: err.Error()})
		return
	}

	// or better to choose 301 Moved Permanently
	http.Redirect(w, r, originalUrl, http.StatusTemporaryRedirect)
}
