# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Validate []byte
         json.Validate                6424 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj.Validate                2258 ns/op            0 B/op            0 allocs/op
     fastjson.Validate                2697 ns/op            0 B/op            0 allocs/op
     jsoniter.Validate                4756 ns/op         2188 B/op          100 allocs/op
   jsonparser >>> not supported <<<
        goccy.Validate               19492 ns/op        27654 B/op          463 allocs/op
      segment.Validate                4240 ns/op            0 B/op            0 allocs/op
     simdjson >>> not supported <<<
        gjson.Validate                2361 ns/op            0 B/op            0 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                2084 ns/op            0 B/op            0 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<

        sonic ████████████▎ 3.08
           oj ███████████▍ 2.84
        gjson ██████████▉ 2.72
     fastjson █████████▌ 2.38
      segment ██████  1.52
     jsoniter █████▍ 1.35
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

Validate string
         json.Validate                6351 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Validate                3288 ns/op         4096 B/op            1 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Validate                2965 ns/op         4096 B/op            1 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                2669 ns/op         4139 B/op            1 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<

        sonic █████████▌ 2.38
        gjson ████████▌ 2.14
     fastjson ███████▋ 1.93
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     jsoniter >>> not supported <<<
           oj >>> not supported <<<
      gjson-v >>> not supported <<<
        json2 >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<

Unmarshal single record (2kb), read few keys generically
         json.Unmarshal              18212 ns/op        12488 B/op           21 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal               5324 ns/op        12034 B/op            7 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal               5303 ns/op        12147 B/op           13 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal               6364 ns/op        16368 B/op            9 allocs/op
      gjson-v.Unmarshal               9043 ns/op        16368 B/op            9 allocs/op
        sonic.Unmarshal               3533 ns/op         4962 B/op           12 allocs/op
      sonic-v.Unmarshal               5676 ns/op         4935 B/op           12 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal               5578 ns/op        12040 B/op            8 allocs/op
        jason.Unmarshal              38050 ns/op        55998 B/op          495 allocs/op

        sonic ████████████████████▌ 5.15
   jsonparser █████████████▋ 3.43
     fastjson █████████████▋ 3.42
          jin █████████████  3.26
      sonic-v ████████████▊ 3.21
        gjson ███████████▍ 2.86
      gjson-v ████████  2.01
         json ▓▓▓▓ 1.00
        jason █▉ 0.48
     jsoniter >>> not supported <<<
     simdjson >>> not supported <<<
      segment >>> not supported <<<
        goccy >>> not supported <<<
        codec >>> not supported <<<
           oj >>> not supported <<<
        json2 >>> not supported <<<

Unmarshal single record (2kb), read few keys into struct
         json.Unmarshal              18231 ns/op        12488 B/op           21 allocs/op
        json2.Unmarshal               6211 ns/op          128 B/op            6 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal               8603 ns/op        14087 B/op           84 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               6106 ns/op        16296 B/op            7 allocs/op
      segment.Unmarshal              10579 ns/op        12064 B/op           12 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               5726 ns/op         9058 B/op           11 allocs/op
      sonic-v.Unmarshal               7737 ns/op         9269 B/op           11 allocs/op
        codec.Unmarshal               7725 ns/op        13240 B/op           15 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<

        sonic ████████████▋ 3.18
        goccy ███████████▉ 2.99
        json2 ███████████▋ 2.94
        codec █████████▍ 2.36
      sonic-v █████████▍ 2.36
     jsoniter ████████▍ 2.12
      segment ██████▉ 1.72
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     fastjson >>> not supported <<<
           oj >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<

