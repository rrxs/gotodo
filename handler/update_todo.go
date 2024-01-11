package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrxs/gotodo/types"
)

func UpdateTodoHandler(c *gin.Context) {
	request := UpdateTodoRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

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

	if request.Title != "" {
		todo.Title = request.Title
	}
	if request.IsDone != nil {
		todo.IsDone = *request.IsDone
	}

	if err := db.Save(&todo).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Errorf("error updating todo item").Error())
		return
	}

	sendSuccess(c, "updating todo", todo)
}
