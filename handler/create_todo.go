package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrxs/gotodo/types"
)

// @BasePath /api
// @Summary Create todo
// @Description Create a new todo item
// @Tags Todos
// @Accept json
// Produce json
// @Param request body CreateTodoRequest true "Request body"
// @Success 200 {object} CreateTodoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /todo [post]
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
