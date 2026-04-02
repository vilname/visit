// Package model модели
package model

import "time"

// User структура пользователя
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Phone     string    `json:"phone"`
	BirthDate string    `json:"birthDate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// UserRegisterRequest запрос на регистрацию
type UserRegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

// UserLoginRequest запрос на авторизацию
type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UserUpdateRequest запрос на обновление профиля
type UserUpdateRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	BirthDate string `json:"birthDate"`
}

// UserProfileResponse ответ с профилем пользователя
type UserProfileResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	BirthDate string `json:"birthDate"`
}

// TokenResponse ответ с токеном
type TokenResponse struct {
	Token string `json:"token"`
}
