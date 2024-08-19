package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MusicRequest struct {
	Id string `json:"id"`
}

func (api ApiAdapter) DownloadMusicByID(c *gin.Context) {
	var req MusicRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := api.service.DownloadVideo(req.Id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

}
