package main

import "fmt"

func main() {

	fmt.Println(1)

	s1 := []int{0, 1, 2, 3, 4}
	e5 := s1[5]
	_ = e5
}

//panic: runtime error: index out of range   类型  runtime error  说明  index out of range
