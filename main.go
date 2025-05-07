package main

import (
	"belajar-gin/config"
	"belajar-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	routes.UserRoutes(r)
	r.Run(":8000")
}