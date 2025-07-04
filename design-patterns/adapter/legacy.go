package main

type Sub struct {
	id               string
	country          string
	region           string
	phone            string
	email            string
	createdOn        string
	subscriptionType string
	billingDate      string
	paymentTerm      string
	creditCard       string
}

func (s *Sub) getAccount() Account {
	return Account{
		id:        s.id,
		country:   s.country,
		region:    s.region,
		phone:     s.phone,
		email:     s.email,
		createdOn: s.createdOn,
	}
}

func (s *Sub) getBilling() Billing {
	return Billing{
		subscriptionType: s.subscriptionType,
		billingDate:      s.billingDate,
		paymentTerm:      s.paymentTerm,
		creditCard:       s.creditCard,
	}
}
