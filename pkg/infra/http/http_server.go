package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func StartServer(r chi.Router) {	
	fmt.Println("Server started")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}