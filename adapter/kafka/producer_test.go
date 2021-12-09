package kafka

import (
	"testing"

	"github.com/Jailtons7/imersao-gateway/adapter/presenter/transaction"
	"github.com/Jailtons7/imersao-gateway/domain/entity"
	processtransation "github.com/Jailtons7/imersao-gateway/usecase/process_transation"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
)

func TestProducerPublish(t *testing.T) {
	expectedOutput := processtransation.TransactionDTOOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "You don't have limit for this transaction",
	}
	// outputJson, _ := json.Marshal(expectedOutput)
	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}
	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())
	err := producer.Publish(expectedOutput, []byte("1"), "test")
	assert.Nil(t, err)
}
