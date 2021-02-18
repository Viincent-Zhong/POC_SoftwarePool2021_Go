package main

import (
	"SoftwareGoDay4/routes"
	"SoftwareGoDay4/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	user.Index = 0
	routes.ApplyRoute(router)
	router.Run(":8080")
}
