package response

import "github.com/gin-gonic/gin"

func GenerateErrResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
