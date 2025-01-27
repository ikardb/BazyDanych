package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateSale_WithMock(t *testing.T) {
	// stworzenie mocka bazy
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	// podlaczenie mocka do GORMa
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}), &gorm.Config{})
	assert.NoError(t, err)

	handler := &SaleHandler{DB: gormDB}

	app := fiber.New()
	app.Post("/sales", handler.CreateSale)

	// przykladowe dane wejsciowe
	saleData := map[string]interface{}{
		"id_sklepu":        1,
		"id_uzytkownika":   2,
		"kwota_transakcji": 200.00,
	}
	body, _ := json.Marshal(saleData)

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "sprzedaz"`).
		WithArgs(
			int16(1),
			int16(2),
			200.00,
			sqlmock.AnyArg(),
			// 0.0,
			// false,
		).
		WillReturnRows(sqlmock.NewRows([]string{"id_sprzedazy"}).AddRow(1))
	mock.ExpectCommit()

	req := httptest.NewRequest("POST", "/sales", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	assert.NoError(t, mock.ExpectationsWereMet())
}
