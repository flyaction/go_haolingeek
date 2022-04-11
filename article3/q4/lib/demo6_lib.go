package lib

import (
	"os"
	in "study/article3/q4/lib/internal" // in => alias name
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
