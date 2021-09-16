package handler

import (
	"github.com/gofiber/fiber/v2"
	"petcard/internal/controllers"
	"petcard/pkg/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(app *fiber.App) {
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
	ad.Post("/create", h.CreateAd)
	ad.Get("/", h.GetAllAds)
	ad.Get("/:id", h.GetListByTitle)
	ad.Delete("/:id", h.Delete)
	ad.Put("/:id", h.Update)

	// specify endpoints
	spec.Post("/create", controllers.CreateSpecify)
	spec.Get("/", controllers.SpecifyGetAll)
	spec.Get("/:id", controllers.SpecifyGetById)
	spec.Delete("/:id", controllers.SpecifyDeleteById)
	spec.Put("/:id", controllers.SpecifyUpdate)
}
