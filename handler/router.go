package handler

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func NewRouter(h *Handler) *mux.Router {
	r := mux.NewRouter()
	r.Use(routingMiddleware)

	mountRoutes(r, h)

	return r
}

func mountRoutes(r *mux.Router, h *Handler) {
	r.HandleFunc("/", makeHandlerFunc(h.HandleHome)).Methods(http.MethodGet)
	r.HandleFunc("/contact", makeHandlerFunc(h.HandleContact)).Methods(http.MethodGet)
}

// Middleware that takes care of extra "/" symbol at the end of URL path (if provided)
func routingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := *r.URL
		url.Path = strings.TrimSuffix(r.URL.Path, "/")
		r.URL = &url

		h.ServeHTTP(w, r)
	})
}

// Render Templ Component
func render(c *Ctx, component templ.Component) error {
	return component.Render(c.Context(), c.Response)
}
