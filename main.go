package main

import (
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	routes.CourierRoutes()
	// router.Run(":8080")
	router.Run()
}
