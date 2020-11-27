package main

import "testing"


func TestWallet(t *testing.T) {

    assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
        t.Helper()
        got := wallet.Balance()

        if got != want {
            t.Errorf("got %s want %s", got, want)
        }
    }

    t.Run("Deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(10))
    })

    t.Run("Withdraw", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(20)}
        wallet.Withdraw(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(10))
    })


   assertError := func(t *testing.T, err error) {
   	t.Helper()
   	if err == nil {
   		t.Error("wanted an error but didn't get one")
	}
   }
   
   t.Run("Withdraw insufficient funds", func(t *testing.T) {
   	wallet := Wallet{Bitcoin(20)}
	err := wallet.Withdraw(Bitcoin(100))

	assertBalance(t, wallet, Bitcoin(20))
    	assertError(t, err)
  })

}











