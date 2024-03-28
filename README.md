# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Validate []byte
     json.Validate            6272 ns/op            0 B/op            0 allocs/op
       oj.Validate            2275 ns/op            0 B/op            0 allocs/op
 fastjson.Validate            2554 ns/op            0 B/op            0 allocs/op
 jsoniter.Validate            4635 ns/op         2187 B/op          100 allocs/op
    buger >>> not supported <<<
    goccy.Validate           19547 ns/op        27647 B/op          463 allocs/op
  segment.Validate            4323 ns/op            0 B/op            0 allocs/op
 simdjson >>> not supported <<<
    gjson.Validate            2401 ns/op            0 B/op            0 allocs/op
  gjson-v >>> not supported <<<
    sonic.Validate            2018 ns/op            0 B/op            0 allocs/op
  sonic-v >>> not supported <<<

    sonic █████████████████████▊ 3.11
       oj ███████████████████▎ 2.76
    gjson ██████████████████▎ 2.61
 fastjson █████████████████▏ 2.46
  segment ██████████▏ 1.45
 jsoniter █████████▍ 1.35
     json ▓▓▓▓▓▓▓ 1.00
    goccy ██▏ 0.32
    buger >>> not supported <<<
 simdjson >>> not supported <<<
  gjson-v >>> not supported <<<
  sonic-v >>> not supported <<<

Validate string
     json.Validate            6106 ns/op            0 B/op            0 allocs/op
       oj >>> not supported <<<
 fastjson.Validate            3122 ns/op         4096 B/op            1 allocs/op
 jsoniter >>> not supported <<<
    buger >>> not supported <<<
    goccy >>> not supported <<<
  segment >>> not supported <<<
 simdjson >>> not supported <<<
    gjson.Validate            2866 ns/op         4096 B/op            1 allocs/op
  gjson-v >>> not supported <<<
    sonic.Validate            2589 ns/op         4137 B/op            1 allocs/op
  sonic-v >>> not supported <<<

    sonic ████████████████▌ 2.36
    gjson ██████████████▉ 2.13
 fastjson █████████████▋ 1.96
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 jsoniter >>> not supported <<<
    buger >>> not supported <<<
    goccy >>> not supported <<<
  segment >>> not supported <<<
 simdjson >>> not supported <<<
  gjson-v >>> not supported <<<
  sonic-v >>> not supported <<<

Unmarshal single JSON record, read few keys
     json.Unmarshal          18458 ns/op        15088 B/op           23 allocs/op
       oj >>> not supported <<<
 fastjson.Unmarshal           5206 ns/op        12034 B/op            7 allocs/op
 jsoniter.Unmarshal           8249 ns/op        14084 B/op           84 allocs/op
    buger.Unmarshal           4975 ns/op        12147 B/op           13 allocs/op
    goccy.Unmarshal           6011 ns/op        16303 B/op            7 allocs/op
  segment.Unmarshal          10247 ns/op        12064 B/op           12 allocs/op
 simdjson >>> not supported <<<
    gjson.Unmarshal           6232 ns/op        16368 B/op            9 allocs/op
  gjson-v.Unmarshal           8855 ns/op        16368 B/op            9 allocs/op
    sonic.Unmarshal           3472 ns/op         4952 B/op           12 allocs/op
  sonic-v.Unmarshal           5613 ns/op         4927 B/op           12 allocs/op

    sonic █████████████████████████████████████▏ 5.32
    buger █████████████████████████▉ 3.71
 fastjson ████████████████████████▊ 3.55
  sonic-v ███████████████████████  3.29
    goccy █████████████████████▍ 3.07
    gjson ████████████████████▋ 2.96
 jsoniter ███████████████▋ 2.24
  gjson-v ██████████████▌ 2.08
  segment ████████████▌ 1.80
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 simdjson >>> not supported <<<

