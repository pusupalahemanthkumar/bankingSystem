postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d --rm postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root banking_system

dropdb:
	docker exec -it postgres12 dropdb banking_system


migrateinit: 
	migrate create -ext sql -dir db/migration -seq init_schema

migratecreate:
    # make migratecreate name=<migrationName>
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/banking_system?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/banking_system?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/banking_system?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/banking_system?sslmode=disable" -verbose down 1


sqlcinit: 
	sqlc init

sqlcgenerate: 
	sqlc generate

test: 
	go test -v cover "./..."

server:
	go run main.go 

mock: 
	mockgen -package mockdb -destination  db/mock/store.go github.com/pusupalahemanthkumar/bankingsystem/db/sqlc Store 


.PHONY: postgres createdb dropdb migrateinit migratecreate migrateup migrateup1 migratedown migratedown1 sqlcinit sqlcgenerate test server mock









