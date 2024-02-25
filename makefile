include .env
GOPATH ?= $(shell go env GOPATH)

api:
	go run ./cmd/*.go

tests:
	go test ./...

db:
	docker-compose up -d postgres

GOOSE := goose -dir ./migrations postgres "host=$(DB_HOST) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable"
migrate-up:
	$(GOOSE) up

swag:
	@$(GOPATH)/bin/swag init -g cmd/main.go -q

mocks:
	go generate ./internal/...
	