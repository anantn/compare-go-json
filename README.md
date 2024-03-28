# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Validate []byte
     json.Validate            6363 ns/op            0 B/op            0 allocs/op
       oj.Validate            2267 ns/op            0 B/op            0 allocs/op
 fastjson.Validate            2686 ns/op            0 B/op            0 allocs/op
 jsoniter.Validate            4629 ns/op         2187 B/op          100 allocs/op
    goccy >>> not supported <<<
 simdjson >>> not supported <<<
    gjson.Validate            2351 ns/op            0 B/op            0 allocs/op
  gjson-v >>> not supported <<<
    sonic.Validate            1956 ns/op            0 B/op            0 allocs/op
  sonic-v >>> not supported <<<

    sonic ██████████████████████▊ 3.25
       oj ███████████████████▋ 2.81
    gjson ██████████████████▉ 2.71
 fastjson ████████████████▌ 2.37
 jsoniter █████████▌ 1.37
     json ▓▓▓▓▓▓▓ 1.00
    goccy >>> not supported <<<
 simdjson >>> not supported <<<
  gjson-v >>> not supported <<<
  sonic-v >>> not supported <<<

Validate string
     json.Validate            6118 ns/op            0 B/op            0 allocs/op
       oj >>> not supported <<<
 fastjson.Validate            3137 ns/op         4096 B/op            1 allocs/op
 jsoniter >>> not supported <<<
    goccy >>> not supported <<<
 simdjson >>> not supported <<<
    gjson.Validate            3008 ns/op         4096 B/op            1 allocs/op
  gjson-v >>> not supported <<<
    sonic.Validate            2652 ns/op         4136 B/op            1 allocs/op
  sonic-v >>> not supported <<<

    sonic ████████████████▏ 2.31
    gjson ██████████████▏ 2.03
 fastjson █████████████▋ 1.95
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 jsoniter >>> not supported <<<
    goccy >>> not supported <<<
 simdjson >>> not supported <<<
  gjson-v >>> not supported <<<
  sonic-v >>> not supported <<<

Unmarshal single JSON record, read few keys
     json.Unmarshal          18397 ns/op        15088 B/op           23 allocs/op
       oj >>> not supported <<<
 fastjson.Unmarshal           5184 ns/op        12034 B/op            7 allocs/op
 jsoniter.Unmarshal           8193 ns/op        14084 B/op           84 allocs/op
    goccy.Unmarshal           6007 ns/op        16302 B/op            7 allocs/op
 simdjson >>> not supported <<<
    gjson.Unmarshal           6067 ns/op        16368 B/op            9 allocs/op
  gjson-v.Unmarshal           8785 ns/op        16368 B/op            9 allocs/op
    sonic.Unmarshal           3714 ns/op         4954 B/op           12 allocs/op
  sonic-v.Unmarshal           5877 ns/op         4923 B/op           12 allocs/op

    sonic ██████████████████████████████████▋ 4.95
 fastjson ████████████████████████▊ 3.55
  sonic-v █████████████████████▉ 3.13
    goccy █████████████████████▍ 3.06
    gjson █████████████████████▏ 3.03
 jsoniter ███████████████▋ 2.25
  gjson-v ██████████████▋ 2.09
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 simdjson >>> not supported <<<

Unmarshal single JSON record, read all keys
     json.Unmarshal          24595 ns/op        17216 B/op           82 allocs/op
       oj.Unmarshal          22730 ns/op         9348 B/op          456 allocs/op
 fastjson.Unmarshal           8403 ns/op        13608 B/op           73 allocs/op
 jsoniter.Unmarshal           9441 ns/op        14534 B/op           84 allocs/op
    goccy.Unmarshal           7619 ns/op        16697 B/op            7 allocs/op
 simdjson.Unmarshal          19913 ns/op       130664 B/op           17 allocs/op
    gjson.Unmarshal          14295 ns/op        18336 B/op           76 allocs/op
  gjson-v.Unmarshal          16826 ns/op        18336 B/op           76 allocs/op
    sonic.Unmarshal          14553 ns/op        27859 B/op           59 allocs/op
  sonic-v.Unmarshal          16715 ns/op        27844 B/op           59 allocs/op

    goccy ██████████████████████▌ 3.23
 fastjson ████████████████████▍ 2.93
 jsoniter ██████████████████▏ 2.61
    gjson ████████████  1.72
    sonic ███████████▊ 1.69
  sonic-v ██████████▎ 1.47
  gjson-v ██████████▏ 1.46
 simdjson ████████▋ 1.24
       oj ███████▌ 1.08
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal many JSON records from small file (100MB), read few keys
     json.Unmarshal      467372922 ns/op    617757429 B/op        39516 allocs/op
       oj >>> not supported <<<
 fastjson.Unmarshal       80409422 ns/op      2536360 B/op        39514 allocs/op
 jsoniter.Unmarshal      203534150 ns/op     89303361 B/op      3984434 allocs/op
    goccy.Unmarshal       95470568 ns/op    106485005 B/op        39865 allocs/op
 simdjson >>> not supported <<<
    gjson.Unmarshal       89710635 ns/op    106043480 B/op        39454 allocs/op
  gjson-v.Unmarshal      178120903 ns/op    106043456 B/op        39454 allocs/op
    sonic.Unmarshal       36274602 ns/op      2889905 B/op        78902 allocs/op
  sonic-v.Unmarshal      112625593 ns/op      2877725 B/op        78902 allocs/op

    sonic ██████████████████████████████████████████████████████████████████████████████████████████▏ 12.88
 fastjson ████████████████████████████████████████▋ 5.81
    gjson ████████████████████████████████████▍ 5.21
    goccy ██████████████████████████████████▎ 4.90
  sonic-v █████████████████████████████  4.15
  gjson-v ██████████████████▎ 2.62
 jsoniter ████████████████  2.30
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 simdjson >>> not supported <<<

