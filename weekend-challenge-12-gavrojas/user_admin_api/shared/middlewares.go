package shared

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"trucode.app/user_admin_api/database"
	"trucode.app/user_admin_api/models"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func AuthenticateSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Con esto aseguramos que el usuario esté autenticado.
		sessionToken, err := c.Cookie("session")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You don't have permission to access this resource"})
			return
		}

		// Obtener la información del usuario. Validamos que haya una sesión y que no haya expirado.
		userData := Sessions[sessionToken]
		if userData.ExpiryTime.Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session expired"})
			return
		}

		var user models.User
		dbConn := database.CreateDbConnection()
		tx := dbConn.Where("id = ?", userData.Uid).First(&user)
		if tx.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Error getting user, you don't have permission to access this resource"})
			return
		}

		c.Set("authorizedUser", user)

		c.Next()
	}
}
