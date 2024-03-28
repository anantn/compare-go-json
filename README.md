# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Validate []byte
     json.Validate            6026 ns/op            0 B/op            0 allocs/op
       oj.Validate            2273 ns/op            0 B/op            0 allocs/op
 fastjson.Validate            2584 ns/op            0 B/op            0 allocs/op
 jsoniter.Validate            4575 ns/op         2187 B/op          100 allocs/op
    goccy.Validate           19374 ns/op        27649 B/op          463 allocs/op
  segment.Validate            4391 ns/op            0 B/op            0 allocs/op
 simdjson >>> not supported <<<
    gjson.Validate            2411 ns/op            0 B/op            0 allocs/op
  gjson-v >>> not supported <<<
    sonic.Validate            1999 ns/op            0 B/op            0 allocs/op
  sonic-v >>> not supported <<<

    sonic █████████████████████  3.01
       oj ██████████████████▌ 2.65
    gjson █████████████████▍ 2.50
 fastjson ████████████████▎ 2.33
  segment █████████▌ 1.37
 jsoniter █████████▏ 1.32
     json ▓▓▓▓▓▓▓ 1.00
    goccy ██▏ 0.31
 simdjson >>> not supported <<<
  gjson-v >>> not supported <<<
  sonic-v >>> not supported <<<

Validate string
     json.Validate            6051 ns/op            0 B/op            0 allocs/op
       oj >>> not supported <<<
 fastjson.Validate            3148 ns/op         4096 B/op            1 allocs/op
 jsoniter >>> not supported <<<
    goccy >>> not supported <<<
  segment >>> not supported <<<
 simdjson >>> not supported <<<
    gjson.Validate            2895 ns/op         4096 B/op            1 allocs/op
  gjson-v >>> not supported <<<
    sonic.Validate            2602 ns/op         4136 B/op            1 allocs/op
  sonic-v >>> not supported <<<

    sonic ████████████████▎ 2.33
    gjson ██████████████▋ 2.09
 fastjson █████████████▍ 1.92
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 jsoniter >>> not supported <<<
    goccy >>> not supported <<<
  segment >>> not supported <<<
 simdjson >>> not supported <<<
  gjson-v >>> not supported <<<
  sonic-v >>> not supported <<<

Unmarshal single JSON record, read few keys
     json.Unmarshal          18453 ns/op        15088 B/op           23 allocs/op
       oj >>> not supported <<<
 fastjson.Unmarshal           5174 ns/op        12034 B/op            7 allocs/op
 jsoniter.Unmarshal           8202 ns/op        14084 B/op           84 allocs/op
    goccy.Unmarshal           5995 ns/op        16301 B/op            7 allocs/op
  segment.Unmarshal          10193 ns/op        12064 B/op           12 allocs/op
 simdjson >>> not supported <<<
    gjson.Unmarshal           6098 ns/op        16368 B/op            9 allocs/op
  gjson-v.Unmarshal           8916 ns/op        16368 B/op            9 allocs/op
    sonic.Unmarshal           3496 ns/op         4958 B/op           12 allocs/op
  sonic-v.Unmarshal           5631 ns/op         4938 B/op           12 allocs/op

    sonic ████████████████████████████████████▉ 5.28
 fastjson ████████████████████████▉ 3.57
  sonic-v ██████████████████████▉ 3.28
    goccy █████████████████████▌ 3.08
    gjson █████████████████████▏ 3.03
 jsoniter ███████████████▋ 2.25
  gjson-v ██████████████▍ 2.07
  segment ████████████▋ 1.81
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 simdjson >>> not supported <<<

