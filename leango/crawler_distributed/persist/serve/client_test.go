package main

import (
	"leango/crawler/engine"
	"leango/crawler/model"
	"leango/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)

	if err != nil {
		panic(err)
	}

	expected := engine.Item{
		Url:  "",
		Type: "zhenai",
		Id:   "1089",
		Payload: model.Profile{
			Name:       "安静的雪",
			Gender:     "女",
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Marriage:   "离异",
			Education:  "大学本科",
			Occupation: "人事/行政",
			Hokou:      "山东菏泽",
			Xinzou:     "牡羊座",
			House:      "已购房",
			Car:        "未购车",
		},
	}
	result := ""
	err = client.Call("ItemSaverService.Save", expected, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
