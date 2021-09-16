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
	user.Get("/", h.GetAllUsers)
	user.Get("/:id", h.GetUserById)
	user.Delete("/:id", h.DeleteUser)
	user.Put("/:id", h.UpdateUser)

	// ad endpoints
	ad.Post("/create", h.CreateAd)
	ad.Get("/", h.GetAllAds)
	ad.Get("/:id", h.GetUserById)
	ad.Delete("/:id", h.DeleteAd)
	ad.Put("/:id", h.UpdateAd)

	// specify endpoints
	spec.Post("/create", h.CreateSpec)
	spec.Get("/", h.GetAllSpecs)
	spec.Get("/:id", h.GetSpecById)
	spec.Delete("/:id", h.DeleteSpec)
	spec.Put("/:id", h.UpdateSpec)
}
