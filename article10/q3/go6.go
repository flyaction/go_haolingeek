package main

import (
	"fmt"
	"time"
)

func producemy(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i, "--", len(p))
	}
}
func consumermy(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v, "--", len(c))
	}
}
func main() {
	ch := make(chan int, 2)
	go producemy(ch)
	go consumermy(ch)

	time.Sleep(1 * time.Second)
}
