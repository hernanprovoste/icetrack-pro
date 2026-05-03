package product

import "fmt"

type Product struct {
	ID          uint
	Name        string
	Description string
	WeightKg    float64
	Price       int64
	Active      bool
	Barcode     string
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
