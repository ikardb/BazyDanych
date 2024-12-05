package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/handlers"
	"gorm.io/gorm"
)

func ProductRoutes(app *fiber.App, db *gorm.DB) {
	h := handlers.ProductHandler{DB: db}
	api := app.Group("/api")
	api.Post("/createProduct", h.CreateProduct)
	api.Get("/products", h.GetProducts)
	api.Get("/getProductById/:id", h.GetProductById)
}
