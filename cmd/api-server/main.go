package main

import (
	"fmt"
	"net/http"

	"github.com/sentry-lab/fabdb/internal/api"
	"github.com/sentry-lab/fabdb/internal/db"
)

func main() {
	dbPool := db.InitDB()
	defer dbPool.Close()

	app := api.App{DB: dbPool}

	server := http.Server{
		Addr:    ":8080",
		Handler: api.InitRouter(&app),
	}

	err := server.ListenAndServe()
	fmt.Println("err:", err)
}
