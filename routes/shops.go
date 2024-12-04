package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/handlers"
	"gorm.io/gorm"
)

func ShopRoutes(app *fiber.App, db *gorm.DB) {
	h := handlers.ShopHandler{DB: db}
	api := app.Group("/api")
	api.Post("/createShop", h.CreateShop)
	api.Get("/shops", h.GetShops)
	api.Get("/getShopById/:id", h.GetShopById)
}
