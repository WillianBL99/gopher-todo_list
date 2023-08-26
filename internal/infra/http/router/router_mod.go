package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"github.com/willianbl99/todo-list_api/config"
	"github.com/willianbl99/todo-list_api/internal/infra/db"
	"github.com/willianbl99/todo-list_api/internal/infra/http/middleware"
)

type RouterMod struct {
	DBModule *db.DbMod
}

func NewRouterMod(dbmod *db.DbMod) *RouterMod {
	return &RouterMod{
		DBModule: dbmod,
	}
}

func (rm *RouterMod) Start(dbmod *db.DbMod) chi.Router {
	ur := dbmod.UserRepository
	tr := dbmod.TaskRepository
	md := middleware.NewMiddleware(ur, tr)
	filesDir := config.NewAppConf().API.WorkDir
	handlerFiles := http.FileServer(http.Dir(filesDir))

	r := chi.NewRouter()
	r.Handle("/docs/*", http.StripPrefix("/docs/", handlerFiles))

	// redirect / to /swagger/*
	r.Get("/*", swagger())
	r.Get("/health", healthCheck)
	r.Get("/swagger/*", swagger())
	TaskRouter(r, tr, md)
	UserRouter(r, ur)

	return r
}

func swagger() http.HandlerFunc {
	url := config.NewAppConf().API.URL
	return httpSwagger.Handler(
		httpSwagger.URL(url + "/docs/swagger.yml"),
	)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"status": "ok"}`)
}
