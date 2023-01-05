package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func MainRouter() chi.Router {
	r := chi.NewRouter()
	UserRouter(r)
	TaskRouter(r)

	r.Get("/health", HealthCheck)

	return r
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}