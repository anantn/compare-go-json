# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Validate []byte
         json.Validate                6606 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj.Validate                2342 ns/op            0 B/op            0 allocs/op
     fastjson.Validate                2619 ns/op            0 B/op            0 allocs/op
     jsoniter.Validate                4705 ns/op         2188 B/op          100 allocs/op
   jsonparser >>> not supported <<<
        goccy.Validate               20128 ns/op        27655 B/op          463 allocs/op
      segment.Validate                4399 ns/op            0 B/op            0 allocs/op
     simdjson >>> not supported <<<
        gjson.Validate                2232 ns/op            0 B/op            0 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                1970 ns/op            0 B/op            0 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic █████████████▍ 3.35
        gjson ███████████▊ 2.96
           oj ███████████▎ 2.82
     fastjson ██████████  2.52
      segment ██████  1.50
     jsoniter █████▌ 1.40
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
     easyjson >>> not supported <<<

Validate string
         json.Validate                6641 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Validate                3294 ns/op         4096 B/op            1 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Validate                2845 ns/op         4096 B/op            1 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                2614 ns/op         4138 B/op            1 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ██████████▏ 2.54
        gjson █████████▎ 2.33
     fastjson ████████  2.02
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
     easyjson >>> not supported <<<

Unmarshal single record (2kb), read few keys generically
         json.Unmarshal              18298 ns/op        12488 B/op           21 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal               5367 ns/op        12034 B/op            7 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal               5263 ns/op        12147 B/op           13 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal               6167 ns/op        16368 B/op            9 allocs/op
      gjson-v.Unmarshal               8647 ns/op        16368 B/op            9 allocs/op
        sonic.Unmarshal               3542 ns/op         4963 B/op           12 allocs/op
      sonic-v.Unmarshal               5672 ns/op         4948 B/op           12 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal               5361 ns/op        12040 B/op            8 allocs/op
        jason.Unmarshal              38233 ns/op        55996 B/op          495 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ████████████████████▋ 5.17
   jsonparser █████████████▉ 3.48
          jin █████████████▋ 3.41
     fastjson █████████████▋ 3.41
      sonic-v ████████████▉ 3.23
        gjson ███████████▊ 2.97
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
     easyjson >>> not supported <<<

Unmarshal single record (2kb), read few keys into struct
         json.Unmarshal              18335 ns/op        12488 B/op           21 allocs/op
        json2.Unmarshal               6153 ns/op          128 B/op            6 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal               8680 ns/op        14087 B/op           84 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               6246 ns/op        16297 B/op            7 allocs/op
      segment.Unmarshal              10565 ns/op        12064 B/op           12 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               5702 ns/op         9046 B/op           11 allocs/op
      sonic-v.Unmarshal               7771 ns/op         9297 B/op           11 allocs/op
        codec.Unmarshal               7602 ns/op        13240 B/op           15 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal              16043 ns/op        14280 B/op           36 allocs/op
     easyjson.Unmarshal              13910 ns/op        12347 B/op           16 allocs/op

        sonic ████████████▊ 3.22
        json2 ███████████▉ 2.98
        goccy ███████████▋ 2.94
        codec █████████▋ 2.41
      sonic-v █████████▍ 2.36
     jsoniter ████████▍ 2.11
      segment ██████▉ 1.74
     easyjson █████▎ 1.32
       ffjson ████▌ 1.14
         json ▓▓▓▓ 1.00
     fastjson >>> not supported <<<
      gjson-v >>> not supported <<<
   jsonparser >>> not supported <<<
           oj >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<

