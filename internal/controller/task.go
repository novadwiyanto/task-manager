package controller

import (
	"net/http"
	"strconv"
	"task-manager/internal/data/request"
	"task-manager/internal/service"
	"task-manager/pkg/utils"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService service.TaskService
}

func NewTaskController(taskService service.TaskService) *TaskController {
	return &TaskController{taskService: taskService}
}

func (t *TaskController) Create(ctx *gin.Context) {
	var task request.CreateTaskRequest
	if err := ctx.ShouldBindJSON(&task); err != nil {
		utils.JsonResponse(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	if err := t.taskService.Create(task); err != nil {
		utils.JsonResponse(ctx, http.StatusInternalServerError, "Error : Failed to create task", err.Error())
		return
	}

	utils.JsonResponse(ctx, http.StatusCreated, "Created", nil)
}

func (t *TaskController) Update(ctx *gin.Context) {
	var task request.UpdateTaskRequest
	if err := ctx.ShouldBindJSON(&task); err != nil {
		utils.JsonResponse(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	taskID := ctx.Param("id")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		utils.JsonResponse(ctx, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}
	task.ID = uint(id)

	if err := t.taskService.Update(task); err != nil {
		utils.JsonResponse(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.JsonResponse(ctx, http.StatusOK, "OK", nil)
}

func (t *TaskController) Delete(ctx *gin.Context) {
	taskID := ctx.Param("id")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		utils.JsonResponse(ctx, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	if err := t.taskService.Delete(uint(id)); err != nil {
		utils.JsonResponse(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.JsonResponse(ctx, http.StatusOK, "OK", nil)
}

func (t *TaskController) FindById(ctx *gin.Context) {
	taskID := ctx.Param("id")
	id, err := strconv.Atoi(taskID)
	if err != nil {
		utils.JsonResponse(ctx, http.StatusBadRequest, "Invalid ID", err.Error())
		return
	}

	data, err := t.taskService.FindByID(uint(id))
	if err != nil {
		utils.JsonResponse(ctx, http.StatusNotFound, "Task Not Found", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (t *TaskController) FindAll(ctx *gin.Context) {
	data, err := t.taskService.FindAll()
	if err != nil {
		utils.JsonResponse(ctx, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.JsonResponse(ctx, http.StatusOK, "OK", data)
}
