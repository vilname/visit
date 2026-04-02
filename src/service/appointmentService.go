// Package service сервисный слой
package service

import (
	"fmt"
	"time"
	"visit/src/model"
	"visit/src/repository"

	"github.com/google/uuid"
)

// CreateAppointment создание записи на приём
func CreateAppointment(userID string, req model.AppointmentCreateRequest) (model.AppointmentResponse, error) {
	dateTime, err := time.Parse(time.RFC3339, req.DateTime)
	if err != nil {
		return model.AppointmentResponse{}, fmt.Errorf("неверный формат даты, используйте RFC3339: %w", err)
	}

	if dateTime.Before(time.Now()) {
		return model.AppointmentResponse{}, fmt.Errorf("нельзя записаться на прошедшую дату")
	}

	doctor, err := repository.GetDoctorByID(req.DoctorID)
	if err != nil {
		return model.AppointmentResponse{}, fmt.Errorf("доктор не найден")
	}

	now := time.Now()
	appointment := model.Appointment{
		ID:        uuid.New().String(),
		UserID:    userID,
		DoctorID:  req.DoctorID,
		DateTime:  dateTime,
		Status:    model.AppointmentStatusCreated,
		Comment:   req.Comment,
		CreatedAt: now,
		UpdatedAt: now,
	}

	createdAppointment, err := repository.CreateAppointment(appointment)
	if err != nil {
		return model.AppointmentResponse{}, err
	}

	return model.AppointmentResponse{
		ID:        createdAppointment.ID,
		Doctor:    doctor,
		DateTime:  createdAppointment.DateTime,
		Status:    createdAppointment.Status,
		Comment:   createdAppointment.Comment,
		CreatedAt: createdAppointment.CreatedAt,
	}, nil
}

// GetUserAppointments получение записей пользователя
func GetUserAppointments(userID string) ([]model.AppointmentResponse, error) {
	appointments, err := repository.GetAppointmentsByUserID(userID)
	if err != nil {
		return nil, err
	}

	responses := make([]model.AppointmentResponse, 0, len(appointments))
	for _, a := range appointments {
		doctor, err := repository.GetDoctorByID(a.DoctorID)
		if err != nil {
			doctor = model.DoctorResponse{ID: a.DoctorID, Name: "Неизвестный доктор"}
		}

		responses = append(responses, model.AppointmentResponse{
			ID:        a.ID,
			Doctor:    doctor,
			DateTime:  a.DateTime,
			Status:    a.Status,
			Comment:   a.Comment,
			CreatedAt: a.CreatedAt,
		})
	}

	return responses, nil
}

// CancelAppointment отмена записи на приём
func CancelAppointment(appointmentID string, userID string) error {
	return repository.CancelAppointment(appointmentID, userID)
}
