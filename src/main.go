package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		showMenu()

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		// switch choice {
		// case "1":
		// 	addTransaction(&transactions)
		// case "2":
		// 	listTransactions(transactions)
		// case "3":
		// 	fmt.Println("Exiting program...")
		// 	return
		// default:
		// 	fmt.Println("Invalid choice. Please try again.")
		// }
	}
}

func showMenu() {
	fmt.Println("---- Transaction Tracker ----")
	fmt.Println("1. Add Transaction")
	fmt.Println("2. List Transactions")
	fmt.Println("3. Exit")
}
