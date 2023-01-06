package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/middleware"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func MainRouter() chi.Router {
	ur := inmemory.UserRepositoryInMemory{}
	tr := inmemory.TaskRepositoryInMemory{}	
	r := chi.NewRouter()

	r.Get("/health", HealthCheck)
	UserRouter(r, &ur)
	TaskRouter(r, &tr, middleware.AuthHandler(&ur))

	return r
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status": "ok"}`)
}