Unmarshal single record (2kb), read all keys generically
         json.Unmarshal              26930 ns/op        29776 B/op          340 allocs/op
        json2.Unmarshal              19831 ns/op        15572 B/op          295 allocs/op
           oj.Unmarshal              18179 ns/op        19301 B/op          465 allocs/op
     fastjson.Unmarshal               8577 ns/op        13608 B/op           73 allocs/op
     jsoniter.Unmarshal              20213 ns/op        31753 B/op          457 allocs/op
   jsonparser.Unmarshal              14210 ns/op        14528 B/op          135 allocs/op
        goccy.Unmarshal              19545 ns/op        34667 B/op          398 allocs/op
      segment.Unmarshal              65139 ns/op       173242 B/op          475 allocs/op
     simdjson.Unmarshal              23519 ns/op       132224 B/op           82 allocs/op
        gjson.Unmarshal              12791 ns/op        18336 B/op           76 allocs/op
      gjson-v.Unmarshal              15421 ns/op        18336 B/op           76 allocs/op
        sonic.Unmarshal              14027 ns/op        27733 B/op           59 allocs/op
      sonic-v.Unmarshal              16191 ns/op        27737 B/op           59 allocs/op
        codec.Unmarshal              30431 ns/op        19456 B/op          435 allocs/op
          jin.Unmarshal              22954 ns/op        25344 B/op          435 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal              14588 ns/op        32168 B/op          183 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     fastjson ████████████▌ 3.14
        gjson ████████▍ 2.11
        sonic ███████▋ 1.92
   jsonparser ███████▌ 1.90
        djson ███████▍ 1.85
      gjson-v ██████▉ 1.75
      sonic-v ██████▋ 1.66
           oj █████▉ 1.48
        goccy █████▌ 1.38
        json2 █████▍ 1.36
     jsoniter █████▎ 1.33
          jin ████▋ 1.17
     simdjson ████▌ 1.15
         json ▓▓▓▓ 1.00
        codec ███▌ 0.88
      segment █▋ 0.41
        jason >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal single record (2kb), read all keys into struct
         json.Unmarshal              23588 ns/op        14608 B/op           80 allocs/op
        json2.Unmarshal              10117 ns/op         3110 B/op           35 allocs/op
           oj.Unmarshal              21938 ns/op         9340 B/op          453 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal               8947 ns/op        14413 B/op           77 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               7312 ns/op        16597 B/op            7 allocs/op
      segment.Unmarshal              17343 ns/op        13680 B/op           73 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               6160 ns/op         9576 B/op            9 allocs/op
      sonic-v.Unmarshal               8319 ns/op         9870 B/op            9 allocs/op
        codec.Unmarshal              10582 ns/op        14536 B/op           72 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal              12541 ns/op        16884 B/op          113 allocs/op
     easyjson.Unmarshal              10235 ns/op        15888 B/op           94 allocs/op

        sonic ███████████████▎ 3.83
        goccy ████████████▉ 3.23
      sonic-v ███████████▎ 2.84
     jsoniter ██████████▌ 2.64
        json2 █████████▎ 2.33
     easyjson █████████▏ 2.30
        codec ████████▉ 2.23
       ffjson ███████▌ 1.88
      segment █████▍ 1.36
           oj ████▎ 1.08
         json ▓▓▓▓ 1.00
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
     fastjson >>> not supported <<<
   jsonparser >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys generically
         json.Unmarshal          453842977 ns/op     21755808 B/op       394052 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal           84975855 ns/op      2533404 B/op        39470 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal           58451703 ns/op      4614600 B/op       177366 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal           59775735 ns/op    618710974 B/op        39635 allocs/op
        gjson.Unmarshal           86707333 ns/op    105925548 B/op        39414 allocs/op
      gjson-v.Unmarshal          170051495 ns/op    105925296 B/op        39411 allocs/op
        sonic.Unmarshal           34974294 ns/op      2886974 B/op        78814 allocs/op
      sonic-v.Unmarshal          112645529 ns/op      2878038 B/op        78814 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal           72959803 ns/op      2652152 B/op        78811 allocs/op
        jason.Unmarshal         1023329936 ns/op   1083126032 B/op     16167139 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ███████████████████████████████████████████████████▉ 12.98
   jsonparser ███████████████████████████████  7.76
     simdjson ██████████████████████████████▎ 7.59
          jin ████████████████████████▉ 6.22
     fastjson █████████████████████▎ 5.34
        gjson ████████████████████▉ 5.23
      sonic-v ████████████████  4.03
      gjson-v ██████████▋ 2.67
         json ▓▓▓▓ 1.00
        jason █▊ 0.44
      segment >>> not supported <<<
        goccy >>> not supported <<<
     jsoniter >>> not supported <<<
           oj >>> not supported <<<
        codec >>> not supported <<<
        json2 >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys into struct
         json.Unmarshal          471453063 ns/op     21755760 B/op       394051 allocs/op
        json2.Unmarshal          194177667 ns/op        16782 B/op           64 allocs/op
           oj.Unmarshal          521491279 ns/op    246894216 B/op     15880781 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          211599264 ns/op     89158494 B/op      3979983 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal           89783212 ns/op    106275284 B/op        39740 allocs/op
      segment.Unmarshal          217857374 ns/op      2526060 B/op        39406 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal          113128627 ns/op    110345965 B/op        78922 allocs/op
      sonic-v.Unmarshal          193043472 ns/op    113028357 B/op        79039 allocs/op
        codec.Unmarshal          119163951 ns/op     48866457 B/op       157622 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal          277558531 ns/op     68132050 B/op       591323 allocs/op
     easyjson.Unmarshal          356383367 ns/op      2530973 B/op        39410 allocs/op

        goccy █████████████████████  5.25
        sonic ████████████████▋ 4.17
        codec ███████████████▊ 3.96
      sonic-v █████████▊ 2.44
        json2 █████████▋ 2.43
     jsoniter ████████▉ 2.23
      segment ████████▋ 2.16
       ffjson ██████▊ 1.70
     easyjson █████▎ 1.32
         json ▓▓▓▓ 1.00
           oj ███▌ 0.90
     fastjson >>> not supported <<<
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys generically
         json.Unmarshal          863644199 ns/op    743947080 B/op     14315151 allocs/op
        json2.Unmarshal          755441394 ns/op    693842868 B/op     11948610 allocs/op
           oj.Unmarshal          863149932 ns/op    809267744 B/op     19911303 allocs/op
     fastjson.Unmarshal          228615857 ns/op     66540809 B/op      2876749 allocs/op
     jsoniter.Unmarshal          739932089 ns/op    820420616 B/op     19201961 allocs/op
   jsonparser.Unmarshal          501272120 ns/op    105248952 B/op      5437894 allocs/op
        goccy.Unmarshal          686762657 ns/op    879512024 B/op     16801738 allocs/op
      segment.Unmarshal         2883925390 ns/op   6770752432 B/op     19978778 allocs/op
     simdjson.Unmarshal          169464188 ns/op    688369877 B/op      2758590 allocs/op
        gjson.Unmarshal          418246869 ns/op    196083408 B/op      2994781 allocs/op
      gjson-v.Unmarshal          511813198 ns/op    196083488 B/op      2994782 allocs/op
        sonic.Unmarshal          400274878 ns/op    890456192 B/op      2285573 allocs/op
      sonic-v.Unmarshal          481235843 ns/op    890442608 B/op      2285573 allocs/op
        codec.Unmarshal         1138297081 ns/op    312883264 B/op     18323501 allocs/op
          jin.Unmarshal          886679913 ns/op    563459296 B/op     18165708 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal          393459547 ns/op    770767992 B/op      7537309 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     simdjson ████████████████████▍ 5.10
     fastjson ███████████████  3.78
        djson ████████▊ 2.20
        sonic ████████▋ 2.16
        gjson ████████▎ 2.06
      sonic-v ███████▏ 1.79
   jsonparser ██████▉ 1.72
      gjson-v ██████▋ 1.69
        goccy █████  1.26
     jsoniter ████▋ 1.17
        json2 ████▌ 1.14
           oj ████  1.00
         json ▓▓▓▓ 1.00
          jin ███▉ 0.97
        codec ███  0.76
      segment █▏ 0.30
        jason >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys into struct
         json.Unmarshal          707508746 ns/op    110917608 B/op      3113016 allocs/op
        json2.Unmarshal          402163594 ns/op    129775826 B/op      1801322 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          215757748 ns/op     97217896 B/op      2955391 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal          148224935 ns/op    106194325 B/op        39623 allocs/op
      segment.Unmarshal          567826241 ns/op     68879048 B/op      2797773 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal          143633373 ns/op    133298163 B/op       157637 allocs/op
      sonic-v.Unmarshal          224078317 ns/op    133559950 B/op       157654 allocs/op
        codec.Unmarshal          298652524 ns/op    102057804 B/op      2718954 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal          399488609 ns/op    217859392 B/op      4689274 allocs/op
     easyjson.Unmarshal          255244203 ns/op    170488164 B/op      3822312 allocs/op

        sonic ███████████████████▋ 4.93
        goccy ███████████████████  4.77
     jsoniter █████████████  3.28
      sonic-v ████████████▋ 3.16
     easyjson ███████████  2.77
        codec █████████▍ 2.37
       ffjson ███████  1.77
        json2 ███████  1.76
      segment ████▉ 1.25
         json ▓▓▓▓ 1.00
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
           oj >>> not supported <<<
   jsonparser >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys generically
         json.Unmarshal        20800786995 ns/op    116289104 B/op      2105577 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal         3423202329 ns/op     14463336 B/op       218594 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal         1889455864 ns/op     24696416 B/op       947400 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal         1357192785 ns/op   2331033504 B/op       217428 allocs/op
        gjson.Unmarshal         3004360983 ns/op   5740687968 B/op       210566 allocs/op
      gjson-v.Unmarshal         6907491382 ns/op   5740688448 B/op       210571 allocs/op
        sonic.Unmarshal          788136111 ns/op     15285628 B/op       421123 allocs/op
      sonic-v.Unmarshal         4418717452 ns/op     15308392 B/op       421123 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal         2762020795 ns/op     14210880 B/op       421118 allocs/op
        jason.Unmarshal        38604032969 ns/op  51080377408 B/op    711851207 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic █████████████████████████████████████████████████████████████████████████████████████████████████████████▌ 26.39
     simdjson █████████████████████████████████████████████████████████████▎ 15.33
   jsonparser ████████████████████████████████████████████  11.01
          jin ██████████████████████████████  7.53
        gjson ███████████████████████████▋ 6.92
     fastjson ████████████████████████▎ 6.08
      sonic-v ██████████████████▊ 4.71
      gjson-v ████████████  3.01
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
     easyjson >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys into struct
         json.Unmarshal        20859629140 ns/op    116289008 B/op      2105576 allocs/op
        json2.Unmarshal         9029062157 ns/op        21080 B/op           72 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal         9122245432 ns/op   4650047680 B/op    210767856 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         3242205251 ns/op   5742216192 B/op       211769 allocs/op
      segment.Unmarshal         8446919698 ns/op     13537216 B/op       210563 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal         4420801805 ns/op   5758300632 B/op       421558 allocs/op
      sonic-v.Unmarshal         8047327728 ns/op   5770238848 B/op       422135 allocs/op
        codec.Unmarshal         4165879245 ns/op    261152232 B/op       842234 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal        10494894497 ns/op    362386512 B/op      3158490 allocs/op
     easyjson.Unmarshal        17620027028 ns/op     13542760 B/op       210581 allocs/op

        goccy █████████████████████████▋ 6.43
        codec ████████████████████  5.01
        sonic ██████████████████▊ 4.72
      sonic-v ██████████▎ 2.59
      segment █████████▉ 2.47
        json2 █████████▏ 2.31
     jsoniter █████████▏ 2.29
       ffjson ███████▉ 1.99
     easyjson ████▋ 1.18
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
         json.Unmarshal        36745543842 ns/op  37172682272 B/op    701321850 allocs/op
        json2.Unmarshal        30663492103 ns/op  35261735552 B/op    579262299 allocs/op
           oj.Unmarshal        34595182232 ns/op  40944594104 B/op    988126985 allocs/op
     fastjson.Unmarshal         9967256564 ns/op   3342950136 B/op    142344582 allocs/op
     jsoniter.Unmarshal        31727823673 ns/op  41731425592 B/op    960519760 allocs/op
   jsonparser.Unmarshal        23781044885 ns/op   5294193160 B/op    273513557 allocs/op
        goccy.Unmarshal        28848710756 ns/op  44770003224 B/op    833148538 allocs/op
      segment.Unmarshal       116716317041 ns/op 341651510616 B/op    995514479 allocs/op
     simdjson.Unmarshal         6699238726 ns/op  27143091576 B/op    137924740 allocs/op
        gjson.Unmarshal        20677191279 ns/op  10406633104 B/op    146758242 allocs/op
      gjson-v.Unmarshal        25099363758 ns/op  10406634224 B/op    146758253 allocs/op
        sonic.Unmarshal        18565350820 ns/op  44978148728 B/op    111177867 allocs/op
      sonic-v.Unmarshal        22363505928 ns/op  44976910760 B/op    111177979 allocs/op
        codec.Unmarshal        55028185032 ns/op  13510682144 B/op    905186218 allocs/op
          jin.Unmarshal        43753709217 ns/op  28265592600 B/op    910027464 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal        19051410994 ns/op  39791446096 B/op    375169899 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     simdjson █████████████████████▉ 5.49
     fastjson ██████████████▋ 3.69
        sonic ███████▉ 1.98
        djson ███████▋ 1.93
        gjson ███████  1.78
      sonic-v ██████▌ 1.64
   jsonparser ██████▏ 1.55
      gjson-v █████▊ 1.46
        goccy █████  1.27
        json2 ████▊ 1.20
     jsoniter ████▋ 1.16
           oj ████▏ 1.06
         json ▓▓▓▓ 1.00
          jin ███▎ 0.84
        codec ██▋ 0.67
      segment █▎ 0.31
        jason >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys into struct
         json.Unmarshal        32640135957 ns/op   4443558224 B/op    137915229 allocs/op
        json2.Unmarshal        20037101043 ns/op   6895391408 B/op     94347015 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal        10001270299 ns/op   5003845592 B/op    150338398 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         6540961686 ns/op   5754911704 B/op       223863 allocs/op
      segment.Unmarshal        25776321918 ns/op   3491109576 B/op    141915744 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal         6270915260 ns/op   7135841976 B/op      4633088 allocs/op
      sonic-v.Unmarshal         9926410359 ns/op   7152570248 B/op      4633824 allocs/op
        codec.Unmarshal        13775232351 ns/op   3062218816 B/op    133914575 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal        17440112923 ns/op  10026583896 B/op    215402572 allocs/op
     easyjson.Unmarshal        12246161549 ns/op   8947230320 B/op    198556552 allocs/op

        sonic ████████████████████▊ 5.21
        goccy ███████████████████▉ 4.99
      sonic-v █████████████▏ 3.29
     jsoniter █████████████  3.26
     easyjson ██████████▋ 2.67
        codec █████████▍ 2.37
       ffjson ███████▍ 1.87
        json2 ██████▌ 1.63
      segment █████  1.27
         json ▓▓▓▓ 1.00
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
           oj >>> not supported <<<
     simdjson >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
     fastjson >>> not supported <<<
   jsonparser >>> not supported <<<

