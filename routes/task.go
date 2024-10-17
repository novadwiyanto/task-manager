package routes

import (
	"task-manager/internal/controller" // Make sure to import the service package
	"task-manager/internal/repository"
	"task-manager/internal/service"
	"task-manager/pkg/auth"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func TaskRoutes(r *gin.Engine, db *gorm.DB) {
	taskRepo := repository.NewTaskRepository(db)
	validate := validator.New()
	taskService := service.NewTaskService(taskRepo, validate)

	t := controller.NewTaskController(taskService)

	task := r.Group("/task", auth.RequireAuth)
	{
		task.GET("", t.FindAll)
		task.GET("/:id", t.FindById)
		task.POST("", t.Create)
		task.PUT("/:id", t.Update)
		task.DELETE("/:id", t.Delete)
	}
}
