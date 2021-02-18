package main

import (
	"SoftwareGoDay4/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.ApplyRoute(router)
	router.Run(":8080")
}
