package api

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/gin-gonic/gin"

	"github.com/mniudanri/go-auth-paseto/api/model"
)

func StartService(server *model.Server) {
	err := server.Router.Run("0.0.0.0:8080")

	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}

func DefineRoutes(server *model.Server) {
	// list Routes
	server.Router.GET("/users", sampleProcess)
}

func InitServer() (*model.Server) {

	server := &model.Server{
		Router: gin.Default(),
	}

	DefineRoutes(server)

	StartService(server)

	return server
}

func sampleProcess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "sample message!"})
	return
}