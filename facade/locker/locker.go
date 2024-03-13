package locker

import "company.com/facade/finance"

type Box struct {
	balance float32
}

func (a *Box) GetBalance() float32 {
	return a.balance
}

func (a *Box) SetBalance(amount float32) {
	a.balance = amount
}

type Locker struct {
}

func (a *Locker) Deposit(account *Box, amount float32) {
	account.balance += amount
}

func (a *Locker) Withdrawl(account *Box, amount float32) {
	account.balance -= amount
}

func (a *Locker) StoreAmount(store finance.Store, amount float32) {
	a.Deposit(store.(*Box), amount)
}
