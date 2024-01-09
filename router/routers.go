package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rrxs/gotodo/handler"
)

func setupRouters(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/todo", handler.GetTodoHandler)
	}
}
