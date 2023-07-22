package api

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/gin-gonic/gin"

	"github.com/mniudanri/go-auth-paseto/api/models"
)

func InitServer() (*model.Server) {

	router := gin.Default()
	
	// list routes
	router.GET("/users", sampleProcess)

	Start(router)

	server := &model.Server{
		Router: router,
	}
	return server
}

func Start(router *gin.Engine) {
	err := router.Run("0.0.0.0:8080")

	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}

func sampleProcess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "sample message!"})
	return
}