package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrxs/gotodo/types"
)

func sendError(c *gin.Context, code int, msg string) {
	c.Header("content-type", "application/json")
	c.JSON(code, gin.H{
		"error":     msg,
		"errorCode": code,
	})
}

func sendSuccess(c *gin.Context, op string, data interface{}) {
	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Error     string `json:"error"`
	ErrorCode string `json:"errorCode"`
}

type CreateTodoResponse struct {
	Message string             `json:"message"`
	Data    types.TodoResponse `json:"data"`
}
