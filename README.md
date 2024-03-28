# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Validate []byte
     json.Valid               6598 ns/op            0 B/op            0 allocs/op
       oj.Validate            2342 ns/op            0 B/op            0 allocs/op
 fastjson.Validate            2822 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               4674 ns/op         2187 B/op          100 allocs/op
 simdjson.Validate           19610 ns/op       118632 B/op           11 allocs/op
    gjson.Valid               2333 ns/op            0 B/op            0 allocs/op
    sonic.Valid               2032 ns/op            0 B/op            0 allocs/op

    sonic ██████████████████████▋ 3.25
    gjson ███████████████████▊ 2.83
       oj ███████████████████▋ 2.82
 fastjson ████████████████▎ 2.34
 jsoniter █████████▉ 1.41
     json ▓▓▓▓▓▓▓ 1.00
 simdjson ██▎ 0.34

Validate string
     json.Valid               6627 ns/op            0 B/op            0 allocs/op
       oj >>> not supported <<<
 fastjson.Validate            3355 ns/op         4096 B/op            1 allocs/op
 jsoniter >>> not supported <<<
 simdjson >>> not supported <<<
    gjson.Valid               2759 ns/op         4096 B/op            1 allocs/op
    sonic.Valid               2601 ns/op         4137 B/op            1 allocs/op

    sonic █████████████████▊ 2.55
    gjson ████████████████▊ 2.40
 fastjson █████████████▊ 1.98
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 jsoniter >>> not supported <<<
 simdjson >>> not supported <<<

