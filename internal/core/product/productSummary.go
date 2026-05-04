package product

type ProductSummary struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Price  int64  `json:"price"`
	Active bool   `json:"active"`
}
