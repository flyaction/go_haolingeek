package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0

	trigger := func(i int, fn func()) {
		for {
			if n := count; n == i {
				fn()
				count++
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}

	trigger(10, func() {})

}
