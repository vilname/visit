package storage

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
	"visit/src/model"
	"visit/test/fixtures"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

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
