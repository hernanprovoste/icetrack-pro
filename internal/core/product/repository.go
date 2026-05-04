package product

import (
	"errors"

	"gorm.io/gorm"
)

func ListAll(db *gorm.DB) ([]Product, error) {
	var products []Product

	result := db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func FindByID(db *gorm.DB, id uint) (*Product, error) {
	var p Product

	result := db.First(&p, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &p, nil
}
