package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucchesisp/go-clean-arch-docker/src/config"
	"log"
)

func Run(port string) {
	ginMode, err := config.GetEnvVariable("GIN_MODE")

	if err != nil {
		log.Fatal("GIN_MODE environment variable not set")
	}

	gin.SetMode(ginMode)

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	GetRoutes(router)

	router.Run(":" + port)
}

func GetRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	addUserRoutes(v1)
}
