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

// type HabitProgress struct {
// 	ID        int       `db:"id"`
// 	HabitId   int       `db:"habit_id"`
// 	Progress  int       `db:"progress"`
// 	Date      time.Time `db:"date"`
// }

type Trigger struct {
	ID          string    `db:"id"`
	HabitID     string    `db:"habit_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type HabitLog struct {
	ID        string    `db:"id"`
	HabitID   string    `db:"habit_id"`
	UserID    string    `db:"user_id"`
	Date      time.Time `db:"date"`
	Status    bool      `db:"status"`
	Notes     string    `db:"notes"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type HabitReminder struct {
	ID            string    `db:"id"`
	HabitID       string    `db:"habit_id"`
	UserID        string    `db:"user_id"`
	ReminderTime  time.Time `db:"reminder_time"`
	RepeatFrequency string    `db:"repeat_frequency"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type Reward struct {
	ID           string    `db:"id"`
	UserID       string    `db:"user_id"`
	RewardName   string    `db:"reward_name"`
	DateAchieved time.Time `db:"date_achieved"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}