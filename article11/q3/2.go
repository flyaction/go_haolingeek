package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 1)
	time.AfterFunc(3*time.Second, func() {
		close(ch1)
	})

loop:
	for {
		select {
		case _, ok := <-ch1:
			if !ok {
				fmt.Println("111111")
				break loop
			}
			fmt.Println("ch1")

		}
	}
	fmt.Println("End!")
}
