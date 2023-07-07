package habits

import (
	"context"

	"github.com/KadirbekSharau/Byte/internal/models"
)

type UserRepository interface {
	CreateHabit(ctx context.Context, user *models.Habit) (string, error)
	CreateHabitProgressByDay(ctx context.Context, progress string) (*models.HabitProgress, error)
}