package transactions

type TransactionType int64

const (
	DEBIT TransactionType = iota
	CREDIT
)
