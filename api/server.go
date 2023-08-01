package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	db "github.com/mniudanri/go-auth-paseto/db/sqlc"
	"github.com/mniudanri/go-auth-paseto/token"
	"github.com/mniudanri/go-auth-paseto/util"
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
	// list Routes
	server.Router.POST("/auth/login", server.LoginUser)
	server.Router.POST("/user", server.CreateUser)
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
