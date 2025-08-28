// main.go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Bank represents a simple bank account.
// It contains a balance and a Mutex to protect against race conditions.
type Bank struct {
	balance int
	mu      sync.Mutex
}

// Deposit adds an amount to the account balance.
// It locks the mutex before updating the balance to ensure thread safety.
func (b *Bank) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure WaitGroup counter is decremented when the goroutine finishes

	b.mu.Lock() // Lock the mutex
	// Simulate some work being done for the transaction
	time.Sleep(100 * time.Millisecond)
	b.balance += amount
	b.mu.Unlock() // Unlock the mutex

	fmt.Printf("Deposited Rs.%d. Current balance: Rs.%d\n", amount, b.balance)
}

// Withdraw subtracts an amount from the account balance.
// It first checks if the withdrawal will result in a negative balance.
// It locks the mutex before checking and updating the balance.
func (b *Bank) Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	b.mu.Lock()
	// Check if there are sufficient funds before proceeding
	if b.balance-amount < 0 {
		fmt.Printf("Withdrawal of Rs.%d failed. Insufficient funds. Current balance: Rs.%d\n", amount, b.balance)
		b.mu.Unlock() // Unlock the mutex before returning
		return
	}
	// Simulate some work
	time.Sleep(100 * time.Millisecond)
	b.balance -= amount
	b.mu.Unlock()

	fmt.Printf("Withdrew Rs.%d. Current balance: Rs.%d\n", amount, b.balance)
}

// GetBalance returns the current account balance.
// It also uses the mutex to ensure a consistent read.
func (b *Bank) GetBalance() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.balance
}

func main() {
	// Initialize the bank account with a starting balance
	account := Bank{balance: 500}
	var wg sync.WaitGroup

	fmt.Println("Starting balance:", account.GetBalance())

	// Launch multiple concurrent deposit and withdrawal operations
	// A WaitGroup is used to wait for all Goroutines to complete.
	operations := []struct {
		op     string
		amount int
	}{
		{"deposit", 200},
		{"withdraw", 150},
		{"deposit", 100},
		{"withdraw", 600}, // This withdrawal should fail
		{"deposit", 50},
		{"withdraw", 250},
	}

	wg.Add(len(operations))
	for _, op := range operations {
		switch op.op {
		case "deposit":
			go account.Deposit(op.amount, &wg)
		case "withdraw":
			go account.Withdraw(op.amount, &wg)
		}
	}

	// Wait for all Goroutines to finish their work
	wg.Wait()

	fmt.Println("\nAll transactions complete.")
	fmt.Printf("Final balance: Rs.%d\n", account.GetBalance())
}
