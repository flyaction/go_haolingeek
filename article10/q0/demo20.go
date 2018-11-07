package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 0
	ch1 <- 1
	ch1 <- 3
	//ch1 <- 4
	fmt.Println(len(ch1))

	elem1 := <-ch1 //先进先出
	fmt.Println(len(ch1))

	elem1 = <-ch1 //先进先出
	fmt.Println(len(ch1))
	elem1 = <-ch1 //先进先出
	fmt.Println(len(ch1))

	//ch1 <- 4

	//elem1 = <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n", elem1)

}
