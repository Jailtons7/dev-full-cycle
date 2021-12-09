package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidTransaction(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 937.0
	assert.Nil(t, transaction.IsValid())
}

func TestAmountTransaction(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1000.01
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "You don't have limit for this transaction", err.Error())

	transaction.Amount = 0.0
	err = transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "The amount must be greater than 1", err.Error())
}
