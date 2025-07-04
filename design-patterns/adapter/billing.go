package main

type Billing struct {
	subscriptionType string
	billingDate      string
	paymentTerm      string
	creditCard       string
}

func (b *Billing) getBilling() Billing {
	return *b
}
