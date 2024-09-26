package service

import (
	"fmt"
	"task-manager/db/models"
	"task-manager/internal/data/request"
	"task-manager/internal/data/response"
	"task-manager/internal/repository"
	"task-manager/pkg/utils"

	"github.com/go-playground/validator/v10"
)

type TaskService interface {
	Create(task request.CreateTaskRequest) error
	Update(task request.UpdateTaskRequest) error
	Delete(taskId uint) error
	FindAll() ([]response.Task, error)
	FindByID(taskId uint) (response.Task, error)
}

type TaskServiceImpl struct {
	TaskRepository repository.TaskRepository
	Validate       *validator.Validate
}

func NewTaskService(taskRepository repository.TaskRepository, validate *validator.Validate) TaskService {
	return &TaskServiceImpl{
		TaskRepository: taskRepository,
		Validate:       validate,
	}
}

func (t *TaskServiceImpl) Create(task request.CreateTaskRequest) error {
	fmt.Println("Create Task Service")
	err := t.Validate.Struct(task)
	if err != nil {
		return utils.ReturnError(err)
	}

	var newTask = models.Task{
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		DueDate:     task.DueDate,
	}

	return utils.ReturnError(t.TaskRepository.Save(newTask))
}

func (t *TaskServiceImpl) Update(task request.UpdateTaskRequest) error {
	err := t.Validate.Struct(task)
	if err != nil {
		return utils.ReturnError(err)
	}

	taskToUpdate, err := t.TaskRepository.FindByID(task.ID)
	if err != nil {
		return utils.ReturnError(err)
	}

	taskToUpdate.Title = task.Title
	taskToUpdate.Description = task.Description
	taskToUpdate.Status = task.Status
	taskToUpdate.Priority = task.Priority
	taskToUpdate.DueDate = task.DueDate

	return utils.ReturnError(t.TaskRepository.Update(taskToUpdate))
}

func (t *TaskServiceImpl) Delete(taskId uint) error {

	return utils.ReturnError(t.TaskRepository.Delete(taskId))
}

func (t *TaskServiceImpl) FindAll() ([]response.Task, error) {
	tasks, err := t.TaskRepository.FindAll()
	if err != nil {
		return nil, utils.ReturnError(err)
	}

	var responseTasks []response.Task
	for _, task := range tasks {
		responseTasks = append(responseTasks, response.Task{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			Priority:    task.Priority,
			DueDate:     task.DueDate,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	return responseTasks, nil
}

func (t *TaskServiceImpl) FindByID(taskId uint) (response.Task, error) {
	result, err := t.TaskRepository.FindByID(taskId)
	if err != nil {
		return response.Task{}, utils.ReturnError(err)
	}

	return response.Task{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Status:      result.Status,
		Priority:    result.Priority,
		DueDate:     result.DueDate,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}, nil
}
