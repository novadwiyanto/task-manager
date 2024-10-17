package controller

import (
	"net/http"
	"task-manager/internal/data/request"
	"task-manager/internal/service"
	"task-manager/pkg/auth"
	"task-manager/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (a *AuthController) Signup(c *gin.Context) {
	var body request.CreateUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.JsonResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		utils.JsonResponse(c, http.StatusBadRequest, "Error : Failed to hash password", nil)
		return
	}

	body.Password = string(hash)
	if err := a.authService.Create(body); err != nil {
		utils.JsonResponse(c, http.StatusInternalServerError, "Error : Failed to create user", err.Error())
		return
	}

	utils.JsonResponse(c, http.StatusCreated, "Created", body)
}

func (a *AuthController) Signin(c *gin.Context) {
	var body request.LoginRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.JsonResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	user, err := a.authService.FindByEmail(body.Email)
	if err != nil {
		utils.JsonResponse(c, http.StatusUnauthorized, "Error : Account not found", err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))
	if err != nil {
		utils.JsonResponse(c, http.StatusUnauthorized, "Error : Invalid password", err.Error())
		return
	}

	token, err := auth.GenerateToken(c, user.ID)
	if err != nil {
		utils.JsonResponse(c, http.StatusInternalServerError, "Error : Failed to generate token", err.Error())
		return
	}

	utils.JsonResponse(c, http.StatusOK, "OK", gin.H{"token": token})
}

func (a *AuthController) Signout(c *gin.Context) {
	c.SetCookie("Authorization", "", 0, "", "", false, true)
	utils.JsonResponse(c, http.StatusOK, "OK", nil)
}
