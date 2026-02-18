.PHONY: up down test run

up:
	docker compose up -d

down:
	docker compose down
	
test:
	go test ./... -v -count=1

run:
	go run cmd/api/main.go