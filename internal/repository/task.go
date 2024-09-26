package repository

import (
	"errors"
	"task-manager/db/models"
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

type TaskRepositoryImpl struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{DB: db}
}

func (t *TaskRepositoryImpl) Save(task models.Task) error {
	result := t.DB.Create(&task)
	return utils.ReturnError(result.Error)
}

func (t *TaskRepositoryImpl) Update(task models.Task) error {
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

func (t *TaskRepositoryImpl) Delete(taskId uint) error {
	result := t.DB.Delete(&models.Task{}, taskId)
	return utils.ReturnError(result.Error)
}

func (t *TaskRepositoryImpl) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	result := t.DB.Find(&tasks)
	if err := utils.ReturnError(result.Error); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskRepositoryImpl) FindByID(id uint) (models.Task, error) {
	var task models.Task
	result := t.DB.First(&task, id)
	if result.Error != nil {
		return task, errors.New("record not found")
	}

	return task, nil
}
