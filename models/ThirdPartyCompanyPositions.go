package models

type ThirdPartyCompanyPositions struct {
	Id_oferty   int16 `gorm:"primaryKey;autoIncrement" json:"id_oferty"`
	Id_firmy    int16 `json:"id_firmy"`
	Id_produktu int16 `json:"id_produktu"`
}

func (ThirdPartyCompanyPositions) TableName() string {
	return "oferta_firmy_zewnetrznej"
}
