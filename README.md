# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Validate []byte
         json.Validate                6606 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj.Validate                2235 ns/op            0 B/op            0 allocs/op
     fastjson.Validate                2771 ns/op            0 B/op            0 allocs/op
     jsoniter.Validate                4685 ns/op         2188 B/op          100 allocs/op
   jsonparser >>> not supported <<<
        goccy.Validate               19932 ns/op        27655 B/op          463 allocs/op
      segment.Validate                4332 ns/op            0 B/op            0 allocs/op
     simdjson >>> not supported <<<
        gjson.Validate                2182 ns/op            0 B/op            0 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                2025 ns/op            0 B/op            0 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<

        sonic █████████████  3.26
        gjson ████████████  3.03
           oj ███████████▊ 2.96
     fastjson █████████▌ 2.38
      segment ██████  1.52
     jsoniter █████▋ 1.41
         json ▓▓▓▓ 1.00
        goccy █▎ 0.33
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        json2 >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<

Validate string
         json.Validate                6303 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Validate                3342 ns/op         4096 B/op            1 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Validate                2781 ns/op         4096 B/op            1 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                2619 ns/op         4138 B/op            1 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<

        sonic █████████▋ 2.41
        gjson █████████  2.27
     fastjson ███████▌ 1.89
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        json2 >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     jsoniter >>> not supported <<<
           oj >>> not supported <<<
      gjson-v >>> not supported <<<
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<

Unmarshal single record (2kb), read few keys generically
         json.Unmarshal              18244 ns/op        12488 B/op           21 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal               5292 ns/op        12034 B/op            7 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal               5252 ns/op        12147 B/op           13 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal               6068 ns/op        16368 B/op            9 allocs/op
      gjson-v.Unmarshal               8589 ns/op        16368 B/op            9 allocs/op
        sonic.Unmarshal               3553 ns/op         4965 B/op           12 allocs/op
      sonic-v.Unmarshal               5725 ns/op         4940 B/op           12 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal               5379 ns/op        12040 B/op            8 allocs/op
        jason.Unmarshal              38224 ns/op        56000 B/op          495 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<

        sonic ████████████████████▌ 5.13
   jsonparser █████████████▉ 3.47
     fastjson █████████████▊ 3.45
          jin █████████████▌ 3.39
      sonic-v ████████████▋ 3.19
        gjson ████████████  3.01
      gjson-v ████████▍ 2.12
         json ▓▓▓▓ 1.00
        jason █▉ 0.48
     jsoniter >>> not supported <<<
     simdjson >>> not supported <<<
      segment >>> not supported <<<
        goccy >>> not supported <<<
        codec >>> not supported <<<
           oj >>> not supported <<<
        json2 >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<

Unmarshal single record (2kb), read few keys into struct
         json.Unmarshal              18255 ns/op        12488 B/op           21 allocs/op
        json2.Unmarshal               5954 ns/op          128 B/op            6 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal               8582 ns/op        14087 B/op           84 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               6220 ns/op        16296 B/op            7 allocs/op
      segment.Unmarshal              10615 ns/op        12064 B/op           12 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               5863 ns/op         9043 B/op           11 allocs/op
      sonic-v.Unmarshal               7826 ns/op         9327 B/op           11 allocs/op
        codec.Unmarshal               7595 ns/op        13240 B/op           15 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal              15923 ns/op        14280 B/op           36 allocs/op

        sonic ████████████▍ 3.11
        json2 ████████████▎ 3.07
        goccy ███████████▋ 2.93
        codec █████████▌ 2.40
      sonic-v █████████▎ 2.33
     jsoniter ████████▌ 2.13
      segment ██████▉ 1.72
       ffjson ████▌ 1.15
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     fastjson >>> not supported <<<
           oj >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

