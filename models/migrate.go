package models

import "gorm.io/gorm"

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(
		&Users{},
		&Shops{},
		&Orders{},
		&Kitchens{},
		&DiningMenus{},
		&DiningMenuPositions{},
		&Products{},
		&OrderPositions{},
		&StockLevels{},
		&Sales{},
		&SalePositions{},
		&ThirdPartyCompanies{},
		&ThirdPartyCompanyPositions{},
	)
}
