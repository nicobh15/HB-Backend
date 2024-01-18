postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root homebuddy

dropdb:
	docker exec -it postgres dropdb homebuddy

migrateup:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/homebuddy?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/homebuddy?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb migrateup migratedown postgres sqlc test