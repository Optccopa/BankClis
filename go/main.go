package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var balance float64
	fmt.Printf("Welcome to GoBank\nBalance: %.2f\n\n1. Deposit\n2. Withdraw\n3. Check Balance\n4. Clear\n5. Exit\n", balance)
	for {
		fmt.Print("@GoBank> ")
		var userInput string
		fmt.Scan(&userInput)

		switch userInput {
		case "1": // Deposit
			var deposit float64
			fmt.Print("Deposit amount: ")
			fmt.Scan(&deposit)
			balance += deposit
			fmt.Printf("Deposited %.2f\n", deposit)
			continue

		case "2": // Withdraw
			var withdraw float64
			fmt.Print("Withdraw amount: ")
			fmt.Scan(&withdraw)
			if withdraw > balance {
				fmt.Print("Cannot withdraw more than your balance.\n")
				continue
			}
			balance -= withdraw
			fmt.Printf("Withdrew %.2f\n", withdraw)

		case "3": // Check balance
			fmt.Printf("Your balance is: %.2f\n", balance)

		case "4": // Clear
			cmd := exec.Command("cmd", "/C", "cls")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()

		case "5": // Exit
			return
		}
	}
}
