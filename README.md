# Project Title
Go + Gin REST APIs sample application

## Description
Go simple application for multi integration tools

## Directory Hierarchy
```
go-auth-paseto
├─ api
│  ├─ middleware
│  │  └─ auth_middleware.go
│  ├─ payload
│  │  └─ user_payload.go
│  ├─ response
│  │  ├─ error_response.go
│  │  └─ user_response.go
│  ├─ server.go
│  └─ user.go
├─ app.env
├─ db
│  ├─ migration
│  │  ├─ 000001_add_user.down.sql
│  │  ├─ 000001_add_user.up.sql
│  │  ├─ 000002_add_sessions.down.sql
│  │  └─ 000002_add_sessions.up.sql
│  ├─ query
│  │  ├─ session.sql
│  │  └─ user.sql
│  └─ sqlc
│     ├─ create_user_tx.go
│     ├─ db.go
│     ├─ main_test.go
│     ├─ models.go
│     ├─ querier.go
│     ├─ session.sql.go
│     ├─ store.go
│     ├─ transactional_tx.go
│     ├─ user.sql.go
│     └─ user_test.go
├─ docker-compose.yaml
├─ Dockerfile
├─ docs
│  ├─ docs.go
│  ├─ swagger.json
│  └─ swagger.yaml
├─ go.mod
├─ go.sum
├─ main.go
├─ Makefile
├─ readme.md
├─ sqlc.yaml
├─ token
│  ├─ maker.go
│  ├─ paseto_maker.go
│  ├─ paseto_maker_test.go
│  └─ payload.go
└─ util
   ├─ config.go
   ├─ password.go
   └─ random.go

```

## Getting Started
### Dependencies
   * [Golang Gin](github.com/gin-gonic/gin)
   * Postgre Interactive tool like pgAdmin or Navicat (optional)
   * [Docker](https://www.docker.com) version 24.0.2
   * [Golang Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
   * [SQLC](https://docs.sqlc.dev/en/latest/overview/install.html)
   * [Gin Swagger](https://github.com/swaggo/gin-swagger)
   * Basic tool for Code Editor, e.g. VSCode for golang setup steps [here](https://medium.com/backend-habit/setting-golang-plugin-on-vscode-for-autocomplete-and-auto-import-30bf5c58138a)
   * This project tested and run with Windows 11 64-bit

### Installing
   * WSL2 (windows user only, basically installation already included when installing Docker)
   * Install Docker
   * Golang migrate (for developing usage)
      - For windows, before installing Go Migrate need to install [Scoop](https://scoop.sh), follow steps for installing Go Migrate [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
   * Install SQLC (for developing usage) [here](https://docs.sqlc.dev/en/latest/overview/install.html)
   * Install Swagger CLI (for developing usage) [here](https://github.com/swaggo/swag)

#### Optional Installation
   For windows, there is no default usage for ```make``` command, therefore need to install ```make``` separately,
   * Install chocolatey for windows [here the guidance](https://chocolatey.org/install)
   * Install ```make``` for windows:
   ```
      choco install make
   ```
   * See available choco distribution packages [here](https://community.chocolatey.org/packages)

### Executing program
Steps for running program:
* Install Docker
* Run ```go mod tidy``` or run ```make install_library``` to install all dependencies, see Makefile for details
* Run ```docker-compose up``` to run application
* Open http://localhost:8080/swagger/index.html for APIs Specification

## Acknowledgments
Tutorial
* Big thanks for [Techschool](https://www.youtube.com/@TECHSCHOOLGURU) for complete [tutorial](https://www.youtube.com/watch?v=rx6CPDK_5mU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&pp=iAQB) :smile: :pray: