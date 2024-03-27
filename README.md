# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Parse string/[]byte to simple go types ([]interface{}, int64, string, etc)
     json.Unmarshal          21343 ns/op        17744 B/op          334 allocs/op
       oj.Parse              10203 ns/op         5691 B/op          364 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal          15530 ns/op        19682 B/op          451 allocs/op
 simdjson.Parse              29447 ns/op       140635 B/op          361 allocs/op
    gjson.ParseBytes         17849 ns/op        20041 B/op          175 allocs/op
    sonic.Unmarshal           9419 ns/op        23630 B/op          168 allocs/op

    sonic ███████████████▊ 2.27
       oj ██████████████▋ 2.09
 jsoniter █████████▌ 1.37
    gjson ████████▎ 1.20
     json ▓▓▓▓▓▓▓ 1.00
 simdjson █████  0.72
 fastjson >>> not supported <<<

Validate string/[]byte
     json.Valid               6402 ns/op            0 B/op            0 allocs/op
       oj.Validate            2272 ns/op            0 B/op            0 allocs/op
 fastjson.Valid               2524 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               4619 ns/op         2187 B/op          100 allocs/op
 simdjson.Validate           18364 ns/op       118632 B/op           11 allocs/op
    gjson.Validate            2338 ns/op            0 B/op            0 allocs/op
    sonic.Valid               1979 ns/op            0 B/op            0 allocs/op

    sonic ██████████████████████▋ 3.23
       oj ███████████████████▋ 2.82
    gjson ███████████████████▏ 2.74
 fastjson █████████████████▊ 2.54
 jsoniter █████████▋ 1.39
     json ▓▓▓▓▓▓▓ 1.00
 simdjson ██▍ 0.35

Iterate tokens in a string/[]byte
     json.Decode             40356 ns/op        22568 B/op         1175 allocs/op
       oj.Tokenize            4624 ns/op         1976 B/op          156 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             16274 ns/op        20360 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Decode             10265 ns/op        24065 B/op          170 allocs/op

       oj █████████████████████████████████████████████████████████████  8.73
    sonic ███████████████████████████▌ 3.93
 jsoniter █████████████████▎ 2.48
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Unmarshal string/[]byte to a struct
     json.Unmarshal          19986 ns/op         2576 B/op           74 allocs/op
       oj.Unmarshal          23665 ns/op         9354 B/op          456 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal           6958 ns/op         3141 B/op          170 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Unmarshal           4702 ns/op         4949 B/op           11 allocs/op

    sonic █████████████████████████████▊ 4.25
 jsoniter ████████████████████  2.87
     json ▓▓▓▓▓▓▓ 1.00
       oj █████▉ 0.84
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal simple types to string/[]byte
     json.Marshal            13613 ns/op         9873 B/op          216 allocs/op
       oj.JSON                4133 ns/op            0 B/op            0 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             7569 ns/op         7052 B/op           94 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             4299 ns/op         2829 B/op            4 allocs/op

       oj ███████████████████████  3.29
    sonic ██████████████████████▏ 3.17
 jsoniter ████████████▌ 1.80
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal a struct to string/[]byte
     json.Marshal             3838 ns/op         3462 B/op            1 allocs/op
       oj.Marshal             4799 ns/op         5178 B/op           45 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             3661 ns/op         3471 B/op            2 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             1551 ns/op         3304 B/op            4 allocs/op

    sonic █████████████████▎ 2.47
 jsoniter ███████▎ 1.05
     json ▓▓▓▓▓▓▓ 1.00
       oj █████▌ 0.80
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Read from single JSON file
     json.Decode             26646 ns/op        32385 B/op          342 allocs/op
       oj.ParseReader        15390 ns/op        21228 B/op          376 allocs/op
 fastjson.Decode              5341 ns/op        12034 B/op            7 allocs/op
 jsoniter.Decode             18775 ns/op        20328 B/op          456 allocs/op
 simdjson.Decode             21649 ns/op       130664 B/op           17 allocs/op
    gjson.Decode              9238 ns/op        16368 B/op            9 allocs/op
    sonic.Decode              7421 ns/op        12516 B/op           21 allocs/op

 fastjson ██████████████████████████████████▉ 4.99
    sonic █████████████████████████▏ 3.59
    gjson ████████████████████▏ 2.88
       oj ████████████  1.73
 jsoniter █████████▉ 1.42
 simdjson ████████▌ 1.23
     json ▓▓▓▓▓▓▓ 1.00

