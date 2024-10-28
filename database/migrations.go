package database

import (
	"task-manager/database/models"
	"task-manager/pkg/utils"
)

func MigrateDatabase() {
	db := DatabaseConnection()
	Query()

	err := db.AutoMigrate(
		&models.Task{},
		&models.User{},
		&models.Team{},
		&models.Category{},
		&models.TeamMember{},
		&models.TaskHistory{},
		&models.Notification{},
	)

	utils.ErrorPanic(err)
}
