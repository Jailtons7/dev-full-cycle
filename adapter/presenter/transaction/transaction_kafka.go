package transaction

import (
	"encoding/json"

	processtransation "github.com/Jailtons7/imersao-gateway/usecase/process_transation"
)

type KafkaPresenter struct {
	ID           string `json:"id"`
	Status       string `json:status`
	ErrorMessage string `json:error_message`
}

func NewTransactionKafkaPresenter() *KafkaPresenter {
	return &KafkaPresenter{}
}

func (t *KafkaPresenter) Bind(input interface{}) error {
	t.ID = input.(processtransation.TransactionDTOOutput).ID
	t.Status = input.(processtransation.TransactionDTOOutput).Status
	t.ErrorMessage = input.(processtransation.TransactionDTOOutput).ErrorMessage
	return nil
}

func (t *KafkaPresenter) Show() ([]byte, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return j, nil
}