Unmarshal many JSON records from small file (100MB), read all keys
     json.Unmarshal      757494156 ns/op   1342729920 B/op     14015563 allocs/op
       oj.Unmarshal      501369057 ns/op    240450688 B/op     15464546 allocs/op
 fastjson.Unmarshal      220172734 ns/op     66615596 B/op      2879963 allocs/op
 jsoniter.Unmarshal      755346954 ns/op    819583812 B/op     19105381 allocs/op
    goccy.Unmarshal      685287052 ns/op    878252572 B/op     16664028 allocs/op
 simdjson.Unmarshal      410207186 ns/op   1470567216 B/op     14804656 allocs/op
    gjson.Unmarshal      439777900 ns/op    196303728 B/op      2998133 allocs/op
  gjson-v.Unmarshal      538387229 ns/op    196303792 B/op      2998135 allocs/op
    sonic.Unmarshal      489889139 ns/op    904561938 B/op      2249997 allocs/op
  sonic-v.Unmarshal      579777682 ns/op    903219324 B/op      2249947 allocs/op

 fastjson ████████████████████████  3.44
 simdjson ████████████▉ 1.85
    gjson ████████████  1.72
    sonic ██████████▊ 1.55
       oj ██████████▌ 1.51
  gjson-v █████████▊ 1.41
  sonic-v █████████▏ 1.31
    goccy ███████▋ 1.11
 jsoniter ███████  1.00
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal many JSON records from semi-large file (5GB), read few keys
     json.Unmarshal     7028787487 ns/op   9001289600 B/op       585114 allocs/op
       oj >>> not supported <<<
 fastjson.Unmarshal     1196543679 ns/op     37538544 B/op       585886 allocs/op
 jsoniter.Unmarshal     3111312902 ns/op   1328907144 B/op     59089570 allocs/op
    goccy.Unmarshal     1463383591 ns/op   1579021392 B/op       590757 allocs/op
 simdjson >>> not supported <<<
    gjson.Unmarshal     1396538882 ns/op   1572579712 B/op       585133 allocs/op
  gjson-v.Unmarshal     2666036228 ns/op   1572577696 B/op       585112 allocs/op
    sonic.Unmarshal      545581639 ns/op     42888076 B/op      1170121 allocs/op
  sonic-v.Unmarshal     1663471749 ns/op     42665216 B/op      1170116 allocs/op

    sonic ██████████████████████████████████████████████████████████████████████████████████████████▏ 12.88
 fastjson █████████████████████████████████████████  5.87
    gjson ███████████████████████████████████▏ 5.03
    goccy █████████████████████████████████▌ 4.80
  sonic-v █████████████████████████████▌ 4.23
  gjson-v ██████████████████▍ 2.64
 jsoniter ███████████████▊ 2.26
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 simdjson >>> not supported <<<

