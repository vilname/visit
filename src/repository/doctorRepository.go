// Package repository репозитории для работы с базой данных
package repository

import (
	"context"
	"fmt"
	"visit/config/storage"
	"visit/src/model"
)

// GetAllDoctors получение списка всех докторов
func GetAllDoctors() ([]model.DoctorResponse, error) {
	db := storage.GetDB()

	rows, err := db.Query(
		context.Background(),
		`SELECT id, name, specialization, experience, COALESCE(description, '')
		FROM doctors ORDER BY name`,
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении списка докторов: %w", err)
	}
	defer rows.Close()

	doctors := make([]model.DoctorResponse, 0)
	for rows.Next() {
		var doctor model.DoctorResponse
		err = rows.Scan(&doctor.ID, &doctor.Name, &doctor.Specialization, &doctor.Experience, &doctor.Description)
		if err != nil {
			return nil, fmt.Errorf("ошибка при чтении доктора: %w", err)
		}
		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

// GetDoctorByID получение доктора по ID
func GetDoctorByID(doctorID string) (model.DoctorResponse, error) {
	db := storage.GetDB()
	var doctor model.DoctorResponse

	err := db.QueryRow(
		context.Background(),
		`SELECT id, name, specialization, experience, COALESCE(description, '')
		FROM doctors WHERE id = $1`,
		doctorID,
	).Scan(&doctor.ID, &doctor.Name, &doctor.Specialization, &doctor.Experience, &doctor.Description)

	if err != nil {
		return doctor, fmt.Errorf("доктор не найден: %w", err)
	}

	return doctor, nil
}

// GetDoctorsBySpecialization получение докторов по специализации
func GetDoctorsBySpecialization(specialization string) ([]model.DoctorResponse, error) {
	db := storage.GetDB()

	rows, err := db.Query(
		context.Background(),
		`SELECT id, name, specialization, experience, COALESCE(description, '')
		FROM doctors WHERE specialization = $1 ORDER BY name`,
		specialization,
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении докторов: %w", err)
	}
	defer rows.Close()

	doctors := make([]model.DoctorResponse, 0)
	for rows.Next() {
		var doctor model.DoctorResponse
		err = rows.Scan(&doctor.ID, &doctor.Name, &doctor.Specialization, &doctor.Experience, &doctor.Description)
		if err != nil {
			return nil, fmt.Errorf("ошибка при чтении доктора: %w", err)
		}
		doctors = append(doctors, doctor)
	}

	return doctors, nil
}
