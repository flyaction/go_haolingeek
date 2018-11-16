package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			fn()
		}(i)
	}

	time.Sleep(2 * time.Second)

	fmt.Println()
	var count uint32
	n := atomic.LoadUint32(&count)
	fmt.Println(n)

	atomic.AddUint32(&count, 1)
	n = atomic.LoadUint32(&count)
	fmt.Println(n)

	atomic.AddUint32(&count, 1)
	n = atomic.LoadUint32(&count)
	fmt.Println(n)

}