Unmarshal single JSON record, read few keys
     json.Unmarshal          26451 ns/op        32386 B/op          342 allocs/op
       oj.Unmarshal          14673 ns/op        21228 B/op          376 allocs/op
 fastjson.Unmarshal           5112 ns/op        12034 B/op            7 allocs/op
 jsoniter.Unmarshal          18082 ns/op        20328 B/op          456 allocs/op
 simdjson.Unmarshal          20281 ns/op       130664 B/op           17 allocs/op
    gjson.Unmarshal           5730 ns/op        16368 B/op            9 allocs/op
    sonic.Unmarshal           6853 ns/op        12408 B/op           21 allocs/op

 fastjson ████████████████████████████████████▏ 5.17
    gjson ████████████████████████████████▎ 4.62
    sonic ███████████████████████████  3.86
       oj ████████████▌ 1.80
 jsoniter ██████████▏ 1.46
 simdjson █████████▏ 1.30
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal single JSON record, read all keys
     json.Unmarshal          25935 ns/op        32383 B/op          342 allocs/op
       oj.Unmarshal          23761 ns/op         9356 B/op          456 allocs/op
 fastjson.Unmarshal           6890 ns/op        12032 B/op            6 allocs/op
 jsoniter.Unmarshal           6851 ns/op         3141 B/op          170 allocs/op
 simdjson.Unmarshal          20079 ns/op       130664 B/op           17 allocs/op
    gjson.Unmarshal          11716 ns/op        17280 B/op            9 allocs/op
    sonic.Unmarshal          13978 ns/op        27571 B/op           59 allocs/op

 jsoniter ██████████████████████████▍ 3.79
 fastjson ██████████████████████████▎ 3.76
    gjson ███████████████▍ 2.21
    sonic ████████████▉ 1.86
 simdjson █████████  1.29
       oj ███████▋ 1.09
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal many JSON records from small file (100MB), read few keys
     json.Unmarshal      810659142 ns/op   1342678080 B/op     14015383 allocs/op
       oj.Unmarshal      506171036 ns/op    332606728 B/op     15622359 allocs/op
 fastjson.Unmarshal       86550090 ns/op      2536936 B/op        39519 allocs/op
 jsoniter.Unmarshal      756062388 ns/op    823291840 B/op     19383942 allocs/op
 simdjson.Unmarshal      442852817 ns/op   1470571053 B/op     14804665 allocs/op
    gjson.Unmarshal       85467019 ns/op    106043248 B/op        39452 allocs/op
    sonic.Unmarshal      136770409 ns/op    153261989 B/op       197417 allocs/op

    gjson ██████████████████████████████████████████████████████████████████▍ 9.49
 fastjson █████████████████████████████████████████████████████████████████▌ 9.37
    sonic █████████████████████████████████████████▍ 5.93
 simdjson ████████████▊ 1.83
       oj ███████████▏ 1.60
 jsoniter ███████▌ 1.07
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal many JSON records from small file (100MB), read all keys
     json.Unmarshal      849435275 ns/op   1342724600 B/op     14015546 allocs/op
       oj.Unmarshal      506252941 ns/op    332603816 B/op     15622344 allocs/op
 fastjson.Unmarshal      158685550 ns/op        19146 B/op          130 allocs/op
 jsoniter.Unmarshal      750909327 ns/op    823272992 B/op     19383878 allocs/op
 simdjson.Unmarshal      446542247 ns/op   1470558480 B/op     14804617 allocs/op
    gjson.Unmarshal      356248564 ns/op    151488816 B/op       118354 allocs/op
    sonic.Unmarshal      481333540 ns/op    897379128 B/op      2249740 allocs/op

 fastjson █████████████████████████████████████▍ 5.35
    gjson ████████████████▋ 2.38
 simdjson █████████████▎ 1.90
    sonic ████████████▎ 1.76
       oj ███████████▋ 1.68
 jsoniter ███████▉ 1.13
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal many JSON records from semi-large file (5GB), read few keys
     json.Unmarshal    12421884560 ns/op  10789397232 B/op    207849860 allocs/op
       oj.Unmarshal     7414648207 ns/op   4935951936 B/op    231680237 allocs/op
 fastjson.Unmarshal     1254794281 ns/op     37538776 B/op       585888 allocs/op
 jsoniter.Unmarshal    11155524576 ns/op  12209002696 B/op    287454216 allocs/op
 simdjson.Unmarshal     5549876536 ns/op  14849041800 B/op    219552747 allocs/op
    gjson.Unmarshal     1204381003 ns/op   1572570160 B/op       585034 allocs/op
    sonic.Unmarshal     1874895856 ns/op   2242112320 B/op      2925179 allocs/op

    gjson ████████████████████████████████████████████████████████████████████████▏ 10.31
 fastjson █████████████████████████████████████████████████████████████████████▎ 9.90
    sonic ██████████████████████████████████████████████▍ 6.63
 simdjson ███████████████▋ 2.24
       oj ███████████▋ 1.68
 jsoniter ███████▊ 1.11
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal many JSON records from semi-large file (5GB), read all keys
     json.Unmarshal    11414433798 ns/op  10789514064 B/op    207850047 allocs/op
       oj.Unmarshal     7007193695 ns/op   4935965408 B/op    231680189 allocs/op
 fastjson.Unmarshal     2365265811 ns/op        96344 B/op          852 allocs/op
 jsoniter.Unmarshal    10187945091 ns/op  12209108136 B/op    287454348 allocs/op
 simdjson.Unmarshal     5571789436 ns/op  14702181320 B/op    219553048 allocs/op
    gjson.Unmarshal     1225498088 ns/op   1572570160 B/op       585034 allocs/op
    sonic.Unmarshal     1876273615 ns/op   2242020304 B/op      2925169 allocs/op

    gjson █████████████████████████████████████████████████████████████████▏ 9.31
    sonic ██████████████████████████████████████████▌ 6.08
 fastjson █████████████████████████████████▊ 4.83
 simdjson ██████████████▎ 2.05
       oj ███████████▍ 1.63
 jsoniter ███████▊ 1.12
     json ▓▓▓▓▓▓▓ 1.00

Marshal custom data through an object builder
     json.Marshal             1040 ns/op          688 B/op           17 allocs/op
       oj.Builder              867 ns/op         1128 B/op           17 allocs/op
 fastjson.Marshal              606 ns/op          679 B/op            6 allocs/op
 jsoniter.Marshal              544 ns/op          432 B/op            7 allocs/op
 simdjson >>> not supported <<<
    gjson.Marshal             1886 ns/op         3600 B/op           32 allocs/op
    sonic.Marshal              735 ns/op         1803 B/op           14 allocs/op

 jsoniter █████████████▍ 1.91
 fastjson ████████████  1.72
    sonic █████████▉ 1.41
       oj ████████▍ 1.20
     json ▓▓▓▓▓▓▓ 1.00
    gjson ███▊ 0.55
 simdjson >>> not supported <<<

 Higher values (longer bars) are better in all cases. The bar graph compares the
 parsing performance. The lighter colored bar is the reference, the go json
 package.

