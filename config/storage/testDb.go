package storage

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
	"unicode"
	"visit/dbmigrate"
	"visit/src/model"
	"visit/test/fixtures"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ensureTestDatabase(ctx context.Context, dbURL string) error {
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL_TEST is empty")
	}
	u, err := url.Parse(dbURL)
	if err != nil {
		return err
	}
	dbName := strings.TrimPrefix(u.Path, "/")
	if dbName == "" {
		return fmt.Errorf("database name missing in DATABASE_URL_TEST")
	}
	for _, r := range dbName {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return fmt.Errorf("invalid database name in URL")
		}
	}
	u.Path = "/postgres"
	adminURL := u.String()

	conn, err := pgx.Connect(ctx, adminURL)
	if err != nil {
		return fmt.Errorf("connect to postgres maintenance db: %w", err)
	}
	defer conn.Close(ctx)

	var exists bool
	err = conn.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)`, dbName).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	_, err = conn.Exec(ctx, fmt.Sprintf("CREATE DATABASE %q", dbName))
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			// Параллельные пакеты go test могут создать БД одновременно.
			return nil
		}
		return err
	}
	return nil
}

// SetupTestDB - настройка тестовой базы данных
func SetupTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()

	// Ждем готовности БД с ретраями
	var db *pgxpool.Pool
	var err error

	//err = godotenv.Load(".env.test")
	//if err != nil {
	//	t.Fatalf("Failed load enf file: %v", err)
	//}

	if err := godotenv.Load("../../.env.test"); err != nil {
		t.Fatalf("Failed load enf file: %v", err)
	}

	dbUrl := os.Getenv("DATABASE_URL_TEST")

	fmt.Println("DATABASE_URL_TEST: " + dbUrl)

	ctxEnsure, cancelEnsure := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelEnsure()
	if err := ensureTestDatabase(ctxEnsure, dbUrl); err != nil {
		t.Fatalf("ensure test database: %v", err)
	}
	if err := dbmigrate.Up(); err != nil {
		t.Fatalf("test migrations: %v", err)
	}

	for i := 0; i < 10; i++ {
		err = InitDB(dbUrl)
		if err == nil {
			db = GetDB()
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Проверяем подключение
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.Ping(ctx); err != nil {
		t.Fatalf("Database ping failed: %v", err)
	}

	return db
}

// CleanupTestDB - очистка тестовой базы данных
func CleanupTestDB(t *testing.T, db *pgxpool.Pool) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Очищаем все таблицы
	tables := []string{"doctors", "appointments", "patients"} // добавьте все таблицы
	for _, table := range tables {
		_, err := db.Exec(ctx, fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
		if err != nil {
			t.Logf("Warning: failed to truncate table %s: %v", table, err)
		}
	}
}

// InsertTestDoctors - вставка тестовых данных из фикстур
func InsertTestDoctors(t *testing.T, db *pgxpool.Pool) {
	t.Helper()

	ctx := context.Background()
	for _, doctor := range fixtures.TestDoctorInsertData {
		_, err := db.Exec(ctx,
			`INSERT INTO doctors (id, name, specialization, experience, description) 
             VALUES ($1, $2, $3, $4, $5)`,
			doctor.ID, doctor.Name, doctor.Specialization, doctor.Experience, doctor.Description)
		if err != nil {
			t.Fatalf("Failed to insert test doctor: %v", err)
		}
	}
}

// CreateTestDoctor - создание одного тестового доктора
func CreateTestDoctor(t *testing.T, db *pgxpool.Pool, doctor model.DoctorResponse) {
	t.Helper()

	ctx := context.Background()
	_, err := db.Exec(ctx,
		`INSERT INTO doctors (id, name, specialization, experience, description) 
         VALUES ($1, $2, $3, $4, $5)`,
		doctor.ID, doctor.Name, doctor.Specialization, doctor.Experience, doctor.Description)
	if err != nil {
		t.Fatalf("Failed to insert test doctor: %v", err)
	}
}
