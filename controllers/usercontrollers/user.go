package usercontrollers

import (
	"net/http"

	"example.com/web-service-gin/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var user []models.User
	models.DB.Find(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}
