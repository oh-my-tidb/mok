package main

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildRawKVKey(t *testing.T) {
	assert := require.New(t)

	keyMode, prefix, err := getKeyPrefix("rawkv", 0)
	assert.NoError(err)
	assert.Equal(KeyModeRaw, *keyMode)

	tests := []struct {
		rawKey   string
		format   string
		expected string
		err      bool
	}{
		{
			rawKey:   "",
			format:   "str",
			expected: "7200000000000000fb",
			err:      false,
		},

		{
			rawKey:   "test",
			format:   "str",
			expected: "7200000074657374ff0000000000000000f7",
			err:      false,
		},
		{
			rawKey:   "testtest",
			format:   "str",
			expected: "7200000074657374ff7465737400000000fb",
			err:      false,
		},
		{
			rawKey:   "74657374",
			format:   "hex",
			expected: "7200000074657374ff0000000000000000f7",
			err:      false,
		},
		{
			rawKey:   "test",
			format:   "invalid",
			expected: "",
			err:      true,
		},
	}

	for _, tt := range tests {
		actualKey, err := buildRawKVKey(prefix, tt.rawKey, tt.format)

		if tt.err && err == nil {
			t.Errorf("expected error, but got nil")
		}

		if !tt.err && err != nil {
			t.Errorf("expected nil error, but got %v", err)
		}

		if hex.EncodeToString(actualKey) != tt.expected {
			t.Errorf("expected key %s, but got %s", tt.expected, hex.EncodeToString(actualKey))
		}
	}
}

func TestParseRawKVKey(t *testing.T) {
	// ❯ ./mok 7200000074657374ff0000000000000000f7
	// "7200000074657374ff0000000000000000f7"
	// └─## decode hex key
	//   └─"r\000\000\000test\377\000\000\000\000\000\000\000\000\367"
	//     ├─## decode mvcc key
	//     │ └─"r\000\000\000test"
	//     │   └─## decode keyspace
	//     │     ├─key mode: rawkv
	//     │     ├─keyspace: 0
	//     │     └─"test"
	//     └─## decode keyspace
	//       ├─key mode: rawkv
	//       ├─keyspace: 0
	//       └─"test\377\000\000\000\000\000\000\000\000\367"
	assert := require.New(t)

	tests := []struct {
		encoded  string
		expected string
	}{
		{
			encoded:  "7200000000000000fb",
			expected: "",
		},
		{
			encoded:  "7200000074657374ff0000000000000000f7",
			expected: "test",
		},
		{
			encoded:  "7200000074657374ff7465737400000000fb",
			expected: "testtest",
		},
	}

	for _, tt := range tests {
		n := N("key", []byte(tt.encoded)).Expand()
		assert.Len(n.variants, 1) // decode hex key
		assert.Equal("decode hex key", n.variants[0].method)
		assert.Len(n.variants[0].children, 1)

		n = n.variants[0].children[0]
		assert.Len(n.variants, 2) // with and without mvcc

		variant := n.variants[0]
		assert.Equal("decode mvcc key", variant.method)
		assert.Len(variant.children, 1)

		n = variant.children[0]
		assert.Len(n.variants, 1)
		variant = n.variants[0]
		assert.Equal("decode keyspace", variant.method)
		assert.Len(variant.children, 3)

		assert.Equal("key_mode", variant.children[0].typ)
		assert.Equal("r", string(variant.children[0].val))

		assert.Equal("keyspace_id", variant.children[1].typ)
		assert.Equal([]byte{0, 0, 0}, variant.children[1].val)

		assert.Equal("raw_key", variant.children[2].typ)
		assert.Equal(tt.expected, string(variant.children[2].val))
	}
}
