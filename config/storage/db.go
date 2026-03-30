// Package storage подключение к базе данных
package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

// InitDB Инициализируем подключение к базе данных
func InitDB() {

	var err error

	db, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}

// GetDB отдает пул подключений с базой
func GetDB() *pgxpool.Pool {
	return db
}
