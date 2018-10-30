package main

import (
	"flag"
	"fmt"
	"os"
)

var keyFormat = flag.String("format", "go", "output format (go/hex/base64/proto)")

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("usage:\nmok {flags} {key}")
		flag.PrintDefaults()
		os.Exit(1)
	}

	N("key", []byte(flag.Arg(0))).Expand().Print()
}
