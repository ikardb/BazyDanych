package handlers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ikardb/BazyDanych/models"
)

func TestGetOrders_WithMock(t *testing.T) {
	// stworzenie mocka bazy
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	// podlaczenie mocka do GORMa
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}), &gorm.Config{})
	assert.NoError(t, err)

	handler := &OrderHandler{DB: gormDB}

	app := fiber.New()
	app.Get("/orders", handler.GetOrders)

	// przykladowe dane (maja zostac zwrocone przez mock)
	var idUser1 int16 = 2
	var idUser2 int16 = 1
	mockOrders := []models.Orders{
		{Id_zamowienia: 1, Id_sklepu: 1, Id_uzytkownika: &idUser1, Data_zamowienia: time.Now(), Koszt_zamowienia: 10.5, Wczytane_do_stanu: false},
		{Id_zamowienia: 2, Id_sklepu: 2, Id_uzytkownika: &idUser2, Data_zamowienia: time.Now().Add(time.Hour), Koszt_zamowienia: 25.0, Wczytane_do_stanu: true},
	}

	rows := sqlmock.NewRows([]string{"id_zamowienia", "id_sklepu", "id_uzytkownika", "data_zamowienia", "koszt_zamowienia", "wczytane_do_stanu"}).
		AddRow(mockOrders[0].Id_zamowienia, mockOrders[0].Id_sklepu, mockOrders[0].Id_uzytkownika, mockOrders[0].Data_zamowienia, mockOrders[0].Koszt_zamowienia, mockOrders[0].Wczytane_do_stanu).
		AddRow(mockOrders[1].Id_zamowienia, mockOrders[1].Id_sklepu, mockOrders[1].Id_uzytkownika, mockOrders[1].Data_zamowienia, mockOrders[1].Koszt_zamowienia, mockOrders[1].Wczytane_do_stanu)

	mock.ExpectQuery(`SELECT \* FROM "zamowienie"`).WillReturnRows(rows)

	// stworzenie zadania do aplikacji fiber
	req := httptest.NewRequest("GET", "/orders", nil)
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)

	// weryfikacja odpowiedzi
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// dekodowanie odpowiedzi (json)
	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	assert.NoError(t, err)

	assert.Equal(t, "orders loaded successfully", responseBody["message"])

	// sprawdzenie danych z odpowiedzi z danymi oczekiwanymi
	var orders []models.Orders
	data, _ := json.Marshal(responseBody["data"])
	json.Unmarshal(data, &orders)

	assert.Equal(t, len(mockOrders), len(orders))
	for i := range orders {
		assert.Equal(t, mockOrders[i].Id_zamowienia, orders[i].Id_zamowienia)
		assert.Equal(t, mockOrders[i].Id_sklepu, orders[i].Id_sklepu)
		assert.Equal(t, mockOrders[i].Id_uzytkownika, orders[i].Id_uzytkownika)
		assert.True(t, mockOrders[i].Data_zamowienia.Round(time.Second).Equal(orders[i].Data_zamowienia.Round(time.Second)))
		assert.Equal(t, mockOrders[i].Koszt_zamowienia, orders[i].Koszt_zamowienia)
		assert.Equal(t, mockOrders[i].Wczytane_do_stanu, orders[i].Wczytane_do_stanu)
	}

	assert.NoError(t, mock.ExpectationsWereMet())
}
