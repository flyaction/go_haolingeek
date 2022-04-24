package main

import "fmt"

func main() {
	aMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	k := "two"

	//v,_:= aMap[k]
	//
	//if v > 0 {
	//	fmt.Println("%q--%d",k,v)
	//}else{
	//	fmt.Println("not found")
	//}

	v, ok := aMap[k]

	if ok {
		fmt.Printf("%q--%d\n", k, v)
	} else {
		fmt.Println("not found\n")
	}

	//数组可以哈
	var test = map[[2]int]int{
		[2]int{1, 2}: 1,
		[2]int{2, 1}: 2,
	}
	fmt.Printf("%v\n", test)

}
