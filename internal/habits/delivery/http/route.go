package http

import (
	"github.com/KadirbekSharau/Byte/internal/habits"
	"github.com/gin-gonic/gin"
)

func RegisterHabitHTTPEndpoints(router *gin.Engine, uc habits.UseCase) {
	h := NewHandler(uc)
	habitEndpoints := router.Group("/habits")
	{
		habitEndpoints.POST("/create", h.CreateHabit)
	}
}