package models

type StockLevels struct {
	Id_stanu    int16 `gorm:"primaryKey;autoIncrement" json:"id_stanu"`
	Id_sklepu   int16 `json:"id_sklepu"`
	Id_produktu int16 `json:"id_produktu"`
	Ilosc       int16 `json:"ilosc"`
}

func (StockLevels) TableName() string {
	return "stan_magazynowy"
}
