package transactions

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func ReadTransactions(folderPath string) []Transaction {
	listOfCSVFiles, err := FindCSVFiles(folderPath)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	transactions := make([]Transaction, 0)
	for _, csvFile := range listOfCSVFiles {
		transactions = append(transactions, ReadCSVFile(csvFile)...)
	}

	return transactions
}

func FindCSVFiles(folderPath string) ([]string, error) {
	//Given the path to a folder, read all the csv files in the folder and return a slice of transactions
	//List all files in the folder
	log.Println("Looking for files in: " + folderPath)
	var filePaths []string

	// Open the folder
	dir, err := os.Open(folderPath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer dir.Close()

	// Read the folder contents
	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Iterate over the files
	for _, file := range files {
		if !file.IsDir() {
			// Check if the file has a .csv extension
			if strings.HasSuffix(file.Name(), ".csv") {
				// Build the full file path
				filePath := filepath.Join(folderPath, file.Name())
				filePaths = append(filePaths, filePath)
				log.Println("Found file: " + filePath)
			}
		}
	}
	return filePaths, nil

}

func ReadCSVFile(filename string) []Transaction {
	transactions := make([]Transaction, 0)

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	counter := 0
	for _, line := range lines {
		transaction, err := HandleLine(line)
		if err != nil {
			log.Printf("Error parsing line: %s with error: %s", line, err.Error())
		} else {
			transactions = append(transactions, transaction)
			counter++
		}
	}
	log.Printf("Read %d transactions from %s", counter, filename)

	return transactions
}

func HandleLine(line []string) (Transaction, error) {
	//line[0] = date
	//line[1] = description
	//line[2] = original description
	//line[3] = amount
	//line[4] = transaction type
	//line[5] = category
	//line[6] = account name
	//line[7] = labels
	//line[8] = notes

	// Parse date
	date, err := time.Parse("1/02/2006", line[0])
	if err != nil {
		return Transaction{}, errors.New("Error parsing date: " + err.Error())
	}

	description := line[1]

	//Parse amount
	amount, err := strconv.ParseFloat(line[3], 64)
	if err != nil {
		return Transaction{}, errors.New("Error parsing amount: " + err.Error())
	}

	//Parse transaction type
	var tType TransactionType
	if line[4] == "debit" {
		tType = DEBIT
	} else {
		tType = CREDIT
	}

	//Parse category
	var category TransactionCategory
	switch line[5] {
	case "Auto & Transport":
		category = AUTO_AND_TRANSPORT
	case "Bills & Utilities":
		category = BILLS_AND_UTILITIES
	case "Business Services":
		category = BUSINESS_SERVICES
	case "Education":
		category = EDUCATION
	case "Entertainment":
		category = ENTERTAINMENT
	case "Fees & Charges":
		category = FEES_AND_CHARGES
	case "Food & Drink":
		category = FOOD_AND_DRINK
	case "Gifts & Donations":
		category = GIFTS_AND_DONATIONS
	case "Health & Fitness":
		category = HEALTH_AND_FITNESS
	case "Home":
		category = HOME
	case "Income":
		category = INCOME
	case "Insurance":
		category = INSURANCE
	case "Investments":
		category = INVESTMENTS
	case "Kids":
		category = KIDS
	case "Loans":
		category = LOANS
	case "Personal Care":
		category = PERSONAL_CARE
	case "Pets":
		category = PETS
	case "Shopping":
		category = SHOPPING
	case "Taxes":
		category = TAXES
	case "Transfer":
		category = TRANSFER
	case "Travel":
		category = TRAVEL
	default:
		category = UNCATEGORIZED
	}

	return Transaction{date, description, tType, category, amount, ""}, nil

}
