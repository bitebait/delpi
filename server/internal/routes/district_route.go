package routes

import (
	"github.com/gin-gonic/gin"

	"apiGo/internal/config"
	"apiGo/internal/database"
	"apiGo/internal/handlers"
	"apiGo/internal/services"
)

type DistrictRoute struct {
	districtHandler *handlers.DistrictHandler
}

func NewDistrictRoute() *DistrictRoute {
	districtService := services.NewDistrictService(database.DB)
	districtHandler := handlers.NewDistrictHandler(districtService)

	return &DistrictRoute{districtHandler}
}

func (r *DistrictRoute) SetupRoutes(g *gin.RouterGroup) {
	g.GET("/barrios", r.districtHandler.GetAll)
	g.GET("/barrios/:ciudadID", r.districtHandler.GetAllByCity)

	if config.IsDevMode() {
		g.POST("/barrios/:ciudadID", r.districtHandler.Create)
	}
}