Marshal custom data through an object builder
         json.Marshal                 1022 ns/op          688 B/op           17 allocs/op
        json2.Marshal                 1069 ns/op          424 B/op           10 allocs/op
           oj.Builder                  880 ns/op         1128 B/op           17 allocs/op
     fastjson.Marshal                  572 ns/op          679 B/op            6 allocs/op
     jsoniter.Marshal                  502 ns/op          432 B/op            7 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                  552 ns/op          416 B/op            2 allocs/op
      segment.Marshal                  561 ns/op          176 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson.Marshal                 1631 ns/op         3024 B/op           27 allocs/op
      gjson-v >>> not supported <<<
        sonic.Marshal                  694 ns/op         1803 B/op           14 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                  603 ns/op          952 B/op            0 allocs/op
          jin.Marshal                 1450 ns/op         2544 B/op           56 allocs/op
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Marshal                  145 ns/op          144 B/op            2 allocs/op
     easyjson.Marshal                  109 ns/op          288 B/op            3 allocs/op

     easyjson █████████████████████████████████████▌ 9.38
       ffjson ████████████████████████████▏ 7.05
     jsoniter ████████▏ 2.04
        goccy ███████▍ 1.85
      segment ███████▎ 1.82
     fastjson ███████▏ 1.79
        codec ██████▊ 1.69
        sonic █████▉ 1.47
           oj ████▋ 1.16
         json ▓▓▓▓ 1.00
        json2 ███▊ 0.96
          jin ██▊ 0.70
        gjson ██▌ 0.63
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
      gjson-v >>> not supported <<<
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
         json.Validate               10430 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj.Validate                3730 ns/op            0 B/op            0 allocs/op
     fastjson.Validate                4868 ns/op            0 B/op            0 allocs/op
     jsoniter.Validate                6758 ns/op         2184 B/op          100 allocs/op
   jsonparser >>> not supported <<<
        goccy.Validate               24413 ns/op        27603 B/op          463 allocs/op
      segment.Validate                3580 ns/op            0 B/op            0 allocs/op
     simdjson >>> not supported <<<
        gjson.Validate                3520 ns/op            0 B/op            0 allocs/op
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        gjson ███████████▊ 2.96
      segment ███████████▋ 2.91
           oj ███████████▏ 2.80
     fastjson ████████▌ 2.14
     jsoniter ██████▏ 1.54
         json ▓▓▓▓ 1.00
        goccy █▋ 0.43
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
        json2 >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Validate string
         json.Validate               10426 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Validate                5246 ns/op         4096 B/op            1 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Validate                3887 ns/op         4096 B/op            1 allocs/op
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        gjson ██████████▋ 2.68
     fastjson ███████▉ 1.99
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
      gjson-v >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
           oj >>> not supported <<<
        json2 >>> not supported <<<
     jsoniter >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal single record (2kb), read few keys generically
         json.Unmarshal              27461 ns/op        12488 B/op           21 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal               8197 ns/op        12034 B/op            7 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal               8014 ns/op        12147 B/op           13 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal               9311 ns/op        16368 B/op            9 allocs/op
      gjson-v.Unmarshal              12779 ns/op        16368 B/op            9 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin.Unmarshal               7798 ns/op        12040 B/op            8 allocs/op
        jason.Unmarshal              51099 ns/op        55997 B/op          495 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

          jin ██████████████  3.52
   jsonparser █████████████▋ 3.43
     fastjson █████████████▍ 3.35
        gjson ███████████▊ 2.95
      gjson-v ████████▌ 2.15
         json ▓▓▓▓ 1.00
        jason ██▏ 0.54
     simdjson >>> not supported <<<
      segment >>> not supported <<<
        goccy >>> not supported <<<
     jsoniter >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
           oj >>> not supported <<<
        json2 >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal single record (2kb), read few keys into struct
         json.Unmarshal              27485 ns/op        12488 B/op           21 allocs/op
        json2.Unmarshal               8911 ns/op          128 B/op            6 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal              12302 ns/op        14063 B/op           84 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               8980 ns/op        16171 B/op            7 allocs/op
      segment.Unmarshal               9547 ns/op        12064 B/op           12 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal              10529 ns/op        13240 B/op           15 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal              21973 ns/op        14206 B/op           36 allocs/op
     easyjson.Unmarshal              20963 ns/op        12326 B/op           16 allocs/op

        json2 ████████████▎ 3.08
        goccy ████████████▏ 3.06
      segment ███████████▌ 2.88
        codec ██████████▍ 2.61
     jsoniter ████████▉ 2.23
     easyjson █████▏ 1.31
       ffjson █████  1.25
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
   jsonparser >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
     fastjson >>> not supported <<<
           oj >>> not supported <<<

