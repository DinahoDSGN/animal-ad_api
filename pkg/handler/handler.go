package handler

import (
	"github.com/gofiber/fiber/v2"
	"petcard/pkg/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	app.Post("/api/signup", h.SignUp)
	app.Post("/api/signin", h.SignIn)

	api := app.Group("/api")

	user := api.Group("/user")
	ad := api.Group("/ad")
	spec := api.Group("/animal")
	parser := api.Group("/parser")
	breed := api.Group("/breed")

	// user endpoints
	user.Get("/", h.GetAllUsers)
	user.Get("/:id", h.GetUserById)
	user.Delete("/:id", h.DeleteUser)
	user.Put("/:id", h.UpdateUser)

	// ad endpoints
	ad.Post("/create", h.CreateAd)
	ad.Get("/", h.GetAllAds)
	ad.Get("/:id", h.GetAdById)
	ad.Delete("/:id", h.DeleteAd)
	ad.Put("/:id", h.UpdateAd)

	// animal endpoints
	spec.Post("/create", h.CreateAnimal)
	spec.Get("/", h.GetAllAnimals)
	spec.Get("/:id", h.GetAnimalById)
	spec.Delete("/:id", h.DeleteAnimal)
	spec.Put("/:id", h.UpdateAnimal)

	// breed endpoints
	breed.Post("/create", h.CreateBreed)
	breed.Get("/", h.GetAllBreeds)
	breed.Get("/:id", h.GetBreedById)
	breed.Delete("/:id", h.DeleteBreed)
	breed.Put("/:id", h.UpdateBreed)

	ad.Get("/sort/by=:by", h.SortBy)
	ad.Get("/sort/by=:by", h.SortBy)

	parser.Post("/push", h.Push)
}
