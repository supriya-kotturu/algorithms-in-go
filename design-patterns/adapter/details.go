package main

import "fmt"

type Details struct {
	accountId   string
	country     string
	contact     string
	email       string
	paymentTerm string
	creditCard  string
}

func (d *Details) printDetails() {
	fmt.Println("---DETAILS---")
	fmt.Println("account id -", d.accountId)
	fmt.Println("country - ", d.country)
	fmt.Println("contact - ", d.contact)
	fmt.Println("email - ", d.email)
	fmt.Println("paymentTerm - ", d.paymentTerm)
	fmt.Println("creditCard - ", d.creditCard)
}
