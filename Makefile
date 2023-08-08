# base running application for docker
network:
	docker network create testenv

# delete network docker
rm_network:
	docker network rm testenv

# setup postgresql full documentation see Docker Hub (https://hub.docker.com/_/postgres)
# use alpine for small image
start_postgres:
	docker run --name --network testenv postgres -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:13-alpine

# stop running postgre
stop_postgres:
	docker stop postgres

# remove postgre image in docker
rm_postgres:
	docker stop postgres
	docker rm postgres

# query create db for postgre in docker image
create_db:
	docker exec -it postgres createdb --username=root --owner=root testdb

# remove db from postgre image
drop_db:
	docker exec -it postgres13 dropdb testdb

# run go migrate, create table structure and indexing to postgre, see query at ./db/migration/*.up.sql
migrate_up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/testdb?sslmode=disable" -verbose up

# run go migrate, rm table structure and indexing in postgre, see query at ./db/migration/*.down.sql
migrate_down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/testdb?sslmode=disable" -verbose down

# docker build no cache, for managing resource memory
build_app:
	docker-compose build --no-cache

# remove all unused docker data: volumes, image, container, build cache
rm_unused_data:
	docker system prune --all --volumes

# remove cache after build docker and go
rm_cache_data:
	docker builder prune
	go clean -cache

# generate sqlc.yaml
sqlc_init:
	@if [["$(shell docker images -q sqlc/sqlc:latest)" == ""]]; then \
		($(shell docker pull sqlc/sqlc)) \
	fi \
	docker run --rm -v $(PWD):/src -w /src sqlc/sqlc init

# generate query and code for integration to db
sqlc:
	@if [["$(shell docker images -q sqlc/sqlc:latest)" == ""]]; then \
		($(shell docker pull sqlc/sqlc)) \
	fi \
	docker run --rm -v $(PWD):/src -w /src sqlc/sqlc generate

# APIs specification
install_library:
	gom mod tidy

# APIs specification
swagger:
	gom mod tidy
	swag init

.PHONY: sqlc swagger network