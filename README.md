# mok
mok: master of keys

A tool for decoding keys for TiDB projects.

https://github.com/pingcap/tidb

https://github.com/tikv/tikv

https://github.com/pingcap/pd

## Install

Download the [latest release](https://github.com/disksing/mok/releases) then unzip the binary to your PATH:

```
    $ unzip mok-vX.X-OS-ARCH.zip
    $ mv mok-vX.X-OS-ARCH/mok /usr/local/bin/
```

You can also manually build it by yourself:

```
    $ git clone https://github.com/disksing/mok.git
    $ cd mok
    $ GO111MODULE=on go build -o /usr/local/bin/mok
```

## Usage

Parse a given key:
```
$ mok 7A7480000000000007FF8F5F728000000000FF083BBA0000000000FAFA6C400A6673FFFE
"7A7480000000000007FF8F5F728000000000FF083BBA0000000000FAFA6C400A6673FFFE"
├─hex
│ └─"zt\200\000\000\000\000\000\007\377\217_r\200\000\000\000\000\377\010;\272\000\000\000\000\000\372\372l@\nfs\377\376"
│   └─rocksdb
│     └─"t\200\000\000\000\000\000\007\377\217_r\200\000\000\000\000\377\010;\272\000\000\000\000\000\372\372l@\nfs\377\376"
│       └─comparable
│         ├─"t\200\000\000\000\000\000\007\217_r\200\000\000\000\000\010;\272"
│         │ └─rowkey
│         │   ├─table: 1935
│         │   └─row: 539578
│         └─ts: 401875853330087937 (2018-07-31 18:58:38.819 +0800 CST)
└─base64
  └─"\354\016\370\363M4\323M4\323M4\323\261E\360^E\357o4\323M4\323M4\024]<\334\020@\323M4\323M4\323A@\024\016\202\343M\000\353\256\367\024QD"
```

Build a key for a given record:
```
$ ./mok --table-id 43 --row-value 81934
built key: 7480000000000000FF2B5F728000000000FF01400E0000000000FA
```

Build a key for a given index:
```
$ ./mok --table-id 43 --index-id 50 --index-value 81934
built key: 7480000000000000FF2B5F698000000000FF0000328000000000FF01400E0000000000FA
```

Build a key for a given record under given keyspace
```
$ ./mok --keyspace-id 255 --table-id 43 --row-value 81934
built key: 780000FF74800000FF000000002B5F7280FF0000000001400E00FE
```

Build a RawKV key for a given raw key under given keyspace
```
$ ./mok --keyspace-id 0 --key-mode rawkv --raw-key test
built key: 7200000074657374FF0000000000000000F7

$ ./mok --keyspace-id 1 --key-mode rawkv --raw-key ''
built key: 7200000100000000FB
```

## TODO

- [x] build keys
- [ ] decode tidb meta keys
