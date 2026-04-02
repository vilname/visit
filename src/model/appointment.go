// Package model модели
package model

import "time"

// Appointment структура записи на приём
type Appointment struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	DoctorID  string    `json:"doctorId"`
	DateTime  time.Time `json:"dateTime"`
	Status    string    `json:"status"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// AppointmentCreateRequest запрос на создание записи на приём
type AppointmentCreateRequest struct {
	DoctorID string `json:"doctorId" binding:"required"`
	DateTime string `json:"dateTime" binding:"required"`
	Comment  string `json:"comment"`
}

// AppointmentResponse ответ с информацией о записи на приём
type AppointmentResponse struct {
	ID        string         `json:"id"`
	Doctor    DoctorResponse `json:"doctor"`
	DateTime  time.Time      `json:"dateTime"`
	Status    string         `json:"status"`
	Comment   string         `json:"comment"`
	CreatedAt time.Time      `json:"createdAt"`
}

// AppointmentStatus статусы записи на приём
const (
	AppointmentStatusCreated   = "created"
	AppointmentStatusConfirmed = "confirmed"
	AppointmentStatusCanceled  = "canceled"
	AppointmentStatusCompleted = "completed"
)
