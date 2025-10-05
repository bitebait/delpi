package routes

import (
	"github.com/gin-gonic/gin"

	"apiGo/internal/config"
	"apiGo/internal/database"
	"apiGo/internal/handlers"
	"apiGo/internal/services"
)

type CityRoute struct {
	cityHandler *handlers.CityHandler
}

func NewCityRoute() *CityRoute {
	cityService := services.NewCityService(database.DB)
	cityHandler := handlers.NewCityHandler(cityService)

	return &CityRoute{cityHandler}
}

func (r *CityRoute) SetupRoutes(g *gin.RouterGroup) {
	g.GET("/ciudades", r.cityHandler.GetAll)
	g.GET("/ciudades/:departamentoID", r.cityHandler.GetAllByState)

	if config.IsDevMode() {
		g.POST("/ciudades/:departamentoID", r.cityHandler.Create)
	}

}
