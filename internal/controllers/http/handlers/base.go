package handlers

import (
	"net/http"

	"github.com/Ergos1/url-shortener.git/internal/controllers/http/dtos"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type BaseHandler struct {
}

func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

func (h *BaseHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.Check)

	return r
}

func (h *BaseHandler) Check(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, dtos.Response{Data: "OK"})
}
