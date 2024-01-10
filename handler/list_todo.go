package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrxs/gotodo/types"
)

func ListTodoHandler(c *gin.Context) {
	todos := []types.Todo{}

	if err := db.Find(&todos).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Errorf("error listing todos").Error())
		return
	}

	sendSuccess(c, "listing todo", todos)
}
