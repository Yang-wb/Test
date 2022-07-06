package main

import (
	"fmt"
	"leango/downloader/moke"
	"leango/downloader/real"
)

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

type Retriever interface {
	Get(url string) string
}

//新增
type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

//组合接口
type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "http://www.imooc.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {

	retriever := &moke.Retriever{Contest: "this is a fake imooc.com"}
	inspect(retriever)
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *moke.Retriever:
		fmt.Println("Contents:", v.Contest)
	case *real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
}
