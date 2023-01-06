package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/controller"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/middleware"
)

func TaskRouter(
	r chi.Router,
	tr repository.TaskRepository,
	md *middleware.Middleware,
) {
	tc := controller.NewTaskController(tr)

	r.With(md.Auth).Route("/task", func(r chi.Router) {
		r.Post("/", tc.SaveTask)
		r.Route("/{taskId}", func(r chi.Router) {
			r.Use(md.TaskParams)
			r.Put("/", tc.UpdateTask)
		})
	})
}
