package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s != want %s", got, want)
		}
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)

		if err != nil {
			t.Errorf("error returned but it should not!")
		}
		if got != want {
			t.Errorf("got %s != want %s", got, want)
		}
	})
	t.Run("Withdraw more thant we have in deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		if err == nil {
			t.Error("Wanted error but got nil")
		}
		if err.Error() != lowBalance.Error() {
			t.Errorf("Inncorect error, want: %s, got: %s", lowBalance.Error(), err.Error())
		}
	})
}
