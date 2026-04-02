// Package repository репозитории для работы с базой данных
package repository

import (
	"context"
	"fmt"
	"time"
	"visit/config/storage"
	"visit/src/model"
)

// CreateAppointment создание записи на приём
func CreateAppointment(appointment model.Appointment) (model.Appointment, error) {
	db := storage.GetDB()

	err := db.QueryRow(
		context.Background(),
		`INSERT INTO appointments (id, user_id, doctor_id, date_time, status, comment, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, user_id, doctor_id, date_time, status, COALESCE(comment, ''), created_at, updated_at`,
		appointment.ID, appointment.UserID, appointment.DoctorID, appointment.DateTime,
		appointment.Status, appointment.Comment, appointment.CreatedAt, appointment.UpdatedAt,
	).Scan(&appointment.ID, &appointment.UserID, &appointment.DoctorID, &appointment.DateTime,
		&appointment.Status, &appointment.Comment, &appointment.CreatedAt, &appointment.UpdatedAt)

	if err != nil {
		return appointment, fmt.Errorf("ошибка при создании записи: %w", err)
	}

	return appointment, nil
}

// GetAppointmentsByUserID получение записей пользователя
func GetAppointmentsByUserID(userID string) ([]model.Appointment, error) {
	db := storage.GetDB()

	rows, err := db.Query(
		context.Background(),
		`SELECT id, user_id, doctor_id, date_time, status, COALESCE(comment, ''), created_at, updated_at
		FROM appointments WHERE user_id = $1 ORDER BY date_time DESC`,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении записей: %w", err)
	}
	defer rows.Close()

	appointments := make([]model.Appointment, 0)
	for rows.Next() {
		var a model.Appointment
		err = rows.Scan(&a.ID, &a.UserID, &a.DoctorID, &a.DateTime, &a.Status, &a.Comment, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("ошибка при чтении записи: %w", err)
		}
		appointments = append(appointments, a)
	}

	return appointments, nil
}

// GetAppointmentByID получение записи по ID
func GetAppointmentByID(appointmentID string) (model.Appointment, error) {
	db := storage.GetDB()
	var appointment model.Appointment

	err := db.QueryRow(
		context.Background(),
		`SELECT id, user_id, doctor_id, date_time, status, COALESCE(comment, ''), created_at, updated_at
		FROM appointments WHERE id = $1`,
		appointmentID,
	).Scan(&appointment.ID, &appointment.UserID, &appointment.DoctorID, &appointment.DateTime,
		&appointment.Status, &appointment.Comment, &appointment.CreatedAt, &appointment.UpdatedAt)

	if err != nil {
		return appointment, fmt.Errorf("запись не найдена: %w", err)
	}

	return appointment, nil
}

// CancelAppointment отмена записи на приём
func CancelAppointment(appointmentID string, userID string) error {
	db := storage.GetDB()
	now := time.Now()

	result, err := db.Exec(
		context.Background(),
		`UPDATE appointments SET status = $1, updated_at = $2 WHERE id = $3 AND user_id = $4 AND status != $5`,
		model.AppointmentStatusCanceled, now, appointmentID, userID, model.AppointmentStatusCanceled,
	)
	if err != nil {
		return fmt.Errorf("ошибка при отмене записи: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("запись не найдена или уже отменена")
	}

	return nil
}
