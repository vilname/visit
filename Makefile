init:
	docker compose up -d --build

docker-down-clear:
	docker compose down -v --remove-orphans

lint:
	golangci-lint run ./...

lint-fast:
	golangci-lint run --fast ./...