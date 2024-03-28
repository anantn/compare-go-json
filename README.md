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
NOTE: sonic is not supported on ARM64, falls back to `encoding/json`.

```

```
