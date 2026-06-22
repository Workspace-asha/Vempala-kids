run:
	docker compose up -d postgres
	DB_HOST=localhost DB_PORT=5432 DB_NAME=vempala_kids DB_USER=admin DB_PASSWORD=admin go run ./cmd/api

stop:
	lsof -ti tcp:8080 | xargs -r kill -9 || true

restart: stop run

build:
	go build -o bin/server ./cmd/api

docker-up:
	docker compose up -d

docker-down:
	docker compose down