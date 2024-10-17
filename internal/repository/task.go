package repository

import (
	"task-manager/database/models"
	"task-manager/internal/data/request"
	"task-manager/pkg/utils"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Save(task models.Task) error
	Update(task models.Task) error
	Delete(taskId uint) error
	FindAll() ([]models.Task, error)
	FindByID(id uint) (models.Task, error)
}

type taskRepositoryImpl struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepositoryImpl {
	return &taskRepositoryImpl{DB: db}
}

func (t *taskRepositoryImpl) Save(task models.Task) error {
	result := t.DB.Create(&task)
	return utils.ReturnError(result.Error)
}

func (t *taskRepositoryImpl) Update(task models.Task) error {
	var updateTask = request.UpdateTaskRequest{
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		DueDate:     task.DueDate,
	}

	result := t.DB.Model(&task).Updates(updateTask)
	return utils.ReturnError(result.Error)
}

func (t *taskRepositoryImpl) Delete(taskId uint) error {
	result := t.DB.Delete(&models.Task{}, taskId)
	return utils.ReturnError(result.Error)
}

func (t *taskRepositoryImpl) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	result := t.DB.Find(&tasks)
	return tasks, utils.ReturnError(result.Error)
}

func (t *taskRepositoryImpl) FindByID(id uint) (models.Task, error) {
	var task models.Task
	result := t.DB.First(&task, id)
	return task, utils.ReturnError(result.Error)
}
