package handlers

import (
	"math"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ikardb/BazyDanych/models"
	"gorm.io/gorm"
)

type SaleHandler struct {
	DB *gorm.DB
}

type SumSalesRequest struct {
	TimeStart time.Time `json:"od_kiedy"`
	TimeEnd   time.Time `json:"do_kiedy"`
	StoreID   int16     `json:"id_sklepu"`
	UserID    int16     `json:"id_uzytkownika"`
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func (h *SaleHandler) CreateSale(context *fiber.Ctx) error {
	sale := models.Sales{}
	sale.Data_sprzedazy = time.Now()

	err := context.BodyParser(&sale)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request failed"})
		return err
	}

	err = h.DB.Create(&sale).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create sale"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "sale has been added"})
	return nil
}

func (h *SaleHandler) GetSales(context *fiber.Ctx) error {
	sales := &[]models.Sales{}

	err := h.DB.Find(sales).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get sales"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "sales loaded successfully", "data": sales})
	return nil
}

func (h *SaleHandler) GetSaleById(context *fiber.Ctx) error {
	sale := &models.Sales{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := h.DB.Where("id_sprzedazy = ?", id).First(sale).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "sale not found"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "sale ID fetched successfully", "data": sale})
	return nil
}

func (h *SaleHandler) GetSalePositions(context *fiber.Ctx) error {
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	salePositions := []models.SalePositionsWithProductName{}
	query := `
		SELECT 
			p.id_pozycji,
			p.id_sprzedazy,
			p.id_produktu,
			p.ilosc,
			p.cena_jednostkowa,
			produkt.nazwa
		FROM pozycja_sprzedazy p
		JOIN sprzedaz s ON s.id_sprzedazy = p.id_sprzedazy
		JOIN produkt ON produkt.id_produktu = p.id_produktu
		WHERE s.id_sprzedazy = ?
	`

	err := h.DB.Raw(query, id).Scan(&salePositions).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "order positions not found"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "order positions fetched successfully", "data": salePositions})
	return nil
}

func (h *SaleHandler) SumSalesFromGivenTimeStoreUser(context *fiber.Ctx) error {
	// parametry beda mialy domyslne wartosci jesli sie parametry wyrzuci z jsona z requesta
	sumSalesRequest := SumSalesRequest{}
	err := context.BodyParser(&sumSalesRequest)

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Couldn't parse the given sumSalesRequest"})
		return err
	}

	if sumSalesRequest.UserID < 0 || sumSalesRequest.StoreID < 0 {
		return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "User ID and Store ID must be non-negative"})
	}

	// trzeba usunac przy wykonywaniu requesta zeby zadzialalo
	// domyslnie sprawdza do chwili obecnej od poczatku miesiaca
	if sumSalesRequest.TimeEnd.IsZero() {
		sumSalesRequest.TimeEnd = time.Now()
	}
	if sumSalesRequest.TimeStart.IsZero() {
		sumSalesRequest.TimeStart = time.Date(
			sumSalesRequest.TimeEnd.Year(), sumSalesRequest.TimeEnd.Month(), 1, 0, 0, 0, 0, sumSalesRequest.TimeEnd.Location())
	}

	// if the sumSalesRequest is starting after it ends - error
	if sumSalesRequest.TimeStart.After(sumSalesRequest.TimeEnd) {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "The given time window is wrong"})
		return nil
	}

	sales := []models.Sales{}

	query := h.DB.Model(&models.Sales{}).Where("data_sprzedazy BETWEEN ? AND ?", sumSalesRequest.TimeStart, sumSalesRequest.TimeEnd)
	// inne query w zaleznosci od podanych przy requescie parametrow
	// ID SKLEPU I UZYTKOWNIKA NIE MOGA BYC 0, JESLI MA SIE NA MYSLI KONKRETNY OBIEKT
	if sumSalesRequest.UserID != 0 {
		query = query.Where("id_uzytkownika = ?", sumSalesRequest.UserID)
	}

	if sumSalesRequest.StoreID != 0 {
		query = query.Where("id_sklepu = ?", sumSalesRequest.StoreID)
	}

	err = query.Find(&sales).Error

	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "not found any sales in a given sumSalesRequest"})
		return err
	}

	var salesSum float64
	for _, sale := range sales {
		salesSum += float64(sale.Kwota_transakcji)
		salesSum = roundFloat(salesSum, 2)
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "sales fetched successfully", "utarg ze sprzedazy:": salesSum, "data": sales})
	return nil
}

func (h *SaleHandler) GetSalesByShopId(context *fiber.Ctx) error {
	sales := &[]models.Sales{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := h.DB.Where("id_sklepu = ?", id).Find(sales).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{"message": "order not found"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "order ID fetched successfully", "data": sales})
	return nil
}
