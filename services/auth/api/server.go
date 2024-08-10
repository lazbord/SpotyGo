package api

import (
	"github.com/gin-gonic/gin"
)

func NewAPI() {
	router := gin.Default()

	router.POST("/login", Login)
	//router.POST("/login/2fa", DoublefactorLogin)

	router.Run("localhost:8080")
}
