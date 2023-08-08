package main

import (
	"github.com/mniudanri/go-auth-paseto/api"
	"github.com/mniudanri/go-auth-paseto/util"
)

// @title           API Specification
// @version         1.0
// @description     Specification for APIs
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8080
// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Config field for swagger
	util.SetupSwagger()

	// Load config from env
	config := util.LoadConfig(".")

	// Define routes, create connection to db,
	// and start service
	server := api.InitServer(config)

	// blank identifier for further used
	_ = server

	// TODO:
	// Connect multi db or engine (e.g ELK)?
	// Connect monitor (e.g. New Relic)?
}
