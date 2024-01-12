package partycontrollers

import (
	"net/http"

	"example.com/web-service-gin/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var party []models.Party
	models.DB.Find(&party)

	c.JSON(http.StatusOK, gin.H{"party": party})
}

