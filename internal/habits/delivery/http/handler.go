package http

import (
	"net/http"

	"github.com/KadirbekSharau/Byte/internal/habits"
	"github.com/KadirbekSharau/Byte/internal/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase habits.UseCase
}

func NewHandler(useCase habits.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) CreateHabit(c *gin.Context) {
	var habit models.Habit

	if err := c.BindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.useCase.CreateHabit(c.Request.Context(), &habit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
