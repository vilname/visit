package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"visit/config"
	"visit/config/storage"
	"visit/src/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DoctorAPITestSuite struct {
	suite.Suite
	db     *pgxpool.Pool
	router *gin.Engine
}

func (s *DoctorAPITestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	s.db = storage.SetupTestDB(s.T())
	s.router = config.InitRoute()
}

func (s *DoctorAPITestSuite) SetupTest() {
	storage.CleanupTestDB(s.T(), s.db)
	storage.InsertTestDoctors(s.T(), s.db)
}

// TestGetDoctorsAPI_ShouldReturnAllDoctors тест GET /api/doctors
func (s *DoctorAPITestSuite) TestGetDoctorsAPI_ShouldReturnAllDoctors() {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/api/doctors", nil)
	w := httptest.NewRecorder()

	// Act
	s.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(s.T(), http.StatusOK, w.Code)

	var response model.DoctorListResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), 4, response.Total)
	assert.Len(s.T(), response.Doctors, 4)
	assert.Equal(s.T(), "Dr. John Smith", response.Doctors[0].Name)
}

// TestGetDoctorsAPI_WithSpecializationFilter тест фильтрации по специализации
func (s *DoctorAPITestSuite) TestGetDoctorsAPI_WithSpecializationFilter() {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/api/doctors?specialization=Cardiologist", nil)
	w := httptest.NewRecorder()

	// Act
	s.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(s.T(), http.StatusOK, w.Code)

	var response model.DoctorListResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), 2, response.Total)
	for _, doc := range response.Doctors {
		assert.Equal(s.T(), "Cardiologist", doc.Specialization)
	}
}

// TestGetDoctorByIDAPI_ShouldReturnDoctor тест GET /api/doctors/:id
func (s *DoctorAPITestSuite) TestGetDoctorByIDAPI_ShouldReturnDoctor() {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/api/doctors/doc-001", nil)
	w := httptest.NewRecorder()

	// Act
	s.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(s.T(), http.StatusOK, w.Code)

	var doctor model.DoctorResponse
	err := json.Unmarshal(w.Body.Bytes(), &doctor)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), "doc-001", doctor.ID)
	assert.Equal(s.T(), "Dr. John Smith", doctor.Name)
	assert.Equal(s.T(), "Cardiologist", doctor.Specialization)
}

// TestGetDoctorByIDAPI_NotFound тест GET с несуществующим ID
func (s *DoctorAPITestSuite) TestGetDoctorByIDAPI_NotFound() {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/api/doctors/non-existent", nil)
	w := httptest.NewRecorder()

	// Act
	s.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(s.T(), http.StatusNotFound, w.Code)
}

// TestGetDoctorByIDAPI_InvalidID тест GET с пустым ID
func (s *DoctorAPITestSuite) TestGetDoctorByIDAPI_InvalidID() {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/api/doctors/", nil)
	w := httptest.NewRecorder()

	// Act
	s.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

func TestDoctorAPISuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping API integration tests in short mode")
	}
	suite.Run(t, new(DoctorAPITestSuite))
}
