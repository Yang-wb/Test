package parser

import (
	"leango/crawler/engine"
	"leango/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d])+岁</td>`)

var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<])+</td>`)

//猜你喜欢
var guessRe = regexp.MustCompile(``)

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {

	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}

	profile.Marriage = extractString(contents, ageRe)

	result := engine.ParseResult{
		Requests: nil,
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:        url,
			ParserFunc: ProfileParser(string(m[2])),
		})
	}

	return result
}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url, name)
	}
}
