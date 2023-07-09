package models

import "time"

type Habit struct {
	ID          string    `db:"id"`
	UserID      string    `db:"user_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Frequency   string    `db:"frequency"`
	Goal        string    `db:"goal"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type HabitProgress struct {
	ID        int       `db:"id"`
	HabitId   int       `db:"habit_id"`
	Progress  int       `db:"progress"`
	Date      time.Time `db:"date"`
}