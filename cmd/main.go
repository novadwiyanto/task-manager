package main

import (
	"net/http"
	"task-manager/database"
	"task-manager/pkg/utils"
	"task-manager/routes"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.DatabaseConnection()
	r := gin.Default()
	database.MigrateDatabase()

	routes.RegisterRoute(r, db)

	server := &http.Server{
		Addr:           ":8081",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	utils.ErrorPanic(err)
}