Tests run on:
 OS:              Ubuntu 22.04.4 LTS
 Processor:       13th Gen Intel(R) Core(TM) i9-13900KF
 Cores:           32
 Processor Speed: 3.00GHz
 Memory:          32 GBValidate []byte
     json.Valid               6598 ns/op            0 B/op            0 allocs/op
       oj.Validate            2342 ns/op            0 B/op            0 allocs/op
 fastjson.Validate            2822 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               4674 ns/op         2187 B/op          100 allocs/op
 simdjson.Validate           19610 ns/op       118632 B/op           11 allocs/op
    gjson.Valid               2333 ns/op            0 B/op            0 allocs/op
    sonic.Valid               2032 ns/op            0 B/op            0 allocs/op

    sonic ██████████████████████▋ 3.25
    gjson ███████████████████▊ 2.83
       oj ███████████████████▋ 2.82
 fastjson ████████████████▎ 2.34
 jsoniter █████████▉ 1.41
     json ▓▓▓▓▓▓▓ 1.00
 simdjson ██▎ 0.34

Validate string
     json.Valid               6627 ns/op            0 B/op            0 allocs/op
       oj >>> not supported <<<
 fastjson.Validate            3355 ns/op         4096 B/op            1 allocs/op
 jsoniter >>> not supported <<<
 simdjson >>> not supported <<<
    gjson.Valid               2759 ns/op         4096 B/op            1 allocs/op
    sonic.Valid               2601 ns/op         4137 B/op            1 allocs/op

    sonic █████████████████▊ 2.55
    gjson ████████████████▊ 2.40
 fastjson █████████████▊ 1.98
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 jsoniter >>> not supported <<<
 simdjson >>> not supported <<<

