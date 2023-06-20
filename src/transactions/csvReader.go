package transactions

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func ReadCSVFile(filename string) []Transaction {
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

func HandleLine(line []string) Transaction {
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
	date, err := time.Parse("01/02/2006", line[0])
	if err != nil {
		panic("Error parsing date: " + err.Error())
	}

	description := line[1]

	//Parse amount
	amount, err := strconv.ParseFloat(line[3], 64)
	if err != nil {
		panic("Error parsing amount: " + err.Error())
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

	return Transaction{date, description, tType, category, amount, ""}

}
