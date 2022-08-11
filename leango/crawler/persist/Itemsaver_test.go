package persist

import (
	"context"
	"encoding/json"
	"leango/crawler/engine"
	"leango/crawler/model"
	"testing"

	"github.com/olivere/elastic/v7"
)

func TestSave(t *testing.T) {
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

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	err = Save(client, index, expected)

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
