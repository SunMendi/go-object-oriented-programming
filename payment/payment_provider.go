// payment/payment_provider.go
// Defines the PaymentProvider interface for abstraction and polymorphism.

package payment

// PaymentProvider abstracts payment processing, enabling polymorphism for different providers.
type PaymentProvider interface {
	Pay(amount float64) error
}
