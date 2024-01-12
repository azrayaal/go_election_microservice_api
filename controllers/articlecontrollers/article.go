package articlecontrollers

import (
	"net/http"

	"example.com/web-service-gin/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var article []models.Article
	models.DB.Find(&article)

	c.JSON(http.StatusOK, gin.H{"article": article})
}

// func ShowDetail(c *gin.Context) {
// 	var article models.Article
// 	id := c.Param("id")

// 	if err := models.DB.First(&article, id).Error; err != nil {
// 		switch err {
// 		case gorm.ErrRecordNotFound:
// 			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
// 			return
// 		default:
// 			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{"article": article})
// }
