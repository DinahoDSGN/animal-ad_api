package handler

import (
	"github.com/gin-gonic/gin"
	"petcard/pkg/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(r *gin.Engine) *gin.Engine {
	router := r

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/signin", h.SignIn)
	}

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.GET("/me", h.GetUser)
			user.GET("/", h.GetAllUsers)
			user.GET("/:id", h.GetUserById)
			user.DELETE("/:id", h.DeleteUser)
			user.PUT("/:id", h.UpdateUser)
		}
		ad := api.Group("/adv")
		{
			ad.POST("/create", h.CreateAd)
			ad.GET("/", h.GetAllAds)
			ad.GET("/:id", h.GetAdById)
			ad.DELETE("/:id", h.DeleteAd)
			ad.PUT("/:id", h.UpdateAd)
		}
		spec := api.Group("/animal")
		{
			spec.POST("/create", h.CreateAnimal)
			spec.GET("/", h.GetAllAnimals)
			spec.GET("/:id", h.GetAnimalById)
			spec.DELETE("/:id", h.DeleteAnimal)
			spec.PUT("/:id", h.UpdateAnimal)
		}
		parser := api.Group("/parser")
		{
			parser.POST("/push", h.Push)
		}
		breed := api.Group("/breed")
		{
			breed.POST("/create", h.CreateBreed)
			breed.GET("/", h.GetAllBreeds)
			breed.GET("/:id", h.GetBreedById)
			breed.DELETE("/:id", h.DeleteBreed)
			breed.PUT("/:id", h.UpdateBreed)
		}
	}

	return router
}
