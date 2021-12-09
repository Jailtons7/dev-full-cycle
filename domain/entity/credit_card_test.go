package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("4000000000000000", "José da Silva", 12, 2024, 123)
	assert.Equal(t, "Invalid credit card number", err.Error())

	_, err = NewCreditCard("6362970000457013", "José da Silva", 12, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("6362970000457013", "José da Silva", 13, 2024, 123)
	assert.Equal(t, "Invalid expiration month", err.Error())

	_, err = NewCreditCard("6362970000457013", "José da Silva", 0, 2024, 123)
	assert.Equal(t, "Invalid expiration month", err.Error())

	_, err = NewCreditCard("6362970000457013", "José da Silva", 5, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)
	_, err := NewCreditCard("6362970000457013", "José da Silva", 5, lastYear.Year(), 123)
	assert.Equal(t, "Invalid expiration year", err.Error())

	nextYear := time.Now().AddDate(+1, 0, 0)
	_, err = NewCreditCard("6362970000457013", "José da Silva", 5, nextYear.Year(), 123)
	assert.Nil(t, err)
}