Unmarshal single record (2kb), read all keys generically
         json.Unmarshal              25811 ns/op        29776 B/op          340 allocs/op
        json2.Unmarshal              19701 ns/op        15573 B/op          295 allocs/op
           oj.Unmarshal              18152 ns/op        19301 B/op          465 allocs/op
     fastjson.Unmarshal               8430 ns/op        13608 B/op           73 allocs/op
     jsoniter.Unmarshal              20503 ns/op        31752 B/op          457 allocs/op
   jsonparser.Unmarshal              14454 ns/op        14528 B/op          135 allocs/op
        goccy.Unmarshal              18991 ns/op        34668 B/op          398 allocs/op
      segment.Unmarshal              65249 ns/op       173242 B/op          475 allocs/op
     simdjson.Unmarshal              23209 ns/op       132224 B/op           82 allocs/op
        gjson.Unmarshal              14147 ns/op        18336 B/op           76 allocs/op
      gjson-v.Unmarshal              17221 ns/op        18336 B/op           76 allocs/op
        sonic.Unmarshal              14219 ns/op        27726 B/op           59 allocs/op
      sonic-v.Unmarshal              16518 ns/op        27722 B/op           59 allocs/op
        codec.Unmarshal              29700 ns/op        19456 B/op          435 allocs/op
          jin.Unmarshal              23636 ns/op        25344 B/op          435 allocs/op
        jason >>> not supported <<<

     fastjson ████████████▏ 3.06
        gjson ███████▎ 1.82
        sonic ███████▎ 1.82
   jsonparser ███████▏ 1.79
      sonic-v ██████▎ 1.56
      gjson-v █████▉ 1.50
           oj █████▋ 1.42
        goccy █████▍ 1.36
        json2 █████▏ 1.31
     jsoniter █████  1.26
     simdjson ████▍ 1.11
          jin ████▎ 1.09
         json ▓▓▓▓ 1.00
        codec ███▍ 0.87
      segment █▌ 0.40
        jason >>> not supported <<<