Unmarshal single record (2kb), read all keys generically
         json.Unmarshal              26208 ns/op        29777 B/op          340 allocs/op
        json2.Unmarshal              19417 ns/op        15573 B/op          295 allocs/op
           oj.Unmarshal              18526 ns/op        19300 B/op          465 allocs/op
     fastjson.Unmarshal               8506 ns/op        13608 B/op           73 allocs/op
     jsoniter.Unmarshal              20265 ns/op        31752 B/op          457 allocs/op
   jsonparser.Unmarshal              14521 ns/op        14528 B/op          135 allocs/op
        goccy.Unmarshal              19330 ns/op        34668 B/op          398 allocs/op
      segment.Unmarshal              64914 ns/op       173242 B/op          475 allocs/op
     simdjson.Unmarshal              22783 ns/op       132224 B/op           82 allocs/op
        gjson.Unmarshal              12843 ns/op        18336 B/op           76 allocs/op
      gjson-v.Unmarshal              15402 ns/op        18336 B/op           76 allocs/op
        sonic.Unmarshal              14088 ns/op        27707 B/op           59 allocs/op
      sonic-v.Unmarshal              16149 ns/op        27697 B/op           59 allocs/op
        codec.Unmarshal              30234 ns/op        19456 B/op          435 allocs/op
          jin.Unmarshal              22700 ns/op        25344 B/op          435 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal              14323 ns/op        32168 B/op          183 allocs/op
       ffjson >>> not supported <<<

     fastjson ████████████▎ 3.08
        gjson ████████▏ 2.04
        sonic ███████▍ 1.86
        djson ███████▎ 1.83
   jsonparser ███████▏ 1.80
      gjson-v ██████▊ 1.70
      sonic-v ██████▍ 1.62
           oj █████▋ 1.41
        goccy █████▍ 1.36
        json2 █████▍ 1.35
     jsoniter █████▏ 1.29
          jin ████▌ 1.15
     simdjson ████▌ 1.15
         json ▓▓▓▓ 1.00
        codec ███▍ 0.87
      segment █▌ 0.40
        jason >>> not supported <<<
       ffjson >>> not supported <<<

