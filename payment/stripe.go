// payment/stripe.go
// Implements the PaymentProvider interface for Stripe payment method.

package payment

import "fmt"

// Stripe struct represents the Stripe payment provider.
type Stripe struct{}

// Pay processes the payment using Stripe, demonstrating polymorphism via interface implementation.
func (s Stripe) Pay(amount float64) error {
	fmt.Println("Paid with Stripe:", amount)
	return nil
}
