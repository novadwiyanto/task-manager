package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoute(r *gin.Engine, db *gorm.DB) *gin.Engine {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Page Not Found",
		})
	})

	TaskRoutes(r, db)
	AuthRoutes(r, db)

	return r
}
