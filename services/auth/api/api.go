package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "it works")
}

func NewAPI() {
	router := gin.Default()
	router.GET("/test", test)

	router.Run("localhost:8080")
}
