FROM golang:1.24-alpine AS development

WORKDIR /app

# Устанавливаем delve и air
RUN apk add --no-cache make && \
    go install github.com/go-delve/delve/cmd/dlv@v1.26 && \
    go install github.com/cosmtrek/air@v1.61.2 && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Проверяем, что dlv установлен и доступен
RUN which dlv && dlv version

# Устанавливаем порты
EXPOSE 8081 2345

CMD dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient --continue ./main.go

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
RUN go install github.com/jstemmer/go-junit-report@latest

# Команда для запуска тестов
# Можно переопределить в docker-compose при необходимости
CMD ["go", "test", "-v", "-cover", "./..."]