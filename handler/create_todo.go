package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrxs/gotodo/types"
)

func CreateTodoHandler(c *gin.Context) {
	request := CreateTodoRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	todo := types.Todo{
		Title:  request.Title,
		IsDone: false,
	}

	if err := db.Create(&todo).Error; err != nil {
		logger.Errorf("error creating item on database: %+v", err.Error())
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "creating todo", todo)
}
