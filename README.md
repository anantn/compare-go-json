# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

The master branch is mostly unchanged from the original repository, see the 
[dynamicValues](https://github.com/anantn/compare-go-json/blob/dynamicSuite/README.md) branch for more benchmark results.

# Benchmarks

## x86_64 (Linux)
```
Parse string/[]byte to simple go types ([]interface{}, int64, string, etc)
     json.Unmarshal          21261 ns/op        17744 B/op          334 allocs/op
       oj.Parse              10385 ns/op         5691 B/op          364 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal          15396 ns/op        19682 B/op          451 allocs/op
 simdjson.Parse              28500 ns/op       140635 B/op          361 allocs/op
    gjson.ParseBytes         16639 ns/op        20040 B/op          175 allocs/op
 gjson-nv >>> not supported <<<
    sonic.Unmarshal           9516 ns/op        23652 B/op          168 allocs/op

    sonic ███████████████▋ 2.23
       oj ██████████████▎ 2.05
 jsoniter █████████▋ 1.38
    gjson ████████▉ 1.28
     json ▓▓▓▓▓▓▓ 1.00
 simdjson █████▏ 0.75
 fastjson >>> not supported <<<
 gjson-nv >>> not supported <<<

Validate string/[]byte
     json.Valid               6587 ns/op            0 B/op            0 allocs/op
       oj.Validate            2260 ns/op            0 B/op            0 allocs/op
 fastjson.Valid               2772 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               4617 ns/op         2187 B/op          100 allocs/op
 simdjson.Validate           16397 ns/op       118633 B/op           11 allocs/op
    gjson.Validate            2253 ns/op            0 B/op            0 allocs/op
 gjson-nv >>> not supported <<<
    sonic.Valid               2091 ns/op            0 B/op            0 allocs/op

    sonic ██████████████████████  3.15
    gjson ████████████████████▍ 2.92
       oj ████████████████████▍ 2.91
 fastjson ████████████████▋ 2.38
 jsoniter █████████▉ 1.43
     json ▓▓▓▓▓▓▓ 1.00
 simdjson ██▊ 0.40
 gjson-nv >>> not supported <<<

Iterate tokens in a string/[]byte
     json.Decode             39294 ns/op        22568 B/op         1175 allocs/op
       oj.Tokenize            4650 ns/op         1976 B/op          156 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             16044 ns/op        20360 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<
    sonic.Decode             10173 ns/op        24088 B/op          170 allocs/op

       oj ███████████████████████████████████████████████████████████▏ 8.45
    sonic ███████████████████████████  3.86
 jsoniter █████████████████▏ 2.45
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<

Unmarshal string/[]byte to a struct
     json.Unmarshal          19498 ns/op         2576 B/op           74 allocs/op
       oj.Unmarshal          22971 ns/op         9348 B/op          456 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal           6710 ns/op         3141 B/op          170 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<
    sonic.Unmarshal           4709 ns/op         4974 B/op           11 allocs/op

    sonic ████████████████████████████▉ 4.14
 jsoniter ████████████████████▎ 2.91
     json ▓▓▓▓▓▓▓ 1.00
       oj █████▉ 0.85
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<

Marshal simple types to string/[]byte
     json.Marshal            13984 ns/op         9873 B/op          216 allocs/op
       oj.JSON                4225 ns/op            0 B/op            0 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             7670 ns/op         7052 B/op           94 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<
    sonic.Marshal             4356 ns/op         2827 B/op            4 allocs/op

       oj ███████████████████████▏ 3.31
    sonic ██████████████████████▍ 3.21
 jsoniter ████████████▊ 1.82
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<

Marshal a struct to string/[]byte
     json.Marshal             3943 ns/op         3462 B/op            1 allocs/op
       oj.Marshal             4765 ns/op         5178 B/op           45 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             3603 ns/op         3471 B/op            2 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<
    sonic.Marshal             1523 ns/op         3282 B/op            4 allocs/op

    sonic ██████████████████  2.59
 jsoniter ███████▋ 1.09
     json ▓▓▓▓▓▓▓ 1.00
       oj █████▊ 0.83
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<

Marshal custom data with a builder
     json.Marshal             1066 ns/op          689 B/op           17 allocs/op
       oj.Builder             1021 ns/op         1131 B/op           17 allocs/op
 fastjson.Marshal              639 ns/op          679 B/op            6 allocs/op
 jsoniter.Marshal              577 ns/op          432 B/op            7 allocs/op
 simdjson >>> not supported <<<
    gjson.Marshal             2153 ns/op         3608 B/op           32 allocs/op
 gjson-nv >>> not supported <<<
    sonic.Marshal              927 ns/op         1808 B/op           14 allocs/op

 jsoniter ████████████▉ 1.85
 fastjson ███████████▋ 1.67
    sonic ████████  1.15
       oj ███████▎ 1.04
     json ▓▓▓▓▓▓▓ 1.00
    gjson ███▍ 0.50
 simdjson >>> not supported <<<
 gjson-nv >>> not supported <<<

Read from single JSON file
     json.Decode             27268 ns/op        32385 B/op          342 allocs/op
       oj.ParseReader        15493 ns/op        21228 B/op          376 allocs/op
 fastjson.Decode              5391 ns/op        12034 B/op            7 allocs/op
 jsoniter.Decode             18583 ns/op        20329 B/op          456 allocs/op
 simdjson.Decode             21218 ns/op       130664 B/op           17 allocs/op
    gjson.Decode              8723 ns/op        16368 B/op            9 allocs/op
 gjson-nv.Decode              6290 ns/op        16368 B/op            9 allocs/op
    sonic.Decode              7299 ns/op        12479 B/op           21 allocs/op

 fastjson ███████████████████████████████████▍ 5.06
 gjson-nv ██████████████████████████████▎ 4.34
    sonic ██████████████████████████▏ 3.74
    gjson █████████████████████▉ 3.13
       oj ████████████▎ 1.76
 jsoniter ██████████▎ 1.47
 simdjson ████████▉ 1.29
     json ▓▓▓▓▓▓▓ 1.00

Read multiple JSON in a small log file, read few keys (100MB)
     json.Decode         809696366 ns/op   1342691224 B/op     14015430 allocs/op
       oj.ParseReader    518351929 ns/op    332604952 B/op     15622348 allocs/op
 fastjson.Decode          82186665 ns/op      2536368 B/op        39514 allocs/op
 jsoniter.Decode         801723826 ns/op    823285864 B/op     19383905 allocs/op
 simdjson.ParseReader    423053908 ns/op   1470575360 B/op     14804689 allocs/op
    gjson.Decode         169356125 ns/op    106043200 B/op        39452 allocs/op
 gjson-nv.Decode          89132714 ns/op    106043155 B/op        39451 allocs/op
    sonic.Decode         139986526 ns/op    154191010 B/op       197467 allocs/op

 fastjson ████████████████████████████████████████████████████████████████████▉ 9.85
 gjson-nv ███████████████████████████████████████████████████████████████▌ 9.08
    sonic ████████████████████████████████████████▍ 5.78
    gjson █████████████████████████████████▍ 4.78
 simdjson █████████████▍ 1.91
       oj ██████████▉ 1.56
 jsoniter ███████  1.01
     json ▓▓▓▓▓▓▓ 1.00

Read multiple JSON in a semi large log file, read few keys (5GB)
     json.Decode       12899878826 ns/op  10789510592 B/op    207850090 allocs/op
       oj.ParseReader   7579681771 ns/op   4935972816 B/op    231680194 allocs/op
 fastjson.Decode        1205493747 ns/op     37538440 B/op       585884 allocs/op
 jsoniter.Decode       11696739704 ns/op  12208954040 B/op    287453779 allocs/op
 simdjson.ParseReader   5530973328 ns/op  14775589328 B/op    219552725 allocs/op
    gjson.Decode        2375472677 ns/op   1572570112 B/op       585033 allocs/op
 gjson-nv.Decode        1178721970 ns/op   1572570112 B/op       585033 allocs/op
    sonic.Decode        4755546118 ns/op  13499292312 B/op    106054237 allocs/op

 gjson-nv ████████████████████████████████████████████████████████████████████████████▌ 10.94
 fastjson ██████████████████████████████████████████████████████████████████████████▉ 10.70
    gjson ██████████████████████████████████████  5.43
    sonic ██████████████████▉ 2.71
 simdjson ████████████████▎ 2.33
       oj ███████████▉ 1.70
 jsoniter ███████▋ 1.10
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
```
Parse string/[]byte to simple go types ([]interface{}, int64, string, etc)
     json.Unmarshal          29261 ns/op        17744 B/op          334 allocs/op
       oj.Parse              13478 ns/op         5691 B/op          364 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal          19626 ns/op        19656 B/op          451 allocs/op
 simdjson.Parse       >>> Unsupported platform <<<
    gjson.ParseBytes         21144 ns/op        20041 B/op          175 allocs/op
 gjson-nv >>> not supported <<<
    sonic.Unmarshal   >>> Unsupported platform <<<

       oj ███████████████▏ 2.17
 jsoniter ██████████▍ 1.49
    gjson █████████▋ 1.38
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> Unsupported platform <<<
 gjson-nv >>> not supported <<<
    sonic >>> Unsupported platform <<<

Validate string/[]byte
     json.Valid              10355 ns/op            0 B/op            0 allocs/op
       oj.Validate            3832 ns/op            0 B/op            0 allocs/op
 fastjson.Valid               4881 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               6583 ns/op         2184 B/op          100 allocs/op
 simdjson.Validate    >>> Unsupported platform <<<
    gjson.Validate            3502 ns/op            0 B/op            0 allocs/op
 gjson-nv >>> not supported <<<
    sonic.Valid       >>> Unsupported platform <<<

    gjson ████████████████████▋ 2.96
       oj ██████████████████▉ 2.70
 fastjson ██████████████▊ 2.12
 jsoniter ███████████  1.57
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<
 gjson-nv >>> not supported <<<
    sonic >>> Unsupported platform <<<

Iterate tokens in a string/[]byte
     json.Decode             47572 ns/op        22568 B/op         1175 allocs/op
       oj.Tokenize            6555 ns/op         1976 B/op          156 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             19926 ns/op        20360 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<
    sonic.Decode      >>> Unsupported platform <<<

       oj ██████████████████████████████████████████████████▊ 7.26
 jsoniter ████████████████▋ 2.39
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<
    sonic >>> Unsupported platform <<<

Unmarshal string/[]byte to a struct
     json.Unmarshal          28148 ns/op         2576 B/op           74 allocs/op
       oj.Unmarshal          27906 ns/op         9351 B/op          456 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal           9085 ns/op         3137 B/op          170 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<
    sonic.Unmarshal   >>> Unsupported platform <<<

 jsoniter █████████████████████▋ 3.10
       oj ███████  1.01
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<
    sonic >>> Unsupported platform <<<

Marshal simple types to string/[]byte
     json.Marshal            16757 ns/op         9860 B/op          216 allocs/op
       oj.JSON                5876 ns/op            0 B/op            0 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             9616 ns/op         7042 B/op           94 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<
    sonic.Marshal     >>> Unsupported platform <<<

       oj ███████████████████▉ 2.85
 jsoniter ████████████▏ 1.74
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<
    sonic >>> Unsupported platform <<<

Marshal a struct to string/[]byte
     json.Marshal             5052 ns/op         3457 B/op            1 allocs/op
       oj.Marshal             6150 ns/op         5170 B/op           45 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             5035 ns/op         3465 B/op            2 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<
    sonic.Marshal     >>> Unsupported platform <<<

 jsoniter ███████  1.00
     json ▓▓▓▓▓▓▓ 1.00
       oj █████▊ 0.82
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
 gjson-nv >>> not supported <<<
    sonic >>> Unsupported platform <<<

Marshal custom data with a builder
     json.Marshal             1210 ns/op          688 B/op           17 allocs/op
       oj.Builder             1280 ns/op         1128 B/op           17 allocs/op
 fastjson.Marshal              818 ns/op          679 B/op            6 allocs/op
 jsoniter.Marshal              700 ns/op          432 B/op            7 allocs/op
 simdjson >>> not supported <<<
    gjson.Marshal             2507 ns/op         3602 B/op           32 allocs/op
 gjson-nv >>> not supported <<<
    sonic.Marshal     >>> Unsupported platform <<<

 jsoniter ████████████  1.73
 fastjson ██████████▎ 1.48
     json ▓▓▓▓▓▓▓ 1.00
       oj ██████▌ 0.95
    gjson ███▍ 0.48
 simdjson >>> not supported <<<
 gjson-nv >>> not supported <<<
    sonic >>> Unsupported platform <<<

Read from single JSON file
     json.Decode             37826 ns/op        32386 B/op          342 allocs/op
       oj.ParseReader        21234 ns/op        21228 B/op          376 allocs/op
 fastjson.Decode              8447 ns/op        12034 B/op            7 allocs/op
 jsoniter.Decode             26572 ns/op        20328 B/op          456 allocs/op
 simdjson.Decode      >>> Unsupported platform <<<
    gjson.Decode             13142 ns/op        16368 B/op            9 allocs/op
 gjson-nv.Decode              9576 ns/op        16368 B/op            9 allocs/op
    sonic.Decode      >>> Unsupported platform <<<

 fastjson ███████████████████████████████▎ 4.48
 gjson-nv ███████████████████████████▋ 3.95
    gjson ████████████████████▏ 2.88
       oj ████████████▍ 1.78
 jsoniter █████████▉ 1.42
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<
    sonic >>> Unsupported platform <<<

Read multiple JSON in a small log file, read few keys (100MB)
     json.Decode         950949458 ns/op   1188217536 B/op     14810436 allocs/op
       oj.ParseReader    822483771 ns/op   1615717792 B/op     17772603 allocs/op
 fastjson.Decode         161158904 ns/op     37919453 B/op       592422 allocs/op
 jsoniter.Decode         859668875 ns/op    649273264 B/op     19389030 allocs/op
 simdjson.ParseReader >>> Unsupported platform <<<
    gjson.Decode         326253395 ns/op    104269904 B/op       592424 allocs/op
 gjson-nv.Decode         231337566 ns/op    104270211 B/op       592427 allocs/op
    sonic.Decode      >>> Unsupported platform <<<

 fastjson █████████████████████████████████████████▎ 5.90
 gjson-nv ████████████████████████████▊ 4.11
    gjson ████████████████████▍ 2.91
       oj ████████  1.16
 jsoniter ███████▋ 1.11
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<
    sonic >>> Unsupported platform <<<

Read multiple JSON in a semi large log file, read few keys (5GB)
     json.Decode       49956521458 ns/op  28649678624 B/op    740522925 allocs/op
       oj.ParseReader  41781025333 ns/op  80785725096 B/op    888632068 allocs/op
 fastjson.Decode        8089876375 ns/op   1895750032 B/op     29620939 allocs/op
 jsoniter.Decode       42987344584 ns/op  32465532248 B/op    969509921 allocs/op
 simdjson.ParseReader >>> Unsupported platform <<<
    gjson.Decode       16329043667 ns/op   5213301920 B/op     29621195 allocs/op
 gjson-nv.Decode       11391665666 ns/op   5213306048 B/op     29621238 allocs/op
    sonic.Decode      >>> Unsupported platform <<<

 fastjson ███████████████████████████████████████████▏ 6.18
 gjson-nv ██████████████████████████████▋ 4.39
    gjson █████████████████████▍ 3.06
       oj ████████▎ 1.20
 jsoniter ████████▏ 1.16
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<
    sonic >>> Unsupported platform <<<

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
