package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		v := <-ch
		fmt.Println(v)
	}()

	ch <- 1

	fmt.Println("2")

}
