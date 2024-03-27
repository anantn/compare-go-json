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
 