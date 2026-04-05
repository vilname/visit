// Package dbmigrate применяет SQL-миграции к БД из переменных окружения DB_*.
package dbmigrate

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func migrationSourceURL() (string, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("runtime.Caller failed")
	}
	dir := filepath.Clean(filepath.Join(filepath.Dir(file), "..", "migration"))
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	path := filepath.ToSlash(absDir)
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return (&url.URL{Scheme: "file", Path: path}).String(), nil
}

// Up поднимает схему БД (DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME).
func Up() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbSQL, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbPass, dbName, dbPort),
	)
	if err != nil {
		return err
	}
	defer dbSQL.Close()

	driver, err := postgres.WithInstance(dbSQL, &postgres.Config{})
	if err != nil {
		return err
	}

	src, err := migrationSourceURL()
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		src,
		"postgres", driver)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration up: %w", err)
	}

	return nil
}
