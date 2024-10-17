package database

import (
	"task-manager/database/models"
	"task-manager/pkg/utils"
)

func MigrateDatabase() {
	db := DatabaseConnection()
	err := db.AutoMigrate(
		&models.Task{},
		&models.User{},
	)

	utils.ErrorPanic(err)
}
