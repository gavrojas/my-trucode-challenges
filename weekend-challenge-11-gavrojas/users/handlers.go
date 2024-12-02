package users

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"trucode.app/apitaskmanager/database"
	"trucode.app/apitaskmanager/models"
)

func Create(c *gin.Context) {
	var engineer models.User
	id, match := c.Params.Get("id")
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get project id"})
		return
	}

	c.BindJSON(&engineer)

	dbconn := database.CreateDbConnection()
	pid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse project id"})
		return
	}

	engineer.ProjectID = uint(pid)

	if tx := dbconn.Create(&engineer); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create a task" + tx.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": engineer})
}

func ReadAll(c *gin.Context) {
	var engineers []map[string]interface{}
	dbconn := database.CreateDbConnection()

	dbconn.Model(models.User{}).Select("id", "name").Find(&engineers)

	if len(engineers) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No users found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Users found", "users": engineers})
}
func ReadByID(c *gin.Context) {
	var engineer models.User
	id, match := c.Params.Get("id")

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	dbconn := database.CreateDbConnection()
	// Precargar proyecto -|
	if tx := dbconn.Preload("Project").First(&engineer, id); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found" + tx.Error.Error()})
		}
		return
	}

	userResponse := map[string]interface{}{
		"id":      engineer.ID,
		"name":    engineer.Name,
		"surname": engineer.Surname,
		"project": map[string]interface{}{
			"id":       engineer.Project.ID,
			"name":     engineer.Project.Name,
			"delivery": engineer.Project.DeliveryDate,
		},
	}

	c.JSON(http.StatusOK, gin.H{"message": "User found", "user": userResponse})
}

func ReasignProject(c *gin.Context) {
	var engineer models.User
	idProject, match := c.Params.Get("idProject")
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get project id"})
		return
	}

	idUser, match := c.Params.Get("idUser")
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get user id"})
		return
	}

	dbconn := database.CreateDbConnection()

	if tx := dbconn.First(&engineer, idUser); tx.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Engineer not found"})
		return
	}

	pid, err := strconv.Atoi(idProject)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse project id"})
		return
	}

	engineer.ProjectID = uint(pid)

	if tx := dbconn.Save(&engineer); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to reassing project" + tx.Error.Error()})
		return
	}

	// Precargar proyecto de nuevo-|
	if tx := dbconn.Preload("Project").First(&engineer, idUser); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found" + tx.Error.Error()})
		}
		return
	}

	// Imprimir respuesta con nuevo proyecto -|
	userResponse := map[string]interface{}{
		"id":      engineer.ID,
		"name":    engineer.Name,
		"surname": engineer.Surname,
		"project": map[string]interface{}{
			"id":       engineer.Project.ID,
			"name":     engineer.Project.Name,
			"delivery": engineer.Project.DeliveryDate,
		},
	}

	c.JSON(http.StatusOK, gin.H{"message": "Engineer reassigned to project", "engineer": userResponse})

}

func Delete(c *gin.Context) {
	var engineer models.User
	id, match := c.Params.Get("id")

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	dbconn := database.CreateDbConnection()

	if tx := dbconn.Preload("Project").First(&engineer, id); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Engineer not found" + tx.Error.Error()})
		}
		return
	}

	if tx := dbconn.Delete(&engineer); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error deleting engineer" + tx.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Engineer deleted"})
}