Unmarshal single record (2kb), read all keys generically
         json.Unmarshal              36170 ns/op        29776 B/op          340 allocs/op
        json2.Unmarshal              23645 ns/op        15545 B/op          295 allocs/op
           oj.Unmarshal              23141 ns/op        19300 B/op          465 allocs/op
     fastjson.Unmarshal              13301 ns/op        13608 B/op           73 allocs/op
     jsoniter.Unmarshal              27625 ns/op        31697 B/op          457 allocs/op
   jsonparser.Unmarshal              20460 ns/op        14528 B/op          135 allocs/op
        goccy.Unmarshal              25519 ns/op        34556 B/op          398 allocs/op
      segment.Unmarshal              62445 ns/op       173242 B/op          475 allocs/op
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x10519945c]

goroutine 88 [running]:
github.com/minio/simdjson-go.(*ParsedJson).ForEach(0x14000172788?, 0x1053719e8)
	/Users/anantn/go/pkg/mod/github.com/minio/simdjson-go@v0.4.5/parsed_json.go:126 +0x3c
main.simdjsonFile1All(0x14000172788)
	/Users/anantn/Code/compare-go-json/simdjson.go:52 +0x1e0
testing.(*B).runN(0x14000172788, 0x1)
	/usr/local/go/src/testing/benchmark.go:193 +0x130
