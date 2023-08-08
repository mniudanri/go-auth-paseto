package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	res "github.com/mniudanri/go-auth-paseto/api/response"
	"github.com/mniudanri/go-auth-paseto/token"
)

const (
	authHeader = "authorization"
	authBearer = "bearer"
)

// AuthMiddleware creates a gin middleware for authorization
func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authHeader)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res.GenerateErrResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res.GenerateErrResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res.GenerateErrResponse(err))
			return
		}

		accessToken := fields[1]
		_, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res.GenerateErrResponse(err))
			return
		}

		ctx.Next()
	}
}
