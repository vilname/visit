package fixtures

import "visit/src/model"

// TestDoctors - тестовые данные докторов
var TestDoctors = []model.DoctorResponse{
	{
		ID:             "doc-001",
		Name:           "Dr. John Smith",
		Specialization: "Cardiologist",
		Experience:     10,
		Description:    "Expert in heart diseases",
	},
	{
		ID:             "doc-002",
		Name:           "Dr. Jane Doe",
		Specialization: "Neurologist",
		Experience:     8,
		Description:    "Brain specialist",
	},
	{
		ID:             "doc-003",
		Name:           "Dr. Bob Wilson",
		Specialization: "Cardiologist",
		Experience:     5,
		Description:    "Heart surgeon",
	},
	{
		ID:             "doc-004",
		Name:           "Dr. Alice Brown",
		Specialization: "Pediatrician",
		Experience:     12,
		Description:    "Child specialist",
	},
}

// TestDoctorInsertData - данные для вставки в БД (без ID, если он генерируется)
var TestDoctorInsertData = []struct {
	ID             string
	Name           string
	Specialization string
	Experience     int
	Description    string
}{
	{"doc-001", "Dr. John Smith", "Cardiologist", 10, "Expert in heart diseases"},
	{"doc-002", "Dr. Jane Doe", "Neurologist", 8, "Brain specialist"},
	{"doc-003", "Dr. Bob Wilson", "Cardiologist", 5, "Heart surgeon"},
	{"doc-004", "Dr. Alice Brown", "Pediatrician", 12, "Child specialist"},
}

// GetTestDoctorByID - получение тестового доктора по ID
func GetTestDoctorByID(id string) *model.DoctorResponse {
	for _, doctor := range TestDoctors {
		if doctor.ID == id {
			return &doctor
		}
	}
	return nil
}

// GetTestDoctorsBySpecialization - получение тестовых докторов по специализации
func GetTestDoctorsBySpecialization(specialization string) []model.DoctorResponse {
	var result []model.DoctorResponse
	for _, doctor := range TestDoctors {
		if doctor.Specialization == specialization {
			result = append(result, doctor)
		}
	}
	return result
}

// GenerateTestDoctor создаёт одного тестового доктора для изолированных тестов
func GenerateTestDoctor(id, name, spec string, exp int, desc string) model.DoctorResponse {
	return model.DoctorResponse{
		ID:             id,
		Name:           name,
		Specialization: spec,
		Experience:     exp,
		Description:    desc,
	}
}

// GenerateTestDoctorsList генерирует N тестовых докторов
func GenerateTestDoctorsList(n int) []model.DoctorResponse {
	doctors := make([]model.DoctorResponse, n)
	for i := 0; i < n; i++ {
		doctors[i] = GenerateTestDoctor(
			string(rune('A'+i)),
			"Test Doctor",
			"General",
			5,
			"Test description",
		)
	}
	return doctors
}
