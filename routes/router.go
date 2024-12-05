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
	DiningMenuRoutes(app, db)
	DiningMenuPositionRoutes(app, db)
	ProductRoutes(app, db)
	OrderPositionRoutes(app, db)
	StockLevelRoutes(app, db)
	SaleRoutes(app, db)
	SalePositionRoutes(app, db)
	ThirdPartyCompanyRoutes(app, db)
	ThirdPartyCompanyPositionRoutes(app, db)
}
