package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouters(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/todo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data": "GET todo",
			})
		})
	}
}
