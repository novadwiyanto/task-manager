package auth

import (
	"fmt"
	"net/http"
	"task-manager/database"
	"task-manager/database/models"
	"task-manager/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {

	fmt.Println("RequireAuth")
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, utils.ReturnError(fmt.Errorf("unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte("secret"), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(claims["exp"].(float64)) < float64(time.Now().Unix()) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		database.DatabaseConnection().Where("id = ?", claims["id"]).First(&user)

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
