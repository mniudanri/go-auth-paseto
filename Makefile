network:
	docker network create testenv

rm_network:
	docker network rm testenv

start_postgres:
	docker run --name --network testenv postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:13-alpine

stop_postgres:
	docker stop postgres

rm_postgres:
	docker stop postgres
	docker rm postgres

create_db:
	docker exec -it postgres createdb --username=root --owner=root testdb

drop_db:
	docker exec -it postgres13 dropdb testdb

migrate_up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/testdb?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/testdb?sslmode=disable" -verbose down

build_app:
	docker-compose build --no-cache

sqlc:
	@if [["$(shell docker images -q kjconroy/sqlc:latest)" == ""]]; then \
		($(shell docker pull kjconroy/sqlc)) \
	fi \
	docker run --rm -v $(PWD):/src -w /src kjconroy/sqlc generate

.PHONY: sqlc