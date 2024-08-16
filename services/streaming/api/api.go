package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MusicRequest struct {
	Id string `json:"id"`
}

func (api ApiAdapter) GetMusicByID(c *gin.Context) {
	var req MusicRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	music, err := api.service.ServiceGetMusicByID(req.Id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"music name": music.Name})
}
