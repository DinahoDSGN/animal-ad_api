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
	//app.Use(cors.Default())
	defer app.Run()
	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	//app.Use(cors.New(cors.Config{
	//	AllowOriginFunc:  func(origin string) bool { return true },
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	//	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
	//	AllowCredentials: true,
	//	MaxAge:           12 * time.Hour,
	//}))

	repos := repository.NewRepository(database.Connect())
	services := services.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes(app)

	//telegramBot := telegram.NewTelegram(database.Connect(), services)
	//telegramBot.InitBot()

}
