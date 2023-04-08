DB_URL=postgresql://root:toor@localhost:5432/gobank?sslmode=disable

postgres:
	docker run --name local-postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=toor -p 5432:5432 -d postgres:latest

createdb:
	docker exec -it local-postgres createdb --username=root --owner=root gobank

dropdb:
	docker exec -it local-postgres dropdb gobank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown