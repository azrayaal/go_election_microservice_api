package candidatecontrollers

import (
	"net/http"

	"example.com/web-service-gin/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var candidate []models.Candidate
	models.DB.Find(&candidate)

	c.JSON(http.StatusOK, gin.H{"candidate": candidate})
}