Unmarshal single JSON record, read few keys
     json.Unmarshal          26451 ns/op        32386 B/op          342 allocs/op
       oj.Unmarshal          14673 ns/op        21228 B/op          376 allocs/op
 fastjson.Unmarshal           5112 ns/op        12034 B/op            7 allocs/op
 jsoniter.Unmarshal          18082 ns/op        20328 B/op          456 allocs/op
 simdjson.Unmarshal          20281 ns/op       130664 B/op           17 allocs/op
    gjson.Unmarshal           5730 ns/op        16368 B/op            9 allocs/op
    sonic.Unmarshal           6853 ns/op        12408 B/op           21 allocs/op

 fastjson ████████████████████████████████████▏ 5.17
    gjson ████████████████████████████████▎ 4.62
    sonic ███████████████████████████  3.86
       oj ████████████▌ 1.80
 jsoniter ██████████▏ 1.46
 simdjson █████████▏ 1.30
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal single JSON record, read all keys
     json.Unmarshal          25935 ns/op        32383 B/op          342 allocs/op
       oj.Unmarshal          23761 ns/op         9356 B/op          456 allocs/op
 fastjson.Unmarshal           6890 ns/op        12032 B/op            6 allocs/op
 jsoniter.Unmarshal           6851 ns/op         3141 B/op          170 allocs/op
 simdjson.Unmarshal          20079 ns/op       130664 B/op           17 allocs/op
    gjson.Unmarshal          11716 ns/op        17280 B/op            9 allocs/op
    sonic.Unmarshal          13978 ns/op        27571 B/op           59 allocs/op

 jsoniter ██████████████████████████▍ 3.79
 fastjson ██████████████████████████▎ 3.76
    gjson ███████████████▍ 2.21
    sonic ████████████▉ 1.86
 simdjson █████████  1.29
       oj ███████▋ 1.09
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal many JSON records from small file (100MB), read few keys
     json.Unmarshal      810659142 ns/op   1342678080 B/op     14015383 allocs/op
       oj.Unmarshal      506171036 ns/op    332606728 B/op     15622359 allocs/op
 fastjson.Unmarshal       86550090 ns/op      2536936 B/op        39519 allocs/op
 jsoniter.Unmarshal      756062388 ns/op    823291840 B/op     19383942 allocs/op
 simdjson.Unmarshal      442852817 ns/op   1470571053 B/op     14804665 allocs/op
    gjson.Unmarshal       85467019 ns/op    106043248 B/op        39452 allocs/op
    sonic.Unmarshal      136770409 ns/op    153261989 B/op       197417 allocs/op

    gjson ██████████████████████████████████████████████████████████████████▍ 9.49
 fastjson █████████████████████████████████████████████████████████████████▌ 9.37
    sonic █████████████████████████████████████████▍ 5.93
 simdjson ████████████▊ 1.83
       oj ███████████▏ 1.60
 jsoniter ███████▌ 1.07
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal many JSON records from small file (100MB), read all keys
     json.Unmarshal      849435275 ns/op   1342724600 B/op     14015546 allocs/op
       oj.Unmarshal      506252941 ns/op    332603816 B/op     15622344 allocs/op
 fastjson.Unmarshal      158685550 ns/op        19146 B/op          130 allocs/op
 jsoniter.Unmarshal      750909327 ns/op    823272992 B/op     19383878 allocs/op
 simdjson.Unmarshal      446542247 ns/op   1470558480 B/op     14804617 allocs/op
    gjson.Unmarshal      356248564 ns/op    151488816 B/op       118354 allocs/op
    sonic.Unmarshal      481333540 ns/op    897379128 B/op      2249740 allocs/op

 fastjson █████████████████████████████████████▍ 5.35
    gjson ████████████████▋ 2.38
 simdjson █████████████▎ 1.90
    sonic ████████████▎ 1.76
       oj ███████████▋ 1.68
 jsoniter ███████▉ 1.13
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal many JSON records from semi-large file (5GB), read few keys
     json.Unmarshal    12421884560 ns/op  10789397232 B/op    207849860 allocs/op
       oj.Unmarshal     7414648207 ns/op   4935951936 B/op    231680237 allocs/op
 fastjson.Unmarshal     1254794281 ns/op     37538776 B/op       585888 allocs/op
 jsoniter.Unmarshal    11155524576 ns/op  12209002696 B/op    287454216 allocs/op
 simdjson.Unmarshal     5549876536 ns/op  14849041800 B/op    219552747 allocs/op
    gjson.Unmarshal     1204381003 ns/op   1572570160 B/op       585034 allocs/op
    sonic.Unmarshal     1874895856 ns/op   2242112320 B/op      2925179 allocs/op

    gjson ████████████████████████████████████████████████████████████████████████▏ 10.31
 fastjson █████████████████████████████████████████████████████████████████████▎ 9.90
    sonic ██████████████████████████████████████████████▍ 6.63
 simdjson ███████████████▋ 2.24
       oj ███████████▋ 1.68
 jsoniter ███████▊ 1.11
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal many JSON records from semi-large file (5GB), read all keys
     json.Unmarshal    11414433798 ns/op  10789514064 B/op    207850047 allocs/op
       oj.Unmarshal     7007193695 ns/op   4935965408 B/op    231680189 allocs/op
 fastjson.Unmarshal     2365265811 ns/op        96344 B/op          852 allocs/op
 jsoniter.Unmarshal    10187945091 ns/op  12209108136 B/op    287454348 allocs/op
 simdjson.Unmarshal     5571789436 ns/op  14702181320 B/op    219553048 allocs/op
    gjson.Unmarshal     1225498088 ns/op   1572570160 B/op       585034 allocs/op
    sonic.Unmarshal     1876273615 ns/op   2242020304 B/op      2925169 allocs/op

    gjson █████████████████████████████████████████████████████████████████▏ 9.31
    sonic ██████████████████████████████████████████▌ 6.08
 fastjson █████████████████████████████████▊ 4.83
 simdjson ██████████████▎ 2.05
       oj ███████████▍ 1.63
 jsoniter ███████▊ 1.12
     json ▓▓▓▓▓▓▓ 1.00

Marshal custom data through an object builder
     json.Marshal             1040 ns/op          688 B/op           17 allocs/op
       oj.Builder              867 ns/op         1128 B/op           17 allocs/op
 fastjson.Marshal              606 ns/op          679 B/op            6 allocs/op
 jsoniter.Marshal              544 ns/op          432 B/op            7 allocs/op
 simdjson >>> not supported <<<
    gjson.Marshal             1886 ns/op         3600 B/op           32 allocs/op
    sonic.Marshal              735 ns/op         1803 B/op           14 allocs/op

 jsoniter █████████████▍ 1.91
 fastjson ████████████  1.72
    sonic █████████▉ 1.41
       oj ████████▍ 1.20
     json ▓▓▓▓▓▓▓ 1.00
    gjson ███▊ 0.55
 simdjson >>> not supported <<<

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
Validate []byte
     json.Valid              10637 ns/op            0 B/op            0 allocs/op
       oj.Validate            3828 ns/op            0 B/op            0 allocs/op
 fastjson.Validate            4918 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               6741 ns/op         2184 B/op          100 allocs/op
 simdjson.Validate    >>> Unsupported platform <<<
    gjson.Valid               3575 ns/op            0 B/op            0 allocs/op
 gjson-nv >>> not supported <<<
    sonic.Valid       >>> unsupported platform <<<

    gjson ████████████████████▊ 2.98
       oj ███████████████████▍ 2.78
 fastjson ███████████████▏ 2.16
 jsoniter ███████████  1.58
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<
 gjson-nv >>> not supported <<<
    sonic >>> unsupported platform <<<

