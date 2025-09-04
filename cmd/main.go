package main

import (
	"github.com/gin-gonic/gin"
	"subagg/internal/routes"
)

func main() {
	router := gin.Default()

	routes.RegisterRoutes(router)
	
	router.Run(":8080")
}
