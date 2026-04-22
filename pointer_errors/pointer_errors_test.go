// package main

// import (
// 	"testing"
// )

// func TestWallet1(t *testing.T) {
// 	wallet := Wallet{}
// 	wallet.Deposit(Bitcoin(10))

// 	got := wallet.Balance()
// 	want := Bitcoin(20)

// 	if got != want {
// 		t.Errorf("got %s want %s", got, want)
// 	}

// 	t.Run("deposit", func(t *testing.T) {
// 		wallet := Wallet{}
// 		wallet.Deposit(Bitcoin(10))

// 		got := wallet.Balance()
// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})

// 	t.Run("withdraw", func(t *testing.T) {
// 		wallet := Wallet{balance: Bitcoin(20)}
// 		wallet.Withdraw(Bitcoin(10))

// 		got := wallet.Balance()
// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})
// }

// func TestWallet(t *testing.T) {
// 	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
// 		t.Helper()
// 		got := wallet.Balance()

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	}

// 	t.Run("deposit", func(t *testing.T) {
// 		wallet := Wallet{}
// 		wallet.Deposit(Bitcoin(10))
// 		assertBalance(t, wallet, Bitcoin(10))
// 	})

// 	t.Run("withdraw", func(t *testing.T) {
// 		wallet := Wallet{balance: Bitcoin(20)}
// 		wallet.Withdraw(Bitcoin(10))
// 		assertBalance(t, wallet, Bitcoin(10))
// 	})
// }

package main

import "testing"

// func TestWallet1(t *testing.T) {
// 	wallet := Wallet{}
// 	wallet.Deposit(Bitcoin(10))

// 	got := wallet.Balance()
// 	want := Bitcoin(11)

// 	if got != want {
// 		t.Errorf("got %s want %s", got, want)
// 	}
// }

// func TestWallet2(t *testing.T) {
// 	t.Run("deposit", func(t *testing.T) {
// 		wallet := Wallet{}
// 		wallet.Deposit(Bitcoin(10))

// 		got := wallet.Balance()
// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})

// 	t.Run("withdraw", func(t *testing.T) {
// 		wallet := Wallet{}
// 		wallet.Withdraw(Bitcoin(10))

// 		got := wallet.Balance()
// 		want := Bitcoin(20)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})
// }

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(15))

		assertBalance(t, wallet, Bitcoin(15))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		err := wallet.Withdraw(Bitcoin(5))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}

	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
