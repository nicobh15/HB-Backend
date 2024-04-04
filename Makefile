postgres:
	docker run --name postgres --network hb-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

backend:
	docker run --name HB-Backend --network hb-network -e DB_SOURCE="postgresql://root:secret@postgres:5432/hb?sslmode=disable" -e GIN_MODE=release -p 8080:8080 hb:latest
	
createdb:
	docker exec -it postgres createdb --username=root --owner=root hb

dropdb:
	docker exec -it postgres dropdb hb

migrateup:
	migrate -path internal/db/migration -database 'postgresql://root:secret@localhost:5432/hb?sslmode=disable' -verbose up

migratedown:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/hb?sslmode=disable" -verbose down

sqlc:
	sqlc generate

mock:
	mockgen -package mockdb -destination internal/db/mock/store.go github.com/nicobh15/hb-backend/internal/db/sqlc Store

test:
	go test -v -cover ./...

server:
	go run cmd/app/main.go

.PHONY: createdb dropdb migrateup migratedown postgres sqlc test server