package main

import (
	"fmt"

	"github.com/willianbl99/todo-list_api/pkg/infra/http"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/router"
)

func main() {
	fmt.Println("Starting server...")
	r := router.MainRouter()
	
	http.StartServer(r)
}
