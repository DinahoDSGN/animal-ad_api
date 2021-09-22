package main

import (
	"github.com/gin-gonic/gin"
	"petcard/pkg/database"
	"petcard/pkg/handler"
	"petcard/pkg/repository"
	"petcard/pkg/services"
	"petcard/telegram"
)

func main() {
	app := gin.Default()

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	repos := repository.NewRepository(database.Connect())
	services := services.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes(app)

	telegramBot := telegram.NewTelegram(database.Connect(), services)
	telegramBot.InitBot()

	go app.Run(":8080")

}
