package models

type DiningMenuPositions struct {
	Id_pozycji_jadlospisu int16 `gorm:"primaryKey;autoIncrement" json:"id_pozycji_jadlospisu"`
	Id_jadlospisu         int16 `json:"id_jadlospisu"`
	Id_produktu           int16 `json:"id_produktu"`
}

func (DiningMenuPositions) TableName() string {
	return "pozycja_jadlospisu"
}