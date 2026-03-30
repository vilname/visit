package constant

// ResponseMessage сообщения ответа сервера
type ResponseMessage string

const (
	// ValidationError ошибки валидации
	ValidationError ResponseMessage = "VALIDATION_ERROR"
)
