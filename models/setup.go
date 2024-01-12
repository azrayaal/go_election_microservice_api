package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := "root:@tcp(localhost:3307)/golang_election_microservice"
	database, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	// database.AutoMigrate(&Article{})

	DB = database
}
