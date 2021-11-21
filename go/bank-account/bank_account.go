package account

import "sync"

// Define the Account type here
type Account struct {
	amount int64
	open   bool
	sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{amount: amount, open: true}
}

func (a *Account) Balance() (int64, bool) {
	if !a.open {
		return 0, false
	}
	return a.amount, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if !a.open {
		return 0, false
	}
	a.Lock()
	defer a.Unlock()
	if a.amount+amount < 0 {
		return 0, false
	}
	a.amount += amount
	return a.amount, true
}

func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if !a.open {
		return 0, false
	}
	a.open = false
	return a.amount, true
}
