package main

import "fmt"

func main() {

	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	MaoPao(list)
	fmt.Println(list)
}

func MaoPao(list []int) {

	length := len(list)

	isTrans := false

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				isTrans = true
			}
		}

		if isTrans == false {
			return
		}

		isTrans = false

	}

}
