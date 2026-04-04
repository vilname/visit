# Запуск тестов в docker-compose
docker-compose -f docker-compose-test.yaml up --build app-test

# Запуск с возможностью просмотра отчета о покрытии
docker-compose -f docker-compose-test.yaml up --build app-test-watch

# Запуск только тестов сервиса докторов
docker-compose -f docker-compose-test.yaml run --rm app-test-runner go test -v ./service -run TestDoctorService

# Запуск в коротком режиме (пропускает интеграционные тесты)
go test -short ./...