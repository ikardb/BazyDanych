package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ikardb/BazyDanych/models"
)

func setupTestPostgresDB() *gorm.DB {
	// polaczenie do TESTOWEJ bazy
	dsn := "host=localhost user=postgres password=1234 dbname=Baza_testy port=5432 sslmode=disable"

	// podlaczenie do bazy
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	err = db.AutoMigrate(&models.SalePositions{}, &models.Sales{}, &models.StockLevels{}, &models.Products{}, &models.Users{}, &models.Shops{})
	if err != nil {
		panic("failed to run migrations: " + err.Error())
	}

	return db
}

// czyszczenie zmian po tescie
func cleanUpDB(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE pozycja_sprzedazy CASCADE")
	db.Exec("TRUNCATE TABLE sprzedaz CASCADE")
	db.Exec("TRUNCATE TABLE stan_magazynowy CASCADE")
	db.Exec("TRUNCATE TABLE produkt CASCADE")
	db.Exec("TRUNCATE TABLE uzytkownicy CASCADE")
	db.Exec("TRUNCATE TABLE sklep CASCADE")
}

func TestCreateSalePosition_Postgres(t *testing.T) {
	// ustawienie bazy (testowej)
	db := setupTestPostgresDB()

	defer cleanUpDB(db)

	app := fiber.New()
	handler := &SalePositionHandler{DB: db}

	// dodanie danych potrzebnych do wykonania zapytania
	ulica := "NYSSka"
	shop := models.Shops{Id_sklepu: 1, Nazwa_sklepu: "sklep pierwszy", Ulica: &ulica}
	db.Create(&shop)

	nazwisko := "Kowal"
	user := models.Users{Id_uzytkownika: 1, Imie: "Marek", Nazwisko: &nazwisko, Administrator: false, Login: "asdasd", Haslo: "xdxdxd", Id_sklepu: 1}
	db.Create(&user)

	sale := models.Sales{Id_sprzedazy: 1, Id_uzytkownika: 1, Id_sklepu: 1, Kwota_transakcji: 0}
	db.Create(&sale)

	product := models.Products{Id_produktu: 1, Cena: 10.5}
	db.Create(&product)

	stockLevels := models.StockLevels{Id_sklepu: 1, Id_produktu: 1, Ilosc: 100}
	db.Create(&stockLevels)

	app.Post("/sale-position", handler.CreateSalePosition)

	salePosition := models.SalePositions{
		Id_sprzedazy:     1,
		Id_produktu:      1,
		Ilosc:            5,
		Cena_jednostkowa: nil,
	}

	// konwersja danych na JSON
	jsonData, err := json.Marshal(salePosition)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/sale-position", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// testowanie odpowiedzi
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// weryfikacja odpowiedzi
	var response map[string]string
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "sale position has been added", response["message"])

	// weryfikowanie czy pozycja sie dodala
	var createdSalePosition models.SalePositions
	db.First(&createdSalePosition, "id_sprzedazy = ? AND id_produktu = ?", salePosition.Id_sprzedazy, salePosition.Id_produktu)
	assert.Equal(t, salePosition.Id_sprzedazy, createdSalePosition.Id_sprzedazy)
	assert.Equal(t, salePosition.Id_produktu, createdSalePosition.Id_produktu)
	assert.Equal(t, salePosition.Ilosc, createdSalePosition.Ilosc)
	assert.Equal(t, product.Cena, *createdSalePosition.Cena_jednostkowa)

	// weryfikowanie czy stan magazynowy sie zaktualizowal
	var updatedStockLevel models.StockLevels
	db.First(&updatedStockLevel, "id_sklepu = ? AND id_produktu = ?", sale.Id_sklepu, salePosition.Id_produktu)
	assert.Equal(t, stockLevels.Ilosc-salePosition.Ilosc, updatedStockLevel.Ilosc)

	// weryfikowanie czy kwota sprzedazy sie zaktualizowala
	var updatedSale models.Sales
	db.First(&updatedSale, "id_sprzedazy = ?", salePosition.Id_sprzedazy)
	expectedTotal := float64(*createdSalePosition.Cena_jednostkowa) * float64(createdSalePosition.Ilosc)
	assert.Equal(t, expectedTotal, updatedSale.Kwota_transakcji)
}
