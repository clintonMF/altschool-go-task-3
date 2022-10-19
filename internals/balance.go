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

	if balance == 0 {
		fmt.Println("Your account is empty, please try to make a deposit before trying to withdraw")
		return
	}
	fmt.Printf("Enter amount: ")
	fmt.Scan(&amount)
	if amount < 0 {
		fmt.Println("invalid amount, enter amount greater than 0")
		return
	}

	if amount > balance {
		fmt.Println("insufficienCt balance")
		fmt.Printf("Currnent balance: %v naira", balance)
		return
	}
	prev_balance := balance
	balance -= amount

	fmt.Printf("\nPrevious balance: %v\n", prev_balance)
	fmt.Printf("Amount withdrawn: %v naira,\nCurrnent balance: %v naira", amount, balance)
}

func DepositFunds() {
	//This function handles the Deposit operations
	newLine(1)
	fmt.Println("deposit funds")
	var amount int
	fmt.Printf("Enter amount: ")
	fmt.Scan(&amount)
	if amount < 0 {
		fmt.Println("invalid amount")
		fmt.Println("You cannot deposit less than 0 naira")
		return
	}
	prev_balance := balance
	balance += amount

	newLine(1)
	fmt.Printf("Previous Balance - %v naira,\nCurrnt balance - %v naira", prev_balance, balance)
}
