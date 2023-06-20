package transactions

import (
	"time"
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
