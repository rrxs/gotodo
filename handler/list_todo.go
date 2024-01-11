package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrxs/gotodo/types"
)

// @BasePath /api
// @Summary List todos
// @Description Returns all todo items
// @Tags Todos
// @Accept json
// Produce json
// @Success 200 {object} ListTodoResponse
// @Failure 500 {object} ErrorResponse
// @Router /todos [get]
func ListTodoHandler(c *gin.Context) {
	todos := []types.Todo{}

	if err := db.Find(&todos).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Errorf("error listing todos").Error())
		return
	}

	sendSuccess(c, "listing todo", todos)
}
