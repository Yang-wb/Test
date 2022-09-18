package main

import (
	"strings"
	"sync"

	"github.com/astaxie/beego/logs"

	"github.com/Shopify/sarama"
)

type KafkaClient struct {
	client sarama.Consumer
	addr   string
	topic  string
	wg     sync.WaitGroup
}

var (
	kafkaClient *KafkaClient
)

func initKafka(addr string, topic string) (err error) {

	kafkaClient = &KafkaClient{}

	consumer, err := sarama.NewConsumer(strings.Split(addr, ","), nil)
	if err != nil {
		logs.Error("Failed to start consumer: %s", err)
		return
	}

	kafkaClient.client = consumer
	kafkaClient.addr = addr
	kafkaClient.topic = topic

	//time.Sleep(time.Hour)
	//wg.Wait()
	consumer.Close()

	return nil
}