Validate string
     json.Valid              10542 ns/op            0 B/op            0 allocs/op
       oj >>> not supported <<<
 fastjson.Validate            5271 ns/op         4096 B/op            1 allocs/op
 jsoniter >>> not supported <<<
 simdjson >>> not supported <<<
    gjson.Valid               3859 ns/op         4096 B/op            1 allocs/op
 gjson-nv >>> not supported <<<
    sonic.Valid       >>> unsupported platform <<<

    gjson ███████████████████  2.73
 fastjson ██████████████  2.00
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 jsoniter >>> not supported <<<
 simdjson >>> not supported <<<
 gjson-nv >>> not supported <<<
    sonic >>> unsupported platform <<<

Unmarshal single JSON record, read few keys
     json.Unmarshal          37327 ns/op        32383 B/op          342 allocs/op
       oj.Unmarshal          20609 ns/op        21228 B/op          376 allocs/op
 fastjson.Unmarshal           8106 ns/op        12034 B/op            7 allocs/op
 jsoniter.Unmarshal          27201 ns/op        31710 B/op          458 allocs/op
 simdjson.Unmarshal   >>> Unsupported platform <<<
    gjson.Unmarshal          13014 ns/op        16368 B/op            9 allocs/op
 gjson-nv.Unmarshal           9164 ns/op        16368 B/op            9 allocs/op
    sonic.Unmarshal   >>> unsupported platform <<<

 fastjson ████████████████████████████████▏ 4.60
 gjson-nv ████████████████████████████▌ 4.07
    gjson ████████████████████  2.87
       oj ████████████▋ 1.81
 jsoniter █████████▌ 1.37
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<
    sonic >>> unsupported platform <<<

Unmarshal single JSON record, read all keys
     json.Unmarshal          36273 ns/op        32384 B/op          342 allocs/op
       oj.Unmarshal          28529 ns/op         9350 B/op          456 allocs/op
 fastjson.Unmarshal          13196 ns/op        13608 B/op           73 allocs/op
 jsoniter.Unmarshal           8235 ns/op         2473 B/op           78 allocs/op
 simdjson.Unmarshal   >>> Unsupported platform <<<
    gjson.Unmarshal          22542 ns/op        18336 B/op           76 allocs/op
 gjson-nv.Unmarshal          18946 ns/op        18336 B/op           76 allocs/op
    sonic.Unmarshal   >>> unsupported platform <<<

 jsoniter ██████████████████████████████▊ 4.40
 fastjson ███████████████████▏ 2.75
 gjson-nv █████████████▍ 1.91
    gjson ███████████▎ 1.61
       oj ████████▉ 1.27
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<
    sonic >>> unsupported platform <<<

Unmarshal many JSON records from small file (100MB), read few keys
     json.Unmarshal     1047842958 ns/op   1342866616 B/op     14015551 allocs/op
       oj.Unmarshal      672275292 ns/op    334991856 B/op     15661829 allocs/op
 fastjson.Unmarshal      127498078 ns/op      2541930 B/op        39562 allocs/op
 jsoniter.Unmarshal      885599187 ns/op    818770040 B/op     19144319 allocs/op
 simdjson.Unmarshal   >>> Unsupported platform <<<
    gjson.Unmarshal      231340466 ns/op    106043814 B/op        39458 allocs/op
 gjson-nv.Unmarshal      118374027 ns/op    106043808 B/op        39458 allocs/op
    sonic.Unmarshal   >>> unsupported platform <<<

 gjson-nv █████████████████████████████████████████████████████████████▉ 8.85
 fastjson █████████████████████████████████████████████████████████▌ 8.22
    gjson ███████████████████████████████▋ 4.53
       oj ██████████▉ 1.56
 jsoniter ████████▎ 1.18
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<
    sonic >>> unsupported platform <<<

