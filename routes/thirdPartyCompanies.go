package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/handlers"
	"gorm.io/gorm"
)

func ThirdPartyCompanyRoutes(app *fiber.App, db *gorm.DB) {
	h := handlers.ThirdPartyCompanyHandler{DB: db}
	api := app.Group("/api")
	api.Post("/createThirdPartyCompany", h.CreateThirdPartyCompany)
	api.Get("/thirdPartyCompanies", h.GetThirdPartyCompanies)
	api.Get("/getThirdPartyCompanyById/:id", h.GetThirdPartyCompanyById)
	api.Get("/getThirdPartyCompanyOfferPositionsById/:id", h.GetThirdPartyCompanyOfferPositionsById)
}
