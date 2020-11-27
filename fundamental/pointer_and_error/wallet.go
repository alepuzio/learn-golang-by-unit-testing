package main

/**
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors
*/

type Bitcoin int

type Wallet struct {
    balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
    return w.balance
}
