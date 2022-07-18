package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

type calculateFunc func(x int, y int) (int, error)

func genCalculator(op operate) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil {

			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {

	op2 := func(x, y int) int {
		return x + y
	}
	x, y := 56, 78
	add := genCalculator(op2)
	result, err := add(x, y)
	fmt.Printf("The result: %d (error: %v)\n", result, err)

}
