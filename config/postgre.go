package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const (
	createUsersTableSQL = `CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			username VARCHAR(50),
			email VARCHAR(50) UNIQUE,
			password VARCHAR(256),
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);`

	createHabitsTableSQL = `CREATE TABLE IF NOT EXISTS habits (
			id UUID PRIMARY KEY,
			user_id UUID REFERENCES users(id),
			name VARCHAR(100),
			description TEXT,
			start_date DATE,
			end_date DATE,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);`

	createTriggersTableSQL = `CREATE TABLE IF NOT EXISTS triggers (
		id UUID PRIMARY KEY,
		habit_id UUID REFERENCES habits(id),
		name VARCHAR(255) NOT NULL,
		description TEXT,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	createHabitLogsTableSQL = `CREATE TABLE IF NOT EXISTS habit_logs (
		id UUID PRIMARY KEY,
		habit_id UUID REFERENCES habits(id),
		user_id UUID REFERENCES users(id),
		date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		status BOOLEAN,
		notes TEXT,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	createHabitRemindersTableSQL = `CREATE TABLE IF NOT EXISTS reminders (
		id UUID PRIMARY KEY,
		habit_id UUID REFERENCES habits(id),
		user_id UUID REFERENCES users(id),
		reminder_time TIMESTAMP WITH TIME ZONE,
		repeat_frequency VARCHAR(255),
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	createHabitRewardsTableSQL = `CREATE TABLE IF NOT EXISTS rewards (
		id UUID PRIMARY KEY,
		user_id UUID REFERENCES users(id),
		reward_name VARCHAR(255),
		date_achieved TIMESTAMP WITH TIME ZONE,
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

	cfg.migrateDB(db)
	return db, nil
}

func (cfg *PostgreConfig) migrateDB(db *sqlx.DB)  {
	executeTable(db, createUsersTableSQL)
	executeTable(db, createHabitsTableSQL)
	executeTable(db, createTriggersTableSQL)
	executeTable(db, createHabitLogsTableSQL)
	executeTable(db, createHabitRemindersTableSQL)
	executeTable(db, createHabitRewardsTableSQL)
}

func executeTable(db *sqlx.DB, table string) error {
	if _, err := db.Exec(table); err != nil {
		return fmt.Errorf("error creating database table: %s", err.Error())
	}
	return nil
}
