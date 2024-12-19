package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/handlers"
	"gorm.io/gorm"
)

func OrderRoutes(app *fiber.App, db *gorm.DB) {
	h := handlers.OrderHandler{DB: db}
	api := app.Group("/api")
	api.Post("/createOrder", h.CreateOrder)
	api.Get("/orders", h.GetOrders)
	api.Get("/getOrderById/:id", h.GetOrderById)
	api.Get("/getOrderPositions/:id", h.GetOrderPositions)
	api.Post("/migrateToStock/:id", h.MigrateToStock)
}
