package main

import (
	"fmt"
	"time"
)

func produce(p chan<- int) { //只读  生产者
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}
func consumer(c <-chan int) { //只写 消费者
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}
func main() {
	ch := make(chan int)

	go produce(ch)
	go consumer(ch)

	time.Sleep(1 * time.Second)
}
