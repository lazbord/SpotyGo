package api

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (api ApiAdapter) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	token, auth, err := api.service.CheckCreditential(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("access_token", token, 7*24*60*60*int(time.Second), "/", os.Getenv("DOMAIN_NAME"), true, true)

	c.JSON(http.StatusOK, gin.H{"userid": auth.ID})
}

func (api ApiAdapter) Test(c *gin.Context) {
	api.service.CreateUser()
}
