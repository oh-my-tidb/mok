package main

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/pingcap/tidb/util/codec"
)

type Node struct {
	typ      string // "key", "table_id", "row_id", "index_id", "index_values", "index_value", "ts"
	val      []byte
	variants []*Variant
}

type Variant struct {
	method   string
	children []*Node
}

func N(t string, v []byte) *Node {
	return &Node{typ: t, val: v}
}

func (n *Node) String() string {
	switch n.typ {
	case "key", "index_values":
		switch *keyFormat {
		case "hex":
			return `"` + strings.ToUpper(hex.EncodeToString(n.val)) + `"`
		case "base64":
			return `"` + base64.StdEncoding.EncodeToString(n.val) + `"`
		case "proto":
			return `"` + formatProto(string(n.val)) + `"`
		default:
			return fmt.Sprintf("%q", n.val)
		}
	case "keyspace_id":
		tmp := []byte{'\x00'}
		t := append(tmp, n.val...)
		id := binary.BigEndian.Uint32(t)
		return fmt.Sprintf("keyspace: %v", id)
	case "table_id":
		_, id, _ := codec.DecodeInt(n.val)
		return fmt.Sprintf("table: %v", id)
	case "row_id":
		_, id, _ := codec.DecodeInt(n.val)
		return fmt.Sprintf("row: %v", id)
	case "index_id":
		_, id, _ := codec.DecodeInt(n.val)
		return fmt.Sprintf("index: %v", id)
	case "index_value":
		_, d, _ := codec.DecodeOne(n.val)
		s, _ := d.ToString()
		return fmt.Sprintf("kind: %v, value: %v", indexTypeToString[d.Kind()], s)
	case "ts":
		_, ts, _ := codec.DecodeUintDesc(n.val)
		return fmt.Sprintf("ts: %v (%v)", ts, GetTimeFromTS(uint64(ts)))
	}
	return fmt.Sprintf("%v:%q", n.typ, n.val)
}

func (n *Node) Expand() *Node {
	for _, fn := range rules {
		if t := fn(n); t != nil {
			for _, child := range t.children {
				child.Expand()
			}
			n.variants = append(n.variants, t)
		}
	}
	return n
}

func (n *Node) Print() {
	fmt.Println(n.String())
	for i, t := range n.variants {
		t.PrintIndent("", i == len(n.variants)-1)
	}
}

func (n *Node) PrintIndent(indent string, last bool) {
	indent = printIndent(indent, last)
	fmt.Println(n.String())
	for i, t := range n.variants {
		t.PrintIndent(indent, i == len(n.variants)-1)
	}
}

func (v *Variant) PrintIndent(indent string, last bool) {
	indent = printIndent(indent, last)
	fmt.Printf("## %s\n", v.method)
	for i, c := range v.children {
		c.PrintIndent(indent, i == len(v.children)-1)
	}
}

func printIndent(indent string, last bool) string {
	if last {
		fmt.Print(indent + "└─")
		return indent + "  "
	}
	fmt.Print(indent + "├─")
	return indent + "│ "
}
