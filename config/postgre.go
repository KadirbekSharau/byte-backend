package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const (
	createUsersTableSQL = `CREATE TABLE IF NOT EXISTS users(
		id serial PRIMARY KEY,
		username VARCHAR (50),
		email VARCHAR (50) UNIQUE NOT NULL,
		password VARCHAR (70) NOT NULL
	);`

	createHabitsTableSQL = `CREATE TABLE IF NOT EXISTS habits(
		id serial PRIMARY KEY,
		user_id INT REFERENCES users(id),
		title VARCHAR (100) NOT NULL,
		description TEXT,
		start_date DATE NOT NULL,
		end_date DATE
	);`

	createTasksTableSQL = `CREATE TABLE IF NOT EXISTS tasks(
		id serial PRIMARY KEY,
		user_id INT REFERENCES users(id),
		title VARCHAR (100) NOT NULL,
		description TEXT,
		date DATE NOT NULL
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

	return nil
}
