package main

import (
	"fmt"
)

func main() {
	str := "Go爱好者"
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))   //%q
	fmt.Printf("  => runes(hex): %x\n", []rune(str))    // %x 16进制
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str)) //
}
