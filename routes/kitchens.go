package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/handlers"
	"gorm.io/gorm"
)

func KitchenRoutes(app *fiber.App, db *gorm.DB) {
	h := handlers.KitchenHandler{DB: db}
	api := app.Group("/api")
	api.Post("/createKitchen", h.CreateKitchen)
	api.Get("/Kitchens", h.GetKitchens)
	api.Get("/getKitchenById/:id", h.GetKitchenById)
}
