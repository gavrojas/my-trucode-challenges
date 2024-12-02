package entries

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"trucode.app/user_admin_api/database"
	"trucode.app/user_admin_api/models"
)

func GetAllEntries(c *gin.Context) {
	sessionUser, _ := c.Get("authorizedUser")
	userID := uint(sessionUser.(models.User).ID)
	var entries []models.Entry

	dbConn := database.CreateDbConnection()
	tx := dbConn.Where("user_id = ?", userID).Find(&entries)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting your entries"})
		return
	}

	if len(entries) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No entries found"})
		return
	}

	c.JSON(http.StatusOK, entries)
}

func GetEntryById(c *gin.Context) {
	id := c.Param("id")
	sessionUser, _ := c.Get("authorizedUser")
	userID := uint(sessionUser.(models.User).ID)

	var entry models.Entry

	dbConn := database.CreateDbConnection()
	tx := dbConn.Where("id = ? AND user_id = ?", id, userID).First(&entry)
	if tx.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}

	c.JSON(http.StatusOK, entry)
}

func CreateEntry(c *gin.Context) {
	var entry models.Entry
	c.BindJSON(&entry)

	sessionUser, _ := c.Get("authorizedUser")
	entry.UserID = uint(sessionUser.(models.User).ID)

	dbConn := database.CreateDbConnection()
	tx := dbConn.Create(&entry)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating entry"})
		return
	}
	c.JSON(http.StatusCreated, entry)
}

func DeleteEntry(c *gin.Context) {
	id := c.Param("id")
	sessionUser, _ := c.Get("authorizedUser")
	userID := uint(sessionUser.(models.User).ID)

	var entry models.Entry

	dbConn := database.CreateDbConnection()
	tx := dbConn.Where("id = ? AND user_id = ?", id, userID).First(&entry)
	if tx.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}
	tx = dbConn.Delete(&entry)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting entry"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Entry deleted successfully"})
}

func UpdateEntry(c *gin.Context) {
	id := c.Param("id")
	sessionUser, _ := c.Get("authorizedUser")
	userID := uint(sessionUser.(models.User).ID)

	var entry models.Entry

	dbConn := database.CreateDbConnection()
	tx := dbConn.Where("id = ? AND user_id = ?", id, userID).First(&entry)
	if tx.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entry not found"})
		return
	}

	c.BindJSON(&entry)
	entry.UserID = userID
	dbConn.Updates(&entry)
	c.JSON(http.StatusOK, entry)
}
