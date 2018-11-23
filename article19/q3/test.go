package main

import (
	"fmt"
)

func tryRecover() {

	defer func() {
		r := recover()

		//fmt.Println(r)
		//fmt.Printf("panic: %v\n", r)
		//fmt.Printf("panic: %v\n", r.(error))

		//这是根据郝林的改的
		if r != nil {
			fmt.Println("Error occurred:", r)
		} else {
			//panic(r)
			panic(fmt.Sprintf("I don't know what to do :%v", r))
		}

		//原先的
		//if err, ok := r.(error); ok {
		//	fmt.Println("Error occurred:", err)
		//} else {
		//	//panic(r)
		//	panic(fmt.Sprintf("I don't know what to do :%v", r))
		//}
	}()

	//panic(errors.New("this is an error"))

	b := 0
	a := 5 / b
	fmt.Println(a)

	//panic(123)

}
func main() {

	tryRecover()
}
