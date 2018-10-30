package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: mok {key}")
		os.Exit(1)
	}
	N("key", []byte(os.Args[1])).Expand().Print()
}
