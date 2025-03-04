package main

import (
	"fmt"
)

func app_V1() {

	var balance = 1000.0
	fmt.Println("Welcome to the GoBank!")
	var choice int
	for true {
		fmt.Println("--------------------------------------")
		fmt.Println("Your current balance is: ", balance)
		fmt.Println("--------------------------------------")
		fmt.Println("Please select an option to continue:")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice here: ")
		fmt.Scan(&choice)

		fmt.Println("You selected option: ", choice)
		switch choice {
		case 1:
			fmt.Println("Your balance is: ", balance)
		case 2:
			fmt.Print("Enter the amount to deposit: ")
			var deposit float64
			fmt.Scan(&deposit)
			balance += deposit
			fmt.Println("Your new balance is: ", balance)
		case 3:
			fmt.Print("Enter the amount to withdraw: ")
			var withdraw float64
			fmt.Scan(&withdraw)
			if withdraw > balance {
				fmt.Println("Insufficient balance")
			} else {
				balance -= withdraw
				fmt.Println("Your new balance is: ", balance)
			}
		case 4:
			fmt.Println("Thank you for using GoBank!")
			return
		default:
			fmt.Println("Invalid choice")
		}
		fmt.Println("")
		fmt.Println("=======================================")
		fmt.Println("")
	}
}
