// Package service сервисный слой
package service

import (
	"fmt"
	"time"
	"visit/src/model"
	"visit/src/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser регистрация нового пользователя
func RegisterUser(req model.UserRegisterRequest) (model.User, error) {
	_, err := repository.GetUserByEmail(req.Email)
	if err == nil {
		return model.User{}, fmt.Errorf("пользователь с таким email уже существует")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, fmt.Errorf("ошибка при хешировании пароля: %w", err)
	}

	now := time.Now()
	user := model.User{
		ID:        uuid.New().String(),
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		CreatedAt: now,
		UpdatedAt: now,
	}

	createdUser, err := repository.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}

	return createdUser, nil
}

// LoginUser авторизация пользователя
func LoginUser(req model.UserLoginRequest) (model.User, error) {
	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		return model.User{}, fmt.Errorf("неверный email или пароль")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return model.User{}, fmt.Errorf("неверный email или пароль")
	}

	return user, nil
}

// GetUserProfile получение профиля пользователя
func GetUserProfile(userID string) (model.UserProfileResponse, error) {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return model.UserProfileResponse{}, err
	}

	return model.UserProfileResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		BirthDate: user.BirthDate,
	}, nil
}

// UpdateUserProfile обновление профиля пользователя
func UpdateUserProfile(userID string, req model.UserUpdateRequest) (model.UserProfileResponse, error) {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return model.UserProfileResponse{}, err
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.BirthDate != "" {
		user.BirthDate = req.BirthDate
	}

	updatedUser, err := repository.UpdateUser(user)
	if err != nil {
		return model.UserProfileResponse{}, err
	}

	return model.UserProfileResponse{
		ID:        updatedUser.ID,
		Email:     updatedUser.Email,
		FirstName: updatedUser.FirstName,
		LastName:  updatedUser.LastName,
		Phone:     updatedUser.Phone,
		BirthDate: updatedUser.BirthDate,
	}, nil
}
