package routes

import (
	"github.com/gofiber/fiber/v2"
	"petcard/controllers"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)

	// USER READ
	app.Get("api/user/get", controllers.UserGetAll)
	app.Get("api/user/get/username=:username", controllers.UserGetByUsername)
	// USER DELETE
	app.Delete("api/user/username=:username", controllers.UserDeleteByUsername)
	// USER UPDATE/PUT
	app.Put("/api/user/name/username=:username", controllers.UserUpdateByName)
	app.Put("/api/user/lastname/username=:username", controllers.UserUpdateByLastname)
	app.Put("/api/user/username/username=:username", controllers.UserUpdateByUsername)
	app.Put("/api/user/email/username=:username", controllers.UserUpdateByEmail)
	app.Put("/api/user/ad/username=:username", controllers.UserUpdateByAd)

	app.Post("/api/ad/create", controllers.CreateAd)

	app.Post("/api/spec/create", controllers.CreateSpecify)

}