migrate_up:
	migrate -path db/migrations -database "postgresql:/postgres:postgres@localhost:5432/simplebank?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migrations -database "postgresql:/postgres:postgres@localhost:5432/simplebank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrate_up migrate_down sqlc