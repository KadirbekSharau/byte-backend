package habits

import (
	"context"

	"github.com/KadirbekSharau/Byte/internal/models"
)

type Repository interface {
	CreateHabit(ctx context.Context, user *models.Habit) error
	//CreateHabitProgressByDay(ctx context.Context, progress string) (*models.HabitProgress, error)
}