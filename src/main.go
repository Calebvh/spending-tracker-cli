package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Calebvh/spending-tracker-cli/config"
	"github.com/Calebvh/spending-tracker-cli/logger"
	"github.com/Calebvh/spending-tracker-cli/transactions"
)

func main() {
	//Initialize logger
	loggerConfig := config.ReadLoggerConfig()
	logger.SetLogConfig(loggerConfig.LogLevel, loggerConfig.DebugMode, loggerConfig.WarningMode)
	logger.Debug("Logger initialized")
	logger.Error("Logger initialized")
	logger.Info("Logger initialized")
	logger.Warning("Logger initialized")
	for {
		showMenu()

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your choice:")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			config := config.ReadTransactionConfig()
			println(config)
			transactions.ReadTransactions(config.TransactionFolder)

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
	fmt.Println("3. Exit")
	fmt.Println("-----------------------------")
}
