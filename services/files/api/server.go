package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lazbord/SpotyGo/common/middleware"
	"github.com/lazbord/SpotyGo/services/files/service"
)

type ApiAdapter struct {
	service *service.FilesService
}

func NewApiAdapter(authService *service.FilesService) *ApiAdapter {
	return &ApiAdapter{
		service: authService,
	}
}

func (api ApiAdapter) NewAPI() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://0.0.0.0:5001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/files/Download", middleware.RequireAuth, api.DownloadMusicByID)
	router.DELETE("/files/Delete", middleware.RequireAuth, api.DeleteMusicByID)

	router.Run("0.0.0.0:5001")
}
