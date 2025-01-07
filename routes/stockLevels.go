package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/handlers"
	"gorm.io/gorm"
)

func StockLevelRoutes(app *fiber.App, db *gorm.DB) {
	h := handlers.StockLevelHandler{DB: db}
	api := app.Group("/api")
	api.Post("/createStockLevel", h.CreateStockLevel)
	api.Get("/stockLevels", h.GetStockLevels)
	api.Get("/getStockLevelById/:id", h.GetStockLevelById)
	api.Get("/getStockLevelByShopID/:id", h.GetStockLevelByShopId)
}
