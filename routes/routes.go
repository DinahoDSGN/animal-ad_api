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
	user.Delete("delete/username=:username", controllers.UserDeleteByUsername)
	// USER UPDATE/PUT
	user.Put("update/name/username=:username", controllers.UserUpdateByName)
	user.Put("update/lastname/username=:username", controllers.UserUpdateByLastname)
	user.Put("update/username/username=:username", controllers.UserUpdateByUsername)
	user.Put("update/email/username=:username", controllers.UserUpdateByEmail)
	user.Put("update/ad/username=:username", controllers.UserUpdateByAd)

	// AD CREATE
	ad.Post("/create", controllers.CreateAd)
	// AD READ
	ad.Get("/get", controllers.AdGetAll)
	ad.Get("/get/title=:title", controllers.AdGetByTitle)
	// USER DELETE
	ad.Delete("delete/title=:title", controllers.AdDeleteByTitle)

	spec.Post("/create", controllers.CreateSpecify)

}
