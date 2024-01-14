package main

import (
	"example.com/web-service-gin/controllers/articlecontrollers"
	"example.com/web-service-gin/controllers/candidatecontrollers"
	"example.com/web-service-gin/controllers/partycontrollers"
	"example.com/web-service-gin/controllers/usercontrollers"
	"example.com/web-service-gin/controllers/votercontrollers"
	"example.com/web-service-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()
	router := gin.Default()

	// ARTICLE
	router.GET("/api/v1/articles", articlecontrollers.Index)
	// PARTY
	router.GET("/api/v1/parties", partycontrollers.Index)
	// CANDIDATE
	router.GET("/api/v1/candidates", candidatecontrollers.Index)
	// VOTER
	router.GET("/api/v1/voters", votercontrollers.Index)
	// USER
	router.GET("/api/v1/users", usercontrollers.Index)
	router.Run()
	// s
}
