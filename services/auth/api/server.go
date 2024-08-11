package api

import (
	"github.com/gin-gonic/gin"
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

	router.POST("/login", api.Login)
	router.POST("/create/user", api.Test)
	// router.POST("/login/2fa", DoublefactorLogin)

	router.Run("localhost:8080")
}
