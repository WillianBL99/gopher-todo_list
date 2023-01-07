package main

import (
	"fmt"

	"github.com/willianbl99/todo-list_api/config"
	"github.com/willianbl99/todo-list_api/pkg/infra/http"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/router"
)

func main() {
	r := router.MainRouter()

	conf := config.NewAppConf()

	fmt.Println("Server running on port: " + conf.API.Port)
	
	http.StartServer(r)
}
