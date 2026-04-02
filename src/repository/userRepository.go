// Package repository репозитории для работы с базой данных
package repository

import (
	"context"
	"fmt"
	"time"
	"visit/config/storage"
	"visit/src/model"
)

// CreateUser создание пользователя в базе данных
func CreateUser(user model.User) (model.User, error) {
	db := storage.GetDB()

	err := db.QueryRow(
		context.Background(),
		`INSERT INTO users (id, email, password, first_name, last_name, phone, birth_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, email, first_name, last_name, phone, birth_date, created_at, updated_at`,
		user.ID, user.Email, user.Password, user.FirstName, user.LastName, user.Phone, user.BirthDate, user.CreatedAt, user.UpdatedAt,
	).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Phone, &user.BirthDate, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return user, fmt.Errorf("ошибка при создании пользователя: %w", err)
	}

	return user, nil
}

// GetUserByEmail получение пользователя по email
func GetUserByEmail(email string) (model.User, error) {
	db := storage.GetDB()
	var user model.User

	err := db.QueryRow(
		context.Background(),
		`SELECT id, email, password, first_name, last_name, phone, COALESCE(birth_date, ''), created_at, updated_at
		FROM users WHERE email = $1`,
		email,
	).Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Phone, &user.BirthDate, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return user, fmt.Errorf("пользователь не найден: %w", err)
	}

	return user, nil
}

// GetUserByID получение пользователя по ID
func GetUserByID(userID string) (model.User, error) {
	db := storage.GetDB()
	var user model.User

	err := db.QueryRow(
		context.Background(),
		`SELECT id, email, password, first_name, last_name, phone, COALESCE(birth_date, ''), created_at, updated_at
		FROM users WHERE id = $1`,
		userID,
	).Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Phone, &user.BirthDate, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return user, fmt.Errorf("пользователь не найден: %w", err)
	}

	return user, nil
}

// UpdateUser обновление данных пользователя
func UpdateUser(user model.User) (model.User, error) {
	db := storage.GetDB()
	now := time.Now()

	err := db.QueryRow(
		context.Background(),
		`UPDATE users SET first_name = $1, last_name = $2, phone = $3, birth_date = $4, updated_at = $5
		WHERE id = $6
		RETURNING id, email, first_name, last_name, phone, COALESCE(birth_date, ''), created_at, updated_at`,
		user.FirstName, user.LastName, user.Phone, user.BirthDate, now, user.ID,
	).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Phone, &user.BirthDate, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return user, fmt.Errorf("ошибка при обновлении пользователя: %w", err)
	}

	return user, nil
}