testing.(*B).run1.func1()
	/usr/local/go/src/testing/benchmark.go:215 +0x50
created by testing.(*B).run1 in goroutine 1
	/usr/local/go/src/testing/benchmark.go:208 +0x90

compare-go-json on  master via 🐹 v1.22.1 took 55s 
❯ code .

compare-go-json on  master via 🐹 v1.22.1 
❯ go build ./... && ./go-json-benchmarks
WARNING: sonic only supports Go1.16~1.22 && CPU amd64, but your environment is not suitable

Validate []byte
         json.Validate               10462 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj.Validate                3737 ns/op            0 B/op            0 allocs/op
     fastjson.Validate                4866 ns/op            0 B/op            0 allocs/op
     jsoniter.Validate                6663 ns/op         2184 B/op          100 allocs/op
   jsonparser >>> not supported <<<
        goccy.Validate               23178 ns/op        27601 B/op          463 allocs/op
      segment.Validate                3553 ns/op            0 B/op            0 allocs/op
     simdjson >>> not supported <<<
        gjson.Validate                3519 ns/op            0 B/op            0 allocs/op
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        gjson ███████████▉ 2.97
      segment ███████████▊ 2.94
           oj ███████████▏ 2.80
     fastjson ████████▌ 2.15
     jsoniter ██████▎ 1.57
         json ▓▓▓▓ 1.00
        goccy █▊ 0.45
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
        json2 >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Validate string
         json.Validate               10425 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Validate                5234 ns/op         4096 B/op            1 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Validate                3885 ns/op         4096 B/op            1 allocs/op
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        gjson ██████████▋ 2.68
     fastjson ███████▉ 1.99
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
      gjson-v >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
           oj >>> not supported <<<
        json2 >>> not supported <<<
     jsoniter >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal single record (2kb), read few keys generically
         json.Unmarshal              27550 ns/op        12488 B/op           21 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal               8240 ns/op        12034 B/op            7 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal               8049 ns/op        12147 B/op           13 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal               9373 ns/op        16368 B/op            9 allocs/op
      gjson-v.Unmarshal              12839 ns/op        16368 B/op            9 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin.Unmarshal               7919 ns/op        12040 B/op            8 allocs/op
        jason.Unmarshal              51388 ns/op        55998 B/op          495 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

          jin █████████████▉ 3.48
   jsonparser █████████████▋ 3.42
     fastjson █████████████▎ 3.34
        gjson ███████████▊ 2.94
      gjson-v ████████▌ 2.15
         json ▓▓▓▓ 1.00
        jason ██▏ 0.54
     simdjson >>> not supported <<<
      segment >>> not supported <<<
        goccy >>> not supported <<<
     jsoniter >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
           oj >>> not supported <<<
        json2 >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal single record (2kb), read few keys into struct
         json.Unmarshal              27479 ns/op        12488 B/op           21 allocs/op
        json2.Unmarshal               8879 ns/op          128 B/op            6 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal              12310 ns/op        14063 B/op           84 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               9010 ns/op        16171 B/op            7 allocs/op
      segment.Unmarshal               9640 ns/op        12064 B/op           12 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal              10552 ns/op        13240 B/op           15 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal              21894 ns/op        14206 B/op           36 allocs/op
     easyjson.Unmarshal              20998 ns/op        12326 B/op           16 allocs/op

        json2 ████████████▍ 3.09
        goccy ████████████▏ 3.05
      segment ███████████▍ 2.85
        codec ██████████▍ 2.60
     jsoniter ████████▉ 2.23
     easyjson █████▏ 1.31
       ffjson █████  1.26
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
   jsonparser >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
     fastjson >>> not supported <<<
           oj >>> not supported <<<