Unmarshal single JSON record, read all keys
     json.Unmarshal          24495 ns/op        17216 B/op           82 allocs/op
       oj.Unmarshal          23419 ns/op         9355 B/op          456 allocs/op
 fastjson.Unmarshal           8561 ns/op        13608 B/op           73 allocs/op
 jsoniter.Unmarshal           9481 ns/op        14534 B/op           84 allocs/op
    goccy.Unmarshal           7538 ns/op        16696 B/op            7 allocs/op
  segment.Unmarshal          18847 ns/op        13680 B/op           73 allocs/op
 simdjson.Unmarshal          20709 ns/op       130664 B/op           17 allocs/op
    gjson.Unmarshal          14628 ns/op        18336 B/op           76 allocs/op
  gjson-v.Unmarshal          17433 ns/op        18336 B/op           76 allocs/op
    sonic.Unmarshal          14289 ns/op        27834 B/op           59 allocs/op
  sonic-v.Unmarshal          16313 ns/op        27801 B/op           59 allocs/op

    goccy ██████████████████████▋ 3.25
 fastjson ████████████████████  2.86
 jsoniter ██████████████████  2.58
    sonic ███████████▉ 1.71
    gjson ███████████▋ 1.67
  sonic-v ██████████▌ 1.50
  gjson-v █████████▊ 1.41
  segment █████████  1.30
 simdjson ████████▎ 1.18
       oj ███████▎ 1.05
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal many JSON records from small file (100MB), read few keys
     json.Unmarshal      470375823 ns/op    617757365 B/op        39515 allocs/op
       oj >>> not supported <<<
 fastjson.Unmarshal       80481329 ns/op      2536367 B/op        39514 allocs/op
 jsoniter.Unmarshal      203944665 ns/op     89308160 B/op      3984436 allocs/op
    goccy.Unmarshal       97387011 ns/op    106490052 B/op        39872 allocs/op
  segment.Unmarshal      213308391 ns/op      2528857 B/op        39450 allocs/op
 simdjson >>> not supported <<<
    gjson.Unmarshal       91750537 ns/op    106043384 B/op        39453 allocs/op
  gjson-v.Unmarshal      177337515 ns/op    106043408 B/op        39454 allocs/op
    sonic.Unmarshal       33406256 ns/op      2898768 B/op        78902 allocs/op
  sonic-v.Unmarshal      108193310 ns/op      2891251 B/op        78902 allocs/op

    sonic ██████████████████████████████████████████████████████████████████████████████████████████████████▌ 14.08
 fastjson ████████████████████████████████████████▉ 5.84
    gjson ███████████████████████████████████▉ 5.13
    goccy █████████████████████████████████▊ 4.83
  sonic-v ██████████████████████████████▍ 4.35
  gjson-v ██████████████████▌ 2.65
 jsoniter ████████████████▏ 2.31
  segment ███████████████▍ 2.21
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 simdjson >>> not supported <<<

Unmarshal many JSON records from small file (100MB), read all keys
     json.Unmarshal      765103746 ns/op   1342703584 B/op     14015473 allocs/op
       oj.Unmarshal      503974730 ns/op    240457208 B/op     15464569 allocs/op
 fastjson.Unmarshal      224019983 ns/op     66615536 B/op      2879962 allocs/op
 jsoniter.Unmarshal      749834133 ns/op    819581284 B/op     19105290 allocs/op
    goccy.Unmarshal      701637344 ns/op    878228408 B/op     16664079 allocs/op
  segment.Unmarshal     2869213640 ns/op   6769774536 B/op     19882552 allocs/op
 simdjson.Unmarshal      410779511 ns/op   1470581666 B/op     14804697 allocs/op
    gjson.Unmarshal      434132718 ns/op    196303808 B/op      2998134 allocs/op
  gjson-v.Unmarshal      532174259 ns/op    196303432 B/op      2998131 allocs/op
    sonic.Unmarshal      485259770 ns/op    902551797 B/op      2249918 allocs/op
  sonic-v.Unmarshal      573220495 ns/op    902307588 B/op      2249926 allocs/op

 fastjson ███████████████████████▉ 3.42
 simdjson █████████████  1.86
    gjson ████████████▎ 1.76
    sonic ███████████  1.58
       oj ██████████▋ 1.52
  gjson-v ██████████  1.44
  sonic-v █████████▎ 1.33
    goccy ███████▋ 1.09
 jsoniter ███████▏ 1.02
     json ▓▓▓▓▓▓▓ 1.00
  segment █▊ 0.27

