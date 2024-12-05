package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/handlers"
	"gorm.io/gorm"
)

func ThirdPartyCompanyPositionRoutes(app *fiber.App, db *gorm.DB) {
	h := handlers.ThirdPartyCompanyPositionHandler{DB: db}
	api := app.Group("/api")
	api.Post("/createThirdPartyCompanyPosition", h.CreateThirdPartyCompanyPosition)
	api.Get("/thirdPartyCompanyPositions", h.GetThirdPartyCompanyPositions)
}
