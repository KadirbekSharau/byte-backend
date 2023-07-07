package habits

import (
	"context"

)

type UseCase interface {
	AddNewHabit(ctx context.Context, name, description string) (string, error)
	TrackHabitProgressByDay(ctx context.Context, habitID, progress string) (string, error)
}