# mok
mok: master of keys

A tool for decoding keys for TiDB projects.

https://github.com/pingcap/tidb

https://github.com/tikv/tikv

https://github.com/pingcap/pd

## Install
```
    $ go get github.com/disksing/mok
    $ go install github.com/disksing/mok
```

## Usage
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

## TODO

- [ ] build keys
- [x] setting output format of keys
- [ ] decode tidb meta keys
- [x] decode index values
