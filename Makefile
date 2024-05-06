postgres:
	docker run --name schedule -p 5432:5432  -e POSTGRES_USER=postgres  -e POSTGRES_PASSWORD=postgres -d postgres:latest

createdb:
	docker exec -it schedule createdb --username=postgres --owner=postgres schedules_data

migratecreate:
	migrate create -ext sql -dir internal/repository/migrations/ -seq init_schema

migrateup:
	 migrate -path internal/repository/migrations/ -database "postgresql://postgres:postgres@$5432:5432/schedules_data?sslmode=disable" -verbose up

dropdb:
	docker exec -it schedule dropdb schedules

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@$5432:5432/schedules_data?sslmode=disable" -verbose down

.PHONY: postgres createdb createtestdb dropdb migrateup migratedown migratecreate