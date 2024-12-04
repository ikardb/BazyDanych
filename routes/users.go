package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/handlers"
	"gorm.io/gorm"
)

func UserRoutes(app *fiber.App, db *gorm.DB) {
	h := handlers.UserHandler{DB: db}
	api := app.Group("/api")
	api.Post("/create_user", h.CreateUser)
	api.Get("/users", h.GetUsers)
}