Unmarshal many JSON records from semi-large file (5GB), read few keys
     json.Unmarshal     7112161593 ns/op   9001289984 B/op       585118 allocs/op
       oj >>> not supported <<<
 fastjson.Unmarshal     1209324930 ns/op     37538448 B/op       585885 allocs/op
 jsoniter.Unmarshal     3112812418 ns/op   1328841600 B/op     59089529 allocs/op
    goccy.Unmarshal     1417517345 ns/op   1578789912 B/op       590371 allocs/op
  segment.Unmarshal     3176187308 ns/op     37446368 B/op       585036 allocs/op
 simdjson >>> not supported <<<
    gjson.Unmarshal     1375831933 ns/op   1572577216 B/op       585107 allocs/op
  gjson-v.Unmarshal     2658360807 ns/op   1572577120 B/op       585106 allocs/op
    sonic.Unmarshal      503286651 ns/op     42933570 B/op      1170122 allocs/op
  sonic-v.Unmarshal     1609387140 ns/op     42788368 B/op      1170122 allocs/op

    sonic ██████████████████████████████████████████████████████████████████████████████████████████████████▉ 14.13
 fastjson █████████████████████████████████████████▏ 5.88
    gjson ████████████████████████████████████▏ 5.17
    goccy ███████████████████████████████████  5.02
  sonic-v ██████████████████████████████▉ 4.42
  gjson-v ██████████████████▋ 2.68
 jsoniter ███████████████▉ 2.28
  segment ███████████████▋ 2.24
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 simdjson >>> not supported <<<

Unmarshal many JSON records from semi-large file (5GB), read all keys
     json.Unmarshal    11300755035 ns/op  19753273488 B/op    207849832 allocs/op
       oj.Unmarshal     7383966153 ns/op   3569357376 B/op    229340181 allocs/op
 fastjson.Unmarshal     3234159040 ns/op    987637832 B/op     42708240 allocs/op
 jsoniter.Unmarshal    10899772721 ns/op  12152888280 B/op    283331973 allocs/op
    goccy.Unmarshal    10258301979 ns/op  13020907864 B/op    247119571 allocs/op
  segment.Unmarshal    39338510744 ns/op 100399629632 B/op    294863183 allocs/op
 simdjson.Unmarshal     5495676617 ns/op  14838625600 B/op    219553001 allocs/op
    gjson.Unmarshal     6052387327 ns/op   2911123376 B/op     44462433 allocs/op
  gjson-v.Unmarshal     7488225508 ns/op   2911123328 B/op     44462433 allocs/op
    sonic.Unmarshal     5343155019 ns/op  13035444144 B/op     33346897 allocs/op
  sonic-v.Unmarshal     6556782846 ns/op  13035607808 B/op     33346899 allocs/op

 fastjson ████████████████████████▍ 3.49
    sonic ██████████████▊ 2.11
 simdjson ██████████████▍ 2.06
    gjson █████████████  1.87
  sonic-v ████████████  1.72
       oj ██████████▋ 1.53
  gjson-v ██████████▌ 1.51
    goccy ███████▋ 1.10
 jsoniter ███████▎ 1.04
     json ▓▓▓▓▓▓▓ 1.00
  segment ██  0.29

Marshal custom data through an object builder
     json.Marshal             1079 ns/op          688 B/op           17 allocs/op
       oj.Builder              837 ns/op         1128 B/op           17 allocs/op
 fastjson.Marshal              609 ns/op          679 B/op            6 allocs/op
 jsoniter.Marshal              491 ns/op          432 B/op            7 allocs/op
    goccy.Marshal              542 ns/op          416 B/op            2 allocs/op
  segment.Marshal              584 ns/op          176 B/op            1 allocs/op
 simdjson >>> not supported <<<
    gjson.Marshal             1791 ns/op         3600 B/op           32 allocs/op
  gjson-v >>> not supported <<<
    sonic.Marshal              711 ns/op         1803 B/op           14 allocs/op
  sonic-v >>> not supported <<<

 jsoniter ███████████████▍ 2.20
    goccy █████████████▉ 1.99
  segment ████████████▉ 1.85
 fastjson ████████████▍ 1.77
    sonic ██████████▌ 1.52
       oj █████████  1.29
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
