package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetWorkflows(db *pgxpool.Pool) []Workflow {
	rows, _ := db.Query(context.Background(), "SELECT id, name FROM workflows")
	var id string
	var name string
	result := make([]Workflow, 0)

	_, err := pgx.ForEachRow(rows, []any{&id, &name}, func() error {
		result = append(result, Workflow{Id: id, Name: name})
		return nil
	})

	if err != nil {
		fmt.Println("get workflows err:", err)
	}

	return result
}
