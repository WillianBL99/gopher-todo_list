package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/controller"
)

func TaskRouter(
	r chi.Router,
	tr repository.TaskRepository,
	auth func(next http.Handler) http.Handler,
) {
	tc := controller.NewTaskController(tr)

	r.With(auth).Route("/tasks", func(r chi.Router) {
		r.Post("/", tc.SaveTask)
	})
}
