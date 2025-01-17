package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/handlers"
	"gorm.io/gorm"
)

func DiningMenuRoutes(app *fiber.App, db *gorm.DB) {
	h := handlers.DiningMenuHandler{DB: db}
	api := app.Group("/api")
	api.Post("/createDiningMenu", h.CreateDiningMenu)
	api.Get("/diningMenus", h.GetDiningMenus)
	api.Get("/getDiningMenuById/:id", h.GetDiningMenuById)
	api.Get("/getDiningMenuPositions/:id", h.GetDiningMenuPositions)
}
