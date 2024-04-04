# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Validate []byte
         json.Validate                6603 ns/op            0 B/op            0 allocs/op
           oj.Validate                2342 ns/op            0 B/op            0 allocs/op
     fastjson.Validate                2764 ns/op            0 B/op            0 allocs/op
     jsoniter.Validate                4722 ns/op         2188 B/op          100 allocs/op
   jsonparser >>> not supported <<<
        goccy.Validate               20374 ns/op        27655 B/op          463 allocs/op
      segment.Validate                4444 ns/op            0 B/op            0 allocs/op
     simdjson >>> not supported <<<
        gjson.Validate                2350 ns/op            0 B/op            0 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                1973 ns/op            0 B/op            0 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<

        sonic █████████████▍ 3.35
           oj ███████████▎ 2.82
        gjson ███████████▏ 2.81
     fastjson █████████▌ 2.39
      segment █████▉ 1.49
     jsoniter █████▌ 1.40
         json ▓▓▓▓ 1.00
        goccy █▎ 0.32
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<

Validate string
         json.Validate                6304 ns/op            0 B/op            0 allocs/op
           oj >>> not supported <<<
     fastjson.Validate                3297 ns/op         4096 B/op            1 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Validate                2953 ns/op         4096 B/op            1 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                2621 ns/op         4140 B/op            1 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<

        sonic █████████▌ 2.41
        gjson ████████▌ 2.13
     fastjson ███████▋ 1.91
         json ▓▓▓▓ 1.00
      segment >>> not supported <<<
        goccy >>> not supported <<<
   jsonparser >>> not supported <<<
     simdjson >>> not supported <<<
     jsoniter >>> not supported <<<
      gjson-v >>> not supported <<<
           oj >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<

Unmarshal single JSON record, read few keys
         json.Unmarshal              19483 ns/op        15088 B/op           23 allocs/op
           oj >>> not supported <<<
     fastjson.Unmarshal               5279 ns/op        12034 B/op            7 allocs/op
     jsoniter.Unmarshal               8604 ns/op        14087 B/op           84 allocs/op
   jsonparser.Unmarshal               5217 ns/op        12147 B/op           13 allocs/op
        goccy.Unmarshal               6132 ns/op        16304 B/op            7 allocs/op
      segment.Unmarshal              10468 ns/op        12064 B/op           12 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal               6412 ns/op        16368 B/op            9 allocs/op
      gjson-v.Unmarshal               9175 ns/op        16368 B/op            9 allocs/op
        sonic.Unmarshal               3560 ns/op         4962 B/op           12 allocs/op
      sonic-v.Unmarshal               5678 ns/op         4936 B/op           12 allocs/op
        codec.Unmarshal               7691 ns/op        13240 B/op           15 allocs/op

        sonic █████████████████████▉ 5.47
   jsonparser ██████████████▉ 3.73
     fastjson ██████████████▊ 3.69
      sonic-v █████████████▋ 3.43
        goccy ████████████▋ 3.18
        gjson ████████████▏ 3.04
        codec ██████████▏ 2.53
     jsoniter █████████  2.26
      gjson-v ████████▍ 2.12
      segment ███████▍ 1.86
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
           oj >>> not supported <<<

Unmarshal single JSON record, read all keys
         json.Unmarshal              25252 ns/op        17216 B/op           82 allocs/op
           oj.Unmarshal              23073 ns/op         9350 B/op          456 allocs/op
     fastjson.Unmarshal               8527 ns/op        13608 B/op           73 allocs/op
     jsoniter.Unmarshal               9718 ns/op        14536 B/op           84 allocs/op
   jsonparser.Unmarshal              14905 ns/op        14528 B/op          135 allocs/op
        goccy.Unmarshal               7448 ns/op        16631 B/op            7 allocs/op
      segment.Unmarshal              18930 ns/op        13680 B/op           73 allocs/op
     simdjson.Unmarshal              18176 ns/op       130664 B/op           17 allocs/op
        gjson.Unmarshal              14547 ns/op        18336 B/op           76 allocs/op
      gjson-v.Unmarshal              17351 ns/op        18336 B/op           76 allocs/op
        sonic.Unmarshal              14247 ns/op        27690 B/op           59 allocs/op
      sonic-v.Unmarshal              16570 ns/op        27692 B/op           59 allocs/op
        codec.Unmarshal               7199 ns/op        13208 B/op            9 allocs/op

        codec ██████████████  3.51
        goccy █████████████▌ 3.39
     fastjson ███████████▊ 2.96
     jsoniter ██████████▍ 2.60
        sonic ███████  1.77
        gjson ██████▉ 1.74
   jsonparser ██████▊ 1.69
      sonic-v ██████  1.52
      gjson-v █████▊ 1.46
     simdjson █████▌ 1.39
      segment █████▎ 1.33
           oj ████▍ 1.09
         json ▓▓▓▓ 1.00

