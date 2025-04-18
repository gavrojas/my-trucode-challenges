package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"trucode.app/user_admin_api/auth"
	"trucode.app/user_admin_api/database"
	"trucode.app/user_admin_api/entries"
	"trucode.app/user_admin_api/models"
	"trucode.app/user_admin_api/shared"
	"trucode.app/user_admin_api/users"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConn := database.CreateDbConnection()
	dbConn.AutoMigrate(&models.User{}, &models.Entry{})

	router := gin.Default()
	router.Use(shared.Cors())
	users.AddUserRoutes(router)
	auth.AddAuthRoutes(router)
	entries.AddEntryRoutes(router)
	router.Run(":3333")
}
