package routes

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"apiGo/internal/config"
)

type Routes interface {
	SetupRoutes(router *gin.RouterGroup)
}

func SetupRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5173"
		},
		MaxAge: 12 * time.Hour,
	}))

	apiGroup := r.Group("/api")

	if config.IsDevMode() {
		apiGroup.GET("/", defaultRoute)
		apiGroup.GET("/test", defaultRoute)
	}

	setup(apiGroup,
		NewStateRoute(),
		NewCityRoute(),
		NewDistrictRoute(),
	)
}

func setup(router *gin.RouterGroup, routes ...Routes) {
	for _, r := range routes {
		r.SetupRoutes(router)
	}
}

func defaultRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hi",
	})
}
