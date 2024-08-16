package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lazbord/SpotyGo/services/streaming/service"
)

type ApiAdapter struct {
	service *service.StreamingService
}

func NewApiAdapter(authService *service.StreamingService) *ApiAdapter {
	return &ApiAdapter{
		service: authService,
	}
}

func (api ApiAdapter) NewAPI() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/streaming/addmusic", api.AddMusic)
	router.GET("/streaming/getmusic", api.GetMusicByID)

	router.Run("localhost:8080")
}
