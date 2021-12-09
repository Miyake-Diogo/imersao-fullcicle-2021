package kafka

import (
	//"encoding/json"
	"testing"

	"github.com/Miyake-Diogo/imersao-fullcicle-2021-gateway/adapter/presenter/transaction"
	"github.com/Miyake-Diogo/imersao-fullcicle-2021-gateway/domain/entity"
	"github.com/Miyake-Diogo/imersao-fullcicle-2021-gateway/usecase/process_transaction"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
)


func TestProducerPublish(t *testing.T) {
	expectedOutput := process_transaction.TransactionDtoOutput{
		ID: "1",
		Status: entity.REJECTED,
		ErrorMessage: "you dont have limit for this transaction",
	}

	// outputJson, _ := json.Marshal(expectedOutput)
	configMap := ckafka.ConfigMap{
		//"bootstrap.servers": "localhost"
		"test.mock.num.brokers": 3,
	
	}

	producer := NewKafkaProducer(&configMap, 
		transaction.NewTransactionKafkaPresenter())
	err := producer.Publish(expectedOutput, []byte("1"), "test")
	assert.Nil(t, err)

}