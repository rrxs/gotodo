package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rrxs/gotodo/handler"
)

func setupRouters(r *gin.Engine) {
	handler.SetupHandlers()

	api := r.Group("/api")
	{
		api.POST("/todo", handler.CreateTodoHandler)
		api.DELETE("/todo", handler.RemoveTodoHandler)
		api.GET("/todos", handler.ListTodoHandler)
		api.GET("/todo", handler.GetTodoHandler)
		api.PUT("/todo", handler.UpdateTodoHandler)
	}
}