Unmarshal single record (2kb), read all keys into struct
         json.Unmarshal              23057 ns/op        14608 B/op           80 allocs/op
        json2.Unmarshal              10134 ns/op         3111 B/op           35 allocs/op
           oj.Unmarshal              21250 ns/op         9340 B/op          453 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal               8990 ns/op        14413 B/op           77 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               6910 ns/op        16593 B/op            7 allocs/op
      segment.Unmarshal              17234 ns/op        13680 B/op           73 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               6152 ns/op         9596 B/op            9 allocs/op
      sonic-v.Unmarshal               8301 ns/op         9845 B/op            9 allocs/op
        codec.Unmarshal              10269 ns/op        14536 B/op           72 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<

        sonic ██████████████▉ 3.75
        goccy █████████████▎ 3.34
      sonic-v ███████████  2.78
     jsoniter ██████████▎ 2.56
        json2 █████████  2.28
        codec ████████▉ 2.25
      segment █████▎ 1.34
           oj ████▎ 1.09
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
     fastjson >>> not supported <<<
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys generically
         json.Unmarshal          453965446 ns/op     21755850 B/op       394052 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal           79541315 ns/op      2533410 B/op        39470 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal           60243154 ns/op      4614629 B/op       177366 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal           55637601 ns/op    616510395 B/op        39635 allocs/op
        gjson.Unmarshal           92203466 ns/op    105925336 B/op        39412 allocs/op
      gjson-v.Unmarshal          178984332 ns/op    105925072 B/op        39409 allocs/op
        sonic.Unmarshal           33596257 ns/op      2885739 B/op        78814 allocs/op
      sonic-v.Unmarshal          107582878 ns/op      2870736 B/op        78813 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal           73278149 ns/op      2652131 B/op        78811 allocs/op
        jason.Unmarshal         1007836621 ns/op   1083136112 B/op     16167166 allocs/op

        sonic ██████████████████████████████████████████████████████  13.51
     simdjson ████████████████████████████████▋ 8.16
   jsonparser ██████████████████████████████▏ 7.54
          jin ████████████████████████▊ 6.20
     fastjson ██████████████████████▊ 5.71
        gjson ███████████████████▋ 4.92
      sonic-v ████████████████▉ 4.22
      gjson-v ██████████▏ 2.54
         json ▓▓▓▓ 1.00
        jason █▊ 0.45
      segment >>> not supported <<<
        goccy >>> not supported <<<
     jsoniter >>> not supported <<<
           oj >>> not supported <<<
        codec >>> not supported <<<
        json2 >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys into struct
         json.Unmarshal          451847037 ns/op     21755738 B/op       394051 allocs/op
        json2.Unmarshal          193902196 ns/op        16804 B/op           64 allocs/op
           oj.Unmarshal          514749953 ns/op    246888960 B/op     15880754 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          209244972 ns/op     89157912 B/op      3979979 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal           91690540 ns/op    106276016 B/op        39735 allocs/op
      segment.Unmarshal          212961030 ns/op      2526041 B/op        39406 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal          113072896 ns/op    109293568 B/op        78909 allocs/op
      sonic-v.Unmarshal          185787762 ns/op    111190938 B/op        79021 allocs/op
        codec.Unmarshal          125209654 ns/op     48866478 B/op       157623 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<

        goccy ███████████████████▋ 4.93
        sonic ███████████████▉ 4.00
        codec ██████████████▍ 3.61
      sonic-v █████████▋ 2.43
        json2 █████████▎ 2.33
     jsoniter ████████▋ 2.16
      segment ████████▍ 2.12
         json ▓▓▓▓ 1.00
           oj ███▌ 0.88
   jsonparser >>> not supported <<<
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     fastjson >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys generically
         json.Unmarshal          853030813 ns/op    743920808 B/op     14315066 allocs/op
        json2.Unmarshal          751023001 ns/op    693895536 B/op     11948721 allocs/op
           oj.Unmarshal          844970174 ns/op    809200064 B/op     19911068 allocs/op
     fastjson.Unmarshal          212417524 ns/op     66540800 B/op      2876749 allocs/op
     jsoniter.Unmarshal          740738827 ns/op    820436296 B/op     19201970 allocs/op
   jsonparser.Unmarshal          493972581 ns/op    105249133 B/op      5437897 allocs/op
        goccy.Unmarshal          673082487 ns/op    879547060 B/op     16801817 allocs/op
      segment.Unmarshal         2759080012 ns/op   6770754960 B/op     19978795 allocs/op
     simdjson.Unmarshal          168459557 ns/op    688366890 B/op      2758585 allocs/op
        gjson.Unmarshal          447002429 ns/op    196083445 B/op      2994781 allocs/op
      gjson-v.Unmarshal          544296704 ns/op    196083400 B/op      2994781 allocs/op
        sonic.Unmarshal          382891388 ns/op    890224144 B/op      2285568 allocs/op
      sonic-v.Unmarshal          467888710 ns/op    890389458 B/op      2285571 allocs/op
        codec.Unmarshal         1087268206 ns/op    312883536 B/op     18323502 allocs/op
          jin.Unmarshal          856353722 ns/op    563459136 B/op     18165707 allocs/op
        jason >>> not supported <<<

     simdjson ████████████████████▎ 5.06
     fastjson ████████████████  4.02
        sonic ████████▉ 2.23
        gjson ███████▋ 1.91
      sonic-v ███████▎ 1.82
   jsonparser ██████▉ 1.73
      gjson-v ██████▎ 1.57
        goccy █████  1.27
     jsoniter ████▌ 1.15
        json2 ████▌ 1.14
           oj ████  1.01
         json ▓▓▓▓ 1.00
          jin ███▉ 1.00
        codec ███▏ 0.78
      segment █▏ 0.31
        jason >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys into struct
         json.Unmarshal          650979750 ns/op    110917592 B/op      3113016 allocs/op
        json2.Unmarshal          385190577 ns/op    129775741 B/op      1801322 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          215254515 ns/op     97217886 B/op      2955390 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal          132362573 ns/op    106193377 B/op        39604 allocs/op
      segment.Unmarshal          545140284 ns/op     68879056 B/op      2797773 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal          138159140 ns/op    133213858 B/op       157635 allocs/op
      sonic-v.Unmarshal          218069909 ns/op    133543633 B/op       157654 allocs/op
        codec.Unmarshal          274707620 ns/op    102057756 B/op      2718954 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<

        goccy ███████████████████▋ 4.92
        sonic ██████████████████▊ 4.71
     jsoniter ████████████  3.02
      sonic-v ███████████▉ 2.99
        codec █████████▍ 2.37
        json2 ██████▊ 1.69
      segment ████▊ 1.19
         json ▓▓▓▓ 1.00
     fastjson >>> not supported <<<
   jsonparser >>> not supported <<<
           oj >>> not supported <<<
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys generically
         json.Unmarshal        20416758744 ns/op    116289008 B/op      2105576 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal         3257919443 ns/op     14463336 B/op       218594 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal         2063474393 ns/op     24696416 B/op       947400 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal         1197716945 ns/op   2281348272 B/op       217476 allocs/op
        gjson.Unmarshal         3130362111 ns/op   5740687680 B/op       210563 allocs/op
      gjson-v.Unmarshal         7106292630 ns/op   5740688352 B/op       210570 allocs/op
        sonic.Unmarshal          767140790 ns/op     15285540 B/op       421122 allocs/op
      sonic-v.Unmarshal         4402580143 ns/op     15226472 B/op       421121 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal         2819503202 ns/op     14210896 B/op       421118 allocs/op
        jason.Unmarshal        37419686130 ns/op  51079786960 B/op    711849141 allocs/op

        sonic ██████████████████████████████████████████████████████████████████████████████████████████████████████████▍ 26.61
     simdjson ████████████████████████████████████████████████████████████████████▏ 17.05
   jsonparser ███████████████████████████████████████▌ 9.89
          jin ████████████████████████████▉ 7.24
        gjson ██████████████████████████  6.52
     fastjson █████████████████████████  6.27
      sonic-v ██████████████████▌ 4.64
      gjson-v ███████████▍ 2.87
         json ▓▓▓▓ 1.00
        jason ██▏ 0.55
      segment >>> not supported <<<
        goccy >>> not supported <<<
     jsoniter >>> not supported <<<
           oj >>> not supported <<<
        codec >>> not supported <<<
        json2 >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys into struct
         json.Unmarshal        20523180568 ns/op    116289008 B/op      2105576 allocs/op
        json2.Unmarshal         8974770622 ns/op        21080 B/op           72 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal         9122540622 ns/op   4650049008 B/op    210767866 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         3354452555 ns/op   5742214504 B/op       211736 allocs/op
      segment.Unmarshal         8329811144 ns/op     13537216 B/op       210563 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal         4284596747 ns/op   5757022760 B/op       421547 allocs/op
      sonic-v.Unmarshal         7715692039 ns/op   5767576768 B/op       422106 allocs/op
        codec.Unmarshal         4419457415 ns/op    261152328 B/op       842235 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<

        goccy ████████████████████████▍ 6.12
        sonic ███████████████████▏ 4.79
        codec ██████████████████▌ 4.64
      sonic-v ██████████▋ 2.66
      segment █████████▊ 2.46
        json2 █████████▏ 2.29
     jsoniter ████████▉ 2.25
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     fastjson >>> not supported <<<
           oj >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys generically
         json.Unmarshal        35191494836 ns/op  37172869776 B/op    701322482 allocs/op
        json2.Unmarshal        30040811890 ns/op  35262401952 B/op    579264595 allocs/op
           oj.Unmarshal        34221262978 ns/op  40944337240 B/op    988126089 allocs/op
     fastjson.Unmarshal         9265004423 ns/op   3342949640 B/op    142344577 allocs/op
     jsoniter.Unmarshal        30688364118 ns/op  41731641272 B/op    960520484 allocs/op
   jsonparser.Unmarshal        23947441260 ns/op   5294193176 B/op    273513559 allocs/op
        goccy.Unmarshal        27144010293 ns/op  44770355424 B/op    833149040 allocs/op
      segment.Unmarshal       115416889078 ns/op 341651494456 B/op    995514396 allocs/op
     simdjson.Unmarshal         6806750050 ns/op  27235632312 B/op    137924693 allocs/op
        gjson.Unmarshal        23015943451 ns/op  10406634000 B/op    146758251 allocs/op
      gjson-v.Unmarshal        26838849198 ns/op  10406635312 B/op    146758263 allocs/op
        sonic.Unmarshal        18474024888 ns/op  44978956416 B/op    111177914 allocs/op
      sonic-v.Unmarshal        22174702866 ns/op  44976325800 B/op    111177946 allocs/op
        codec.Unmarshal        53443847031 ns/op  13510682912 B/op    905186229 allocs/op
          jin.Unmarshal        42698143100 ns/op  28265590280 B/op    910027439 allocs/op
        jason >>> not supported <<<

     simdjson ████████████████████▋ 5.17
     fastjson ███████████████▏ 3.80
        sonic ███████▌ 1.90
      sonic-v ██████▎ 1.59
        gjson ██████  1.53
   jsonparser █████▉ 1.47
      gjson-v █████▏ 1.31
        goccy █████▏ 1.30
        json2 ████▋ 1.17
     jsoniter ████▌ 1.15
           oj ████  1.03
         json ▓▓▓▓ 1.00
          jin ███▎ 0.82
        codec ██▋ 0.66
      segment █▏ 0.30
        jason >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys into struct
         json.Unmarshal        30688048908 ns/op   4443558208 B/op    137915230 allocs/op
        json2.Unmarshal        19413076668 ns/op   6895391488 B/op     94347021 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal        10085930194 ns/op   5003847000 B/op    150338412 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         5768555079 ns/op   5754886528 B/op       223352 allocs/op
      segment.Unmarshal        25067040057 ns/op   3491109624 B/op    141915748 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal         6170504488 ns/op   7137545272 B/op      4633097 allocs/op
      sonic-v.Unmarshal         9940230262 ns/op   7153562152 B/op      4633843 allocs/op
        codec.Unmarshal        12709439569 ns/op   3062219744 B/op    133914583 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<

        goccy █████████████████████▎ 5.32
        sonic ███████████████████▉ 4.97
      sonic-v ████████████▎ 3.09
     jsoniter ████████████▏ 3.04
        codec █████████▋ 2.41
        json2 ██████▎ 1.58
      segment ████▉ 1.22
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
     fastjson >>> not supported <<<
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
           oj >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<

