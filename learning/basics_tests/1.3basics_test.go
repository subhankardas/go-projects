package basics_tests

import (
	"errors"
	"testing"
)

/* Test Driven Development (TDD)
1. Write unit tests first */
func TestWallet(test *testing.T) {
	assertBalance := func(test *testing.T, wallet Wallet, expected Money) {
		test.Helper()
		actual := wallet.Balance()
		if expected != actual {
			test.Errorf("expected %d and got %d", expected, actual)
		}
	}

	assertError := func(test *testing.T, err error, errmsg string) {
		test.Helper()
		if err == nil { // Check for error
			test.Fatal("expected", errmsg, "error")
		}
		if err.Error() != errmsg { // Check for error message
			test.Errorf("expected %q and got %q", errmsg, err)
		}
	}

	test.Run("deposit", func(test *testing.T) {
		wallet := Wallet{}
		expected := Money(10)                 // Given
		wallet.Deposit(Money(10))             // When
		assertBalance(test, wallet, expected) // Then
	})

	test.Run("withdraw", func(test *testing.T) {
		wallet := Wallet{balance: Money(10)}
		expected := Money(0)                  // Given
		wallet.Withdraw(Money(10))            // When
		assertBalance(test, wallet, expected) // Then
	})

	test.Run("withdraw insufficient funds error", func(test *testing.T) {
		initialBalance := Money(10)
		wallet := Wallet{balance: initialBalance}   // Given
		err := wallet.Withdraw(Money(20))           // When
		assertBalance(test, wallet, initialBalance) // Then
		assertError(test, err, "insufficient funds")
	})
}

/* 2. Then write the code to pass the tests */
type Money int
type Wallet struct {
	balance Money
}

func (wallet *Wallet) Deposit(amount Money) {
	wallet.balance += amount
}

func (wallet *Wallet) Withdraw(amount Money) error {
	if amount > wallet.balance {
		return errors.New("insufficient funds")
	}
	wallet.balance -= amount
	return nil
}

func (wallet *Wallet) Balance() Money {
	return wallet.balance
}
