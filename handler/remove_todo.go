package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrxs/gotodo/types"
)

// @BasePath /api
// @Summary Remove todo
// @Description Remove a todo item
// @Tags Todos
// @Accept json
// Produce json
// @Param id query string true "Todo ID"
// @Success 200 {object} RemoveTodoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /todo [delete]
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
