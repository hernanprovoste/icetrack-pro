package main

import (
	"github.com/hernanprovoste/icetrack-pro/internal/core/product"
)

func main() {
	p := product.Product{
		ID:          1,
		Name:        "Hielo 2kg",
		Description: "Saco de hielo para distribución mayorista",
		WeightKg:    2,
		Price:       3500,
		Active:      true,
		Barcode:     "1234567890123",
	}

	product.Print(p)
}
