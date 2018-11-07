package main

import (
	"fmt"
	"time"
)

func main() {
	i := 3
	go func(a int) {
		fmt.Println(a)
		fmt.Println("1")
	}(i)

	fmt.Println("2")

	time.Sleep(1 * time.Second)
}
