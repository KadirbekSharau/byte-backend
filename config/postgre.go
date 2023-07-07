package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const (
	createUsersTableSQL = `CREATE TABLE users (
			id UUID PRIMARY KEY,
			username VARCHAR(50),
			email VARCHAR(50) UNIQUE,
			password VARCHAR(256),
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);`

	createHabitsTableSQL = `CREATE TABLE habits (
			id UUID PRIMARY KEY,
			user_id UUID REFERENCES users(id),
			name VARCHAR(100),
			description TEXT,
			start_date DATE,
			end_date DATE,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);`

	createHabitProgressTableSQL = `CREATE TABLE habit_progress (
			id UUID PRIMARY KEY,
			habit_id UUID REFERENCES habits(id),
			date DATE,
			status BOOLEAN,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);`

	createTasksTableSQL = `CREATE TABLE tasks (
			id UUID PRIMARY KEY,
			user_id UUID REFERENCES users(id),
			name VARCHAR(100),
			description TEXT,
			start_time TIME,
			end_time TIME,
			date DATE,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);`
)

func NewPostgresDB() (*sqlx.DB, error) {
	cfg := PostgreConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: os.Getenv("POSTGRES_USERNAME"),
		DBName:   os.Getenv("POSTGRES_DBNAME"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("POSTGRES_PASSWORD")}
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = migrateDB(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrateDB(db *sqlx.DB) error {
	if _, err := db.Exec(createUsersTableSQL); err != nil {
		return err
	}
	if _, err := db.Exec(createHabitsTableSQL); err != nil {
		return err
	}
	if _, err := db.Exec(createTasksTableSQL); err != nil {
		return err
	}
	if _, err := db.Exec(createHabitProgressTableSQL); err != nil {
		return err
	}

	return nil
}
