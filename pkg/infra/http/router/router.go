package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func MainRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/health", HealthCheck)
	UserRouter(r)
	TaskRouter(r)

	return r
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status": "ok"}`)
}