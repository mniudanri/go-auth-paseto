package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/mniudanri/go-auth-paseto/api/middleware"
	db "github.com/mniudanri/go-auth-paseto/db/sqlc"
	"github.com/mniudanri/go-auth-paseto/token"
	"github.com/mniudanri/go-auth-paseto/util"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Router     *gin.Engine
	Config     util.Config
	Store      db.Store
	TokenMaker token.Maker
}

func StartService(server *Server, config util.Config) {
	err := server.Router.Run(config.Host)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
func DefineRoutes(server *Server) {
	v1 := server.Router.Group("/v1")
	{
		// list Routes
		v1.POST("/auth/login", server.LoginUser)
		v1.POST("/user", server.CreateUser)

		authRoutes := v1.Group("/").Use(middleware.AuthMiddleware(server.TokenMaker))
		authRoutes.GET("/user/:username", server.GetUserByUsername)
	}
	server.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
func GenerateToken(config util.Config) token.Maker {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Warn().Msg(fmt.Sprint("cannot create token maker: %w", err))
	}

	return tokenMaker
}
func InitServer(config util.Config) *Server {
	server := &Server{
		Router:     gin.Default(),
		TokenMaker: GenerateToken(config),
		Store:      db.CreateConnection(config),
		Config:     config,
	}

	DefineRoutes(server)
	StartService(server, config)

	return server
}
