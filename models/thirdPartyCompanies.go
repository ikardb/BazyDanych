package models

type ThirdPartyCompanies struct {
	Id_firmy int16  `gorm:"primaryKey;autoIncrement" json:"id_firmy"`
	Nazwa    string `json:"nazwa"`
}

func (ThirdPartyCompanies) TableName() string {
	return "firma_zewnetrzna"
}
