package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/handlers"
	"gorm.io/gorm"
)

func SaleRoutes(app *fiber.App, db *gorm.DB) {
	h := handlers.SaleHandler{DB: db}
	api := app.Group("/api")
	api.Post("/createSale", h.CreateSale)
	api.Get("/sales", h.GetSales)
	api.Get("/getSaleById/:id", h.GetSaleById)
	api.Get("/getSalePositions/:id", h.GetSalePositions)
	api.Post("/sumSalesFromGivenTime", h.SumSalesFromGivenTimeStoreUser)
}
