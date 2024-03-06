package errors

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func HandlerError(c *gin.Context, code string, message string, httpStatus int) {
	c.JSON(httpStatus, ErrorResponse{
		Code:    code,
		Message: message,
	})
}
