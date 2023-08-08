package response

import "github.com/gin-gonic/gin"

// Error 400
type Error400 struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// Error 500
type Error500 struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"server error"`
}

func GenerateErrResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
