dbname := "listings"
dsn := "postgresql://root:secret@localhost:5432/${dbname}?sslmode=disable"

postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

stoprun:
	docker kill postgres15

credb:
	docker exec -it postgres15 createdb --username=root --owner=root ${dbname}

dropdb:
	docker exec -it postgres15 dropdb ${dbname}

## cremig: create migration; needs an argument name=custom_migration_name
cremig:
	migrate create -ext sql -dir db/migration -seq ${name}

mup:
	migrate -path db/migration -database ${dsn} -verbose up

mdown:
	migrate -path db/migration -database ${dsn} -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

itdb:
	docker exec -it postgres15 psql -U root

.PHONY: postgres credb dropdb cremig mup mdown sqlc test itdb