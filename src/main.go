package main

import (
	"fmt"

	"github.com/Calebvh/spending-tracker-cli/transactions"
)

func main() {
	fmt.Println("Hello World")

	transactions.PrintHello()

	transactions.ReadCSVFile("/media/caleb/Backup/Documents/SpendingData/01_2023_Transactions.csv")

}
