package db

import (
	"fmt"
	"task-manager/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "task_manager"
)

func DatabaseConnection() (db *gorm.DB) {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	utils.ErrorPanic(err)
	return
}
