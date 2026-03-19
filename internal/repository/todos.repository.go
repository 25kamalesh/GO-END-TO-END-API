package repository

import (
	"context"
	"github.com/25Kamalesh/go_todo_api/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTodo(ctx context.Context, pool *pgxpool.Pool, title string, completed bool) (*models.Todos, error) {
	var todo models.Todos
	query :=  "INSERT INTO TODOS (title, completed) VALUES ($1, $2) RETURNING id, title, completed, created_at, updated_at"
	err := pool.QueryRow(ctx , query , title ,completed).Scan(
		&todo.ID, 
		&todo.Title, 
		&todo.Completed, 
		&todo.CreatedAt, 
		&todo.UpdatedAt,
	)

	if err != nil {
		return nil , err
	}
	return &todo , nil

}
