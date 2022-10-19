package internals

import (
	"fmt"
)

var balance int

func AccountBalance() {
	fmt.Printf("Your account Balance is %v naira", balance)
	newLine(1)
}

func WithdrawFunds() {
	newLine(1)
	fmt.Println("Withdraw funds")
	var amount int
	fmt.Printf("Enter amount: ")
	fmt.Scan(&amount)
	if amount < 0 {
		fmt.Println("invalid amount, enter amount greater than 0")
		WithdrawFunds()
	}

	if amount > balance {
		fmt.Println("insufficienCt balance")
		fmt.Printf("Currnebalance - %v naira", balance)
		WithdrawFunds()
	}
	balance -= amount

	fmt.Printf("Amount withdrawn - %v naira,\nCurrnebalance - %v naira", amount, balance)
}

func DepositFunds() {
	fmt.Println("deposit funds")
	var amount int
	fmt.Printf("Enter amount: ")
	fmt.Scan(&amount)
	if amount < 0 {
		fmt.Println("invalid amount")
	}
	prev_balance := balance
	balance += amount

	fmt.Printf("Previous Balance - %v naira,\nCurrnt balance - %v naira", prev_balance, balance)
}
