package main

import (
	"fmt"
	"github.com/projectdiscovery/blackrock"
)

func main() {
	//test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	bl := blackrock.New(10, 3)
	for i := 0; i < 10; i++ {
		idx := bl.Shuffle(int64(i))
		unIdx := bl.UnShuffle(idx)
		fmt.Println(fmt.Sprintf("{bl.A:%v}-{bl.B:%v}-{bl.Rounds:%v}-{bl.Range:%v}-{i:%v}-{idx:%v}-{unIdx:%v}", bl.A, bl.B, bl.Rounds, bl.Range, i, idx, unIdx))
	}
}