Unmarshal single JSON record, read all keys
     json.Unmarshal          24452 ns/op        17216 B/op           82 allocs/op
       oj.Unmarshal          23375 ns/op         9355 B/op          456 allocs/op
 fastjson.Unmarshal           8585 ns/op        13608 B/op           73 allocs/op
 jsoniter.Unmarshal           9525 ns/op        14534 B/op           84 allocs/op
    buger.Unmarshal          14411 ns/op        14528 B/op          135 allocs/op
    goccy.Unmarshal           7505 ns/op        16680 B/op            7 allocs/op
  segment.Unmarshal          18744 ns/op        13680 B/op           73 allocs/op
 simdjson.Unmarshal          18659 ns/op       130664 B/op           17 allocs/op
    gjson.Unmarshal          14705 ns/op        18336 B/op           76 allocs/op
  gjson-v.Unmarshal          17240 ns/op        18336 B/op           76 allocs/op
    sonic.Unmarshal          14229 ns/op        27759 B/op           59 allocs/op
  sonic-v.Unmarshal          16422 ns/op        27745 B/op           59 allocs/op

    goccy ██████████████████████▊ 3.26
 fastjson ███████████████████▉ 2.85
 jsoniter █████████████████▉ 2.57
    sonic ████████████  1.72
    buger ███████████▉ 1.70
    gjson ███████████▋ 1.66
  sonic-v ██████████▍ 1.49
  gjson-v █████████▉ 1.42
 simdjson █████████▏ 1.31
  segment █████████▏ 1.30
       oj ███████▎ 1.05
     json ▓▓▓▓▓▓▓ 1.00

Unmarshal many JSON records from small file (100MB), read few keys
     json.Unmarshal      467502956 ns/op    617757397 B/op        39515 allocs/op
       oj >>> not supported <<<
 fastjson.Unmarshal       81009536 ns/op      2536360 B/op        39514 allocs/op
 jsoniter.Unmarshal      207915079 ns/op     89307064 B/op      3984433 allocs/op
    buger.Unmarshal       59628107 ns/op      4815730 B/op       275541 allocs/op
    goccy.Unmarshal       96344940 ns/op    106489478 B/op        39878 allocs/op
  segment.Unmarshal      213685315 ns/op      2528876 B/op        39450 allocs/op
 simdjson >>> not supported <<<
    gjson.Unmarshal       91747605 ns/op    106043352 B/op        39453 allocs/op
  gjson-v.Unmarshal      178791035 ns/op    106043536 B/op        39455 allocs/op
    sonic.Unmarshal       34111506 ns/op      2895209 B/op        78902 allocs/op
  sonic-v.Unmarshal      108529406 ns/op      2870777 B/op        78902 allocs/op

    sonic ███████████████████████████████████████████████████████████████████████████████████████████████▉ 13.71
    buger ██████████████████████████████████████████████████████▉ 7.84
 fastjson ████████████████████████████████████████▍ 5.77
    gjson ███████████████████████████████████▋ 5.10
    goccy █████████████████████████████████▉ 4.85
  sonic-v ██████████████████████████████▏ 4.31
  gjson-v ██████████████████▎ 2.61
 jsoniter ███████████████▋ 2.25
  segment ███████████████▎ 2.19
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 simdjson >>> not supported <<<

Unmarshal many JSON records from small file (100MB), read all keys
     json.Unmarshal      768111625 ns/op   1342693800 B/op     14015438 allocs/op
       oj.Unmarshal      505221100 ns/op    240453384 B/op     15464556 allocs/op
 fastjson.Unmarshal      225281196 ns/op     66615568 B/op      2879962 allocs/op
 jsoniter.Unmarshal      751836707 ns/op    819593676 B/op     19105312 allocs/op
    buger.Unmarshal      420675719 ns/op    105502056 B/op      5443966 allocs/op
    goccy.Unmarshal      704949027 ns/op    878218264 B/op     16663939 allocs/op
  segment.Unmarshal     2832206430 ns/op   6769765944 B/op     19882461 allocs/op
 simdjson.Unmarshal      413451615 ns/op   1470559696 B/op     14804622 allocs/op
    gjson.Unmarshal      440284293 ns/op    196303904 B/op      2998136 allocs/op
  gjson-v.Unmarshal      535801981 ns/op    196303488 B/op      2998131 allocs/op
    sonic.Unmarshal      475658980 ns/op    900466402 B/op      2249876 allocs/op
  sonic-v.Unmarshal      555284988 ns/op    899187564 B/op      2249875 allocs/op

 fastjson ███████████████████████▊ 3.41
 simdjson █████████████  1.86
    buger ████████████▊ 1.83
    gjson ████████████▏ 1.74
    sonic ███████████▎ 1.61
       oj ██████████▋ 1.52
  gjson-v ██████████  1.43
  sonic-v █████████▋ 1.38
    goccy ███████▋ 1.09
 jsoniter ███████▏ 1.02
     json ▓▓▓▓▓▓▓ 1.00
  segment █▉ 0.27

