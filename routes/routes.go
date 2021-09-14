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

	// user endpoints
	user.Get("/", controllers.UserGetAll)
	user.Get("/:username", controllers.UserGetByUsername)
	user.Delete("/:username", controllers.UserDeleteByUsername)
	user.Put("/:username", controllers.UserUpdate)

	// ad endpoints
	ad.Post("/create", controllers.CreateAd)
	ad.Get("/", controllers.AdGetAll)
	ad.Get("/:title", controllers.AdGetByTitle)
	ad.Delete("/:title", controllers.AdDeleteByTitle)
	ad.Put("/:title", controllers.AdUpdate)

	// specify endpoints
	spec.Post("/create", controllers.CreateSpecify)
	spec.Get("/", controllers.SpecifyGetAll)
	spec.Get("/:id", controllers.SpecifyGetById)
	spec.Delete("/:id", controllers.SpecifyDeleteById)
	spec.Put("/:id", controllers.SpecifyUpdate)

}
