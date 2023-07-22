package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	router := gin.Default()

	router.GET("/users", sampleProcess)
	
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}

func sampleProcess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "sample message!"})
	return
}