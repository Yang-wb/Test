package main

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

func run() (err error) {
	consumer := kafkaClient.client

	partitionList, err := consumer.Partitions(kafkaClient.topic)
	if err != nil {
		logs.Error("Failed to get the list of partitions: ", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(kafkaClient.topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			logs.Error("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer) {
			kafkaClient.wg.Add(1)
			for msg := range pc.Messages() {
				logs.Debug("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				err = sendToES(kafkaClient.topic, msg.Value)
				if err != nil {
					logs.Warn("send to se failed, err:%v", err)
				}
			}
			kafkaClient.wg.Done()
		}(pc)
	}

	kafkaClient.wg.Wait()

	return nil
}
