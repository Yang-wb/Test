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

func ParseProfile(contents []byte, name string) engine.ParseResult {

	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}

	profile.Marriage = extractString(contents, ageRe)

	result := engine.ParseResult{
		Requests: nil,
		Items:    []interface{}{profile},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		name := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
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
