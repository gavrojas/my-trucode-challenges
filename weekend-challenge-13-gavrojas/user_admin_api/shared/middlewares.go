package shared

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"trucode.app/user_admin_api/database"
	"trucode.app/user_admin_api/models"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

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
		tokenStr := GetTokenFromRequest(c)
		token, err := jwt.ParseWithClaims(tokenStr, &Payload{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, _ := token.Claims.(*Payload)
		fmt.Println(claims)

		userData := Sessions[claims.Session]

		if userData.ExpiryTime.Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session expired"})
			return
		}

		var user models.User
		dbConn := database.CreateDbConnection()
		tx := dbConn.Where("id = ?", userData.Uid).First(&user)
		if tx.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": tx.Error.Error()})
			return
		}

		c.Set("authorizedUser", user)

		c.Next()
	}
}