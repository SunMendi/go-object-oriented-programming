// main.go
// Entry point of the application demonstrating OOP concepts in Go.
// It creates products, sets up a payment provider, and processes a cart checkout.

package main

import (
	"oop_with_go/models"
	"oop_with_go/payment"
	"oop_with_go/service"
)

func main() {
	// Create product instances using encapsulation and constructor pattern
	shirt := *models.NewProduct(1, "Shirt", 500, 10)
	pant := *models.NewProduct(2, "Pant", 800, 5)

	// Use abstraction and polymorphism for payment provider
	bkash := payment.Bkash{}

	// Create a cart and add products, demonstrating composition and method usage
	cart := service.NewCart(bkash)
	cart.AddItem(shirt)
	cart.AddItem(pant)
	cart.Checkout() // Checkout uses the payment provider abstraction
}
