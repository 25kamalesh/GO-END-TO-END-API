package repository

import (
	"context"

	"github.com/25Kamalesh/go_todo_api/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

// CreateUser inserts a new user into the database
func CreateUser(
	ctx context.Context,
	pool *pgxpool.Pool,
	name string,
	email string,
	password string,
) (*models.User, error) {

	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, name, email, password, created_at
	`

	var user models.User

	err := pool.QueryRow(ctx, query, name, email, password).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByEmail fetches a user using email (used for login)
func GetUserByEmail(
	ctx context.Context,
	pool *pgxpool.Pool,
	email string,
) (*models.User, error) {

	query := `
		SELECT id, name, email, password, created_at
		FROM users
		WHERE email = $1
	`

	var user models.User

	err := pool.QueryRow(ctx, query, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
