package main

import (
	"github.com/sentry-lab/fabdb/internal/api"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: api.InitRouter(),
	}

	server.ListenAndServe()
}
