package db

import (
	"task-manager/db/models"
	"task-manager/pkg/utils"
)

func MigrateDatabase() {
	db := DatabaseConnection()
	err := db.AutoMigrate(
		&models.Task{},
	)

	utils.ErrorPanic(err)
}