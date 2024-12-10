package api

import (
	"fmt"
	"net/http"
)

func InitRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) { fmt.Println("test") })

	return router
}
