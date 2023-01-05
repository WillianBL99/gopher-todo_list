package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/middleware"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TaskRouter(r chi.Router) {
	r.With(middleware.AuthHandler(&inmemory.UserRepositoryInMemory{})).Route("/tasks", func(r chi.Router) {
		r.Get("/", GetTasks)
	})
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all tasks"))
}