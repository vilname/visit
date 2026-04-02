init:
	docker compose up -d --build

docker-down-clear:
	docker compose down -v --remove-orphans

lint:
	docker compose exec -it app golangci-lint run ./...

lint-fast:
	docker compose exec -it app golangci-lint run --fast ./...

.PHONY: test test-up test-down test-clean

test-up:
	docker-compose -f docker-compose-test.yaml --env-file .env.test up --build

test-down:
	docker-compose -f docker-compose-test.yaml down

test-clean:
	docker-compose -f docker-compose-test.yaml down -v
	rm -rf test-reports/

test-run:
	docker-compose -f docker-compose-test.yaml --env-file .env.test run --rm app-test

test-watch:
	docker-compose -f docker-compose-test.yaml --env-file .env.test run --rm app-test-watch

test-coverage:
	@echo "Opening coverage report..."
	open test-reports/coverage.html || xdg-open test-reports/coverage.html