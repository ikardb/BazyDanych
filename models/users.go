package models

type Users struct {
	Id_uzytkownika int16   `gorm:"primaryKey;autoIncrement" json:"id_uzytkownika"`
	Imie           string  `json:"imie"`
	Nazwisko       *string `json:"nazwisko"`
	Administrator  bool    `json:"administrator"`
	Login          string  `json:"login"`
	Haslo          string  `json:"haslo"`
	Id_sklepu      uint    `json:"id_sklepu"`
}

func (Users) TableName() string {
	return "uzytkownicy"
}
