package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api ApiAdapter) DownloadMusicByID(c *gin.Context) {

	err := api.service.ServiceDownloadVideo(c.Query("videoId"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Music added to db"})
}

func (api ApiAdapter) DeleteMusicByID(c *gin.Context) {

	err := api.service.ServiceDeleteMusicById(c.Query("videoId"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Music deleted"})
}
