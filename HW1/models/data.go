package models

var DB []Product

type Product struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Amount int     `json:"amount"`
	Price  float64 `json:"price"`
}

func init() {
	product1 := Product{
		ID:     1,
		Title:  "Table",
		Amount: 10,
		Price:  400.50,
	}
	DB = append(DB, product1)
}

func FindProductById(id int) (Product, bool) {
	var product Product
	var found bool
	for _, b := range DB {
		if b.ID == id {
			product = b
			found = true
			break
		}
	}
	return product, found
}
