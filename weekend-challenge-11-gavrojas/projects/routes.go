package projects

import "github.com/gin-gonic/gin"

func AddProjectsRoutes(router *gin.Engine) {
	router.POST("/projects", Create)
	router.GET("/projects", ReadAll)
	router.GET("/projects/:id", ReadByID)
	router.PUT("/projects/:id", EditByID)
	router.DELETE("/projects/:id", Delete)
}
