package http

import (
	"fmt"
	"net/http"

	"github.com/willianbl99/todo-list_api/config"
	"github.com/willianbl99/todo-list_api/internal/infra/db"
	"github.com/willianbl99/todo-list_api/internal/infra/http/router"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type HttpMod struct {
	DBModule  *db.DbMod
	RouterMod *router.RouterMod
}

func NewHttpMod(dbmod *db.DbMod) *HttpMod {
	return &HttpMod{
		DBModule:  dbmod,
		RouterMod: router.NewRouterMod(dbmod),
	}
}

func (hm *HttpMod) Start() {
	apicf := config.NewAppConf().API
	port := fmt.Sprint(apicf.Port)

	rt := hm.RouterMod.Start(hm.DBModule)

	fmt.Println("Server running on port: " + port)
	err := http.ListenAndServe(":"+port, rt)
	e.New().InfraHttpErr(e.ErrorOnStartServer, err.Error()).Fatal()
}
