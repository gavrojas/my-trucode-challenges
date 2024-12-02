package projects

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"trucode.app/apitaskmanager/database"
	"trucode.app/apitaskmanager/models"
)

func Create(c *gin.Context) {
	log.Println("Iniciando la creación del proyecto")
	var project models.Project
	c.BindJSON(&project)

	dbconn := database.CreateDbConnection()
	log.Println("Conexión a la base de datos establecida")

	if tx := dbconn.Create(&project); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create a project" + tx.Error.Error()})
		return
	}

	log.Println("Proyecto creado:", project.ID)

	c.JSON(http.StatusCreated, gin.H{"message": "Project created successfully", "project": project})
}

func ReadAll(c *gin.Context) {
	var projects []map[string]interface{}
	dbconn := database.CreateDbConnection()

	dbconn.Model(models.Project{}).Select("id", "name").Find(&projects)

	if len(projects) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No projects found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Projects found", "projects": projects})
}

func ReadByID(c *gin.Context) {
	var project models.Project
	id, match := c.Params.Get("id")

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}
	dbconn := database.CreateDbConnection()
	// Precargar tareas e ingenieros -|
	if tx := dbconn.Preload("Tasks").Preload("Engineers").First(&project, id); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found" + tx.Error.Error()})
		}
		return
	}

	projectTasks := make([]map[string]interface{}, 0)
	for _, task := range project.Tasks {
		projectTasks = append(projectTasks, map[string]interface{}{
			"id":        task.ID,
			"name":      task.Name,
			"estimate":  task.EstimateHours,
			"projectId": task.ProjectID,
		})
	}

	projectEnginers := make([]map[string]interface{}, 0)
	for _, engineer := range project.Engineers {
		projectEnginers = append(projectEnginers, map[string]interface{}{
			"id":        engineer.ID,
			"name":      engineer.Name,
			"surname":   engineer.Surname,
			"projectId": engineer.ProjectID,
		})
	}

	projectResponse := map[string]interface{}{
		"id":        project.ID,
		"name":      project.Name,
		"delivery":  project.DeliveryDate,
		"engineers": projectEnginers,
		"tasks":     projectTasks,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project found", "project": projectResponse})
}

func Delete(c *gin.Context) {
	var project models.Project
	id, match := c.Params.Get("id")

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	dbconn := database.CreateDbConnection()

	if tx := dbconn.Preload("Tasks").Preload("Engineers").First(&project, id); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found" + tx.Error.Error()})
		}
		return
	}

	if tx := dbconn.Select(clause.Associations).Delete(&project); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error deleting project" + tx.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}

func EditByID(c *gin.Context) {
	var project models.Project
	id, match := c.Params.Get("id")

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	c.BindJSON(&project)

	dbconn := database.CreateDbConnection()

	if tx := dbconn.First(&project, id); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found" + tx.Error.Error()})
		}
		return
	}

	if tx := dbconn.Save(&project); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to edit project" + tx.Error.Error()})
		return
	}

	// Precargar tareas e ingenieros -|
	if tx := dbconn.Preload("Tasks").Preload("Engineers").First(&project, id); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found" + tx.Error.Error()})
		}
		return
	}

	projectTasks := make([]map[string]interface{}, 0)
	for _, task := range project.Tasks {
		projectTasks = append(projectTasks, map[string]interface{}{
			"id":        task.ID,
			"name":      task.Name,
			"estimate":  task.EstimateHours,
			"projectId": task.ProjectID,
		})
	}

	projectEnginers := make([]map[string]interface{}, 0)
	for _, engineer := range project.Engineers {
		projectEnginers = append(projectEnginers, map[string]interface{}{
			"id":        engineer.ID,
			"name":      engineer.Name,
			"surname":   engineer.Surname,
			"projectId": engineer.ProjectID,
		})
	}

	projectResponse := map[string]interface{}{
		"id":        project.ID,
		"name":      project.Name,
		"delivery":  project.DeliveryDate,
		"engineers": projectEnginers,
		"tasks":     projectTasks,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully", "project": projectResponse})
}