Unmarshal many JSON records from semi-large file (5GB), read few keys
     json.Unmarshal     7110308081 ns/op   9001289888 B/op       585117 allocs/op
       oj >>> not supported <<<
 fastjson.Unmarshal     1213712483 ns/op     37538448 B/op       585885 allocs/op
 jsoniter.Unmarshal     3167632856 ns/op   1328825192 B/op     59089506 allocs/op
    buger.Unmarshal      884535424 ns/op     71377368 B/op      4094501 allocs/op
    goccy.Unmarshal     1407353726 ns/op   1578817848 B/op       590473 allocs/op
  segment.Unmarshal     3174418208 ns/op     37446464 B/op       585037 allocs/op
 simdjson >>> not supported <<<
    gjson.Unmarshal     1368111643 ns/op   1572576832 B/op       585103 allocs/op
  gjson-v.Unmarshal     2678022070 ns/op   1572577312 B/op       585108 allocs/op
    sonic.Unmarshal      498267505 ns/op     42933565 B/op      1170122 allocs/op
  sonic-v.Unmarshal     1594407087 ns/op     42664912 B/op      1170113 allocs/op

    sonic ███████████████████████████████████████████████████████████████████████████████████████████████████▉ 14.27
    buger ████████████████████████████████████████████████████████▎ 8.04
 fastjson █████████████████████████████████████████  5.86
    gjson ████████████████████████████████████▍ 5.20
    goccy ███████████████████████████████████▎ 5.05
  sonic-v ███████████████████████████████▏ 4.46
  gjson-v ██████████████████▌ 2.66
 jsoniter ███████████████▋ 2.24
  segment ███████████████▋ 2.24
     json ▓▓▓▓▓▓▓ 1.00
       oj >>> not supported <<<
 simdjson >>> not supported <<<

Unmarshal many JSON records from semi-large file (5GB), read all keys
     json.Unmarshal    11447057361 ns/op  19753086928 B/op    207849182 allocs/op
       oj.Unmarshal     7468401881 ns/op   3569330624 B/op    229340098 allocs/op
 fastjson.Unmarshal     3291232362 ns/op    987634840 B/op     42708214 allocs/op
 jsoniter.Unmarshal    10846217954 ns/op  12153011152 B/op    283332719 allocs/op
    buger.Unmarshal     6118403317 ns/op   1562015712 B/op     80734454 allocs/op
    goccy.Unmarshal    10343276132 ns/op  13020661864 B/op    247121033 allocs/op
  segment.Unmarshal    39543224498 ns/op 100399606128 B/op    294862851 allocs/op
 simdjson.Unmarshal     5470108761 ns/op  14817620112 B/op    219552921 allocs/op
    gjson.Unmarshal     6084229629 ns/op   2911123360 B/op     44462433 allocs/op
  gjson-v.Unmarshal     7532090808 ns/op   2911123344 B/op     44462433 allocs/op
    sonic.Unmarshal     5331276049 ns/op  13035812896 B/op     33346907 allocs/op
  sonic-v.Unmarshal     6554266534 ns/op  13035730160 B/op     33346896 allocs/op

 fastjson ████████████████████████▎ 3.48
    sonic ███████████████  2.15
 simdjson ██████████████▋ 2.09
    gjson █████████████▏ 1.88
    buger █████████████  1.87
  sonic-v ████████████▏ 1.75
       oj ██████████▋ 1.53
  gjson-v ██████████▋ 1.52
    goccy ███████▋ 1.11
 jsoniter ███████▍ 1.06
     json ▓▓▓▓▓▓▓ 1.00
  segment ██  0.29

Marshal custom data through an object builder
     json.Marshal             1010 ns/op          688 B/op           17 allocs/op
       oj.Builder              836 ns/op         1128 B/op           17 allocs/op
 fastjson.Marshal              585 ns/op          679 B/op            6 allocs/op
 jsoniter.Marshal              489 ns/op          432 B/op            7 allocs/op
    buger >>> not supported <<<
    goccy.Marshal              550 ns/op          416 B/op            2 allocs/op
  segment.Marshal              560 ns/op          176 B/op            1 allocs/op
 simdjson >>> not supported <<<
    gjson.Marshal             1852 ns/op         3600 B/op           32 allocs/op
  gjson-v >>> not supported <<<
    sonic.Marshal              689 ns/op         1803 B/op           14 allocs/op
  sonic-v >>> not supported <<<

 jsoniter ██████████████▍ 2.07
    goccy ████████████▊ 1.84
  segment ████████████▋ 1.80
 fastjson ████████████  1.73
    sonic ██████████▎ 1.47
       oj ████████▍ 1.21
     json ▓▓▓▓▓▓▓ 1.00
    gjson ███▊ 0.55
    buger >>> not supported <<<
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
