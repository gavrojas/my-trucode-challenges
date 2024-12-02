package tasks

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
	var task models.Task
	id, match := c.Params.Get("id")
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get project id"})
		return
	}

	c.BindJSON(&task)

	dbconn := database.CreateDbConnection()
	pid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse project id"})
		return
	}

	task.ProjectID = uint(pid)

	if tx := dbconn.Create(&task); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create a task" + tx.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func ReadAll(c *gin.Context) {
	var tasks []map[string]interface{}
	dbconn := database.CreateDbConnection()

	dbconn.Model(models.Task{}).Select("id", "name").Find(&tasks)

	if len(tasks) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No tasks found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tasks found", "tasks": tasks})
}

func ReadByID(c *gin.Context) {
	var task models.Task
	id, match := c.Params.Get("id")

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	dbconn := database.CreateDbConnection()
	// Precargar proyecto -|
	if tx := dbconn.Preload("Project").First(&task, id); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found" + tx.Error.Error()})
		}
		return
	}

	taskResponse := map[string]interface{}{
		"id":       task.ID,
		"name":     task.Name,
		"estimate": task.EstimateHours,
		"project": map[string]interface{}{
			"id":       task.Project.ID,
			"name":     task.Project.Name,
			"delivery": task.Project.DeliveryDate,
		},
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task found", "task": taskResponse})
}

func ReasignProject(c *gin.Context) {
	var task models.Task
	idProject, match := c.Params.Get("idProject")

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	idTask, match := c.Params.Get("idTask")
	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get task id"})
		return
	}

	dbconn := database.CreateDbConnection()

	if tx := dbconn.First(&task, idTask); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found" + tx.Error.Error()})
		}
		return
	}

	pid, err := strconv.Atoi(idProject)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse project id"})
		return
	}

	task.ProjectID = uint(pid)

	if tx := dbconn.Save(&task); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to reassign project" + tx.Error.Error()})
		return
	}

	if tx := dbconn.Preload("Project").First(&task, idTask); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found" + tx.Error.Error()})
		}
		return
	}

	taskResponse := map[string]interface{}{
		"id":       task.ID,
		"name":     task.Name,
		"estimate": task.EstimateHours,
		"project": map[string]interface{}{
			"id":       task.Project.ID,
			"name":     task.Project.Name,
			"delivery": task.Project.DeliveryDate,
		},
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task edited successfully", "task": taskResponse})
}

func Delete(c *gin.Context) {
	var task models.Task
	id, match := c.Params.Get("id")

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	dbconn := database.CreateDbConnection()

	if tx := dbconn.Preload("Project").First(&task, id); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found" + tx.Error.Error()})
		}
		return
	}

	if tx := dbconn.Delete(&task); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error deleting task" + tx.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
