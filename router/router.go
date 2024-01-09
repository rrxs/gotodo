package router

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.Default()

	setupRouters(r)

	r.Run()
}
