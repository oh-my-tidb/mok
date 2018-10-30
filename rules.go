package main

import (
	"encoding/base64"
	"encoding/hex"
	"strconv"
	"strings"
)

type Rule func(*Node) *Variant

var rules = []Rule{
	DecodeHex,
	DecodeComparableKey,
	DecodeRocksDBKey,
	DecodeTablePrefix,
	DecodeTableRow,
	DecodeTableIndex,
	UnQuote,
	DecodeBase64,
}

func DecodeHex(n *Node) *Variant {
	if n.typ != "key" {
		return nil
	}
	decoded, err := hex.DecodeString(string(n.val))
	if err != nil {
		return nil
	}
	return &Variant{
		method:   "hex",
		children: []*Node{N("key", decoded)},
	}
}

func DecodeComparableKey(n *Node) *Variant {
	if n.typ != "key" {
		return nil
	}
	b, decoded, err := DecodeBytes(n.val, nil)
	if err != nil {
		return nil
	}
	children := []*Node{N("key", decoded)}
	switch len(b) {
	case 0:
	case 8:
		children = append(children, N("ts", b))
	default:
		return nil
	}
	return &Variant{
		method:   "comparable",
		children: children,
	}
}

func DecodeRocksDBKey(n *Node) *Variant {
	if n.typ != "key" {
		return nil
	}
	if len(n.val) > 0 && n.val[0] == 'z' {
		return &Variant{
			method:   "rocksdb",
			children: []*Node{N("key", n.val[1:])},
		}
	}
	return nil
}

func DecodeTablePrefix(n *Node) *Variant {
	if n.typ == "key" && len(n.val) == 9 && n.val[0] == 't' {
		return &Variant{
			method:   "table_prefix",
			children: []*Node{N("table_id", n.val[1:])},
		}
	}
	return nil
}

func DecodeTableRow(n *Node) *Variant {
	if n.typ == "key" && len(n.val) == 19 && n.val[0] == 't' && n.val[9] == '_' && n.val[10] == 'r' {
		return &Variant{
			method:   "rowkey",
			children: []*Node{N("table_id", n.val[1:9]), N("row_id", n.val[11:])},
		}
	}
	return nil
}

func DecodeTableIndex(n *Node) *Variant {
	if n.typ == "key" && len(n.val) >= 19 && n.val[0] == 't' && n.val[9] == '_' && n.val[10] == 'i' {
		return &Variant{
			method:   "indexkey",
			children: []*Node{N("table_id", n.val[1:9]), N("index_id", n.val[11:19])},
		}
	}
	return nil
}

func UnQuote(n *Node) *Variant {
	if n.typ != "key" {
		return nil
	}
	if strings.Index(string(n.val), `\`) == -1 {
		return nil
	}
	s, err := strconv.Unquote(`'` + string(n.val) + `'`)
	if err != nil {
		return nil
	}
	return &Variant{
		method:   "unquote",
		children: []*Node{N("key", []byte(s))},
	}
}

func DecodeBase64(n *Node) *Variant {
	if n.typ != "key" {
		return nil
	}
	s, err := base64.StdEncoding.DecodeString(string(n.val))
	if err != nil {
		return nil
	}
	return &Variant{
		method:   "base64",
		children: []*Node{N("key", []byte(s))},
	}
}
