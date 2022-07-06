package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func() { // race condition !
			for {
				a[i]++
				//交出控制权
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Microsecond)
	fmt.Println(a)
}
