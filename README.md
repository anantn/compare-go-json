# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Validate []byte
     json.Validate            6594 ns/op            0 B/op            0 allocs/op
       oj.Validate            2271 ns/op            0 B/op            0 allocs/op
 fastjson.Validate            2630 ns/op            0 B/op            0 allocs/op
 jsoniter.Validate            4487 ns/op         2187 B/op          100 allocs/op
 simdjson >>> not supported <<<
    gjson.Validate            2216 ns/op            0 B/op            0 allocs/op
  gjson-v >>> not supported <<<
    sonic.Validate            1979 ns/op            0 B/op            0 allocs/op
  sonic-v >>> not supported <<<

    sonic ███████████████████████▎ 3.33
    gjson ████████████████████▊ 2.98
       oj ████████████████████▎ 2.90
 fastjson █████████████████▌ 2.51
 jsoniter ██████████▎ 1.47
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> not supported <<<
  gjson-v >>> not supported <<<
  sonic-v >>> not supported <<<

Validate string
     json.Validate            6289 ns/op            0 B/op            0 allocs/op
       oj >>> not supported <<<
 fastjson.Validate            3145 ns/op         4096 B/op            1 allocs/op
 jsoniter >>> not supported <<<
 simdjson >>> not supported <<<
    gjson.Validate            2709 ns/op         4096 B/op            1 allocs/op
  gjson-v >>> not supported <<<
    sonic.Validate            2526 ns/op         4130 B/op            1 allocs/op
  sonic-v >>> not supported <<<

    sonic █████████████████▍ 2.49
    gjson ████████████████▎ 2.32
 fastjson █████████████▉ 2.00
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 jsoniter >>> not supported <<<
 simdjson >>> not supported <<<
  gjson-v >>> not supported <<<
  sonic-v >>> not supported <<<

Unmarshal single JSON record, read few keys
     json.Unmarshal          18540 ns/op        15088 B/op           23 allocs/op
       oj >>> not supported <<<
 fastjson.Unmarshal           4888 ns/op        12034 B/op            7 allocs/op
 jsoniter.Unmarshal           8142 ns/op        14080 B/op           84 allocs/op
 simdjson >>> not supported <<<
    gjson.Unmarshal           5665 ns/op        16368 B/op            9 allocs/op
  gjson-v.Unmarshal           8112 ns/op        16368 B/op            9 allocs/op
    sonic.Unmarshal           3365 ns/op         4945 B/op           12 allocs/op
  sonic-v.Unmarshal           5490 ns/op         4901 B/op           12 allocs/op

    sonic ██████████████████████████████████████▌ 5.51
 fastjson ██████████████████████████▌ 3.79
  sonic-v ███████████████████████▋ 3.38
    gjson ██████████████████████▉ 3.27
  gjson-v ███████████████▉ 2.29
 jsoniter ███████████████▉ 2.28
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 simdjson >>> not supported <<<

Unmarshal single JSON record, read all keys
     json.Unmarshal          24506 ns/op        17216 B/op           82 allocs/op
       oj.Unmarshal          22215 ns/op         9354 B/op          456 allocs/op
 fastjson.Unmarshal           8175 ns/op        13608 B/op           73 allocs/op
 jsoniter.Unmarshal           9249 ns/op        14530 B/op           84 allocs/op
 simdjson.Unmarshal          19734 ns/op       130665 B/op           17 allocs/op
    gjson.Unmarshal          12397 ns/op        18336 B/op           76 allocs/op
  gjson-v.Unmarshal          15443 ns/op        18336 B/op           76 allocs/op
    sonic.Unmarshal          13726 ns/op        27671 B/op           59 allocs/op
  sonic-v.Unmarshal          15804 ns/op        27643 B/op           59 allocs/op

 fastjson ████████████████████▉ 3.00
 jsoniter ██████████████████▌ 2.65
    gjson █████████████▊ 1.98
    sonic ████████████▍ 1.79
  gjson-v ███████████  1.59
  sonic-v ██████████▊ 1.55
 simdjson ████████▋ 1.24
       oj ███████▋ 1.10
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal many JSON records from small file (100MB), read few keys
     json.Unmarshal      482090147 ns/op    617757333 B/op        39515 allocs/op
       oj >>> not supported <<<
 fastjson.Unmarshal       79797412 ns/op      2536360 B/op        39514 allocs/op
 jsoniter.Unmarshal      202497253 ns/op     89274800 B/op      3984420 allocs/op
 simdjson >>> not supported <<<
    gjson.Unmarshal       83467680 ns/op    106043665 B/op        39456 allocs/op
  gjson-v.Unmarshal      165406999 ns/op    106043460 B/op        39454 allocs/op
    sonic.Unmarshal       33466102 ns/op      2885460 B/op        78902 allocs/op
  sonic-v.Unmarshal      108460594 ns/op      2873907 B/op        78901 allocs/op

    sonic ████████████████████████████████████████████████████████████████████████████████████████████████████▊ 14.41
 fastjson ██████████████████████████████████████████▎ 6.04
    gjson ████████████████████████████████████████▍ 5.78
  sonic-v ███████████████████████████████  4.44
  gjson-v ████████████████████▍ 2.91
 jsoniter ████████████████▋ 2.38
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 simdjson >>> not supported <<<

