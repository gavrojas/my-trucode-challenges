package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"trucode.app/apitaskmanager/database"
	"trucode.app/apitaskmanager/models"
	"trucode.app/apitaskmanager/projects"
	"trucode.app/apitaskmanager/tasks"
	"trucode.app/apitaskmanager/users"
)

func main() {
	// Cargar variables de entorno -|
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env file with vars")
	}

	// Conectarme a la base de datos -|
	dbconn := database.CreateDbConnection()
	// Crear si es que no existen mis tablas en la base de datos -|
	dbconn.AutoMigrate(&models.Project{}, &models.Task{}, &models.User{})

	router := gin.Default()
	projects.AddProjectsRoutes(router)
	tasks.AddTasksRoutes(router)
	users.AddUsersRoutes(router)
	router.Run("0.0.0.0:3333")
}
