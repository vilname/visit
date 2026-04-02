// Package service сервисный слой
package service

import (
	"visit/src/model"
	"visit/src/repository"
)

// GetAllDoctors получение списка всех докторов
func GetAllDoctors() (model.DoctorListResponse, error) {
	doctors, err := repository.GetAllDoctors()
	if err != nil {
		return model.DoctorListResponse{}, err
	}

	return model.DoctorListResponse{
		Doctors: doctors,
		Total:   len(doctors),
	}, nil
}

// GetDoctorByID получение доктора по ID
func GetDoctorByID(doctorID string) (model.DoctorResponse, error) {
	return repository.GetDoctorByID(doctorID)
}

// GetDoctorsBySpecialization получение докторов по специализации
func GetDoctorsBySpecialization(specialization string) (model.DoctorListResponse, error) {
	doctors, err := repository.GetDoctorsBySpecialization(specialization)
	if err != nil {
		return model.DoctorListResponse{}, err
	}

	return model.DoctorListResponse{
		Doctors: doctors,
		Total:   len(doctors),
	}, nil
}
