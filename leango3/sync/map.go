package sync

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	m.Store("cat", "Tom")
	m.Store("mouse", "Jerry")

	val, ok := m.Load("cat")
	if ok {
		fmt.Println(len(val.(string)))
	}
}
