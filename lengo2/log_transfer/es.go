package main

import (
	"fmt"

	"gopkg.in/olivere/elastic.v2"
)

type LogMessage struct {
	App     string
	Topic   string
	Message string
}

var (
	esClient *elastic.Client
)

func initEs(addr string) (err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(addr))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}

	esClient = client

	fmt.Println("conn es succ")
	return nil

	//tweet := Tweet{User: "olivere", Message: "Take Five"}
	//_, err = client.Index().
	//	Index("twitter").
	//	Type("tweet").
	//	Id("1").
	//	BodyJson(tweet).
	//	Do()
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//	return
	//}
	//
	//fmt.Println("insert succ")
}

func sendToES(topic string, data []byte) error {
	msg := &LogMessage{}
	msg.Topic = topic
	msg.Message = string(data)

	_, err := esClient.Index().
		Index(topic).
		Type(topic).
		//Id("1").
		BodyJson(msg).
		Do()
	if err != nil {
		// Handle error
		//panic(err)
		return err
	}
	return nil
}
