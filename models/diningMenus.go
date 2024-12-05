package models

type DiningMenus struct {
	Id_jadlospisu int16  `gorm:"primaryKey;autoIncrement" json:"id_jadlospisu"`
	Id_kuchni     *int16 `json:"id_kuchni"`
	Id_firmy      *int16 `json:"id_firmy"`
}

func (DiningMenus) TableName() string {
	return "jadlospis"
}
