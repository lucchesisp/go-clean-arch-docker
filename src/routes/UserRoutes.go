package routes

import "github.com/gin-gonic/gin"

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "users route ok",
		})
	})
}
