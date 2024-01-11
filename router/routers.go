package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rrxs/gotodo/docs"
	"github.com/rrxs/gotodo/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouters(r *gin.Engine) {
	handler.SetupHandlers()
	basePath := "/api"

	docs.SwaggerInfo.Title = "Gotodo API"
	docs.SwaggerInfo.Description = "This is a simple API created to manage Todo items."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = basePath
	api := r.Group(basePath)
	{
		api.POST("/todo", handler.CreateTodoHandler)
		api.DELETE("/todo", handler.RemoveTodoHandler)
		api.GET("/todos", handler.ListTodoHandler)
		api.GET("/todo", handler.GetTodoHandler)
		api.PUT("/todo", handler.UpdateTodoHandler)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
