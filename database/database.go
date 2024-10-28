package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"task-manager/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() (db *gorm.DB) {

	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port, _ := strconv.Atoi(portStr)

	if host == "" || port == 0 || user == "" || dbname == "" || password == "" {
		log.Fatalf("Database configuration missing or incomplete")
	}

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	utils.ErrorPanic(err)
	return
}
