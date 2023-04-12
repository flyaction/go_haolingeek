package main

import "fmt"

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}

	fmt.Println(QuickSort(arr))
}

func QuickSort(arr []int) []int {

	length := len(arr)
	if length < 2 {
		return arr
	}

	base := arr[0]

	left := make([]int, 0, 0)
	right := make([]int, 0, 0)
	middle := make([]int, 0, 0)
	middle = append(middle, base)

	for i := 1; i < length; i++ {
		if arr[i] > base {
			right = append(right, arr[i])
		} else {
			left = append(left, arr[i])
		}
	}

	right = QuickSort(right)
	left = QuickSort(left)

	return append(append(left, middle...), right...)

}
