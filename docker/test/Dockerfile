# ========== ДОБАВЛЯЕМ СТАДИЮ ДЛЯ ТЕСТОВ ==========
FROM golang:1.24-alpine AS test

WORKDIR /app

# Устанавливаем необходимые утилиты для тестов
RUN apk add --no-cache git make curl

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Устанавливаем инструменты для генерации отчетов о тестах (опционально)
RUN go install github.com/jstemmer/go-junit-report@v1.0.0

# Создаем директорию для отчетов
RUN mkdir -p /app/test-reports

# Точка входа для тестов - используем cmd/test.go
# Аргументы можно передавать через docker-compose
ENTRYPOINT ["go", "run", "cmd/test.go"]

# Дефолтная команда (запуск всех тестов)
CMD ["-pkg=./...", "-v=true", "-cover=true"]