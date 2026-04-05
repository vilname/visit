package repository

import (
	"context"
	"testing"
	"time"
	"visit/config/storage"
	"visit/src/repository"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DoctorRepositoryTestSuite struct {
	suite.Suite
	db     *pgxpool.Pool
	ctx    context.Context
	cancel context.CancelFunc
}

func (s *DoctorRepositoryTestSuite) SetupSuite() {
	s.ctx, s.cancel = context.WithTimeout(context.Background(), 30*time.Second)
	s.db = storage.SetupTestDB(s.T())
}

func (s *DoctorRepositoryTestSuite) TearDownSuite() {
	storage.CleanupTestDB(s.T(), s.db)
	s.cancel()
}

func (s *DoctorRepositoryTestSuite) SetupTest() {
	storage.CleanupTestDB(s.T(), s.db)
	storage.InsertTestDoctors(s.T(), s.db)
}

// TestGetAllDoerts_ShouldReturnAllDoctors проверяет получение всех докторов
func (s *DoctorRepositoryTestSuite) TestGetAllDoctors_ShouldReturnAllDoctors() {
	// Act
	doctors, err := repository.GetAllDoctors()

	// Assert
	require.NoError(s.T(), err)
	assert.Len(s.T(), doctors, 4)
	assert.Equal(s.T(), "Dr. John Smith", doctors[0].Name)
}

// TestGetDoctorByID_ShouldReturnDoctor проверяет получение доктора по ID
func (s *DoctorRepositoryTestSuite) TestGetDoctorByID_ShouldReturnDoctor() {
	// Act
	doctor, err := repository.GetDoctorByID("doc-001")

	// Assert
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "doc-001", doctor.ID)
	assert.Equal(s.T(), "Dr. John Smith", doctor.Name)
	assert.Equal(s.T(), "Cardiologist", doctor.Specialization)
	assert.Equal(s.T(), 10, doctor.Experience)
}

// TestGetDoctorByID_NotFound проверяет ошибку при несуществующем ID
func (s *DoctorRepositoryTestSuite) TestGetDoctorByID_NotFound() {
	// Act
	doctor, err := repository.GetDoctorByID("non-existent-id")

	// Assert
	assert.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "доктор не найден")
	assert.Empty(s.T(), doctor.ID)
}

// TestGetDoctorsBySpecialization_ShouldFilter правильно фильтрует по специализации
func (s *DoctorRepositoryTestSuite) TestGetDoctorsBySpecialization_ShouldFilter() {
	// Act
	cardiologists, err := repository.GetDoctorsBySpecialization("Cardiologist")

	// Assert
	require.NoError(s.T(), err)
	assert.Len(s.T(), cardiologists, 2) // John Smith и Bob Wilson

	for _, doc := range cardiologists {
		assert.Equal(s.T(), "Cardiologist", doc.Specialization)
	}
}

// TestGetDoctorsBySpecialization_EmptyResult возвращает пустой список для несуществующей специализации
func (s *DoctorRepositoryTestSuite) TestGetDoctorsBySpecialization_EmptyResult() {
	// Act
	doctors, err := repository.GetDoctorsBySpecialization("Dermatologist")

	// Assert
	require.NoError(s.T(), err)
	assert.Empty(s.T(), doctors)
}

func TestDoctorRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}
	suite.Run(t, new(DoctorRepositoryTestSuite))
}
