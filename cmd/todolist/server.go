package main

import (
	"github.com/willianbl99/todo-list_api/pkg/infra/http"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/router"
)

func main() {
	r := router.MainRouter()
	
	http.StartServer(r)
}
