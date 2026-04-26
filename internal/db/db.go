package db

import (
	"database/sql"
	"fmt"
	"log"

	"Subly/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB(cfg *config.Config) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	database, err := sql.Open("pgx", connStr)

	if err != nil {
		log.Fatal("Не удалось подключиться к БД", err)
	}

	err = database.Ping()

	if err != nil {
		log.Fatal("Не удалось подключиться к БД", err)
	}

	DB = database

	log.Println("PostgreSQL подключен успешно")
}

func GetDB() *sql.DB {
	return DB
}
