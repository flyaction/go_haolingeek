package main

import "fmt"

func main() {
	s := make([]int, 0)
	fmt.Printf("len(s) = %d, cap(s)=%d, addr=%p\n", len(s), cap(s), s)
	for i := 1; i <= 10; i++ {
		s = append(s, 1)
		fmt.Printf("i:%d, len(s) = %d, cap(s)=%d, addr=%p\n", i, len(s), cap(s), s)
	}
}
