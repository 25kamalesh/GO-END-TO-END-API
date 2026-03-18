package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)


func ConnectPostgres(databaseURI string) (*pgxpool.Pool,error) {
	ctx := context.Background()
	config , err := pgxpool.ParseConfig(databaseURI)
	if err != nil {
		log.Printf("Unable to parse database URI: %v\n", err)
		return nil, err
	}
	pool , err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Printf("Unable to create connection pool: %v\n", err)
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		log.Printf("Unable to ping to database : %v" , err)
		pool.Close()
		return nil, err
		
	}
	return pool, nil
}