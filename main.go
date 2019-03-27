package main

import (
	"fmt"
	"os"

	"github.com/jiangxin/mytest/hello"
)

func main() {
	var msg string
	if len(os.Args) > 1 {
		msg = os.Args[1]
	}
	fmt.Println(hello.Hello(msg))
}
