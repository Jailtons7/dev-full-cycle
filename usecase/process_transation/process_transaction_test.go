package processtransation

import (
	"testing"
	"time"

	"github.com/Jailtons7/imersao-gateway/domain/entity"
	mock_repository "github.com/Jailtons7/imersao-gateway/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestExecuteInvalidCreditCard(t *testing.T) {
	input := TransactionDTOInput{
		ID:                         "1",
		AccountID:                  "1",
		CreditCardNumber:           "4000000000000000",
		CreditCardName:             "Jailton Santos",
		CreditCardExpirationMonth:  10,
		CreditCardExpirationYear:   time.Now().Year(),
		CreditCardVerificationCode: 123,
		CreditCardAmount:           950,
	}
	expectedOutput := TransactionDTOOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "Invalid credit card number",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.CreditCardAmount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestExecuteRejectedTransaction(t *testing.T) {
	input := TransactionDTOInput{
		ID:                         "1",
		AccountID:                  "1",
		CreditCardNumber:           "6362970000457013",
		CreditCardName:             "Jailton Santos",
		CreditCardExpirationMonth:  10,
		CreditCardExpirationYear:   time.Now().Year(),
		CreditCardVerificationCode: 123,
		CreditCardAmount:           1200,
	}
	expectedOutput := TransactionDTOOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "You don't have limit for this transaction",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.CreditCardAmount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestExecuteApprovedTransaction(t *testing.T) {
	input := TransactionDTOInput{
		ID:                         "1",
		AccountID:                  "1",
		CreditCardNumber:           "6362970000457013",
		CreditCardName:             "Jailton Santos",
		CreditCardExpirationMonth:  10,
		CreditCardExpirationYear:   time.Now().Year(),
		CreditCardVerificationCode: 123,
		CreditCardAmount:           900,
	}
	expectedOutput := TransactionDTOOutput{
		ID:           "1",
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.CreditCardAmount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
