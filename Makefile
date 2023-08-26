postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb: 
	docker exec postgres12 createdb --username=root --owner=root go_bank

dropdb: 
	docker exec postgres12 dropdb go_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/go_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/go_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	air

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server