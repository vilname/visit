init:
	docker compose up -d --build

docker-down-clear:
	docker compose down -v --remove-orphans

lint:
	docker compose exec -it app golangci-lint run ./...

lint-fast:
	docker compose exec -it app golangci-lint run --fast ./...

test-api:
	docker compose -f docker-compose-test.yaml run --rm app-test
