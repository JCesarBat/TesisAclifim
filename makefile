postgres:
	 docker run --name acklifim -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:latest

redis:
	docker run --name redis -p 6379:6379 -d redis

createdb:
	docker exec -it acklifim createdb --username=root --owner=root ackilifm_db

dropdb:
	docker exec -it acklifim dropdb ackilifm_db

migrateup:
	migrate -path database/migrations -database "postgresql://root:postgres@localhost:5432/ackilifm_db?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migrations -database "postgresql://root:postgres@localhost:5432/ackilifm_db?sslmode=disable" -verbose down
#migrate create -ext sql -dir database/migrations -seq <migration-name>

sqlc:
	sqlc generate
server:
	go run cmd/main.go

docs:
	swag init -g ./cmd/main.go -o pkg/docs

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock proto evans_linux redis db_schema docs
