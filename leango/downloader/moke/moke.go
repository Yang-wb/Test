package moke

import "fmt"

type Retriever struct {
	Contest string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever:{Contents=%s}", r.Contest)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contest = form["contest"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contest
}
