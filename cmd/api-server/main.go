package main

import (
	"fmt"
	"net/http"

	"github.com/sentry-lab/fabdb/internal/api"
)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: api.InitRouter(),
	}

	err := server.ListenAndServe()
	fmt.Println("err:", err)
}
