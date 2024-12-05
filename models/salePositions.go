package models

type SalePositions struct {
	Id_pozycji       int16    `gorm:"primaryKey;autoIncrement" json:"id_pozycji"`
	Id_sprzedazy     int16    `json:"id_sprzedazy"`
	Id_produktu      int16    `json:"id_produktu"`
	Ilosc            int16    `json:"ilosc"`
	Cena_jednostkowa *float64 `json:"cena_jednostkowa"`
}

func (SalePositions) TableName() string {
	return "pozycja_sprzedazy"
}
