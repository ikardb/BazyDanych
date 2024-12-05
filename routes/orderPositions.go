package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/handlers"
	"gorm.io/gorm"
)

func OrderPositionRoutes(app *fiber.App, db *gorm.DB) {
	h := handlers.OrderPositionHandler{DB: db}
	api := app.Group("/api")
	api.Post("/createOrderPosition", h.CreateOrderPosition)
	api.Get("/orderPositions", h.GetOrderPositions)
}
