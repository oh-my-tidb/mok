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
│ └─"zt\x80\x00\x00\x00\x00\x00\a\xff\x8f_r\x80\x00\x00\x00\x00\xff\b;\xba\x00\x00\x00\x00\x00\xfa\xfal@\nfs\xff\xfe"
│   └─rocksdb
│     └─"t\x80\x00\x00\x00\x00\x00\a\xff\x8f_r\x80\x00\x00\x00\x00\xff\b;\xba\x00\x00\x00\x00\x00\xfa\xfal@\nfs\xff\xfe"
│       └─comparable
│         ├─"t\x80\x00\x00\x00\x00\x00\a\x8f_r\x80\x00\x00\x00\x00\b;\xba"
│         │ └─rowkey
│         │   ├─table: 1935
│         │   └─row: 539578
│         └─ts: 401875853330087937 (2018-07-31 18:58:38.819 +0800 CST)
└─base64
  └─"\xec\x0e\xf8\xf3M4\xd3M4\xd3M4ӱE\xf0^E\xefo4\xd3M4\xd3M4\x14]<\xdc\x10@\xd3M4\xd3M4\xd3A@\x14\x0e\x82\xe3M\x00\xeb\xae\xf7\x14QD"
```

## TODO

- [ ] build keys
- [x] setting output format of keys
- [ ] decode tidb meta keys
- [ ] decode index values
- [ ] support url format