Unmarshal many JSON records from semi-large file (5GB), read all keys
     json.Unmarshal    11291821787 ns/op  19753431152 B/op    207850380 allocs/op
       oj.Unmarshal     7426905731 ns/op   3569351040 B/op    229340164 allocs/op
 fastjson.Unmarshal     3240012592 ns/op    987634936 B/op     42708213 allocs/op
 jsoniter.Unmarshal    11143600359 ns/op  12153694288 B/op    283332233 allocs/op
    goccy.Unmarshal    10126664531 ns/op  13022527440 B/op    247124111 allocs/op
 simdjson.Unmarshal     5451909233 ns/op  14796781456 B/op    219553417 allocs/op
    gjson.Unmarshal     6156039514 ns/op   2911123376 B/op     44462433 allocs/op
  gjson-v.Unmarshal     7503099686 ns/op   2911123344 B/op     44462433 allocs/op
    sonic.Unmarshal     5448013537 ns/op  13035648944 B/op     33346902 allocs/op
  sonic-v.Unmarshal     6740288756 ns/op  13035567200 B/op     33346902 allocs/op

 fastjson ████████████████████████▍ 3.49
    sonic ██████████████▌ 2.07
 simdjson ██████████████▍ 2.07
    gjson ████████████▊ 1.83
  sonic-v ███████████▋ 1.68
       oj ██████████▋ 1.52
  gjson-v ██████████▌ 1.50
    goccy ███████▊ 1.12
 jsoniter ███████  1.01
     json ▓▓▓▓▓▓▓ 1.00

Marshal custom data through an object builder
     json.Marshal             1009 ns/op          688 B/op           17 allocs/op
       oj.Builder              825 ns/op         1128 B/op           17 allocs/op
 fastjson.Marshal              583 ns/op          679 B/op            6 allocs/op
 jsoniter.Marshal              484 ns/op          432 B/op            7 allocs/op
    goccy.Marshal              548 ns/op          416 B/op            2 allocs/op
 simdjson >>> not supported <<<
    gjson.Marshal             1730 ns/op         3600 B/op           32 allocs/op
  gjson-v >>> not supported <<<
    sonic.Marshal              693 ns/op         1803 B/op           14 allocs/op
  sonic-v >>> not supported <<<

 jsoniter ██████████████▌ 2.08
    goccy ████████████▉ 1.84
 fastjson ████████████  1.73
    sonic ██████████▏ 1.46
       oj ████████▌ 1.22
     json ▓▓▓▓▓▓▓ 1.00
    gjson ████  0.58
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
     json.Validate           10425 ns/op            0 B/op            0 allocs/op
       oj.Validate            3766 ns/op            0 B/op            0 allocs/op
 fastjson.Validate            4886 ns/op            0 B/op            0 allocs/op
 jsoniter.Validate            6596 ns/op         2184 B/op          100 allocs/op
 simdjson >>> not supported <<<
    gjson.Validate            3522 ns/op            0 B/op            0 allocs/op
  gjson-v >>> not supported <<<
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

    gjson ████████████████████▋ 2.96
       oj ███████████████████▍ 2.77
 fastjson ██████████████▉ 2.13
 jsoniter ███████████  1.58
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> not supported <<<
  gjson-v >>> not supported <<<
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

Validate string
     json.Validate           10407 ns/op            0 B/op            0 allocs/op
       oj >>> not supported <<<
 fastjson.Validate            5198 ns/op         4096 B/op            1 allocs/op
 jsoniter >>> not supported <<<
 simdjson >>> not supported <<<
    gjson.Validate            3814 ns/op         4096 B/op            1 allocs/op
  gjson-v >>> not supported <<<
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

    gjson ███████████████████  2.73
 fastjson ██████████████  2.00
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 jsoniter >>> not supported <<<
 simdjson >>> not supported <<<
  gjson-v >>> not supported <<<
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

Unmarshal single JSON record, read few keys
     json.Unmarshal          28072 ns/op        15088 B/op           23 allocs/op
       oj >>> not supported <<<
 fastjson.Unmarshal           7903 ns/op        12034 B/op            7 allocs/op
 jsoniter.Unmarshal          11950 ns/op        14061 B/op           84 allocs/op
 simdjson >>> not supported <<<
    gjson.Unmarshal           8983 ns/op        16368 B/op            9 allocs/op
  gjson-v.Unmarshal          12590 ns/op        16368 B/op            9 allocs/op
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

 fastjson ████████████████████████▊ 3.55
    gjson █████████████████████▉ 3.13
 jsoniter ████████████████▍ 2.35
  gjson-v ███████████████▌ 2.23
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 simdjson >>> not supported <<<
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

Unmarshal single JSON record, read all keys
     json.Unmarshal          34766 ns/op        17216 B/op           82 allocs/op
       oj.Unmarshal          28234 ns/op         9356 B/op          456 allocs/op
 fastjson.Unmarshal          13086 ns/op        13608 B/op           73 allocs/op
 jsoniter.Unmarshal          13656 ns/op        14510 B/op           84 allocs/op
 simdjson.Unmarshal   >>> Unsupported platform <<<
    gjson.Unmarshal          18840 ns/op        18336 B/op           76 allocs/op
  gjson-v.Unmarshal          22183 ns/op        18336 B/op           76 allocs/op
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

 fastjson ██████████████████▌ 2.66
 jsoniter █████████████████▊ 2.55
    gjson ████████████▉ 1.85
  gjson-v ██████████▉ 1.57
       oj ████████▌ 1.23
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

