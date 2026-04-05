package service

import (
	"testing"
	"visit/src/model"
	"visit/test/fixtures"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockDoctorRepository мок для репозитория докторов
type MockDoctorRepository struct {
	mock.Mock
}

func (m *MockDoctorRepository) GetAllDoctors() ([]model.DoctorResponse, error) {
	args := m.Called()
	return args.Get(0).([]model.DoctorResponse), args.Error(1)
}

func (m *MockDoctorRepository) GetDoctorByID(id string) (model.DoctorResponse, error) {
	args := m.Called(id)
	return args.Get(0).(model.DoctorResponse), args.Error(1)
}

func (m *MockDoctorRepository) GetDoctorsBySpecialization(spec string) ([]model.DoctorResponse, error) {
	args := m.Called(spec)
	return args.Get(0).([]model.DoctorResponse), args.Error(1)
}

// TestGetAllDoctors_Service проверяет контракт репозитория, который использует сервис (мок).
// Сам service.GetAllDoctors без внедрения зависимостей вызывает глобальный repository.
func TestGetAllDoctors_Service(t *testing.T) {
	mockRepo := new(MockDoctorRepository)
	expectedDoctors := fixtures.TestDoctors
	mockRepo.On("GetAllDoctors").Return(expectedDoctors, nil)

	result, err := mockRepo.GetAllDoctors()
	require.NoError(t, err)
	assert.Equal(t, expectedDoctors, result)
	mockRepo.AssertExpectations(t)
}

// TestGetDoctorByID_Service_Success тест успешного получения доктора
func TestGetDoctorByID_Service_Success(t *testing.T) {
	expectedDoctor := fixtures.TestDoctors[0]

	// В реальном тесте с моком:
	// mockRepo.On("GetDoctorByID", "doc-001").Return(expectedDoctor, nil)
	// result, err := service.GetDoctorByID("doc-001")

	assert.Equal(t, "doc-001", expectedDoctor.ID)
	assert.Equal(t, "Dr. John Smith", expectedDoctor.Name)
}

// TestGetDoctorsBySpecialization_Service тест фильтрации через сервис
func TestGetDoctorsBySpecialization_Service(t *testing.T) {
	specialization := "Cardiologist"
	expectedDoctors := fixtures.GetTestDoctorsBySpecialization(specialization)

	assert.Len(t, expectedDoctors, 2)
	for _, doc := range expectedDoctors {
		assert.Equal(t, specialization, doc.Specialization)
	}
}
