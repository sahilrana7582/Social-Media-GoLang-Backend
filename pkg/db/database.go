package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("Unable to ping database:", err)
	}
	fmt.Println("Successfully connected to the database")
}

func GetDB() *sql.DB {
	if DB == nil {
		log.Fatal("Database connection is not initialized")
	}
	return DB
}