Unmarshal single record (2kb), read all keys into struct
         json.Unmarshal              23460 ns/op        14608 B/op           80 allocs/op
        json2.Unmarshal              10002 ns/op         3110 B/op           35 allocs/op
           oj.Unmarshal              21539 ns/op         9339 B/op          453 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal               8821 ns/op        14413 B/op           77 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               7228 ns/op        16596 B/op            7 allocs/op
      segment.Unmarshal              17371 ns/op        13680 B/op           73 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               6267 ns/op         9580 B/op            9 allocs/op
      sonic-v.Unmarshal               8443 ns/op         9864 B/op            9 allocs/op
        codec.Unmarshal              10540 ns/op        14536 B/op           72 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal              12659 ns/op        16883 B/op          113 allocs/op

        sonic ██████████████▉ 3.74
        goccy ████████████▉ 3.25
      sonic-v ███████████  2.78
     jsoniter ██████████▋ 2.66
        json2 █████████▍ 2.35
        codec ████████▉ 2.23
       ffjson ███████▍ 1.85
      segment █████▍ 1.35
           oj ████▎ 1.09
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     fastjson >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
   jsonparser >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys generically
         json.Unmarshal          447488884 ns/op     21755845 B/op       394052 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal           79221982 ns/op      2533390 B/op        39470 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal           56664580 ns/op      4614627 B/op       177366 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal           58720813 ns/op    610637391 B/op        39635 allocs/op
        gjson.Unmarshal           87576272 ns/op    105924950 B/op        39408 allocs/op
      gjson-v.Unmarshal          171284371 ns/op    105924992 B/op        39408 allocs/op
        sonic.Unmarshal           33906705 ns/op      2895669 B/op        78814 allocs/op
      sonic-v.Unmarshal          108593861 ns/op      2874825 B/op        78813 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal           69890388 ns/op      2652134 B/op        78811 allocs/op
        jason.Unmarshal         1001693170 ns/op   1083095008 B/op     16167016 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<

        sonic ████████████████████████████████████████████████████▊ 13.20
   jsonparser ███████████████████████████████▌ 7.90
     simdjson ██████████████████████████████▍ 7.62
          jin █████████████████████████▌ 6.40
     fastjson ██████████████████████▌ 5.65
        gjson ████████████████████▍ 5.11
      sonic-v ████████████████▍ 4.12
      gjson-v ██████████▍ 2.61
         json ▓▓▓▓ 1.00
        jason █▊ 0.45
      segment >>> not supported <<<
        goccy >>> not supported <<<
     jsoniter >>> not supported <<<
           oj >>> not supported <<<
        codec >>> not supported <<<
        json2 >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys into struct
         json.Unmarshal          448324750 ns/op     21755733 B/op       394051 allocs/op
        json2.Unmarshal          188617182 ns/op        16804 B/op           64 allocs/op
           oj.Unmarshal          511520169 ns/op    246894112 B/op     15880773 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          207203772 ns/op     89156934 B/op      3979979 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal           88736803 ns/op    106274899 B/op        39732 allocs/op
      segment.Unmarshal          214087348 ns/op      2526041 B/op        39406 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal          114211099 ns/op    110333182 B/op        78918 allocs/op
      sonic-v.Unmarshal          190406873 ns/op    112618882 B/op        79039 allocs/op
        codec.Unmarshal          117975921 ns/op     48866415 B/op       157622 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal          272773333 ns/op     68131584 B/op       591320 allocs/op

        goccy ████████████████████▏ 5.05
        sonic ███████████████▋ 3.93
        codec ███████████████▏ 3.80
        json2 █████████▌ 2.38
      sonic-v █████████▍ 2.35
     jsoniter ████████▋ 2.16
      segment ████████▍ 2.09
       ffjson ██████▌ 1.64
         json ▓▓▓▓ 1.00
           oj ███▌ 0.88
   jsonparser >>> not supported <<<
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     fastjson >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys generically
         json.Unmarshal          855439640 ns/op    743911688 B/op     14315018 allocs/op
        json2.Unmarshal          747345670 ns/op    693878908 B/op     11948710 allocs/op
           oj.Unmarshal          851356523 ns/op    809228224 B/op     19911161 allocs/op
     fastjson.Unmarshal          221393105 ns/op     66540784 B/op      2876749 allocs/op
     jsoniter.Unmarshal          735601480 ns/op    820416676 B/op     19201918 allocs/op
   jsonparser.Unmarshal          488556064 ns/op    105248920 B/op      5437894 allocs/op
        goccy.Unmarshal          668978070 ns/op    879545480 B/op     16801923 allocs/op
      segment.Unmarshal         2791130529 ns/op   6770756896 B/op     19978809 allocs/op
     simdjson.Unmarshal          171459897 ns/op    688370062 B/op      2758585 allocs/op
        gjson.Unmarshal          414620608 ns/op    196083413 B/op      2994781 allocs/op
      gjson-v.Unmarshal          512047958 ns/op    196083416 B/op      2994781 allocs/op
        sonic.Unmarshal          387398219 ns/op    890333200 B/op      2285568 allocs/op
      sonic-v.Unmarshal          468098372 ns/op    890389314 B/op      2285569 allocs/op
        codec.Unmarshal         1086018591 ns/op    312883424 B/op     18323502 allocs/op
          jin.Unmarshal          837954065 ns/op    563459016 B/op     18165706 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal          379986856 ns/op    770792648 B/op      7537395 allocs/op
       ffjson >>> not supported <<<

     simdjson ███████████████████▉ 4.99
     fastjson ███████████████▍ 3.86
        djson █████████  2.25
        sonic ████████▊ 2.21
        gjson ████████▎ 2.06
      sonic-v ███████▎ 1.83
   jsonparser ███████  1.75
      gjson-v ██████▋ 1.67
        goccy █████  1.28
     jsoniter ████▋ 1.16
        json2 ████▌ 1.14
          jin ████  1.02
           oj ████  1.00
         json ▓▓▓▓ 1.00
        codec ███▏ 0.79
      segment █▏ 0.31
        jason >>> not supported <<<
       ffjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys into struct
         json.Unmarshal          673320003 ns/op    110917616 B/op      3113016 allocs/op
        json2.Unmarshal          379988884 ns/op    129775837 B/op      1801322 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          210620808 ns/op     97217915 B/op      2955391 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal          143757789 ns/op    106192365 B/op        39577 allocs/op
      segment.Unmarshal          542640995 ns/op     68879072 B/op      2797773 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal          140339148 ns/op    133227128 B/op       157635 allocs/op
      sonic-v.Unmarshal          217703134 ns/op    133535406 B/op       157653 allocs/op
        codec.Unmarshal          285245173 ns/op    102057772 B/op      2718954 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal          379714454 ns/op    217858125 B/op      4689261 allocs/op

        sonic ███████████████████▏ 4.80
        goccy ██████████████████▋ 4.68
     jsoniter ████████████▊ 3.20
      sonic-v ████████████▎ 3.09
        codec █████████▍ 2.36
       ffjson ███████  1.77
        json2 ███████  1.77
      segment ████▉ 1.24
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     fastjson >>> not supported <<<
           oj >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
   jsonparser >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys generically
         json.Unmarshal        20337594444 ns/op    116289088 B/op      2105577 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal         3223455745 ns/op     14463336 B/op       218594 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal         1825897838 ns/op     24696416 B/op       947400 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal         1152251580 ns/op   2299937824 B/op       217371 allocs/op
        gjson.Unmarshal         2933615591 ns/op   5740687488 B/op       210561 allocs/op
      gjson-v.Unmarshal         6698258361 ns/op   5740688448 B/op       210571 allocs/op
        sonic.Unmarshal          754651750 ns/op     15285620 B/op       421123 allocs/op
      sonic-v.Unmarshal         4256549461 ns/op     15267608 B/op       421124 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal         2658307753 ns/op     14210896 B/op       421118 allocs/op
        jason.Unmarshal        37445686933 ns/op  51080071728 B/op    711850146 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<

        sonic ███████████████████████████████████████████████████████████████████████████████████████████████████████████▊ 26.95
     simdjson ██████████████████████████████████████████████████████████████████████▌ 17.65
   jsonparser ████████████████████████████████████████████▌ 11.14
          jin ██████████████████████████████▌ 7.65
        gjson ███████████████████████████▋ 6.93
     fastjson █████████████████████████▏ 6.31
      sonic-v ███████████████████  4.78
      gjson-v ████████████▏ 3.04
         json ▓▓▓▓ 1.00
        jason ██▏ 0.54
      segment >>> not supported <<<
        goccy >>> not supported <<<
     jsoniter >>> not supported <<<
           oj >>> not supported <<<
        codec >>> not supported <<<
        json2 >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys into struct
         json.Unmarshal        20364881011 ns/op    116288992 B/op      2105576 allocs/op
        json2.Unmarshal         8690878044 ns/op        21080 B/op           72 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal         8890747682 ns/op   4650047936 B/op    210767858 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         3090416550 ns/op   5742226608 B/op       211778 allocs/op
      segment.Unmarshal         8250610581 ns/op     13537216 B/op       210563 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal         4303314873 ns/op   5757236760 B/op       421560 allocs/op
      sonic-v.Unmarshal         7833730064 ns/op   5767190560 B/op       422089 allocs/op
        codec.Unmarshal         4079037794 ns/op    261152232 B/op       842234 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal        10451631305 ns/op    362387272 B/op      3158504 allocs/op

        goccy ██████████████████████████▎ 6.59
        codec ███████████████████▉ 4.99
        sonic ██████████████████▉ 4.73
      sonic-v ██████████▍ 2.60
      segment █████████▊ 2.47
        json2 █████████▎ 2.34
     jsoniter █████████▏ 2.29
       ffjson ███████▊ 1.95
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     fastjson >>> not supported <<<
           oj >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys generically
         json.Unmarshal        35207192460 ns/op  37172730544 B/op    701322006 allocs/op
        json2.Unmarshal        29219316963 ns/op  35262205184 B/op    579263924 allocs/op
           oj.Unmarshal        33846669147 ns/op  40944081592 B/op    988125205 allocs/op
     fastjson.Unmarshal         9459179271 ns/op   3342949992 B/op    142344580 allocs/op
     jsoniter.Unmarshal        29935042799 ns/op  41731649432 B/op    960520490 allocs/op
   jsonparser.Unmarshal        23159595610 ns/op   5294193048 B/op    273513557 allocs/op
        goccy.Unmarshal        27488428064 ns/op  44770009000 B/op    833147884 allocs/op
      segment.Unmarshal       114948027079 ns/op 341651490152 B/op    995514319 allocs/op
     simdjson.Unmarshal         6592252483 ns/op  27142925168 B/op    137924689 allocs/op
        gjson.Unmarshal        20526128427 ns/op  10406632208 B/op    146758233 allocs/op
      gjson-v.Unmarshal        25044253499 ns/op  10406634608 B/op    146758257 allocs/op
        sonic.Unmarshal        18513334895 ns/op  44977979256 B/op    111177902 allocs/op
      sonic-v.Unmarshal        22171396480 ns/op  44976987328 B/op    111177975 allocs/op
        codec.Unmarshal        53774598071 ns/op  13510682880 B/op    905186231 allocs/op
          jin.Unmarshal        42179220148 ns/op  28265588552 B/op    910027426 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal        18501551846 ns/op  39791427968 B/op    375169798 allocs/op
       ffjson >>> not supported <<<

     simdjson █████████████████████▎ 5.34
     fastjson ██████████████▉ 3.72
        djson ███████▌ 1.90
        sonic ███████▌ 1.90
        gjson ██████▊ 1.72
      sonic-v ██████▎ 1.59
   jsonparser ██████  1.52
      gjson-v █████▌ 1.41
        goccy █████  1.28
        json2 ████▊ 1.20
     jsoniter ████▋ 1.18
           oj ████▏ 1.04
         json ▓▓▓▓ 1.00
          jin ███▎ 0.83
        codec ██▌ 0.65
      segment █▏ 0.31
        jason >>> not supported <<<
       ffjson >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys into struct
         json.Unmarshal        31484553941 ns/op   4443558448 B/op    137915231 allocs/op
        json2.Unmarshal        19156437400 ns/op   6895391664 B/op     94347020 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal         9666212089 ns/op   5003847032 B/op    150338412 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         6259542046 ns/op   5754914872 B/op       223364 allocs/op
      segment.Unmarshal        25193255880 ns/op   3491109512 B/op    141915745 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal         6196346396 ns/op   7137864776 B/op      4633100 allocs/op
      sonic-v.Unmarshal         9910448860 ns/op   7151267160 B/op      4633814 allocs/op
        codec.Unmarshal        13356329668 ns/op   3062219856 B/op    133914583 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal        16799907110 ns/op  10026588840 B/op    215402597 allocs/op

        sonic ████████████████████▎ 5.08
        goccy ████████████████████  5.03
     jsoniter █████████████  3.26
      sonic-v ████████████▋ 3.18
        codec █████████▍ 2.36
       ffjson ███████▍ 1.87
        json2 ██████▌ 1.64
      segment ████▉ 1.25
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     fastjson >>> not supported <<<
           oj >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
   jsonparser >>> not supported <<<

