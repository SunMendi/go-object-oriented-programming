// payment/bkash.go
// Implements the PaymentProvider interface for Bkash payment method.

package payment

import "fmt"

// Bkash struct represents the Bkash payment provider.
type Bkash struct{}

// Pay processes the payment using Bkash, demonstrating polymorphism via interface implementation.
func (b Bkash) Pay(amount float64) error {
	fmt.Println("Paid With Bkash:", amount)
	return nil
}
