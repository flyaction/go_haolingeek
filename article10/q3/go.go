package main

import "fmt"

func main() {

	var ch1 = make(chan int, 3)

	fmt.Println(len(ch1))

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	fmt.Println(len(ch1))
	elem := <-ch1
	fmt.Println(len(ch1))
	fmt.Println(elem)

}
