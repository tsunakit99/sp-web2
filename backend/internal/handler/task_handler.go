package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tsunakit99/sp-web2/backend/internal/usecase"
)

type TaskHandler struct {
	usecase usecase.TaskUsecase
}

func NewTaskHandler(u usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{usecase: u}
}

func RegisterTaskRoutes(g *echo.Group, h *TaskHandler) {
	g.GET("/tasks", h.GetTasks)
	// 将来的にここに POST /tasks なども追加可能
}

// GET /tasks
func (h *TaskHandler) GetTasks(c echo.Context) error {
	userID := c.Get("userID").(string) // JWTミドルウェアでセットされる前提

	tasks, err := h.usecase.GetTasks(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch tasks")
	}
	return c.JSON(http.StatusOK, tasks)
}