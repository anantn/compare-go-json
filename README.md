# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

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
NOTE: sonic is not supported on ARM64, falls back to `encoding/json`.

```
WARNING: sonic only supports Go1.16~1.22 && CPU amd64, but your environment is not suitable

Parse string/[]byte to simple go types ([]interface{}, int64, string, etc)
     json.Unmarshal          29198 ns/op        17745 B/op          334 allocs/op
       oj.Parse              13634 ns/op         5691 B/op          364 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal          19983 ns/op        19655 B/op          451 allocs/op
 simdjson.Parse       >>> Unsupported platform <<<
    gjson.ParseBytes         21488 ns/op        20042 B/op          175 allocs/op
    sonic.Unmarshal          34436 ns/op        40624 B/op          345 allocs/op

       oj ██████████████▉ 2.14
 jsoniter ██████████▏ 1.46
    gjson █████████▌ 1.36
     json ▓▓▓▓▓▓▓ 1.00
    sonic █████▉ 0.85
 fastjson >>> not supported <<<
 simdjson >>> Unsupported platform <<<

Validate string/[]byte
     json.Valid              10444 ns/op            0 B/op            0 allocs/op
       oj.Validate            3805 ns/op            0 B/op            0 allocs/op
 fastjson.Valid               4968 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               6804 ns/op         2184 B/op          100 allocs/op
 simdjson.Validate    >>> Unsupported platform <<<
    gjson.Validate            3553 ns/op            0 B/op            0 allocs/op
    sonic.Valid              10423 ns/op            0 B/op            0 allocs/op

    gjson ████████████████████▌ 2.94
       oj ███████████████████▏ 2.74
 fastjson ██████████████▋ 2.10
 jsoniter ██████████▋ 1.53
    sonic ███████  1.00
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<

Iterate tokens in a string/[]byte
     json.Decode             48070 ns/op        22568 B/op         1175 allocs/op
       oj.Tokenize            6671 ns/op         1976 B/op          156 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             20086 ns/op        20360 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Decode             33050 ns/op        32433 B/op          343 allocs/op

       oj ██████████████████████████████████████████████████▍ 7.21
 jsoniter ████████████████▊ 2.39
    sonic ██████████▏ 1.45
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Unmarshal string/[]byte to a struct
     json.Unmarshal          28047 ns/op         2576 B/op           74 allocs/op
       oj.Unmarshal          28298 ns/op         9355 B/op          456 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal           9104 ns/op         3137 B/op          170 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Unmarshal          33366 ns/op        25456 B/op           85 allocs/op

 jsoniter █████████████████████▌ 3.08
     json ▓▓▓▓▓▓▓ 1.00
       oj ██████▉ 0.99
    sonic █████▉ 0.84
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal simple types to string/[]byte
     json.Marshal            17038 ns/op         9860 B/op          216 allocs/op
       oj.JSON                5885 ns/op            0 B/op            0 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             9730 ns/op         7043 B/op           94 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal            16855 ns/op         9524 B/op          217 allocs/op

       oj ████████████████████▎ 2.90
 jsoniter ████████████▎ 1.75
    sonic ███████  1.01
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal a struct to string/[]byte
     json.Marshal             5100 ns/op         3457 B/op            1 allocs/op
       oj.Marshal             6285 ns/op         5170 B/op           45 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             5088 ns/op         3465 B/op            2 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             4942 ns/op         3121 B/op            2 allocs/op

    sonic ███████▏ 1.03
 jsoniter ███████  1.00
     json ▓▓▓▓▓▓▓ 1.00
       oj █████▋ 0.81
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal custom data with a builder
     json.Marshal             1241 ns/op          688 B/op           17 allocs/op
       oj.Builder             1310 ns/op         1129 B/op           17 allocs/op
 fastjson.Marshal              846 ns/op          679 B/op            6 allocs/op
 jsoniter.Marshal              705 ns/op          432 B/op            7 allocs/op
 simdjson >>> not supported <<<
    gjson.Marshal             2535 ns/op         3602 B/op           32 allocs/op
    sonic.Marshal             1560 ns/op         1804 B/op           14 allocs/op

 jsoniter ████████████▎ 1.76
 fastjson ██████████▎ 1.47
     json ▓▓▓▓▓▓▓ 1.00
       oj ██████▋ 0.95
    sonic █████▌ 0.80
    gjson ███▍ 0.49
 simdjson >>> not supported <<<

Read from single JSON file
     json.Decode             38297 ns/op        32385 B/op          342 allocs/op
       oj.ParseReader        21580 ns/op        21228 B/op          376 allocs/op
 fastjson.Decode              8632 ns/op        12034 B/op            7 allocs/op
 jsoniter.Decode             27322 ns/op        20328 B/op          456 allocs/op
 simdjson.Decode      >>> Unsupported platform <<<
    gjson.Decode             13492 ns/op        16368 B/op            9 allocs/op
    sonic.Decode             16156 ns/op        12331 B/op           21 allocs/op

 fastjson ███████████████████████████████  4.44
    gjson ███████████████████▊ 2.84
    sonic ████████████████▌ 2.37
       oj ████████████▍ 1.77
 jsoniter █████████▊ 1.40
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<

Read multiple JSON in a small log file (100MB)
     json.Decode         981410417 ns/op   1188213416 B/op     14810435 allocs/op
       oj.ParseReader    858121521 ns/op   1615712496 B/op     17772597 allocs/op
 fastjson.Decode         164539875 ns/op     37919426 B/op       592422 allocs/op
 jsoniter.Decode         868243042 ns/op    649325192 B/op     19390507 allocs/op
 simdjson.ParseReader >>> Unsupported platform <<<
    gjson.Decode         328929343 ns/op    104269928 B/op       592424 allocs/op
    sonic.Decode         886378937 ns/op    795107460 B/op      2963034 allocs/op

 fastjson █████████████████████████████████████████▊ 5.96
    gjson ████████████████████▉ 2.98
       oj ████████  1.14
 jsoniter ███████▉ 1.13
    sonic ███████▊ 1.11
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<

Read multiple JSON in a semi large log file (5GB)
     json.Decode       50736163000 ns/op  28649606296 B/op    740522368 allocs/op
       oj.ParseReader  42201668167 ns/op  80785720560 B/op    888632187 allocs/op
 fastjson.Decode        8259766917 ns/op   1895749840 B/op     29620937 allocs/op
 jsoniter.Decode       43761873125 ns/op  32465600808 B/op    969508800 allocs/op
 simdjson.ParseReader >>> Unsupported platform <<<
    gjson.Decode       16592271042 ns/op   5213301824 B/op     29621194 allocs/op
    sonic.Decode       63902368208 ns/op  68969457312 B/op   1095976245 allocs/op

 fastjson ██████████████████████████████████████████▉ 6.14
    gjson █████████████████████▍ 3.06
       oj ████████▍ 1.20
 jsoniter ████████  1.16
     json ▓▓▓▓▓▓▓ 1.00
    sonic █████▌ 0.79
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
 