Unmarshal single record (2kb), read all keys generically
         json.Unmarshal              36293 ns/op        29776 B/op          340 allocs/op
        json2.Unmarshal              23583 ns/op        15545 B/op          295 allocs/op
           oj.Unmarshal              23403 ns/op        19301 B/op          465 allocs/op
     fastjson.Unmarshal              13352 ns/op        13608 B/op           73 allocs/op
     jsoniter.Unmarshal              27788 ns/op        31698 B/op          457 allocs/op
   jsonparser.Unmarshal              20160 ns/op        14528 B/op          135 allocs/op
        goccy.Unmarshal              25621 ns/op        34555 B/op          398 allocs/op
      segment.Unmarshal              63477 ns/op       173241 B/op          475 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal              18841 ns/op        18336 B/op           76 allocs/op
      gjson-v.Unmarshal              22323 ns/op        18336 B/op           76 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal              37157 ns/op        19456 B/op          435 allocs/op
          jin.Unmarshal              28743 ns/op        25344 B/op          435 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal              20707 ns/op        32167 B/op          183 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     fastjson ██████████▊ 2.72
        gjson ███████▋ 1.93
   jsonparser ███████▏ 1.80
        djson ███████  1.75
      gjson-v ██████▌ 1.63
           oj ██████▏ 1.55
        json2 ██████▏ 1.54
        goccy █████▋ 1.42
     jsoniter █████▏ 1.31
          jin █████  1.26
         json ▓▓▓▓ 1.00
        codec ███▉ 0.98
      segment ██▎ 0.57
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
        jason >>> not supported <<<
        sonic >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal single record (2kb), read all keys into struct
         json.Unmarshal              33099 ns/op        14608 B/op           80 allocs/op
        json2.Unmarshal              14032 ns/op         3106 B/op           35 allocs/op
           oj.Unmarshal              26592 ns/op         9340 B/op          453 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal              13536 ns/op        14392 B/op           77 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal              11101 ns/op        16288 B/op            7 allocs/op
      segment.Unmarshal              15112 ns/op        13680 B/op           73 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal              15330 ns/op        14536 B/op           72 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal              18800 ns/op        16820 B/op          113 allocs/op
     easyjson.Unmarshal              16354 ns/op        15888 B/op           94 allocs/op

        goccy ███████████▉ 2.98
     jsoniter █████████▊ 2.45
        json2 █████████▍ 2.36
      segment ████████▊ 2.19
        codec ████████▋ 2.16
     easyjson ████████  2.02
       ffjson ███████  1.76
           oj ████▉ 1.24
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
   jsonparser >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys generically
         json.Unmarshal          663396479 ns/op     21756096 B/op       394053 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal          126024989 ns/op      2537813 B/op        39513 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal           87215115 ns/op      4614347 B/op       177229 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal          122017069 ns/op    105925440 B/op        39413 allocs/op
      gjson-v.Unmarshal          235885208 ns/op    105925177 B/op        39410 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin.Unmarshal           95803284 ns/op      2652133 B/op        78811 allocs/op
        jason.Unmarshal         1343679500 ns/op   1082892592 B/op     16167067 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

   jsonparser ██████████████████████████████▍ 7.61
          jin ███████████████████████████▋ 6.92
        gjson █████████████████████▋ 5.44
     fastjson █████████████████████  5.26
      gjson-v ███████████▏ 2.81
         json ▓▓▓▓ 1.00
        jason █▉ 0.49
     simdjson >>> not supported <<<
      segment >>> not supported <<<
        goccy >>> not supported <<<
     jsoniter >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
           oj >>> not supported <<<
        json2 >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys into struct
         json.Unmarshal          662870500 ns/op     21755744 B/op       394052 allocs/op
        json2.Unmarshal          253062437 ns/op        16410 B/op           67 allocs/op
           oj.Unmarshal          669727062 ns/op    246953388 B/op     15880755 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          285867979 ns/op     89442694 B/op      3979980 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal          125449421 ns/op    106045892 B/op        39739 allocs/op
      segment.Unmarshal          181779312 ns/op      2526037 B/op        39406 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal          170802631 ns/op     48866458 B/op       157622 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal          387229111 ns/op     67899648 B/op       591339 allocs/op
     easyjson.Unmarshal          531779021 ns/op      2527752 B/op        39414 allocs/op

        goccy █████████████████████▏ 5.28
        codec ███████████████▌ 3.88
      segment ██████████████▌ 3.65
        json2 ██████████▍ 2.62
     jsoniter █████████▎ 2.32
       ffjson ██████▊ 1.71
     easyjson ████▉ 1.25
         json ▓▓▓▓ 1.00
           oj ███▉ 0.99
     simdjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
   jsonparser >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys generically
         json.Unmarshal         1157970375 ns/op    743912176 B/op     14315048 allocs/op
        json2.Unmarshal          949824666 ns/op    693397368 B/op     11939627 allocs/op
           oj.Unmarshal         1157516917 ns/op    809306824 B/op     19911292 allocs/op
     fastjson.Unmarshal          364131194 ns/op     66551413 B/op      2876854 allocs/op
     jsoniter.Unmarshal          990867583 ns/op    819520936 B/op     19202037 allocs/op
   jsonparser.Unmarshal          639996646 ns/op    105469824 B/op      5437894 allocs/op
        goccy.Unmarshal          897758375 ns/op    877665964 B/op     16802547 allocs/op
      segment.Unmarshal         3014437000 ns/op   6770606656 B/op     19978945 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal          641942229 ns/op    196084576 B/op      2994787 allocs/op
      gjson-v.Unmarshal          778361166 ns/op    196084632 B/op      2994789 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal         1434169583 ns/op    313138064 B/op     18323511 allocs/op
          jin.Unmarshal         1147027459 ns/op    563339816 B/op     18165734 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal          647903291 ns/op    770795888 B/op      7537422 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     fastjson ████████████▋ 3.18
   jsonparser ███████▏ 1.81
        gjson ███████▏ 1.80
        djson ███████▏ 1.79
      gjson-v █████▉ 1.49
        goccy █████▏ 1.29
        json2 ████▉ 1.22
     jsoniter ████▋ 1.17
          jin ████  1.01
           oj ████  1.00
         json ▓▓▓▓ 1.00
        codec ███▏ 0.81
      segment █▌ 0.38
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
        jason >>> not supported <<<
        sonic >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys into struct
         json.Unmarshal          933990583 ns/op    111143072 B/op      3113020 allocs/op
        json2.Unmarshal          527895917 ns/op    129790324 B/op      1792490 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          306623958 ns/op     97329814 B/op      2955472 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal          210295891 ns/op    107153896 B/op        42818 allocs/op
      segment.Unmarshal          402844430 ns/op     69133245 B/op      2797769 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal          393172819 ns/op    102290645 B/op      2718959 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal          557960687 ns/op    218509932 B/op      4690382 allocs/op
     easyjson.Unmarshal          425623097 ns/op    170814813 B/op      3822747 allocs/op

        goccy █████████████████▊ 4.44
     jsoniter ████████████▏ 3.05
        codec █████████▌ 2.38
      segment █████████▎ 2.32
     easyjson ████████▊ 2.19
        json2 ███████  1.77
       ffjson ██████▋ 1.67
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
   jsonparser >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
     fastjson >>> not supported <<<
           oj >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys generically
         json.Unmarshal        31256352459 ns/op    116289888 B/op      2105581 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal         5306753792 ns/op     14548640 B/op       218834 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal         2907524375 ns/op     24696736 B/op       947485 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal         4791219583 ns/op   5740738176 B/op       211089 allocs/op
      gjson-v.Unmarshal        10465913625 ns/op   5740728480 B/op       210956 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin.Unmarshal         3547572875 ns/op     14211056 B/op       421119 allocs/op
        jason.Unmarshal        63233934791 ns/op  51092971856 B/op    711854429 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

   jsonparser ███████████████████████████████████████████  10.75
          jin ███████████████████████████████████▏ 8.81
        gjson ██████████████████████████  6.52
     fastjson ███████████████████████▌ 5.89
      gjson-v ███████████▉ 2.99
         json ▓▓▓▓ 1.00
        jason █▉ 0.49
     simdjson >>> not supported <<<
      segment >>> not supported <<<
        goccy >>> not supported <<<
     jsoniter >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
           oj >>> not supported <<<
        json2 >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys into struct
         json.Unmarshal        31213790583 ns/op    116289600 B/op      2105580 allocs/op
        json2.Unmarshal        12104046042 ns/op        17624 B/op           72 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal        13603473209 ns/op   4644531120 B/op    210771942 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         5188465750 ns/op   5748746056 B/op       231824 allocs/op
      segment.Unmarshal         7626735375 ns/op     13537216 B/op       210563 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal         6482178875 ns/op    261153384 B/op       842246 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal        16537171292 ns/op    362913888 B/op      3160013 allocs/op
     easyjson.Unmarshal        26878690042 ns/op     13547936 B/op       210631 allocs/op

        goccy ████████████████████████  6.02
        codec ███████████████████▎ 4.82
      segment ████████████████▎ 4.09
        json2 ██████████▎ 2.58
     jsoniter █████████▏ 2.29
       ffjson ███████▌ 1.89
     easyjson ████▋ 1.16
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
   jsonparser >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
     fastjson >>> not supported <<<
           oj >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys generically
         json.Unmarshal        59162200917 ns/op  37196790208 B/op    701326642 allocs/op
        json2.Unmarshal        48371781417 ns/op  35262026584 B/op    579152682 allocs/op
           oj.Unmarshal        62492061084 ns/op  40969488264 B/op    988134641 allocs/op
     fastjson.Unmarshal        17905291041 ns/op   3343060096 B/op    142344995 allocs/op
     jsoniter.Unmarshal        52463296250 ns/op  41776313344 B/op    960571419 allocs/op
   jsonparser.Unmarshal        32124871375 ns/op   5277725704 B/op    273513675 allocs/op
        goccy.Unmarshal        47378195500 ns/op  44837337656 B/op    833470674 allocs/op
      segment.Unmarshal       146381703709 ns/op 341675898856 B/op    995540014 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal        31415638125 ns/op  10406698784 B/op    146758700 allocs/op
      gjson-v.Unmarshal        37213082084 ns/op  10406686896 B/op    146758588 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal        69553223041 ns/op  13524897808 B/op    905187844 allocs/op
          jin.Unmarshal        56425113875 ns/op  28244853656 B/op    910029005 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal        33990522333 ns/op  39791627456 B/op    375172643 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     fastjson █████████████▏ 3.30
        gjson ███████▌ 1.88
   jsonparser ███████▎ 1.84
        djson ██████▉ 1.74
      gjson-v ██████▎ 1.59
        goccy ████▉ 1.25
        json2 ████▉ 1.22
     jsoniter ████▌ 1.13
          jin ████▏ 1.05
         json ▓▓▓▓ 1.00
           oj ███▊ 0.95
        codec ███▍ 0.85
      segment █▌ 0.40
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
        jason >>> not supported <<<
        sonic >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys into struct
         json.Unmarshal        43829186083 ns/op   4426671440 B/op    137915338 allocs/op
        json2.Unmarshal        27201680958 ns/op   6898796264 B/op     94203710 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal        14564398292 ns/op   4989275008 B/op    150343002 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         9151100916 ns/op   5821500096 B/op       434577 allocs/op
      segment.Unmarshal        19519526583 ns/op   3474104424 B/op    141915804 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal        18220652125 ns/op   3045079632 B/op    133914646 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal        24764420875 ns/op  10030680192 B/op    215459150 allocs/op
     easyjson.Unmarshal        20962012875 ns/op   8935461448 B/op    198579608 allocs/op

        goccy ███████████████████▏ 4.79
     jsoniter ████████████  3.01
        codec █████████▌ 2.41
      segment ████████▉ 2.25
     easyjson ████████▎ 2.09
       ffjson ███████  1.77
        json2 ██████▍ 1.61
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
   jsonparser >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
     fastjson >>> not supported <<<
           oj >>> not supported <<<

