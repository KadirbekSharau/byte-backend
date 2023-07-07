package models

import "time"


type Habit struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
	UserId    int       `db:"user_id"`
}

type HabitProgress struct {
	ID        int       `db:"id"`
	HabitId   int       `db:"habit_id"`
	Progress  int       `db:"progress"`
	Date      time.Time `db:"date"`
}