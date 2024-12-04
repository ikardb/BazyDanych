package models

type Kitchens struct {
	Id_kuchni int16   `gorm:"primaryKey;autoIncrement" json:"id_kuchni"`
	Ulica     *string `json:"ulica"`
}

func (Kitchens) TableName() string {
	return "kuchnia"
}
