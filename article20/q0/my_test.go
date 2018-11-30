package main

import (
	"testing"
)

var mygreeting string
var myerr error

func Test_hello(t *testing.T) {
	tests := []struct{ name, outs string }{
		{"xingdong", "Hello,xingdong"},
		{"", "Hello,"},
	}

	for _, tt := range tests {
		mygreeting, myerr = hello(tt.name)

	}

}
