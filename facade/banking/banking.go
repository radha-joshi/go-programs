package banking

import "company.com/facade/finance"

type Account struct {
	balance float32
}

func (a *Account) GetBalance() float32 {
	return a.balance
}

func (a *Account) SetBalance(amount float32) {
	a.balance = amount
}

type Bank struct {
}

func (a *Bank) Deposit(account *Account, amount float32) {
	account.balance += amount
}

func (a *Bank) Withdrawl(account *Account, amount float32) {
	account.balance -= amount
}

func (a *Bank) Transfer(source *Account, target *Account, amount float32) {
	a.Withdrawl(source, amount)
	a.Deposit(target, amount)
}

func (a *Bank) StoreAmount(store finance.Store, amount float32) {
	a.Deposit(store.(*Account), amount)
}
