// Package model модели
package model

// DoctorExtended расширенная структура доктора
type DoctorExtended struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Specialization string `json:"specialization"`
	Experience     int    `json:"experience"`
	Description    string `json:"description"`
}

// DoctorResponse ответ с информацией о докторе
type DoctorResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Specialization string `json:"specialization"`
	Experience     int    `json:"experience"`
	Description    string `json:"description"`
}

// DoctorListResponse ответ со списком докторов
type DoctorListResponse struct {
	Doctors []DoctorResponse `json:"doctors"`
	Total   int              `json:"total"`
}
