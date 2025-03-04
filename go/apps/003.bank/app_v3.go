package main

import (
	"fmt"

	"fcfarah.dev.br/bank/fileops"
)

func appV3() {
	// var balance = 1000.0
	var balance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println(err)
		fmt.Println("------------")
		// return
	}

	fmt.Println("Welcome to the GoBank!")
	for {
		displayMenu(balance)
		choice := getUserChoice()
		if !processChoice(choice, &balance) {
			break
		}
	}
}

func displayMenu(balance float64) {
	fmt.Println("--------------------------------------")
	fmt.Println("Your current balance is: ", balance)
	fmt.Println("--------------------------------------")
	fmt.Println("Please select an option to continue:")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")
}

func getUserChoice() int {
	var choice int
	fmt.Print("Enter your choice here: ")
	fmt.Scan(&choice)
	fmt.Println("You selected option: ", choice)
	return choice
}

func processChoice(choice int, balance *float64) bool {
	switch choice {
	case 1:
		fmt.Println("Your balance is: ", *balance)
	case 2:
		deposit(balance)
	case 3:
		withdraw(balance)
	case 4:
		fmt.Println("Thank you for using GoBank!")
		return false
	default:
		fmt.Println("Invalid choice")
	}
	fmt.Println("")
	fmt.Println("=======================================")
	fmt.Println("")
	return true
}

// deposit adds the given amount to the balance.
// It prompts the user to enter the deposit amount and updates the balance accordingly.
// The balance is passed as a pointer to allow modifying the original value.
func deposit(balance *float64) {
	fmt.Print("Enter the amount to deposit: ")
	var deposit float64
	fmt.Scan(&deposit)
	*balance += deposit
	fileops.WriteFloatToFile(balance, accountBalanceFile)
	fmt.Println("Your new balance is: ", *balance)
}

func withdraw(balance *float64) {
	fmt.Print("Enter the amount to withdraw: ")
	var withdraw float64
	fmt.Scan(&withdraw)
	if withdraw > *balance {
		fmt.Println("Insufficient balance")
	} else {
		*balance -= withdraw
		fileops.WriteFloatToFile(balance, accountBalanceFile)
		fmt.Println("Your new balance is: ", *balance)
	}
}

const accountBalanceFile = "balance.txt"
