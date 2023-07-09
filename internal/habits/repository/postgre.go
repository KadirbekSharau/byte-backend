package repository

import (
	"context"
	"fmt"

	"github.com/KadirbekSharau/Byte/internal/models"
	"github.com/jmoiron/sqlx"
)

type habitRepository struct {
	db *sqlx.DB
}

func NewHabitRepository(db *sqlx.DB) *habitRepository {
	return &habitRepository{
		db: db,
	}
}

func (r *habitRepository) CreateHabit(ctx context.Context, habit *models.Habit) error {
	query := `INSERT INTO habits (id, user_id, name, description, frequency, goal, created_at, updated_at)
	           VALUES (:id, :user_id, :name, :description, :frequency, :goal, :created_at, :updated_at)`
	_, err := r.db.NamedExecContext(ctx, query, habit)
	if err != nil {
		return fmt.Errorf("failed to insert habit: %w", err)
	}
	return nil
}