Unmarshal many JSON records from small file (100MB), read few keys
         json.Unmarshal          491714342 ns/op    617757461 B/op        39516 allocs/op
           oj >>> not supported <<<
     fastjson.Unmarshal           80379104 ns/op      2536670 B/op        39518 allocs/op
     jsoniter.Unmarshal          206962245 ns/op     89199448 B/op      3984430 allocs/op
   jsonparser.Unmarshal           60552461 ns/op      4818211 B/op       276783 allocs/op
        goccy.Unmarshal           95585303 ns/op    106449228 B/op        39823 allocs/op
      segment.Unmarshal          215111197 ns/op      2528876 B/op        39450 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal           92359684 ns/op    106043472 B/op        39454 allocs/op
      gjson-v.Unmarshal          179712687 ns/op    106043488 B/op        39455 allocs/op
        sonic.Unmarshal           34065879 ns/op      2900826 B/op        78902 allocs/op
      sonic-v.Unmarshal          108783088 ns/op      2878992 B/op        78902 allocs/op
        codec.Unmarshal          125290230 ns/op     48921110 B/op       157799 allocs/op

        sonic █████████████████████████████████████████████████████████▋ 14.43
   jsonparser ████████████████████████████████▍ 8.12
     fastjson ████████████████████████▍ 6.12
        gjson █████████████████████▎ 5.32
        goccy ████████████████████▌ 5.14
      sonic-v ██████████████████  4.52
        codec ███████████████▋ 3.92
      gjson-v ██████████▉ 2.74
     jsoniter █████████▌ 2.38
      segment █████████▏ 2.29
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
           oj >>> not supported <<<

Unmarshal many JSON records from small file (100MB), read all keys
         json.Unmarshal          776141478 ns/op   1342985536 B/op     14015400 allocs/op
           oj.Unmarshal          493253922 ns/op    240713376 B/op     15464527 allocs/op
     fastjson.Unmarshal          215828878 ns/op     66614644 B/op      2879956 allocs/op
     jsoniter.Unmarshal          749169925 ns/op    819714368 B/op     19105243 allocs/op
   jsonparser.Unmarshal          407017103 ns/op    105383133 B/op      5443966 allocs/op
        goccy.Unmarshal          687000095 ns/op    877826840 B/op     16663412 allocs/op
      segment.Unmarshal         2653214418 ns/op   6770074712 B/op     19882657 allocs/op
     simdjson.Unmarshal          416966349 ns/op   1470824533 B/op     14804620 allocs/op
        gjson.Unmarshal          435813434 ns/op    196303413 B/op      2998131 allocs/op
      gjson-v.Unmarshal          535879726 ns/op    196303176 B/op      2998129 allocs/op
        sonic.Unmarshal          471866948 ns/op    898749186 B/op      2249801 allocs/op
      sonic-v.Unmarshal          555538103 ns/op    896904244 B/op      2249785 allocs/op
        codec.Unmarshal         1118816665 ns/op    313002816 B/op     18304550 allocs/op

     fastjson ██████████████▍ 3.60
   jsonparser ███████▋ 1.91
     simdjson ███████▍ 1.86
        gjson ███████  1.78
        sonic ██████▌ 1.64
           oj ██████▎ 1.57
      gjson-v █████▊ 1.45
      sonic-v █████▌ 1.40
        goccy ████▌ 1.13
     jsoniter ████▏ 1.04
         json ▓▓▓▓ 1.00
        codec ██▊ 0.69
      segment █▏ 0.29

