start_postgres:
	docker run --name postgres13 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:13-alpine

stop_postgres:
	docker stop postgres13

remove_postgres:
	docker stop postgres13
	docker rm postgres13

create_db:
	docker exec -it postgres13 createdb --username=root --owner=root simple_bank

drop_db:
	docker exec -it postgres13 dropdb simple_bank

migrate_up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	@if [["$(shell docker images -q kjconroy/sqlc:latest)" == ""]]; then \
		($(shell docker pull kjconroy/sqlc)) \
	fi \
	
	docker run --rm -v $(PWD):/src -w /src kjconroy/sqlc generate

.PHONY: sqlc