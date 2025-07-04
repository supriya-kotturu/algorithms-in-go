package main

type Account struct {
	id        string
	country   string
	region    string
	phone     string
	email     string
	createdOn string
}

func (a *Account) getAccount() Account {
	return *a
}
