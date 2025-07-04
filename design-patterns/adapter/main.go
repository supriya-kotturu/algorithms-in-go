package main

// Target interface - we need to support this new format of data
type Subscription interface {
	getDetails() Details
}

// Legacy interface - the older format which we need to convert to the
// `Details` format to support the new Structure
type LegacySubscription interface {
	getAccount() Account
	getBilling() Billing
}

// Adapter - the new struct which converts the old format to new format
// it should implement (have all the methods implemented, which are
// defined by the interface) the target interface
type Adapter struct {
	legacy LegacySubscription
}

// the method implemented to convert data into new format and satisfies the
// target interface
func (a *Adapter) getDetails() Details {
	account := a.legacy.getAccount()
	billing := a.legacy.getBilling()

	return Details{
		accountId:   account.id,
		country:     account.country,
		contact:     account.phone,
		email:       account.email,
		creditCard:  billing.creditCard,
		paymentTerm: billing.paymentTerm,
	}
}

func main() {
	// the struct which is in old format
	sub := &Sub{
		id:               "account-id-#23",
		country:          "Canada",
		region:           "west-ca-12",
		phone:            "127-789-239",
		email:            "johndoe@gmail.com",
		createdOn:        "12-01-2025",
		subscriptionType: "monthly",
		billingDate:      "01-12-2024",
		paymentTerm:      "annual",
		creditCard:       "2345 8923 3468",
	}

	adapter := Adapter{legacy: sub}
	details := adapter.getDetails()
	details.printDetails()
}
