package models

type Products struct {
	Id_produktu int16   `gorm:"primaryKey;autoIncrement" json:"id_produktu"`
	Nazwa       string  `json:"nazwa"`
	Cena        float64 `json:"cena"`
}

func (Products) TableName() string {
	return "produkt"
}
