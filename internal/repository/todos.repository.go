package repository

import (
	"context"

	"github.com/25Kamalesh/go_todo_api/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTodo(ctx context.Context, pool *pgxpool.Pool, userID int, title string, completed bool) (*models.Todos, error) {
	var todo models.Todos
	query := "INSERT INTO todos (user_id, title, completed) VALUES ($1, $2, $3) RETURNING id, user_id, title, completed, created_at, updated_at"
	err := pool.QueryRow(ctx, query, userID, title, completed).Scan(
		&todo.ID,
		&todo.UserID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &todo, nil

}
