package main

import (
	"fmt"
	"log"
	"net/http"
	"task-manager/database"
	"task-manager/pkg/utils"
	"task-manager/routes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	fmt.Println("Environment variables loaded successfully")

	db := database.DatabaseConnection()
	r := gin.Default()
	database.MigrateDatabase()

	routes.RegisterRoute(r, db)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	utils.ErrorPanic(err)
}
