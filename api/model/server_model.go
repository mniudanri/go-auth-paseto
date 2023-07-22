package model

import (
	"github.com/gin-gonic/gin"

	"github.com/mniudanri/go-auth-paseto/util"
)

// Server used for holding passing data
type Server struct {
	Router     *gin.Engine
	Config     util.Config
}