package response

import (
	"github.com/gin-gonic/gin"
)

type GenericSuccess struct {
	Data  interface{} `json:"data,omitempty"`
	Token interface{} `json:"token,omitempty"`
}

type GenericError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func SendError(ctx *gin.Context, code int, error *GenericError) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"error": error,
	})
}

func SendSuccess(ctx *gin.Context, code int, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, data)
}

type ErrorResponse struct {
	Error GenericError `json:"error"`
}
