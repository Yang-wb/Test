package main

import (
	"fmt"
	"regexp"
)

const text = `My email is ccmouse@gmail.com
email is abc@def.org
email2 is kkkk@qq.com
email2 is ddd@abc.com.cn
`

func main() {
	//re, err := regexp.Compile("ccmous@gmail.com")
	//re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	//match := re.FindString(text)
	//fmt.Println(match)

	// -1 所有
	//allString := re.FindAllString(text, -1)
	//fmt.Println(allString)

	//提取
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	subMatch := re.FindAllStringSubmatch(text, -1)
	for _, m := range subMatch {
		fmt.Println(m)
	}
}
