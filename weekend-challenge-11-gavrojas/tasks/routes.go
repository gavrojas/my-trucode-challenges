package tasks

import "github.com/gin-gonic/gin"

func AddTasksRoutes(router *gin.Engine) {
	router.POST("/projects/:id/tasks", Create)
	router.GET("/tasks", ReadAll)
	router.GET("/tasks/:id", ReadByID)
	router.PUT("/projects/:idProject/tasks/:idTask", ReasignProject)
	router.DELETE("/tasks/:id", Delete)
}
