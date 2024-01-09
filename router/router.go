package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rrxs/gotodo/handler"
)

func Init() {
	r := gin.Default()

	handler.SetupHandlers()

	setupRouters(r)

	r.Run()
}