Marshal custom data through an object builder
         json.Marshal                  954 ns/op          688 B/op           17 allocs/op
        json2.Marshal                 1041 ns/op          424 B/op           10 allocs/op
           oj.Builder                  860 ns/op         1128 B/op           17 allocs/op
     fastjson.Marshal                  532 ns/op          679 B/op            6 allocs/op
     jsoniter.Marshal                  501 ns/op          432 B/op            7 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                  545 ns/op          416 B/op            2 allocs/op
      segment.Marshal                  552 ns/op          176 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson.Marshal                 1617 ns/op         3024 B/op           27 allocs/op
      gjson-v >>> not supported <<<
        sonic.Marshal                  692 ns/op         1803 B/op           14 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                  600 ns/op          953 B/op            0 allocs/op
          jin.Marshal                 1446 ns/op         2544 B/op           56 allocs/op
        jason >>> not supported <<<

     jsoniter ███████▌ 1.90
     fastjson ███████▏ 1.79
        goccy ███████  1.75
      segment ██████▉ 1.73
        codec ██████▎ 1.59
        sonic █████▌ 1.38
           oj ████▍ 1.11
         json ▓▓▓▓ 1.00
        json2 ███▋ 0.92
          jin ██▋ 0.66
        gjson ██▎ 0.59
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
   jsonparser >>> not supported <<<
        jason >>> not supported <<<

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
