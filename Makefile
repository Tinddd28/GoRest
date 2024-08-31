postgres:
	docker-compose up -d

createdb:
	docker exec -it postgres_container createdb -U root -O root simple_app

dropdb:
	docker exec -it postgres_container dropdb simple_app

migrateup:
	migrate -path db/migration -database "postgresql://root:12345678@localhost:5432/simple_app?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:12345678@localhost:5432/simple_app?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown