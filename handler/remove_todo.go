package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrxs/gotodo/types"
)

func RemoveTodoHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, fmt.Errorf("query param 'id' is required").Error())
		return
	}

	todo := types.Todo{}

	if err := db.First(&todo, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("todo not found. 'id': %s", id))
		return
	}

	if err := db.Delete(&todo).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error removing todo item. 'id': %s", id))
		return
	}

	sendSuccess(c, "removing todo", todo)
}
