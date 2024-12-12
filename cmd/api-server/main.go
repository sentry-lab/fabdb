package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sentry-lab/fabdb/internal/api"
	"github.com/sentry-lab/fabdb/internal/db"
)

func main() {
	dbPool := db.InitDB()
	defer dbPool.Close()

	app := api.App{DB: dbPool}

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}
	server := http.Server{
		Addr:    ":" + port,
		Handler: api.InitRouter(&app),
	}

	log.Println("Listening on", server.Addr)
	err := server.ListenAndServe()
	fmt.Println("err:", err)
}
