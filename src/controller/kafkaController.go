package controller

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rayato159/neversitup-kafka-lab/src/config"
	"github.com/Rayato159/neversitup-kafka-lab/src/pkg/queue"

	"github.com/IBM/sarama"
)

type (
	KafkaControllerHandler interface {
		StockProcessor()
		PaymentProcessor()
	}

	kafkaController struct {
		cfg *config.Config
	}
)

func NewKafkaController(cfg *config.Config) KafkaControllerHandler {
	return &kafkaController{
		cfg: cfg,
	}
}

func (k *kafkaController) StockProcessor() {
	worker, err := queue.ConnectConsumer([]string{k.cfg.Kafka.Url})
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer worker.Close()

	consumer, err := worker.ConsumePartition("order", 0, sarama.OffsetOldest)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer consumer.Close()

	fmt.Println("Stock process start...")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	msgCount := 0
	for {
		select {
		case err := <-consumer.Errors():
			fmt.Println(err)
		case msg := <-consumer.Messages():
			msgCount++
			fmt.Printf("Stock Process | Message Count(%d) | Topic(%s) | Message(%s) \n", msgCount, string(msg.Topic), string(msg.Value))
		case <-sigchan:
			fmt.Println("Interrupt is detected")
			return
		}
	}
}

func (k *kafkaController) PaymentProcessor() {
	worker, err := queue.ConnectConsumer([]string{k.cfg.Kafka.Url})
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer worker.Close()

	consumer, err := worker.ConsumePartition("order", 0, sarama.OffsetOldest)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer consumer.Close()

	fmt.Println("Payment process start...")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	msgCount := 0
	for {
		select {
		case err := <-consumer.Errors():
			fmt.Println(err)
		case msg := <-consumer.Messages():
			msgCount++
			fmt.Printf("Payment Process | Message Count(%d) | Topic(%s) | Message(%s) \n", msgCount, string(msg.Topic), string(msg.Value))
		case <-sigchan:
			fmt.Println("Interrupt is detected")
			return
		}
	}
}
