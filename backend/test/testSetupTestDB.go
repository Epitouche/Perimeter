package test

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"area/schemas"
)

func SetupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Action{}, &schemas.Service{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
