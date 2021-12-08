package entity

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestCreditCardValidation(t *testing.T) {
	_, err := NewCreditCard("1234567891234567", "Tião Carneiro", 12, 2028, 456)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("5355119872960768", "Tião Carneiro", 12, 2023, 698)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("5355119872960768", "Tião Carneiro", 13, 2023, 698)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5355119872960768", "Tião Carneiro", 0, 2023, 698)
	assert.Equal(t,"invalid expiration month", err.Error())

	_, err = NewCreditCard("5355119872960768", "Tião Carneiro", 5, 2023, 698)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1,0,0)
	_, err := NewCreditCard("5355119872960768", "Tião Carneiro", 12, lastYear.Year(), 123)
	assert.Equal(t, "invalid expiration year", err.Error())
}