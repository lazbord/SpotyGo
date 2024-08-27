package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lazbord/SpotyGo/common/middleware"
	"github.com/lazbord/SpotyGo/services/auth/service"
)

type ApiAdapter struct {
	service *service.AuthService
}

func NewApiAdapter(authService *service.AuthService) *ApiAdapter {
	return &ApiAdapter{
		service: authService,
	}
}

func (api ApiAdapter) NewAPI() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://0.0.0.0:5000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/login", api.Login)
	router.POST("/create/user", middleware.RequireAuth, api.Test)

	router.Run("0.0.0.0:5000")
}
