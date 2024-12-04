package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	UserRoutes(app, db)
	ShopRoutes(app, db)
	OrderRoutes(app, db)
	KitchenRoutes(app, db)
}
