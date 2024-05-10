package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mcgtrt/gtx/view/home"
)

type Handler struct {
	// Here access to the database, etc
}

func (h *Handler) HandleHome(c *Ctx) error {
	return render(c, home.Home())
}

type Ctx struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func (c *Ctx) Context() context.Context {
	return c.Request.Context()
}

type apiFunc func(ctx *Ctx) error

// Wrap handler function with sendiforce router context signature
func makeHandlerFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := &Ctx{
			Response: w,
			Request:  r,
		}

		if err := fn(c); err != nil {
			writeJSON(w, 500, map[string]any{
				"error":      "internal server error",
				"msg":        err.Error(),
				"statusCode": 500,
			})
		}
	}
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
