package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer
)

func InitKafka(addr string) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error("init kafka producer failed err:", err)
		return err
	}

	defer client.Close()

	return nil
}

func SendToKafka(data, topic string) error {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed,err:%v,data:%v,topic:%v", err)
		return err
	}

	logs.Debug("send succ, pid:%v offset:%v topic:%v\n", pid, offset, topic)

	return nil
}
