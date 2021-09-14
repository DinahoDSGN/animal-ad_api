package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"petcard/models"
)

var DB *gorm.DB

func Connect(){
	connection, err := gorm.Open(mysql.Open("root:root@/petcard"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil{
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.Ad{})
	connection.AutoMigrate(&models.Specify{})
	connection.AutoMigrate(&models.User{})
}