Marshal custom data through an object builder
         json.Marshal                  969 ns/op          688 B/op           17 allocs/op
        json2.Marshal                 1048 ns/op          424 B/op           10 allocs/op
           oj.Builder                  850 ns/op         1128 B/op           17 allocs/op
     fastjson.Marshal                  527 ns/op          679 B/op            6 allocs/op
     jsoniter.Marshal                  480 ns/op          432 B/op            7 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                  540 ns/op          416 B/op            2 allocs/op
      segment.Marshal                  550 ns/op          176 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson.Marshal                 1583 ns/op         3024 B/op           27 allocs/op
      gjson-v >>> not supported <<<
        sonic.Marshal                  682 ns/op         1803 B/op           14 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                  584 ns/op          935 B/op            0 allocs/op
          jin.Marshal                 1474 ns/op         2544 B/op           56 allocs/op
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Marshal                  144 ns/op          144 B/op            2 allocs/op

       ffjson ██████████████████████████▉ 6.73
     jsoniter ████████  2.02
     fastjson ███████▎ 1.84
        goccy ███████▏ 1.79
      segment ███████  1.76
        codec ██████▋ 1.66
        sonic █████▋ 1.42
           oj ████▌ 1.14
         json ▓▓▓▓ 1.00
        json2 ███▋ 0.92
          jin ██▋ 0.66
        gjson ██▍ 0.61
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
   jsonparser >>> not supported <<<

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
