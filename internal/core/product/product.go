package product

import (
	"fmt"
	"time"
)

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null;size:100" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	WeightKg    float64   `gorm:"not null;default:0" json:"weight_kg"`
	Price       int64     `gorm:"not null;default:0" json:"price"`
	Active      bool      `gorm:"not null;default:true" json:"active"`
	Barcode     string    `gorm:"uniqueIndex" json:"barcode"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func Print(p Product) {
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Description: %s\n", p.Description)
	fmt.Printf("Weight (kg): %.2f\n", p.WeightKg)
	fmt.Printf("Price: %d\n", p.Price)
	fmt.Printf("Active: %t\n", p.Active)
	fmt.Printf("Barcode: %s\n", p.Barcode)
}
