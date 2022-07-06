package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			i++
			out <- i
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {

		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0) //nil channel
	//n := 0
	//hasValue := false

	var valuse []int
	//定时器 到时会向times 发送消息
	times := time.After(10 * time.Second)
	//定时 每秒发送一次
	tick := time.Tick(time.Second)
	for {
		//非阻塞式收取chan
		//谁先获取数据则先打印哪个数据
		var activeWorker chan<- int
		var activeValue int
		if len(valuse) > 0 {
			activeWorker = worker
			activeValue = valuse[0]
		}
		select {
		case n := <-c1:
			//hasValue = true
			valuse = append(valuse, n)
		case n := <-c2:
			//hasValue = true
			valuse = append(valuse, n)
		case activeWorker <- activeValue:
			valuse = valuse[1:]
		case <-time.After(800 * time.Millisecond):
			//每两次接收时间差800毫秒还没有接收到数据
			fmt.Println("time.out")
		case <-tick:
			fmt.Println("queue len=", len(valuse))

		case <-times:
			//运行10秒结束
			fmt.Println("bye")
			return
		}
	}

}
