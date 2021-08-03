package main

import "testing"

func TestWallet(t *testing.T) {

	assetBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assetError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Error("wanted error but didnt get one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(50))
		assetBalance(t, wallet, 50)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balanace: Bitcoin(50)}
		wallet.Withdraw(Bitcoin(30))
		assetBalance(t, wallet, 20)
	})

	t.Run("withdraw more then amount", func(t *testing.T) {
		startingBal := Bitcoin(20)
		wallet := Wallet{balanace: startingBal}
		err := wallet.Withdraw(Bitcoin(100))
		assetError(t, err, "cannot withdraw, insufficient funds")
		assetBalance(t, wallet, startingBal)

	})
}
