init:
	docker compose up -d --build

docker-down-clear:
	docker compose down -v --remove-orphans