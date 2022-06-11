before:
	@docker run -itd -p 5432:5432 -e POSTGRES_PASSWORD=${PASSWORD} --rm --name postgres postgres:12.11-alpine

start:
	@go run main.go

migrate:
	@go run database/main.go

build:
	@go build GOOS=linux GOARCH=amd64 -o app

stop:
	@docker stop postgres