Unmarshal many JSON records from semi-large file (5GB), read few keys
         json.Unmarshal        23331130459 ns/op  27371014656 B/op       209834 allocs/op
           oj >>> not supported <<<
     fastjson.Unmarshal         3291353776 ns/op     14556504 B/op       218105 allocs/op
     jsoniter.Unmarshal         9658140216 ns/op   4629848616 B/op    209953670 allocs/op
   jsonparser.Unmarshal         2617922828 ns/op     41703088 B/op      9494678 allocs/op
        goccy.Unmarshal         3981546610 ns/op   5737878128 B/op       225005 allocs/op
      segment.Unmarshal         8444259861 ns/op     13484928 B/op       209746 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal         4168977540 ns/op   5718448032 B/op       210111 allocs/op
      gjson-v.Unmarshal         8465874067 ns/op   5718441216 B/op       210040 allocs/op
        sonic.Unmarshal          840257790 ns/op     15519840 B/op       419505 allocs/op
      sonic-v.Unmarshal         4351764356 ns/op     15438064 B/op       419505 allocs/op
        codec.Unmarshal         4396351445 ns/op    260140112 B/op       838976 allocs/op

        sonic ███████████████████████████████████████████████████████████████████████████████████████████████████████████████  27.77
   jsonparser ███████████████████████████████████▋ 8.91
     fastjson ████████████████████████████▎ 7.09
        goccy ███████████████████████▍ 5.86
        gjson ██████████████████████▍ 5.60
      sonic-v █████████████████████▍ 5.36
        codec █████████████████████▏ 5.31
      segment ███████████  2.76
      gjson-v ███████████  2.76
     jsoniter █████████▋ 2.42
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
           oj >>> not supported <<<

Unmarshal many JSON records from semi-large file (5GB), read all keys
         json.Unmarshal        37690611873 ns/op  64613988864 B/op    698450509 allocs/op
           oj.Unmarshal        24271470614 ns/op  12073547520 B/op    773131913 allocs/op
     fastjson.Unmarshal         9880777882 ns/op   3330145448 B/op    141792743 allocs/op
     jsoniter.Unmarshal        37030676399 ns/op  41916030752 B/op    957721809 allocs/op
   jsonparser.Unmarshal        20061977116 ns/op   5267562672 B/op    272452537 allocs/op
        goccy.Unmarshal        33996892673 ns/op  45095417552 B/op    831610301 allocs/op
      segment.Unmarshal       126505733043 ns/op 340323739440 B/op    992919235 allocs/op
     simdjson.Unmarshal        18478541385 ns/op  50095024528 B/op    744180760 allocs/op
        gjson.Unmarshal        19826681906 ns/op  10366251392 B/op    146188784 allocs/op
      gjson-v.Unmarshal        24345015189 ns/op  10366251328 B/op    146188784 allocs/op
        sonic.Unmarshal        17256337276 ns/op  44723317488 B/op    110533209 allocs/op
      sonic-v.Unmarshal        20977089457 ns/op  44723358368 B/op    110533209 allocs/op
        codec.Unmarshal        53181073851 ns/op  13514943360 B/op    905239480 allocs/op

     fastjson ███████████████▎ 3.81
        sonic ████████▋ 2.18
     simdjson ████████▏ 2.04
        gjson ███████▌ 1.90
   jsonparser ███████▌ 1.88
      sonic-v ███████▏ 1.80
           oj ██████▏ 1.55
      gjson-v ██████▏ 1.55
        goccy ████▍ 1.11
     jsoniter ████  1.02
         json ▓▓▓▓ 1.00
        codec ██▊ 0.71
      segment █▏ 0.30

Marshal custom data through an object builder
         json.Marshal                 1027 ns/op          688 B/op           17 allocs/op
           oj.Builder                  842 ns/op         1128 B/op           17 allocs/op
     fastjson.Marshal                  562 ns/op          679 B/op            6 allocs/op
     jsoniter.Marshal                  501 ns/op          432 B/op            7 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                  560 ns/op          416 B/op            2 allocs/op
      segment.Marshal                  565 ns/op          176 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson.Marshal                 1833 ns/op         3600 B/op           32 allocs/op
      gjson-v >>> not supported <<<
        sonic.Marshal                  659 ns/op         1803 B/op           14 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                  561 ns/op          928 B/op            0 allocs/op

     jsoniter ████████▏ 2.05
        goccy ███████▎ 1.83
        codec ███████▎ 1.83
     fastjson ███████▎ 1.83
      segment ███████▎ 1.82
        sonic ██████▏ 1.56
           oj ████▉ 1.22
         json ▓▓▓▓ 1.00
        gjson ██▏ 0.56
   jsonparser >>> not supported <<<
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
