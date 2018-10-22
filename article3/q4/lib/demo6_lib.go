package lib

import (
	in "haolingeek/article3/q4/lib/internal"
	"os"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
