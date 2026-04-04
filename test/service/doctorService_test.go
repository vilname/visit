package service

import (
	"context"
	"testing"
	"time"
	"visit/config/storage"
	"visit/src/service"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// DoctorServiceTestSuite - набор тестов для сервиса докторов
type DoctorServiceTestSuite struct {
	suite.Suite
	db     *pgxpool.Pool
	ctx    context.Context
	cancel context.CancelFunc
}

// SetupSuite - запускается один раз перед всеми тестами
func (s *DoctorServiceTestSuite) SetupSuite() {
	s.ctx, s.cancel = context.WithTimeout(context.Background(), 30*time.Second)
	s.db = storage.SetupTestDB(s.T())
}

// TearDownSuite - запускается один раз после всех тестов
func (s *DoctorServiceTestSuite) TearDownSuite() {
	storage.CleanupTestDB(s.T(), s.db)
	s.cancel()
}

// SetupTest - запускается перед каждым тестом
func (s *DoctorServiceTestSuite) SetupTest() {
	// Очистка и вставка свежих тестовых данных перед каждым тестом
	storage.CleanupTestDB(s.T(), s.db)
	storage.InsertTestDoctors(s.T(), s.db)
}

// TestGetAllDoctors_ShouldReturnAllDoctors - тест получения всех докторов
func (s *DoctorServiceTestSuite) TestGetAllDoctors_ShouldReturnAllDoctors() {
	// Act
	result, err := service.GetAllDoctors()

	// Assert
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 4, result.Total)
	assert.Len(s.T(), result.Doctors, 4)
}

// ... остальные тесты из предыдущего примера ...

// Запуск всех тестов
func TestDoctorServiceTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	suite.Run(t, new(DoctorServiceTestSuite))
}
