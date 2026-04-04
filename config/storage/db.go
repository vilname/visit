// Package storage подключение к базе данных
package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

// InitDB Инициализируем подключение к базе данных
func InitDB(dbUrl string) error {

	var err error

	db, err = pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return err
	}
	return nil
}

// GetDB отдает пул подключений с базой
func GetDB() *pgxpool.Pool {
	return db
}
