package routes

import (
	"github.com/gofiber/fiber/v2"
	"petcard/controllers"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)

	api := app.Group("/api")

	user := api.Group("/user")
	ad := api.Group("/ad")
	spec := api.Group("/spec")

	// USER READ
	user.Get("/get", controllers.UserGetAll)
	user.Get("/get/username=:username", controllers.UserGetByUsername)
	// USER DELETE
	user.Delete("/username=:username", controllers.UserDeleteByUsername)
	// USER UPDATE/PUT
	user.Put("/name/username=:username", controllers.UserUpdateByName)
	user.Put("/lastname/username=:username", controllers.UserUpdateByLastname)
	user.Put("/username/username=:username", controllers.UserUpdateByUsername)
	user.Put("/email/username=:username", controllers.UserUpdateByEmail)
	user.Put("/ad/username=:username", controllers.UserUpdateByAd)

	ad.Post("/create", controllers.CreateAd)

	spec.Post("/create", controllers.CreateSpecify)

}
