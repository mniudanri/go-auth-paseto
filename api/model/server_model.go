package model

import (
	"github.com/gin-gonic/gin"
)

// Server used for holding passing data
type Server struct {
	Router     *gin.Engine
}