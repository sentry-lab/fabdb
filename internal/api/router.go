package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sentry-lab/fabdb/internal/db"
)

type App struct {
	DB *pgxpool.Pool
}

func InitRouter(app *App) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /workflows", func(w http.ResponseWriter, r *http.Request) {
		SetHeaders(&w)

		data := db.GetWorkflows(app.DB)
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println("err marshaling workflows", err)
		}
		w.Write(jsonData)
	})

	return router
}
