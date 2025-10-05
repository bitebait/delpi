package routes

import (
	"github.com/gin-gonic/gin"

	"apiGo/internal/config"
	"apiGo/internal/database"
	"apiGo/internal/handlers"
	"apiGo/internal/services"
)

type StateRoute struct {
	stateHandler *handlers.StateHandler
}

func NewStateRoute() *StateRoute {
	stateService := services.NewStateService(database.DB)
	stateHandler := handlers.NewStateHandler(stateService)

	return &StateRoute{stateHandler}
}

func (r *StateRoute) SetupRoutes(g *gin.RouterGroup) {
	g.GET("/departamentos", r.stateHandler.GetAll)

	if config.IsDevMode() {
		g.POST("/departamentos", r.stateHandler.Create)
	}

}
