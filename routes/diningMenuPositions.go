package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/handlers"
	"gorm.io/gorm"
)

func DiningMenuPositionRoutes(app *fiber.App, db *gorm.DB) {
	h := handlers.DiningMenuPositionHandler{DB: db}
	api := app.Group("/api")
	api.Post("/createDiningMenuPosition", h.CreateDiningMenuPosition)
	api.Get("/diningMenuPositions", h.GetDiningMenuPositions)
}
