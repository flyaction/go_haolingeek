package main

import "fmt"

// 示例4。
func getIntChannew() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

func main() {
	intChan2 := getIntChannew()
	for elem := range intChan2 {
		fmt.Printf("The element in intChan2: %v\n", elem)
	}
}
