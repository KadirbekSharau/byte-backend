package habits

import (
	"context"

	"github.com/KadirbekSharau/Byte/internal/models"
)

type UseCase interface {
	CreateHabit(ctx context.Context, habit *models.Habit) error
	//TrackHabitProgressByDay(ctx context.Context, habitID, progress string) (string, error)
}