Read multiple JSON in a small log file (100MB)
     json.Decode         744776082 ns/op   1188215368 B/op     14810437 allocs/op
       oj.ParseReader    749100920 ns/op   1615708336 B/op     17772489 allocs/op
 fastjson.Decode         129097027 ns/op     37919390 B/op       592422 allocs/op
 jsoniter.Decode         702803778 ns/op    649331232 B/op     19390803 allocs/op
 simdjson.ParseReader    336034633 ns/op   1388713808 B/op     13625784 allocs/op
    gjson.Decode         269638732 ns/op    104270240 B/op       592427 allocs/op
    sonic.Decode         443481382 ns/op    803131074 B/op      2963365 allocs/op

 fastjson ████████████████████████████████████████▍ 5.77
    gjson ███████████████████▎ 2.76
 simdjson ███████████████▌ 2.22
    sonic ███████████▊ 1.68
 jsoniter ███████▍ 1.06
     json ▓▓▓▓▓▓▓ 1.00
       oj ██████▉ 0.99

Read multiple JSON in a semi large log file (5GB)
     json.Decode       40537043036 ns/op  28649556208 B/op    740520844 allocs/op
       oj.ParseReader  38168848908 ns/op  80785290544 B/op    888625314 allocs/op
 fastjson.Decode        6426564552 ns/op   1895752816 B/op     29620968 allocs/op
 jsoniter.Decode       34873845414 ns/op  32465547800 B/op    969506563 allocs/op
 simdjson.ParseReader  14157583251 ns/op  33795267880 B/op    681285461 allocs/op
    gjson.Decode       12832344728 ns/op   5213263136 B/op     29620791 allocs/op
    sonic.Decode       16017868326 ns/op  37207713080 B/op    444312006 allocs/op

 fastjson ████████████████████████████████████████████▏ 6.31
    gjson ██████████████████████  3.16
 simdjson ████████████████████  2.86
    sonic █████████████████▋ 2.53
 jsoniter ████████▏ 1.16
       oj ███████▍ 1.06
     json ▓▓▓▓▓▓▓ 1.00

 Higher values (longer bars) are better in all cases. The bar graph compares the
 parsing performance. The lighter colored bar is the reference, the go json
 package.

Tests run on:
 OS:              Ubuntu 22.04.4 LTS
 Processor:       13th Gen Intel(R) Core(TM) i9-13900KF
 Cores:           32
 Processor Speed: 3.00GHz
 Memory:          32 GB
```

## ARM64 (Mac - M1 Max)
NOTE: sonic is not supported on ARM64, falls back to `encoding/json`.

```
WARNING: sonic only supports Go1.16~1.22 && CPU amd64, but your environment is not suitable

Parse string/[]byte to simple go types ([]interface{}, int64, string, etc)
     json.Unmarshal          29702 ns/op        17745 B/op          334 allocs/op
       oj.Parse              13462 ns/op         5691 B/op          364 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal          19687 ns/op        19657 B/op          451 allocs/op
 simdjson.Parse       >>> Unsupported platform <<<
    gjson.ParseBytes         21157 ns/op        20041 B/op          175 allocs/op
    sonic.Unmarshal          34205 ns/op        40623 B/op          345 allocs/op

       oj ███████████████▍ 2.21
 jsoniter ██████████▌ 1.51
    gjson █████████▊ 1.40
     json ▓▓▓▓▓▓▓ 1.00
    sonic ██████  0.87
 fastjson >>> not supported <<<
 simdjson >>> Unsupported platform <<<

Validate string/[]byte
     json.Valid              10720 ns/op            0 B/op            0 allocs/op
       oj.Validate            3737 ns/op            0 B/op            0 allocs/op
 fastjson.Valid               4869 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               6554 ns/op         2184 B/op          100 allocs/op
 simdjson.Validate    >>> Unsupported platform <<<
    gjson.Validate            3521 ns/op            0 B/op            0 allocs/op
    sonic.Valid              10699 ns/op            0 B/op            0 allocs/op

    gjson █████████████████████▎ 3.04
       oj ████████████████████  2.87
 fastjson ███████████████▍ 2.20
 jsoniter ███████████▍ 1.64
    sonic ███████  1.00
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<

