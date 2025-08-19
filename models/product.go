// models/product.go
// Demonstrates encapsulation in Go by using private fields and public methods.

package models

// Product struct represents a product with encapsulated fields.
type Product struct {
	ID    int     // Public field for product ID
	name  string  // Private field for product name
	price float64 // Private field for product price
	stock int     // Private field for product stock
}

// NewProduct is a constructor function for Product, enforcing encapsulation.
func NewProduct(id int, name string, price float64, stock int) *Product {
	return &Product{ID: id, name: name, price: price, stock: stock}
}

// Name returns the product's name (getter for private field).
func (p Product) Name() string {
	return p.name
}

// Price returns the product's price (getter for private field).
func (p Product) Price() float64 {
	return p.price
}

// ReduceStock safely modifies the stock, demonstrating controlled access.
func (p *Product) ReduceStock(amount int) bool {
	if amount > p.stock {
		return false
	}
	p.stock -= amount
	return true
}

// The private fields (name, price, stock) are only accessible via methods, demonstrating encapsulation.
