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
