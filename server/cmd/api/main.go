package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"apiGo/internal/config"
	"apiGo/internal/database"
	"apiGo/internal/routes"
)

func main() {
	config.LoadEnv()

	database.Setup()

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "release"
	}

	gin.SetMode(ginMode)

	r := gin.Default()

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error al ejecutar el servidor:", err)
	}
}
