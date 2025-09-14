.PHONY: help
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: run
run: ## Run the application
	go run cmd/server/main.go

.PHONY: build
build: ## Build the application
	go build -o bin/server cmd/server/main.go

.PHONY: test
test: ## Run tests
	go test -v ./...

.PHONY: migrate-up
migrate-up: ## Run database migrations up
	migrate -path migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" up

.PHONY: migrate-down
migrate-down: ## Run database migrations down
	migrate -path migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" down

.PHONY: sqlc
sqlc: ## Generate code from SQL
	sqlc generate

.PHONY: swagger
swagger: ## Generate swagger documentation
	swag init -g cmd/server/main.go -o docs

.PHONY: deps
deps: ## Download dependencies
	go mod download
	go mod tidy

.PHONY: clean
clean: ## Clean build artifacts
	rm -rf bin/ tmp/ logs/