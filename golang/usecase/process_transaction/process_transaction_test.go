package process_transaction

import (
	"testing"
	"time"
	"github.com/golang/mock/gomock"
	"github.com/Miyake-Diogo/imersao-fullcicle-2021-gateway/domain/entity"
	mock_repository "github.com/Miyake-Diogo/imersao-fullcicle-2021-gateway/domain/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	input := TransactionDtoInput{
		ID:		  "1",
		AccountID: "1",
		CreditCardNumber: "1234567890123456",
		CreditCardName: "John Doe",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear: time.Now().Year() + 1,
		CreditCardCVV: 123,
		Amount: 145.50,
	}
	expectedOutput := TransactionDtoOutput{
		ID:		  "1",
		Status: entity.REJECTED,
		ErrorMessage: "invalid credit card number",

	}
    ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.
		EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output , err := usecase.Execute(input)
	assert.Equal(t, expectedOutput, output)
	assert.Nil(t, err)
}

func TestProcessTransaction_ExecuteRejectTransaction(t *testing.T) {
	input := TransactionDtoInput{
		ID:		  "1",
		AccountID: "1",
		CreditCardNumber: "5355119872960768",
		CreditCardName: "John Doe",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear: time.Now().Year() + 1,
		CreditCardCVV: 123,
		Amount: 4145.50,
	}
	expectedOutput := TransactionDtoOutput{
		ID:		  "1",
		Status: entity.REJECTED,
		ErrorMessage: "Invalid credit card number",

	}
    ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.
		EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output , err := usecase.Execute(input)
	assert.Equal(t, expectedOutput, output)
	assert.Nil(t, err)
}