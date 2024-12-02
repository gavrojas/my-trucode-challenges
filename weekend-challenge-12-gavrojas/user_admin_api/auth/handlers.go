package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/gofrs/uuid/v5"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"trucode.app/user_admin_api/database"
	"trucode.app/user_admin_api/models"
	"trucode.app/user_admin_api/shared"
)

func Register(c *gin.Context) {
	var userInput models.UserInput
	c.BindJSON(&userInput)

	user := models.User{
		Username: userInput.Username,
		Password: userInput.Password,
	}

	dbconn := database.CreateDbConnection()

	if tx := dbconn.Create(&user); tx.Error != nil {
		// Error si el user ya existe por error postgress 23505... el error contiene "duplicate key value violates unique constraint"
		if strings.Contains(tx.Error.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"username": user.Username,
		"id":       user.ID,
	})
}

func login(c *gin.Context) {
	var input models.UserInput
	var user models.User
	c.BindJSON(&input)

	dbconn := database.CreateDbConnection()
	dbconn.Where("username = ?", input.Username).Find(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	sessionToken := uuid.NewV5(uuid.UUID{}, "session").String()

	session := shared.Session{
		Uid: user.ID,
		// 24 horas
		ExpiryTime: time.Now().Add(time.Hour * 24),
	}

	shared.Sessions[sessionToken] = session

	c.SetCookie("session", sessionToken, 10*60, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"session": sessionToken,
	})

}
