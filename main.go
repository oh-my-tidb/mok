package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pingcap/tidb/util/codec"
)

var keyFormat = flag.String("format", "proto", "output format (go/hex/base64/proto)")
var tableID = flag.Int64("table-id", 0, "table ID")
var indexID = flag.Int64("index-id", 0, "index ID")
var rowValue = flag.String("row-value", "", "row value")
var indexValue = flag.String("index-value", "", "index value")

func main() {
	flag.Parse()

	if flag.NArg() == 1 { // Decode the given key.
		n := N("key", []byte(flag.Arg(0)))
		n.Expand().Print()
	} else if flag.NArg() == 0 { // Build a key with given flags.
		key := []byte{'t'}
		key = codec.EncodeInt(key, *tableID)
		if *tableID == 0 {
			fmt.Println("table ID shouldn't be 0")
			os.Exit(1)
		}

		if *indexID == 0 {
			if *rowValue != "" {
				key = append(key, []byte("_r")...)
				rowValueInt, err := strconv.ParseInt(*rowValue, 10, 64)
				if err != nil {
					fmt.Printf("invalid row value: %s\n", *rowValue)
					os.Exit(1)
				}
				key = codec.EncodeInt(key, rowValueInt)
			}
		} else {
			key = append(key, []byte("_i")...)
			key = codec.EncodeInt(key, *indexID)
			if *indexValue != "" {
				indexValueInt, err := strconv.ParseInt(*indexValue, 10, 64)
				if err != nil {
					fmt.Printf("invalid index value: %s\n", *indexValue)
					os.Exit(1)
					key = codec.EncodeInt(key, indexValueInt)
				}
			}
		}

		key = codec.EncodeBytes([]byte{}, key)
		fmt.Printf("built key: %s\n", strings.ToUpper(hex.EncodeToString(key)))
	} else {
		fmt.Println("usage:\nmok {flags} [key]")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