Unmarshal many JSON records from small file (100MB), read all keys
     json.Unmarshal     1047734750 ns/op   1342842392 B/op     14015467 allocs/op
       oj.Unmarshal      648902291 ns/op    334986768 B/op     15661816 allocs/op
 fastjson.Unmarshal      360945611 ns/op     66629186 B/op      2880080 allocs/op
 jsoniter.Unmarshal      880051270 ns/op    818788508 B/op     19144378 allocs/op
 simdjson.Unmarshal   >>> Unsupported platform <<<
    gjson.Unmarshal      654816375 ns/op    196303792 B/op      2998137 allocs/op
 gjson-nv.Unmarshal      538108250 ns/op    196304808 B/op      2998147 allocs/op
    sonic.Unmarshal   >>> unsupported platform <<<

 fastjson ████████████████████▎ 2.90
 gjson-nv █████████████▋ 1.95
       oj ███████████▎ 1.61
    gjson ███████████▏ 1.60
 jsoniter ████████▎ 1.19
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<
    sonic >>> unsupported platform <<<

Unmarshal many JSON records from semi-large file (5GB), read few keys
     json.Unmarshal    53056473292 ns/op  36386548424 B/op    700788890 allocs/op
       oj.Unmarshal    32322598084 ns/op  16773135168 B/op    783104397 allocs/op
 fastjson.Unmarshal     6394035333 ns/op    126338512 B/op      1973377 allocs/op
 jsoniter.Unmarshal    44008887250 ns/op  40937448416 B/op    957236531 allocs/op
 simdjson.Unmarshal   >>> Unsupported platform <<<
    gjson.Unmarshal    11654008875 ns/op   5302093504 B/op      1972869 allocs/op
 gjson-nv.Unmarshal     5878325459 ns/op   5302100704 B/op      1972944 allocs/op
    sonic.Unmarshal   >>> unsupported platform <<<

 gjson-nv ███████████████████████████████████████████████████████████████▏ 9.03
 fastjson ██████████████████████████████████████████████████████████  8.30
    gjson ███████████████████████████████▊ 4.55
       oj ███████████▍ 1.64
 jsoniter ████████▍ 1.21
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<
    sonic >>> unsupported platform <<<

Unmarshal many JSON records from semi-large file (5GB), read all keys
     json.Unmarshal    53295911667 ns/op  36386221904 B/op    700787637 allocs/op
       oj.Unmarshal    32432675375 ns/op  16773220768 B/op    783104715 allocs/op
 fastjson.Unmarshal    18188278917 ns/op   3329680480 B/op    143992799 allocs/op
 jsoniter.Unmarshal    43846998292 ns/op  40937603960 B/op    957237086 allocs/op
 simdjson.Unmarshal   >>> Unsupported platform <<<
    gjson.Unmarshal    11597700500 ns/op   5302096960 B/op      1972905 allocs/op
 gjson-nv.Unmarshal     5915174167 ns/op   5302103680 B/op      1972975 allocs/op
    sonic.Unmarshal   >>> unsupported platform <<<

 gjson-nv ███████████████████████████████████████████████████████████████  9.01
    gjson ████████████████████████████████▏ 4.60
 fastjson ████████████████████▌ 2.93
       oj ███████████▌ 1.64
 jsoniter ████████▌ 1.22
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<
    sonic >>> unsupported platform <<<

Marshal custom data through an object builder
     json.Marshal             1182 ns/op          688 B/op           17 allocs/op
       oj.Builder             1230 ns/op         1128 B/op           17 allocs/op
 fastjson.Marshal              803 ns/op          679 B/op            6 allocs/op
 jsoniter.Marshal              653 ns/op          432 B/op            7 allocs/op
 simdjson >>> not supported <<<
    gjson.Marshal             2404 ns/op         3601 B/op           32 allocs/op
 gjson-nv >>> not supported <<<
    sonic.Marshal     >>> unsupported platform <<<

 jsoniter ████████████▋ 1.81
 fastjson ██████████▎ 1.47
     json ▓▓▓▓▓▓▓ 1.00
       oj ██████▋ 0.96
    gjson ███▍ 0.49
 simdjson >>> not supported <<<
 gjson-nv >>> not supported <<<
    sonic >>> unsupported platform <<<

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
