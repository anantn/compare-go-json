# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Parse string/[]byte to simple go types ([]interface{}, int64, string, etc)
     json.Unmarshal          21501 ns/op        17745 B/op          334 allocs/op
       oj.Parse              10243 ns/op         5691 B/op          364 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal          15292 ns/op        19681 B/op          451 allocs/op
 simdjson.Parse              28975 ns/op       140633 B/op          361 allocs/op
    gjson.ParseBytes         16370 ns/op        20041 B/op          175 allocs/op
    sonic.Unmarshal           9407 ns/op        23638 B/op          168 allocs/op

    sonic ███████████████▉ 2.29
       oj ██████████████▋ 2.10
 jsoniter █████████▊ 1.41
    gjson █████████▏ 1.31
     json ▓▓▓▓▓▓▓ 1.00
 simdjson █████▏ 0.74
 fastjson >>> not supported <<<

Validate string/[]byte
     json.Valid               6557 ns/op            0 B/op            0 allocs/op
       oj.Validate            2282 ns/op            0 B/op            0 allocs/op
 fastjson.Valid               2673 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               4541 ns/op         2187 B/op          100 allocs/op
 simdjson.Validate           17501 ns/op       118632 B/op           11 allocs/op
    gjson.Validate            2223 ns/op            0 B/op            0 allocs/op
    sonic.Valid               1996 ns/op            0 B/op            0 allocs/op

    sonic ██████████████████████▉ 3.29
    gjson ████████████████████▋ 2.95
       oj ████████████████████  2.87
 fastjson █████████████████▏ 2.45
 jsoniter ██████████  1.44
     json ▓▓▓▓▓▓▓ 1.00
 simdjson ██▌ 0.37

Iterate tokens in a string/[]byte
     json.Decode             39612 ns/op        22568 B/op         1175 allocs/op
       oj.Tokenize            4622 ns/op         1976 B/op          156 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             15854 ns/op        20361 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Decode             10272 ns/op        24055 B/op          170 allocs/op

       oj ███████████████████████████████████████████████████████████▉ 8.57
    sonic ██████████████████████████▉ 3.86
 jsoniter █████████████████▍ 2.50
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Unmarshal string/[]byte to a struct
     json.Unmarshal          19752 ns/op         2576 B/op           74 allocs/op
       oj.Unmarshal          23386 ns/op         9356 B/op          456 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal           6706 ns/op         3141 B/op          170 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Unmarshal           4616 ns/op         4959 B/op           11 allocs/op

    sonic █████████████████████████████▉ 4.28
 jsoniter ████████████████████▌ 2.95
     json ▓▓▓▓▓▓▓ 1.00
       oj █████▉ 0.84
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal simple types to string/[]byte
     json.Marshal            13945 ns/op         9873 B/op          216 allocs/op
       oj.JSON                4183 ns/op            0 B/op            0 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             7680 ns/op         7052 B/op           94 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             4355 ns/op         2809 B/op            4 allocs/op

       oj ███████████████████████▎ 3.33
    sonic ██████████████████████▍ 3.20
 jsoniter ████████████▋ 1.82
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal a struct to string/[]byte
     json.Marshal             3956 ns/op         3462 B/op            1 allocs/op
       oj.Marshal             4790 ns/op         5178 B/op           45 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             3623 ns/op         3471 B/op            2 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             1528 ns/op         3275 B/op            4 allocs/op

    sonic ██████████████████  2.59
 jsoniter ███████▋ 1.09
     json ▓▓▓▓▓▓▓ 1.00
       oj █████▊ 0.83
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal custom data with a builder
     json.Marshal             1051 ns/op          689 B/op           17 allocs/op
       oj.Builder             1021 ns/op         1131 B/op           17 allocs/op
 fastjson.Marshal              650 ns/op          679 B/op            6 allocs/op
 jsoniter.Marshal              575 ns/op          432 B/op            7 allocs/op
 simdjson >>> not supported <<<
    gjson.Marshal             2144 ns/op         3608 B/op           32 allocs/op
    sonic.Marshal              932 ns/op         1808 B/op           14 allocs/op

 jsoniter ████████████▊ 1.83
 fastjson ███████████▎ 1.62
    sonic ███████▉ 1.13
       oj ███████▏ 1.03
     json ▓▓▓▓▓▓▓ 1.00
    gjson ███▍ 0.49
 simdjson >>> not supported <<<

Read from single JSON file
     json.Decode             27301 ns/op        32384 B/op          342 allocs/op
       oj.ParseReader        15642 ns/op        21227 B/op          376 allocs/op
 fastjson.Decode              5318 ns/op        12034 B/op            7 allocs/op
 jsoniter.Decode             18860 ns/op        20329 B/op          456 allocs/op
 simdjson.Decode             20972 ns/op       130664 B/op           17 allocs/op
    gjson.Decode              8848 ns/op        16368 B/op            9 allocs/op
    sonic.Decode              7420 ns/op        12483 B/op           21 allocs/op

 fastjson ███████████████████████████████████▉ 5.13
    sonic █████████████████████████▊ 3.68
    gjson █████████████████████▌ 3.09
       oj ████████████▏ 1.75
 jsoniter ██████████▏ 1.45
 simdjson █████████  1.30
     json ▓▓▓▓▓▓▓ 1.00

Read multiple JSON in a small log file (100MB)
     json.Decode         770474232 ns/op   1188215256 B/op     14810436 allocs/op
       oj.ParseReader    752801428 ns/op   1615708160 B/op     17772487 allocs/op
 fastjson.Decode         129529450 ns/op     37919330 B/op       592421 allocs/op
 jsoniter.Decode         684797052 ns/op    649329928 B/op     19390790 allocs/op
 simdjson.ParseReader    333817387 ns/op   1388713210 B/op     13625785 allocs/op
    gjson.Decode         255049173 ns/op    104269472 B/op       592419 allocs/op
    sonic.Decode         451498685 ns/op    805408080 B/op      2963370 allocs/op

 fastjson █████████████████████████████████████████▋ 5.95
    gjson █████████████████████▏ 3.02
 simdjson ████████████████▏ 2.31
    sonic ███████████▉ 1.71
 jsoniter ███████▉ 1.13
       oj ███████▏ 1.02
     json ▓▓▓▓▓▓▓ 1.00

Read multiple JSON in a semi large log file (5GB)
     json.Decode       47117880420 ns/op  28649576048 B/op    740521191 allocs/op
       oj.ParseReader  38197687451 ns/op  80785300272 B/op    888625186 allocs/op
 fastjson.Decode        6519485117 ns/op   1895749168 B/op     29620930 allocs/op
 jsoniter.Decode       34551176567 ns/op  32465546024 B/op    969506514 allocs/op
 simdjson.ParseReader  14544017730 ns/op  33763778544 B/op    681285434 allocs/op
    gjson.Decode       12112099170 ns/op   5213263136 B/op     29620791 allocs/op
    sonic.Decode       15772335888 ns/op  37208246168 B/op    444312018 allocs/op

 fastjson ██████████████████████████████████████████████████▌ 7.23
    gjson ███████████████████████████▏ 3.89
 simdjson ██████████████████████▋ 3.24
    sonic ████████████████████▉ 2.99
 jsoniter █████████▌ 1.36
       oj ████████▋ 1.23
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
 