package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Id_uzytkownika int16   `gorm:"primaryKey;autoIncrement" json:"id_uzytkownika"`
	Imie           string  `json:"imie"`
	Nazwisko       *string `json:"nazwisko"`
	Administrator  bool    `json:"administrator"`
	Login          string  `json:"login"`
	Haslo          string  `json:"haslo"`
	Id_sklepu      int16   `json:"id_sklepu"`
}

func (Users) TableName() string {
	return "uzytkownicy"
}

func (Users) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (Users) ComparePasswordWithHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
