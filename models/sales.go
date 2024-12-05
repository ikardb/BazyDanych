package models

import "time"

type Sales struct {
	Id_sprzedazy     int16     `gorm:"primaryKey;autoIncrement" json:"id_sprzedazy"`
	Id_sklepu        int16     `json:"id_sklepu"`
	Id_uzytkownika   int16     `json:"id_uzytkownika"`
	Kwota_transakcji float64   `json:"kwota_transakcji"`
	Data_sprzedazy   time.Time `json:"data_sprzedazy"`
}

func (Sales) TableName() string {
	return "sprzedaz"
}
