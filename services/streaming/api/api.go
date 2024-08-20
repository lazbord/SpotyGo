package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api ApiAdapter) StreamMusicByID(c *gin.Context) {

	music, err := api.service.ServiceGetMusicByID(c.Query("videoId"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "audio/mpeg")
	c.Writer.Write(music.Data)
}
