package transactions

import (
	"encoding/csv"
	"fmt"
	"os"
)


struct Transaction {
	date string
	description string
	tType TransactionType
	category TransactionCategory
)




func ReadCSVFile(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}

func PrintHello() {
	fmt.Println("Hello, World! Imported from utils")
}
