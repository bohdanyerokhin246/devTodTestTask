package main

import (
	"devTodTestTask/internal/config"
	"devTodTestTask/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.ConnectDB()
	r := gin.Default()
	routes.SetupRoutes(r, db)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
