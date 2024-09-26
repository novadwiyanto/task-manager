package main

import (
	"net/http"
	"task-manager/db"
	"task-manager/internal/routes"
	"task-manager/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.DatabaseConnection()
	r := gin.Default()

	routes.RegisterRoute(r, db)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	utils.ErrorPanic(err)
}
