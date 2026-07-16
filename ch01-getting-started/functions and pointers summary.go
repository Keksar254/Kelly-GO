package main

import "fmt"

type Account struct {
	Owner   string
	Balance float64
}

func deposit(acc *Account, amount float64) (float64, error) {
	if amount <= 0 {
		return acc.Balance, fmt.Errorf("deposit amount must be positive, got %.2f", amount)
	}
	acc.Balance += amount
	return acc.Balance, nil
}

func withdraw(acc *Account, amount float64) (float64, error) {
	if amount <= 0 {
		return acc.Balance, fmt.Errorf("withdrawal amount must be positive, got %.2f", amount)
	}
	if amount > acc.Balance {
		return acc.Balance, fmt.Errorf("insufficient funds: balance is %.2f, tried to withdraw %.2f", acc.Balance, amount)
	}
	acc.Balance -= amount
	return acc.Balance, nil
}

func totalDeposits(amounts ...float64) float64 {
	total := 0.0
	for _, amt := range amounts {
		total += amt
	}
	return total
}

func makeTransactionLogger(accountOwner string) func(string, float64) {
	transactionCount := 0

	return func(action string, amount float64) {
		transactionCount++
		fmt.Printf("[%s] Transaction #%d: %s of %.2f\n", accountOwner, transactionCount, action, amount)
	}
}

func processTransaction(acc *Account, action string, amount float64) error {
	fmt.Println("--- Processing started ---")
	defer fmt.Println("--- Processing finished ---")

	var err error
	switch action {
	case "deposit":
		_, err = deposit(acc, amount)
	case "withdraw":
		_, err = withdraw(acc, amount)
	default:
		err = fmt.Errorf("unknown action: %s", action)
	}

	return err
}

func main() {
	account := &Account{Owner: "Kelly", Balance: 1000.0}

	logTransaction := makeTransactionLogger(account.Owner)

	newBalance, err := deposit(account, 500)
	if err != nil {
		fmt.Println("Deposit error:", err)
	} else {
		logTransaction("deposit", 500)
		fmt.Printf("New balance: %.2f\n\n", newBalance)
	}

	newBalance, err = withdraw(account, 200)
	if err != nil {
		fmt.Println("Withdraw error:", err)
	} else {
		logTransaction("withdraw", 200)
		fmt.Printf("New balance: %.2f\n\n", newBalance)
	}

	_, err = withdraw(account, 5000)
	if err != nil {
		fmt.Println("Withdraw error:", err)
		fmt.Println()
	}

	batchTotal := totalDeposits(50, 25.50, 100, 75.25)
	fmt.Printf("Total of batch deposits: %.2f\n\n", batchTotal)

	err = processTransaction(account, "deposit", 300)
	if err != nil {
		fmt.Println("Transaction error:", err)
	} else {
		logTransaction("deposit", 300)
	}

	fmt.Printf("\nFinal balance for %s: %.2f\n", account.Owner, account.Balance)
}