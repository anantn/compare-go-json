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
     json.Unmarshal          29493 ns/op        17743 B/op          334 allocs/op
       oj.Parse              13474 ns/op         5691 B/op          364 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal          19733 ns/op        19657 B/op          451 allocs/op
 simdjson.Parse       >>> Unsupported platform <<<
    gjson.ParseBytes         17462 ns/op        20040 B/op          175 allocs/op
    sonic.Unmarshal          34242 ns/op        40625 B/op          345 allocs/op

       oj ███████████████▎ 2.19
    gjson ███████████▊ 1.69
 jsoniter ██████████▍ 1.49
     json ▓▓▓▓▓▓▓ 1.00
    sonic ██████  0.86
 fastjson >>> not supported <<<
 simdjson >>> Unsupported platform <<<

Validate string/[]byte
     json.Valid              10698 ns/op            0 B/op            0 allocs/op
       oj.Validate            3745 ns/op            0 B/op            0 allocs/op
 fastjson.Valid               4873 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               6560 ns/op         2184 B/op          100 allocs/op
 simdjson.Validate    >>> Unsupported platform <<<
    gjson.Validate            3505 ns/op            0 B/op            0 allocs/op
    sonic.Valid              10684 ns/op            0 B/op            0 allocs/op

    gjson █████████████████████▎ 3.05
       oj ███████████████████▉ 2.86
 fastjson ███████████████▎ 2.20
 jsoniter ███████████▍ 1.63
    sonic ███████  1.00
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<

Iterate tokens in a string/[]byte
     json.Decode             48077 ns/op        22568 B/op         1175 allocs/op
       oj.Tokenize            6561 ns/op         1976 B/op          156 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             19989 ns/op        20362 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Decode             33650 ns/op        32432 B/op          343 allocs/op

       oj ███████████████████████████████████████████████████▎ 7.33
 jsoniter ████████████████▊ 2.41
    sonic ██████████  1.43
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Unmarshal string/[]byte to a struct
     json.Unmarshal          28676 ns/op         2576 B/op           74 allocs/op
       oj.Unmarshal          27899 ns/op         9356 B/op          456 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal           9137 ns/op         3137 B/op          170 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Unmarshal          33528 ns/op        25456 B/op           85 allocs/op

 jsoniter █████████████████████▉ 3.14
       oj ███████▏ 1.03
     json ▓▓▓▓▓▓▓ 1.00
    sonic █████▉ 0.86
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal simple types to string/[]byte
     json.Marshal            16649 ns/op         9860 B/op          216 allocs/op
       oj.JSON                5907 ns/op            0 B/op            0 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             9819 ns/op         7042 B/op           94 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal            16440 ns/op         9524 B/op          217 allocs/op

       oj ███████████████████▋ 2.82
 jsoniter ███████████▊ 1.70
    sonic ███████  1.01
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal a struct to string/[]byte
     json.Marshal             5047 ns/op         3457 B/op            1 allocs/op
       oj.Marshal             6155 ns/op         5170 B/op           45 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             5043 ns/op         3465 B/op            2 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             4886 ns/op         3121 B/op            2 allocs/op

    sonic ███████▏ 1.03
 jsoniter ███████  1.00
     json ▓▓▓▓▓▓▓ 1.00
       oj █████▋ 0.82
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Read from single JSON file
     json.Decode             38715 ns/op        32384 B/op          342 allocs/op
       oj.ParseReader        21324 ns/op        21228 B/op          376 allocs/op
 fastjson.Decode              8400 ns/op        12034 B/op            7 allocs/op
 jsoniter.Decode             26706 ns/op        20327 B/op          456 allocs/op
 simdjson.Decode      >>> Unsupported platform <<<
    gjson.Decode              9541 ns/op        16368 B/op            9 allocs/op
    sonic.Decode             15580 ns/op        12313 B/op           21 allocs/op

 fastjson ████████████████████████████████▎ 4.61
    gjson ████████████████████████████▍ 4.06
    sonic █████████████████▍ 2.48
       oj ████████████▋ 1.82
 jsoniter ██████████▏ 1.45
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<

Read multiple JSON in a small log file (100MB)
     json.Decode         989807958 ns/op   1188214660 B/op     14810438 allocs/op
       oj.ParseReader    836860375 ns/op   1615716488 B/op     17772614 allocs/op
 fastjson.Decode         180654611 ns/op    758001616 B/op       592492 allocs/op
 jsoniter.Decode         766561250 ns/op   1264564856 B/op     19390894 allocs/op
 simdjson.ParseReader >>> Unsupported platform <<<
    gjson.Decode          45652329 ns/op    720083379 B/op           50 allocs/op
    sonic.Decode         878642083 ns/op    794567944 B/op      2962991 allocs/op

    gjson ███████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████▊ 21.68
 fastjson ██████████████████████████████████████▎ 5.48
 jsoniter █████████  1.29
       oj ████████▎ 1.18
    sonic ███████▉ 1.13
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<

Read multiple JSON in a semi large log file (5GB)
     json.Decode       51264181209 ns/op  28649608784 B/op    740522344 allocs/op
       oj.ParseReader  41515130417 ns/op  80785619680 B/op    888631231 allocs/op
 fastjson.Decode       34472489750 ns/op  99289857552 B/op    414701117 allocs/op
 jsoniter.Decode       43521046750 ns/op  32465448104 B/op    969503744 allocs/op
 simdjson.ParseReader >>> Unsupported platform <<<
    gjson.Decode       11664578042 ns/op   5213295968 B/op     29621133 allocs/op
    sonic.Decode       63818727709 ns/op  68969453552 B/op   1095975923 allocs/op

    gjson ██████████████████████████████▊ 4.39
 fastjson ██████████▍ 1.49
       oj ████████▋ 1.23
 jsoniter ████████▏ 1.18
     json ▓▓▓▓▓▓▓ 1.00
    sonic █████▌ 0.80
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
 