Unmarshal many JSON records from small file (100MB), read few keys
     json.Unmarshal      679465187 ns/op    617757264 B/op        39514 allocs/op
       oj >>> not supported <<<
 fastjson.Unmarshal      126690156 ns/op      2541912 B/op        39562 allocs/op
 jsoniter.Unmarshal      284395791 ns/op     89289068 B/op      3984412 allocs/op
 simdjson >>> not supported <<<
    gjson.Unmarshal      120578773 ns/op    106043787 B/op        39458 allocs/op
  gjson-v.Unmarshal      232864691 ns/op    106044121 B/op        39461 allocs/op
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

    gjson ███████████████████████████████████████▍ 5.64
 fastjson █████████████████████████████████████▌ 5.36
  gjson-v ████████████████████▍ 2.92
 jsoniter ████████████████▋ 2.39
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 simdjson >>> not supported <<<
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

Unmarshal many JSON records from small file (100MB), read all keys
     json.Unmarshal     1071344292 ns/op   1342830280 B/op     14015425 allocs/op
       oj.Unmarshal      679421292 ns/op    240318632 B/op     15464599 allocs/op
 fastjson.Unmarshal      373046333 ns/op     66629426 B/op      2880083 allocs/op
 jsoniter.Unmarshal      933623916 ns/op    818169060 B/op     19104942 allocs/op
 simdjson.Unmarshal   >>> Unsupported platform <<<
    gjson.Unmarshal      557989687 ns/op    196304368 B/op      2998142 allocs/op
  gjson-v.Unmarshal      680339375 ns/op    196304048 B/op      2998140 allocs/op
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

 fastjson ████████████████████  2.87
    gjson █████████████▍ 1.92
       oj ███████████  1.58
  gjson-v ███████████  1.57
 jsoniter ████████  1.15
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

Unmarshal many JSON records from semi-large file (5GB), read few keys
     json.Unmarshal    34785656292 ns/op  27483772736 B/op      1972575 allocs/op
       oj >>> not supported <<<
 fastjson.Unmarshal     6327255167 ns/op    126338952 B/op      1973382 allocs/op
 jsoniter.Unmarshal    14282073875 ns/op   4448093216 B/op    199225215 allocs/op
 simdjson >>> not supported <<<
    gjson.Unmarshal     6163322166 ns/op   5302143232 B/op      1973387 allocs/op
  gjson-v.Unmarshal    11875501458 ns/op   5302134240 B/op      1973263 allocs/op
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

    gjson ███████████████████████████████████████▌ 5.64
 fastjson ██████████████████████████████████████▍ 5.50
  gjson-v ████████████████████▌ 2.93
 jsoniter █████████████████  2.44
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 simdjson >>> not supported <<<
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

Unmarshal many JSON records from semi-large file (5GB), read all keys
     json.Unmarshal    52561558709 ns/op  63743253456 B/op    700784347 allocs/op
       oj.Unmarshal    33042919708 ns/op  12039254512 B/op    773242395 allocs/op
 fastjson.Unmarshal    18701755000 ns/op   3329707360 B/op    143993057 allocs/op
 jsoniter.Unmarshal    50629508458 ns/op  40911557312 B/op    955276303 allocs/op
 simdjson.Unmarshal   >>> Unsupported platform <<<
    gjson.Unmarshal    27784123833 ns/op   9815241360 B/op    149910315 allocs/op
  gjson-v.Unmarshal    33542825541 ns/op   9815247696 B/op    149910413 allocs/op
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

 fastjson ███████████████████▋ 2.81
    gjson █████████████▏ 1.89
       oj ███████████▏ 1.59
  gjson-v ██████████▉ 1.57
 jsoniter ███████▎ 1.04
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> Unsupported platform <<<
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

Marshal custom data through an object builder
     json.Marshal             1246 ns/op          688 B/op           17 allocs/op
       oj.Builder             1390 ns/op         1129 B/op           17 allocs/op
 fastjson.Marshal              860 ns/op          679 B/op            6 allocs/op
 jsoniter.Marshal              698 ns/op          432 B/op            7 allocs/op
 simdjson >>> not supported <<<
    gjson.Marshal             2715 ns/op         3602 B/op           32 allocs/op
  gjson-v >>> not supported <<<
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

 jsoniter ████████████▍ 1.79
 fastjson ██████████▏ 1.45
     json ▓▓▓▓▓▓▓ 1.00
       oj ██████▎ 0.90
    gjson ███▏ 0.46
 simdjson >>> not supported <<<
  gjson-v >>> not supported <<<
    sonic >>> not supported <<<
  sonic-v >>> not supported <<<

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
