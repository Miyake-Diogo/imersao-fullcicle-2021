package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTransaction_IsValid(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1.0
	transaction.Status = "success"
	assert.Nil(t, transaction.IsValid())
}

func TestTransaction_IsNotValidWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1001.0
	transaction.Status = "approved"
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "you dont have limit for this transaction", err.Error())
}

func TestTransaction_IsNotValidWithAmountLessThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0.0
	transaction.Status = "approved"
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "the amount must be greater than 1", err.Error())
}

func TestTransaction_IsNotValidWithStatusNotSuccess(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1.0
	transaction.Status = "rejected"
	err := transaction.IsValid()
	assert.Equal(t, "transaction is rejected check with the bank", err.Error())
}
