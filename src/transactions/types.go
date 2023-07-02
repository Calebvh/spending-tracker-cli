package transactions

import (
	"time"
	"fmt"
)

type TransactionType int64

const (
	DEBIT TransactionType = iota
	CREDIT
)

type TransactionCategory int64

const (
	AUTO_AND_TRANSPORT TransactionCategory = iota
	BILLS_AND_UTILITIES
	BUSINESS_SERVICES
	EDUCATION
	ENTERTAINMENT
	FEES_AND_CHARGES
	FOOD_AND_DRINK
	GIFTS_AND_DONATIONS
	HEALTH_AND_FITNESS
	HOME
	INCOME
	INSURANCE
	INVESTMENTS
	KIDS
	LOANS
	PERSONAL_CARE
	PETS
	SHOPPING
	TAXES
	TRANSFER
	TRAVEL
	UNCATEGORIZED
)

type Transaction struct {
	date        time.Time
	description string
	tType       TransactionType
	category    TransactionCategory
	amount      float64
	note        string
}

func printTransaction(transaction Transaction) {
	fmt.Println("Transaction Details:")
	fmt.Println("Date:", transaction.date.Format("02/01/2006"))
	fmt.Println("Description:", transaction.description)
	fmt.Println("Type:", transaction.tType)
	fmt.Println("Category:", transaction.category)
	fmt.Println("Amount:", transaction.amount)
	fmt.Println("Note:", transaction.note)
}
