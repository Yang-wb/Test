package parser

import (
	"leango/crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	//fmt.Printf()
	result := ParseCityList(contents)

	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d request; but had %d", resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d request; but had %d", resultSize, len(result.Items))
	}
}
