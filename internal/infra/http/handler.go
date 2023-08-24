package http

import (
	"fmt"
	"net/http"
)

func AppToHttp(w http.ResponseWriter, err error) {
	fmt.Println(err.Error())
}
