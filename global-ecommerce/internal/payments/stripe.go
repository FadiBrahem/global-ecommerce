package payments

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

func CreatePaymentIntent(amount int64, currency string) (*stripe.PaymentIntent, error) {
	stripe.Key = "your_stripe_secret_key"

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
	}

	return paymentintent.New(params)
}
