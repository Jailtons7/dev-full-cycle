package repository

type TransactionRepository interface {
	Insert(id string, accountID string, amount float64, status string, errorMessage string) error
}
