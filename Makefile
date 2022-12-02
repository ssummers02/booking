-include ./.env

# Default command for prepare code for a commit
default: format lint-fix

build:
	docker-compose build link-app

run:
	docker-compose up link-app

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.1

migrate-new:
	migrate create -ext sql -dir ./migrations "$(name)"
#[CI\CD]
migrate-up:
	migrate -database="postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&&query" -path ./migrations up

# rollback last migration
migrate-down:
	migrate -database="postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&&query" -path ./migrations down 1

# drop all tables into database
migrate-drop:
	migrate -database="postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&&query" -path ./migrations drop -f

migrate-rebase: migrate-drop migrate-up

# [CI\CD] Auto-format code
format:
	gofmt -s -w . && \
	go vet ./... && \
	go mod tidy

### Linter ###

lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.0

# [CI\CD] lint code
lint:
	golangci-lint run
# lint and auto-fix possible problems
lint-fix:
	golangci-lint run --fix
