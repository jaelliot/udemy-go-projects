package main

import (
	"fmt" // for input/output

	"example.com/bank/fileops" // for file operations
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error reading balance file. Defaulting to 1000")
		fmt.Println(err)
		fmt.Println("------------------------------------")
		panic("Cannot continue without balance file")
	}

	fmt.Println("Welcome to the Bank of Golang")

	for {

		presentOptions()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your new balance is: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdraw amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than your balance.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New Amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
