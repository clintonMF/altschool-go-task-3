package main

import (
	"fmt"
	"os"

	g "github.com/clintonMF/altschool-go-task-3/internals"
)

func welcome() {
	fmt.Println("*******{ Welcome to the ATM CLI app 🍾 }*******")
}

func exitProgram() {
	fmt.Println("Goodbye 👋")
	os.Exit(0)
}

func newLine(n int) {
	for n > 0 {
		fmt.Println("")
		n--
	}
}

func anotherOperation() {
	var yesOrNo string
	newLine(1)
	fmt.Println("Do you want to perform another operation")
	fmt.Println("if Yes type y, else type anything")
	fmt.Scan(&yesOrNo)
	if yesOrNo == "y" {
		menu()
	}
}
func menu() {
	newLine(1)
	fmt.Println("Select Operation:")
	fmt.Println("1. Change pin \n2. Account Balance")
	fmt.Println("3. Withdraw funds \n4. Deposit funds")
	fmt.Println("0. Exit program \t")
	newLine(1)

	var menuNumber int
	fmt.Printf("Enter operation number: ")
	_, err := fmt.Scan(&menuNumber)
	if err != nil {
		fmt.Println("Error:, please select the correct menu item")
		//menu()
	}

	switch menuNumber {
	case 1:
		g.ChangePin()
		anotherOperation()
	case 2:
		g.AccountBalance()
		anotherOperation()
	case 3:
		g.WithdrawFunds()
		anotherOperation()
	case 4:
		g.DepositFunds()
		anotherOperation()
	case 0:
		exitProgram()
	default:
		fmt.Println("Invalid input ")
		anotherOperation()
	}
}

func main() {
	g.Login()
	welcome()
	menu()
}
