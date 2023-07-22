package api

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/gin-gonic/gin"

	"github.com/mniudanri/go-auth-paseto/api/model"
	"github.com/mniudanri/go-auth-paseto/util"
)

func StartService(server *model.Server, config util.Config) {
	err := server.Router.Run(config.Host)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}

func DefineRoutes(server *model.Server) {
	// list Routes
	server.Router.GET("/users", sampleProcess)
}

func InitServer(config util.Config) (*model.Server) {

	server := &model.Server{
		Router: gin.Default(),
	}

	DefineRoutes(server)

	StartService(server, config)

	return server
}

func sampleProcess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "sample message!"})
	return
}