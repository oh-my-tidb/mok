package main

import (
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pingcap/tidb/util/codec"
)

var keyFormat = flag.String("format", "proto", "output format (go/hex/base64/proto)")
var NullSpaceID = int64(0xffffffff)
var keyspaceID = flag.Int64("keyspace-id", NullSpaceID, "keyspace ID")
var tableID = flag.Int64("table-id", 0, "table ID")
var indexID = flag.Int64("index-id", 0, "index ID")
var rowValue = flag.String("row-value", "", "row value")
var indexValue = flag.String("index-value", "", "index value")

func getKeyPrefix(keyspaceID int64) ([]byte, error) {
	if keyspaceID == NullSpaceID {
		return []byte{'t'}, nil
	}
	if keyspaceID > 0xffffff {
		return nil, fmt.Errorf("invalid keyspace value: %d", keyspaceID)
	}
	var prefix [4]byte
	binary.BigEndian.PutUint32(prefix[:], uint32(keyspaceID))
	prefix[0] = 'x'
	return append(prefix[:], 't'), nil
}

func main() {
	flag.Parse()

	if flag.NArg() == 1 { // Decode the given key.
		n := N("key", []byte(flag.Arg(0)))
		n.Expand().Print()
	} else if flag.NArg() == 0 { // Build a key with given flags.
		key, err := getKeyPrefix(*keyspaceID)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

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
