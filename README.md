# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Parse string/[]byte to simple go types ([]interface{}, int64, string, etc)
     json.Unmarshal          21466 ns/op        17744 B/op          334 allocs/op
       oj.Parse              10307 ns/op         5691 B/op          364 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal          15518 ns/op        19681 B/op          451 allocs/op
 simdjson.Parse              29127 ns/op       140634 B/op          361 allocs/op
    gjson.ParseBytes         14806 ns/op        20040 B/op          175 allocs/op
    sonic.Unmarshal           9568 ns/op        23664 B/op          168 allocs/op

    sonic ███████████████▋ 2.24
       oj ██████████████▌ 2.08
    gjson ██████████▏ 1.45
 jsoniter █████████▋ 1.38
     json ▓▓▓▓▓▓▓ 1.00
 simdjson █████▏ 0.74
 fastjson >>> not supported <<<

Validate string/[]byte
     json.Valid               6379 ns/op            0 B/op            0 allocs/op
       oj.Validate            2262 ns/op            0 B/op            0 allocs/op
 fastjson.Valid               2567 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               4622 ns/op         2187 B/op          100 allocs/op
 simdjson.Validate           19669 ns/op       118632 B/op           11 allocs/op
    gjson.Validate            2386 ns/op            0 B/op            0 allocs/op
    sonic.Valid               1986 ns/op            0 B/op            0 allocs/op

    sonic ██████████████████████▍ 3.21
       oj ███████████████████▋ 2.82
    gjson ██████████████████▋ 2.67
 fastjson █████████████████▍ 2.49
 jsoniter █████████▋ 1.38
     json ▓▓▓▓▓▓▓ 1.00
 simdjson ██▎ 0.32

Iterate tokens in a string/[]byte
     json.Decode             39695 ns/op        22568 B/op         1175 allocs/op
       oj.Tokenize            4647 ns/op         1976 B/op          156 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             16388 ns/op        20360 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Decode             10400 ns/op        24098 B/op          170 allocs/op

       oj ███████████████████████████████████████████████████████████▊ 8.54
    sonic ██████████████████████████▋ 3.82
 jsoniter ████████████████▉ 2.42
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Unmarshal string/[]byte to a struct
     json.Unmarshal          20120 ns/op         2576 B/op           74 allocs/op
       oj.Unmarshal          23657 ns/op         9352 B/op          456 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal           6884 ns/op         3141 B/op          170 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Unmarshal           4646 ns/op         4965 B/op           11 allocs/op

    sonic ██████████████████████████████▎ 4.33
 jsoniter ████████████████████▍ 2.92
     json ▓▓▓▓▓▓▓ 1.00
       oj █████▉ 0.85
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal simple types to string/[]byte
     json.Marshal            13565 ns/op         9873 B/op          216 allocs/op
       oj.JSON                4162 ns/op            0 B/op            0 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             7461 ns/op         7052 B/op           94 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             4255 ns/op         2796 B/op            4 allocs/op

       oj ██████████████████████▊ 3.26
    sonic ██████████████████████▎ 3.19
 jsoniter ████████████▋ 1.82
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal a struct to string/[]byte
     json.Marshal             3825 ns/op         3462 B/op            1 allocs/op
       oj.Marshal             4842 ns/op         5178 B/op           45 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             3586 ns/op         3471 B/op            2 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             1541 ns/op         3310 B/op            4 allocs/op

    sonic █████████████████▍ 2.48
 jsoniter ███████▍ 1.07
     json ▓▓▓▓▓▓▓ 1.00
       oj █████▌ 0.79
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Read from single JSON file
     json.Decode             26635 ns/op        32383 B/op          342 allocs/op
       oj.ParseReader        15348 ns/op        21228 B/op          376 allocs/op
 fastjson.Decode              5350 ns/op        12034 B/op            7 allocs/op
 jsoniter.Decode             19016 ns/op        20328 B/op          456 allocs/op
 simdjson.Decode             21728 ns/op       130664 B/op           17 allocs/op
    gjson.Decode              6631 ns/op        16368 B/op            9 allocs/op
    sonic.Decode              7363 ns/op        12527 B/op           21 allocs/op

 fastjson ██████████████████████████████████▊ 4.98
    gjson ████████████████████████████  4.02
    sonic █████████████████████████▎ 3.62
       oj ████████████▏ 1.74
 jsoniter █████████▊ 1.40
 simdjson ████████▌ 1.23
     json ▓▓▓▓▓▓▓ 1.00

Read multiple JSON in a small log file (100MB)
     json.Decode         760619982 ns/op   1188215216 B/op     14810436 allocs/op
       oj.ParseReader    772730508 ns/op   1615707280 B/op     17772476 allocs/op
 fastjson.Decode         169581834 ns/op    758001769 B/op       592494 allocs/op
 jsoniter.Decode         615869736 ns/op   1264548416 B/op     19390813 allocs/op
 simdjson.ParseReader    339722034 ns/op   1388717630 B/op     13625793 allocs/op
    gjson.Decode          81924325 ns/op    720083398 B/op           51 allocs/op
    sonic.Decode         449701700 ns/op    804422544 B/op      2963353 allocs/op

    gjson ████████████████████████████████████████████████████████████████▉ 9.28
 fastjson ███████████████████████████████▍ 4.49
 simdjson ███████████████▋ 2.24
    sonic ███████████▊ 1.69
 jsoniter ████████▋ 1.24
     json ▓▓▓▓▓▓▓ 1.00
       oj ██████▉ 0.98

Read multiple JSON in a semi large log file (5GB)
     json.Decode       41145660800 ns/op  28649554480 B/op    740520826 allocs/op
       oj.ParseReader  38856723451 ns/op  80785308272 B/op    888625626 allocs/op
 fastjson.Decode       32093603765 ns/op  99289492656 B/op    414697316 allocs/op
 jsoniter.Decode       35022161164 ns/op  32465530648 B/op    969506510 allocs/op
 simdjson.ParseReader  14220122085 ns/op  33742788096 B/op    681285434 allocs/op
    gjson.Decode        8844119698 ns/op   5213263136 B/op     29620791 allocs/op
    sonic.Decode       16295161291 ns/op  37209417880 B/op    444312030 allocs/op

    gjson ████████████████████████████████▌ 4.65
 simdjson ████████████████████▎ 2.89
    sonic █████████████████▋ 2.53
 fastjson ████████▉ 1.28
 jsoniter ████████▏ 1.17
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
 