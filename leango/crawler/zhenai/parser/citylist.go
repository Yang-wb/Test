package parser

import (
	"fmt"
	"leango/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http//www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>"`

func ParseCityList(contents []byte, url string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:    string(m[1]),
				Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
			})
	}
	fmt.Printf("Matches found: %d\n", len(matches))

	return result
}
