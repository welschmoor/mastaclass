postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root listings

dropdb:
	docker exec -it postgres15 dropdb listings

## cremig: create migration; needs an argument name=custom_migration_name
cremig:
	migrate create -ext sql -dir db/migration -seq ${name}

mup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/listings?sslmode=disable" -verbose up

mdown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/listings?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

itdb:
	docker exec -it postgres15 psql -U root

.PHONY: postgres createdb dropdb cremig mup mdown sqlc test itdb