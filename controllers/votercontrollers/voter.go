package votercontrollers

import (
	"net/http"

	"example.com/web-service-gin/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var voter []models.Voter
	models.DB.Preload("User").Preload("Candidate").Find(&voter)

	c.JSON(http.StatusOK, gin.H{"voter": voter})
}
