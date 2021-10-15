package main

import (
	"github.com/gin-gonic/gin"
	"petcard/pkg/database"
	"petcard/pkg/handler"
	"petcard/pkg/repository"
	"petcard/pkg/services"
)

func main() {
	app := gin.Default()
	defer app.Run()
	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})

	repos := repository.NewRepository(database.Connect())
	services := services.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes(app)

	//telegramBot := telegram.NewTelegram(database.Connect(), services)
	//telegramBot.InitBot()

}
