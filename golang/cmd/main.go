package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"github.com/Miyake-Diogo/imersao-fullcicle-2021-gateway/adapter/factory"
	"github.com/Miyake-Diogo/imersao-fullcicle-2021-gateway/adapter/presenter/transaction"
	"github.com/Miyake-Diogo/imersao-fullcicle-2021-gateway/usecase/process_transaction"
	"github.com/Miyake-Diogo/imersao-fullcicle-2021-gateway/adapter/broker/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// configuração Database

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	// configuração repository
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()
	// ConfigMap Producer
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
	}
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	// producer
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)
	// ConfigMap Consumer
	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "goapp",
		"client.id":         "goapp",
	}

	// topic
	topics := []string{"transactions"}
	// consumer
	consumer := kafka.NewConsumer(configMapConsumer, topics)
	go consumer.Consume(msgChan)
	// usecase
	usecase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")

	for msg := range msgChan {
		var input process_transaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}
}
