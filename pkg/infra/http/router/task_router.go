package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
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
	r.Route("/tasks", func(r chi.Router) {
		r.Use(md.Auth)

		r.Get("/", tc.GetAllTasks)
	})

	r.Route("/task", func(r chi.Router) {
		r.Use(md.Auth)

		r.Post("/", tc.SaveTask)

		r.Route("/{taskId}", func(r chi.Router) {
			r.Use(md.TaskParams)
			r.Put("/", tc.UpdateTask)
			r.Delete("/", tc.DeleteTask)
			r.Patch("/done", tc.MoveTask(entity.Done))
			r.Patch("/undone", tc.MoveTask(entity.Undone))
		})
	})
}
