package api

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sentry-lab/fabdb/internal/db"
)

type App struct {
	DB *pgxpool.Pool
}

func InitRouter(app *App) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/workflows", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprint(db.GetWorkflows(app.DB))))
	})

	return router
}
