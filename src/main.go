package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Calebvh/spending-tracker-cli/config"
)

func main() {

	for {
		showMenu()

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("-----------------------------\n")
		fmt.Print("Enter your choice: ")
		fmt.Print("-----------------------------\n")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			config := config.ReadConfig()
			println(config.TransactionFolder)

		case "2":
			println("List Transactions")
		case "3":
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func showMenu() {
	fmt.Println("---- Transaction Tracker ----")
	fmt.Println("1. Load Transactions")
	fmt.Println("2. List Transactions")
}
