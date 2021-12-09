package processtransation

import (
	"github.com/Jailtons7/imersao-gateway/domain/entity"
	"github.com/Jailtons7/imersao-gateway/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
}

func NewProcessTransaction(repository repository.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDTOInput) (TransactionDTOOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.CreditCardAmount
	cc, invalidCC := entity.NewCreditCard(
		input.CreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear,
		input.CreditCardVerificationCode,
	)
	if invalidCC != nil {
		err := p.Repository.Insert(
			transaction.ID, transaction.AccountID, transaction.Amount, entity.REJECTED,
			invalidCC.Error(),
		)
		if err != nil {
			return TransactionDTOOutput{}, err
		}
		output := TransactionDTOOutput{
			ID:           transaction.ID,
			Status:       entity.REJECTED,
			ErrorMessage: invalidCC.Error(),
		}
		return output, nil
	}
	transaction.SetCreditCard(*cc)
	invalidTransaction := transaction.IsValid()

	if invalidTransaction != nil {
		err := p.Repository.Insert(
			transaction.ID, transaction.AccountID, transaction.Amount, entity.REJECTED,
			invalidTransaction.Error(),
		)
		if err != nil {
			return TransactionDTOOutput{}, err
		}
		output := TransactionDTOOutput{
			ID:           transaction.ID,
			Status:       entity.REJECTED,
			ErrorMessage: invalidTransaction.Error(),
		}
		return output, nil
	}

	err := p.Repository.Insert(
		transaction.ID, transaction.AccountID, transaction.Amount, entity.APPROVED,
		"",
	)
	if err != nil {
		return TransactionDTOOutput{}, err
	}
	output := TransactionDTOOutput{
		ID:           transaction.ID,
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	return output, nil
}
