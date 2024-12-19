package models

type Shops struct {
	Id_sklepu    int16   `gorm:"primaryKey;autoIncrement" json:"id_sklepu"`
	Nazwa_sklepu string  `json:"nazwa_sklepu"`
	Ulica        *string `json:"ulica"`
}

func (Shops) TableName() string {
	return "sklep"
}
