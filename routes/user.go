package routes

import (
	"task-manager/internal/controller"
	"task-manager/internal/repository"
	"task-manager/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func AuthRoutes(r *gin.Engine, db *gorm.DB) {
	authRepository := repository.NewAuthRepository(db)
	validate := validator.New()
	authService := service.NewAuthService(authRepository, validate)
	a := controller.NewAuthController(authService)

	auth := r.Group("/auth")
	{
		auth.POST("/register", a.Signup)
		auth.POST("/login", a.Signin)
		auth.POST("/logout", a.Signout)
	}
}
