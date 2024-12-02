package entries

import (
	"github.com/gin-gonic/gin"
	"trucode.app/user_admin_api/shared"
)

func AddEntryRoutes(r *gin.Engine) {
	group := r.Group("/entries")

	group.Use(shared.AuthenticateSession())

	group.GET("", GetAllEntries)
	group.GET("/:id", GetEntryById)
	group.POST("", CreateEntry)
	group.DELETE("/:id", DeleteEntry)
	group.PUT("/:id", UpdateEntry)
}
