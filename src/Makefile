createdb:
	docker exec -it docker-pg createdb --username=postgres --owner=postgres go_starter

dropdb:
	docker exec -it docker-pg dropdb go_starter

migrateup:
	migrate -path src/db/migration -database "postgresql://postgres:password@localhost:5433/go_starter?sslmode=disable" -verbose up

migratedown:
	migrate -path src/db/migration -database "postgresql://postgres:password@localhost:5433/go_starter?sslmode=disable" -verbose down

sqlc:
	sqlc generate

run:
	go run main.go

.PHONY: createdb dropdb migrateup migratedown sqlc run