Unmarshal many JSON records from small file (100MB), read all keys
     json.Unmarshal      782114089 ns/op   1342711664 B/op     14015499 allocs/op
       oj.Unmarshal      480794580 ns/op    240450586 B/op     15464546 allocs/op
 fastjson.Unmarshal      213254770 ns/op     66615228 B/op      2879959 allocs/op
 jsoniter.Unmarshal      708829163 ns/op    819204196 B/op     19105088 allocs/op
 simdjson.Unmarshal      414385352 ns/op   1470573650 B/op     14804672 allocs/op
    gjson.Unmarshal      394035758 ns/op    196303136 B/op      2998129 allocs/op
  gjson-v.Unmarshal      480439662 ns/op    196303152 B/op      2998129 allocs/op
    sonic.Unmarshal      480761215 ns/op    900554709 B/op      2249814 allocs/op
  sonic-v.Unmarshal      560824459 ns/op    899300576 B/op      2249792 allocs/op

 fastjson █████████████████████████▋ 3.67
    gjson █████████████▉ 1.98
 simdjson █████████████▏ 1.89
  gjson-v ███████████▍ 1.63
    sonic ███████████▍ 1.63
       oj ███████████▍ 1.63
  sonic-v █████████▊ 1.39
 jsoniter ███████▋ 1.10
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal many JSON records from semi-large file (5GB), read few keys
     json.Unmarshal     7331440390 ns/op   9001289600 B/op       585114 allocs/op
       oj >>> not supported <<<
 fastjson.Unmarshal     1177966402 ns/op     37538344 B/op       585883 allocs/op
 jsoniter.Unmarshal     3218532432 ns/op   1328420600 B/op     59089318 allocs/op
 simdjson >>> not supported <<<
    gjson.Unmarshal     1332135754 ns/op   1572574624 B/op       585080 allocs/op
  gjson-v.Unmarshal     2559855514 ns/op   1572576736 B/op       585102 allocs/op
    sonic.Unmarshal      504974529 ns/op     42776172 B/op      1170117 allocs/op
  sonic-v.Unmarshal     1596212163 ns/op     42696208 B/op      1170111 allocs/op

    sonic █████████████████████████████████████████████████████████████████████████████████████████████████████▋ 14.52
 fastjson ███████████████████████████████████████████▌ 6.22
    gjson ██████████████████████████████████████▌ 5.50
  sonic-v ████████████████████████████████▏ 4.59
  gjson-v ████████████████████  2.86
 jsoniter ███████████████▉ 2.28
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 simdjson >>> not supported <<<

Unmarshal many JSON records from semi-large file (5GB), read all keys
     json.Unmarshal    11638924518 ns/op  19753331088 B/op    207850030 allocs/op
       oj.Unmarshal     7193921166 ns/op   3569293936 B/op    229339958 allocs/op
 fastjson.Unmarshal     3197459131 ns/op    987638968 B/op     42708261 allocs/op
 jsoniter.Unmarshal    10662782266 ns/op  12149748248 B/op    283330412 allocs/op
 simdjson.Unmarshal     5502750873 ns/op  14807057816 B/op    219552682 allocs/op
    gjson.Unmarshal     5694964297 ns/op   2911123344 B/op     44462433 allocs/op
  gjson-v.Unmarshal     7001723533 ns/op   2911123376 B/op     44462433 allocs/op
    sonic.Unmarshal     5389681619 ns/op  13035730336 B/op     33346898 allocs/op
  sonic-v.Unmarshal     6592474782 ns/op  13035485456 B/op     33346902 allocs/op

 fastjson █████████████████████████▍ 3.64
    sonic ███████████████  2.16
 simdjson ██████████████▊ 2.12
    gjson ██████████████▎ 2.04
  sonic-v ████████████▎ 1.77
  gjson-v ███████████▋ 1.66
       oj ███████████▎ 1.62
 jsoniter ███████▋ 1.09
     json ▓▓▓▓▓▓▓ 1.00

Marshal custom data through an object builder
     json.Marshal             1014 ns/op          688 B/op           17 allocs/op
       oj.Builder              814 ns/op         1128 B/op           17 allocs/op
 fastjson.Marshal              574 ns/op          679 B/op            6 allocs/op
 jsoniter.Marshal              499 ns/op          432 B/op            7 allocs/op
 simdjson >>> not supported <<<
    gjson.Marshal             1696 ns/op         3600 B/op           32 allocs/op
  gjson-v >>> not supported <<<
    sonic.Marshal              692 ns/op         1803 B/op           14 allocs/op
  sonic-v >>> not supported <<<

 jsoniter ██████████████▏ 2.03
 fastjson ████████████▎ 1.77
    sonic ██████████▎ 1.47
       oj ████████▋ 1.25
     json ▓▓▓▓▓▓▓ 1.00
    gjson ████▏ 0.60
 simdjson >>> not supported <<<
  gjson-v >>> not supported <<<
  sonic-v >>> not supported <<<

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
