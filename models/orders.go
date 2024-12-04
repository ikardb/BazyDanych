package models

import "time"

type Orders struct {
	Id_zamowienia    uint      `gorm:"primaryKey;autoIncrement" json:"id_zamowienia"`
	Id_sklepu        uint      `json:"id_sklepu"`
	Id_uzytkownika   *int16    `json:"id_uzytkownika"`
	Data_zamowienia  time.Time `json:"data_zamowienia"`
	Koszt_zamowienia float64   `json:"koszt_zamowienia"`
}

func (Orders) TableName() string {
	return "zamowienie"
}