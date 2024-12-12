package api

import (
	"net/http"
)

func InitRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) { http.NotFound(w, r) })

	return router
}
