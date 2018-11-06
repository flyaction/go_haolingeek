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
		fmt.Println("%q--%d", k, v)
	} else {
		fmt.Println("not found")
	}

}
