package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/handlers"
	"gorm.io/gorm"
)

func SalePositionRoutes(app *fiber.App, db *gorm.DB) {
	h := handlers.SalePositionHandler{DB: db}
	api := app.Group("/api")
	api.Post("/createSalePosition", h.CreateSalePosition)
	api.Get("/salePositions", h.GetSalePositions)
}
