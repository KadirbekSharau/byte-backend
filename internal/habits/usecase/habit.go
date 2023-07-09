package usecase

import (
	"context"
	"fmt"

	"github.com/KadirbekSharau/Byte/internal/habits"
	"github.com/KadirbekSharau/Byte/internal/models"
)

type habitUseCase struct {
	repo habits.Repository
}

func NewHabitUseCase(repo habits.Repository) *habitUseCase {
	return &habitUseCase{
		repo: repo,
	}
}

func (uc *habitUseCase) CreateHabit(ctx context.Context, habit *models.Habit) error {
	err := uc.repo.CreateHabit(ctx, habit)
	if err != nil {
		return fmt.Errorf("failed to create habit: %w", err)
	}
	return nil
}