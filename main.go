package main

import (
	"example.com/web-service-gin/controllers/articlecontrollers"
	"example.com/web-service-gin/controllers/candidatecontrollers"
	"example.com/web-service-gin/controllers/partycontrollers"
	"example.com/web-service-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()
	router := gin.Default()

	// ARTICLE
	router.GET("/api/v1/articles", articlecontrollers.Index)
	// router.GET("/api/v1/article/:id", articlecontrollers.ShowDetail)
	// sa
	// PARTY
	router.GET("/api/v1/parties", partycontrollers.Index)

	// CANDIDATE
	router.GET("/api/v1/candidates", candidatecontrollers.Index)
	router.Run()

}
