// service/cart_service.go
// Cart service demonstrates composition, encapsulation, and interface usage in Go OOP.

package service

import (
	"fmt"
	"oop_with_go/models"
	"oop_with_go/payment"
)

// Cart struct composes products and a payment provider, demonstrating OOP composition.
type Cart struct {
	items   []models.Product        // List of products in the cart
	payment payment.PaymentProvider // Payment provider abstraction
}

// NewCart creates a new Cart with the given payment provider (constructor pattern).
func NewCart(p payment.PaymentProvider) *Cart {
	return &Cart{payment: p}
}

// AddItem adds a product to the cart, reducing stock if available.
func (c *Cart) AddItem(p models.Product) {
	if p.ReduceStock(1) {
		c.items = append(c.items, p)
		fmt.Println("Added:", p.Name())
	} else {
		fmt.Println("Out of stock:", p.Name())
	}
}

// Checkout calculates the total and processes payment using the provided payment provider.
func (c *Cart) Checkout() {
	total := 0.0
	for _, item := range c.items {
		total += item.Price()
	}
	c.payment.Pay(total)
}