Iterate tokens in a string/[]byte
     json.Decode             47987 ns/op        22568 B/op         1175 allocs/op
       oj.Tokenize            6533 ns/op         1976 B/op          156 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             19998 ns/op        20360 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Decode             33518 ns/op        32432 B/op          343 allocs/op

       oj ███████████████████████████████████████████████████▍ 7.35
 jsoniter ████████████████▊ 2.40
    sonic ██████████  1.43
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Unmarshal string/[]byte to a struct
     json.Unmarshal          28217 ns/op         2576 B/op           74 allocs/op
       oj.Unmarshal          27898 ns/op         9354 B/op          456 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal           9114 ns/op         3137 B/op          170 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Unmarshal          33446 ns/op        25456 B/op           85 allocs/op

 jsoniter █████████████████████▋ 3.10
       oj ███████  1.01
     json ▓▓▓▓▓▓▓ 1.00
    sonic █████▉ 0.84
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal simple types to string/[]byte
     json.Marshal            16955 ns/op         9860 B/op          216 allocs/op
       oj.JSON                5896 ns/op            0 B/op            0 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             9644 ns/op         7042 B/op           94 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal            16746 ns/op         9524 B/op          217 allocs/op

       oj ████████████████████▏ 2.88
 jsoniter ████████████▎ 1.76
    sonic ███████  1.01
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal a struct to string/[]byte
     json.Marshal             5032 ns/op         3457 B/op            1 allocs/op
       oj.Marshal             6120 ns/op         5170 B/op           45 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             5032 ns/op         3465 B/op            2 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             4924 ns/op         3121 B/op            2 allocs/op

    sonic ███████▏ 1.02
     json ▓▓▓▓▓▓▓ 1.00
 jsoniter ███████  1.00
       oj █████▊ 0.82
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Read from single JSON file
     json.Decode             38532 ns/op        32384 B/op          342 allocs/op
       oj.ParseReader        21328 ns/op        21228 B/op          376 allocs/op
 fastjson.Decode              8428 ns/op        12034 B/op            7 allocs/op
 jsoniter.Decode             26736 ns/op        20328 B/op          456 allocs/op
 simdjson.Decode      >>> Unsupported platform <<<
    gjson.Decode             13139 ns/op        16368 B/op            9 allocs/op
    sonic.Decode             15889 ns/op        12338 B/op           21 allocs/op

 fastjson ████████████████████████████████  4.57
    gjson ████████████████████▌ 2.93
    sonic ████████████████▉ 2.43
       oj ████████████▋ 1.81
 jsoniter ██████████  1.44
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<

Read multiple JSON in a small log file (100MB)
     json.Decode        1006640937 ns/op   1188213588 B/op     14810438 allocs/op
       oj.ParseReader    829563791 ns/op   1615713456 B/op     17772605 allocs/op
 fastjson.Decode         165035256 ns/op     37919385 B/op       592421 allocs/op
 jsoniter.Decode         867519500 ns/op    649324144 B/op     19390498 allocs/op
 simdjson.ParseReader >>> Unsupported platform <<<
    gjson.Decode         327220062 ns/op    104269688 B/op       592421 allocs/op
    sonic.Decode         882882500 ns/op    795187660 B/op      2963031 allocs/op

 fastjson ██████████████████████████████████████████▋ 6.10
    gjson █████████████████████▌ 3.08
       oj ████████▍ 1.21
 jsoniter ████████  1.16
    sonic ███████▉ 1.14
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<

Read multiple JSON in a semi large log file (5GB)
     json.Decode       51670198500 ns/op  28649572368 B/op    740522071 allocs/op
       oj.ParseReader  42258644292 ns/op  80785647936 B/op    888631486 allocs/op
 fastjson.Decode        8245936792 ns/op   1895746672 B/op     29620904 allocs/op
 jsoniter.Decode       43650830791 ns/op  32465560120 B/op    969508464 allocs/op
 simdjson.ParseReader >>> Unsupported platform <<<
    gjson.Decode       16478191708 ns/op   5213284736 B/op     29621016 allocs/op
    sonic.Decode       64020213875 ns/op  68969419488 B/op   1095975931 allocs/op

 fastjson ███████████████████████████████████████████▊ 6.27
    gjson █████████████████████▉ 3.14
       oj ████████▌ 1.22
 jsoniter ████████▎ 1.18
     json ▓▓▓▓▓▓▓ 1.00
    sonic █████▋ 0.81
 simdjson >>> Unsupported platform <<<

 Higher values (longer bars) are better in all cases. The bar graph compares the
 parsing performance. The lighter colored bar is the reference, the go json
 package.

Tests run on:
 Machine:         MacBookPro18,2
 OS:              macOS 14.4.1
 Processor:       
 Cores:           proc 10:8:2
 Processor Speed: 
 Memory:          32 GB
 ```
 