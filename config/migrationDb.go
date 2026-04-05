// Package config конфиг
package config

import "visit/dbmigrate"

// InitMigrationDB создаёт подключение к базе и применяет миграции вверх.
func InitMigrationDB() error {
	return dbmigrate.Up()
}
