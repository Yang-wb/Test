package parser

import (
	"fmt"
	"leango/crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http//album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>"`)
	cityUrlRe = regexp.MustCompile(`href="(http//album.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte, url string) engine.ParseResult {

	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:    string(m[1]),
				Parser: NewProfileParser(string(m[2])),
			})
	}
	fmt.Printf("Matches found: %d\n", len(matches))

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return result
}
