package models

import "gorm.io/gorm"

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(
		&Users{},
		// Dodac kolejne modele tutaj
	)
}
