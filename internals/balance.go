package internals

import (
	"fmt"
)

var balance float64

func AccountBalance() {
	fmt.Printf("Your account Balance is %.2f naira", balance)
	newLine(1)
}

func WithdrawFunds() {

	newLine(1)
	fmt.Println("Withdraw funds")
	var amount float64

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
		fmt.Printf("Currnent balance: %.2f naira", balance)
		return
	}
	prev_balance := balance
	balance -= amount

	fmt.Printf("\nPrevious balance: %.2f \n", prev_balance)
	fmt.Printf("Amount withdrawn: %.2f  naira,\nCurrnent balance: %.2f  naira", amount, balance)
}

func DepositFunds() {
	//This function handles the Deposit operations
	newLine(1)
	fmt.Println("deposit funds")
	var amount float64
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
	fmt.Printf("Previous Balance: %.2f  naira,\nCurrnt balance: %.2f naira", prev_balance, balance)
}
