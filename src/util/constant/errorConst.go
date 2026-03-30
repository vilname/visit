// Package constant константы
package constant

// ErrorType тип ошибки
type ErrorType string

// ErrorMessage сообщение ошибки
type ErrorMessage string

const (
	// MaxAttemptGenerateCode сообщение об ошибке
	MaxAttemptGenerateCode ErrorMessage = "The number of attempts exceeded"
)

const (
	// MaxAttempt ошибка количества попыток
	MaxAttempt ErrorType = "MAX_ATTEMPT"

	// DateError ошибка даты
	DateError ErrorType = "DATE_ERROR"
)
