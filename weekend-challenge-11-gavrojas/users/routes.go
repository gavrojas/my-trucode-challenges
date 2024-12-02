package users

import "github.com/gin-gonic/gin"

func AddUsersRoutes(router *gin.Engine) {
	router.POST("/projects/:id/users", Create)
	router.GET("/users", ReadAll)
	router.GET("/users/:id", ReadByID)
	router.PUT("/projects/:idProject/users/:idUser", ReasignProject)
	router.DELETE("/users/:id", Delete)
}
