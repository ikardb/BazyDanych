package models

type OrderPositions struct {
	Id_pozycji_zamowienia int16 `gorm:"primaryKey;autoIncrement" json:"id_pozycji_zamowienia"`
	Id_zamowienia         int16 `json:"id_zamowienia"`
	Id_pozycji_jadlospisu int16 `json:"id_pozycji_jadlospisu"`
	Ilosc_produktu        int16 `json:"ilosc_produktu"`
}

type OrderPositionsWithProductName struct {
	Id_pozycji_zamowienia int16   `gorm:"primaryKey;autoIncrement" json:"id_pozycji_zamowienia"`
	Id_zamowienia         int16   `json:"id_zamowienia"`
	Id_pozycji_jadlospisu int16   `json:"id_pozycji_jadlospisu"`
	Ilosc_produktu        int16   `json:"ilosc_produktu"`
	Nazwa                 string  `json:"nazwa"`
	Cena                  float64 `json:"cena"`
}

func (OrderPositions) TableName() string {
	return "pozycja_zamowienia"
}