Marshal custom data through an object builder
         json.Marshal                 1268 ns/op          688 B/op           17 allocs/op
        json2.Marshal                 1249 ns/op          424 B/op           10 allocs/op
           oj.Builder                 1419 ns/op         1129 B/op           17 allocs/op
     fastjson.Marshal                  865 ns/op          679 B/op            6 allocs/op
     jsoniter.Marshal                  707 ns/op          432 B/op            7 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                  816 ns/op          416 B/op            2 allocs/op
      segment.Marshal                  740 ns/op          176 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson.Marshal                 2599 ns/op         3026 B/op           27 allocs/op
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Marshal                  770 ns/op          969 B/op            0 allocs/op
          jin.Marshal                 2307 ns/op         2546 B/op           56 allocs/op
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Marshal                  207 ns/op          144 B/op            2 allocs/op
     easyjson.Marshal                  185 ns/op          288 B/op            3 allocs/op

     easyjson ███████████████████████████▍ 6.85
       ffjson ████████████████████████▌ 6.13
     jsoniter ███████▏ 1.79
      segment ██████▊ 1.71
        codec ██████▌ 1.65
        goccy ██████▏ 1.55
     fastjson █████▊ 1.47
        json2 ████  1.02
         json ▓▓▓▓ 1.00
           oj ███▌ 0.89
          jin ██▏ 0.55
        gjson █▉ 0.49
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
        sonic >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
      gjson-v >>> not supported <<<
   jsonparser >>> not supported <<<

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
