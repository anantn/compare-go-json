# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Unmarshal single record (2kb), read few keys generically
         json.Unmarshal              18422 ns/op        12488 B/op           21 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal               5310 ns/op        12034 B/op            7 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal               5314 ns/op        12208 B/op           17 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal               6091 ns/op        16368 B/op            9 allocs/op
      gjson-v.Unmarshal               8703 ns/op        16368 B/op            9 allocs/op
        sonic.Unmarshal               3548 ns/op         4960 B/op           12 allocs/op
      sonic-v.Unmarshal               5816 ns/op         4941 B/op           12 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal               5625 ns/op        12344 B/op           10 allocs/op
        jason.Unmarshal              39396 ns/op        55999 B/op          495 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ████████████████████▊ 5.19
     fastjson █████████████▉ 3.47
   jsonparser █████████████▊ 3.47
          jin █████████████  3.28
      sonic-v ████████████▋ 3.17
        gjson ████████████  3.02
      gjson-v ████████▍ 2.12
         json ▓▓▓▓ 1.00
        jason █▊ 0.47
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
         json.Unmarshal              18326 ns/op        12488 B/op           21 allocs/op
        json2.Unmarshal               6223 ns/op          128 B/op            6 allocs/op
           oj >>> not supported <<<
     fastjson.Unmarshal               5485 ns/op        12080 B/op           10 allocs/op
     jsoniter.Unmarshal               8633 ns/op        14087 B/op           84 allocs/op
   jsonparser.Unmarshal               5488 ns/op        12208 B/op           17 allocs/op
        goccy.Unmarshal               6136 ns/op        16296 B/op            7 allocs/op
      segment.Unmarshal              10509 ns/op        12064 B/op           12 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal               6292 ns/op        16400 B/op           10 allocs/op
      gjson-v.Unmarshal               8940 ns/op        16400 B/op           10 allocs/op
        sonic.Unmarshal               5792 ns/op         9023 B/op           11 allocs/op
      sonic-v.Unmarshal               7900 ns/op         9354 B/op           11 allocs/op
        codec.Unmarshal               7568 ns/op        13240 B/op           15 allocs/op
          jin.Unmarshal               5578 ns/op        12344 B/op           10 allocs/op
        jason.Unmarshal              39044 ns/op        55995 B/op          495 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal              15921 ns/op        14280 B/op           36 allocs/op
     easyjson.Unmarshal              13645 ns/op        12346 B/op           16 allocs/op

   * fastjson █████████████▎ 3.34
 * jsonparser █████████████▎ 3.34
        * jin █████████████▏ 3.29
        sonic ████████████▋ 3.16
        goccy ███████████▉ 2.99
        json2 ███████████▊ 2.94
      * gjson ███████████▋ 2.91
        codec █████████▋ 2.42
      sonic-v █████████▎ 2.32
     jsoniter ████████▍ 2.12
    * gjson-v ████████▏ 2.05
      segment ██████▉ 1.74
     easyjson █████▎ 1.34
       ffjson ████▌ 1.15
         json ▓▓▓▓ 1.00
      * jason █▉ 0.47
     simdjson >>> not supported <<<
           oj >>> not supported <<<
        djson >>> not supported <<<

Unmarshal single record (2kb), read all keys generically
         json.Unmarshal              26155 ns/op        29777 B/op          340 allocs/op
        json2.Unmarshal              19575 ns/op        15574 B/op          295 allocs/op
           oj.Unmarshal              18310 ns/op        19299 B/op          465 allocs/op
     fastjson.Unmarshal               8616 ns/op        13608 B/op           73 allocs/op
     jsoniter.Unmarshal              20484 ns/op        31751 B/op          457 allocs/op
   jsonparser.Unmarshal              14631 ns/op        14528 B/op          135 allocs/op
        goccy.Unmarshal              19403 ns/op        34667 B/op          398 allocs/op
      segment.Unmarshal              65652 ns/op       173242 B/op          475 allocs/op
     simdjson.Unmarshal              23437 ns/op       132224 B/op           82 allocs/op
        gjson.Unmarshal              12845 ns/op        18336 B/op           76 allocs/op
      gjson-v.Unmarshal              15422 ns/op        18336 B/op           76 allocs/op
        sonic.Unmarshal              13861 ns/op        27717 B/op           59 allocs/op
      sonic-v.Unmarshal              16082 ns/op        27732 B/op           59 allocs/op
        codec.Unmarshal              30222 ns/op        19456 B/op          435 allocs/op
          jin.Unmarshal              23291 ns/op        25344 B/op          435 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal              14405 ns/op        32169 B/op          183 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     fastjson ████████████▏ 3.04
        gjson ████████▏ 2.04
        sonic ███████▌ 1.89
        djson ███████▎ 1.82
   jsonparser ███████▏ 1.79
      gjson-v ██████▊ 1.70
      sonic-v ██████▌ 1.63
           oj █████▋ 1.43
        goccy █████▍ 1.35
        json2 █████▎ 1.34
     jsoniter █████  1.28
          jin ████▍ 1.12
     simdjson ████▍ 1.12
         json ▓▓▓▓ 1.00
        codec ███▍ 0.87
      segment █▌ 0.40
        jason >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal single record (2kb), read all keys into struct
         json.Unmarshal              23542 ns/op        14608 B/op           80 allocs/op
        json2.Unmarshal              10325 ns/op         3112 B/op           35 allocs/op
           oj.Unmarshal              22050 ns/op         9340 B/op          453 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal               8894 ns/op        14412 B/op           77 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               7295 ns/op        16593 B/op            7 allocs/op
      segment.Unmarshal              17329 ns/op        13680 B/op           73 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               6103 ns/op         9583 B/op            9 allocs/op
      sonic-v.Unmarshal               8333 ns/op         9846 B/op            9 allocs/op
        codec.Unmarshal              10558 ns/op        14536 B/op           72 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal              12445 ns/op        16885 B/op          113 allocs/op
     easyjson.Unmarshal              10192 ns/op        15888 B/op           94 allocs/op

        sonic ███████████████▍ 3.86
        goccy ████████████▉ 3.23
      sonic-v ███████████▎ 2.83
     jsoniter ██████████▌ 2.65
     easyjson █████████▏ 2.31
        json2 █████████  2.28
        codec ████████▉ 2.23
       ffjson ███████▌ 1.89
      segment █████▍ 1.36
           oj ████▎ 1.07
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
         json.Unmarshal          468050367 ns/op     21755872 B/op       394053 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal           82465056 ns/op      2533964 B/op        39475 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal           61157062 ns/op      4614627 B/op       177366 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal           60419757 ns/op    608536160 B/op        39635 allocs/op
        gjson.Unmarshal           85583831 ns/op    105925341 B/op        39412 allocs/op
      gjson-v.Unmarshal          169863404 ns/op    105925472 B/op        39413 allocs/op
        sonic.Unmarshal           34665969 ns/op      2886972 B/op        78814 allocs/op
      sonic-v.Unmarshal          111946963 ns/op      2873458 B/op        78813 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal           70550502 ns/op      2652124 B/op        78811 allocs/op
        jason.Unmarshal         1012626669 ns/op   1083115712 B/op     16167109 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ██████████████████████████████████████████████████████  13.50
     simdjson ██████████████████████████████▉ 7.75
   jsonparser ██████████████████████████████▌ 7.65
          jin ██████████████████████████▌ 6.63
     fastjson ██████████████████████▋ 5.68
        gjson █████████████████████▉ 5.47
      sonic-v ████████████████▋ 4.18
      gjson-v ███████████  2.76
         json ▓▓▓▓ 1.00
        jason █▊ 0.46
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
         json.Unmarshal          454888666 ns/op     21755856 B/op       394052 allocs/op
        json2.Unmarshal          198867073 ns/op        16782 B/op           64 allocs/op
           oj.Unmarshal          521089645 ns/op    246894384 B/op     15880781 allocs/op
     fastjson.Unmarshal           84230234 ns/op      2533957 B/op        39475 allocs/op
     jsoniter.Unmarshal          212750758 ns/op     89158363 B/op      3979981 allocs/op
   jsonparser.Unmarshal           62351832 ns/op      4614626 B/op       177366 allocs/op
        goccy.Unmarshal           88738817 ns/op    106274096 B/op        39733 allocs/op
      segment.Unmarshal          223378935 ns/op      2526060 B/op        39406 allocs/op
     simdjson.Unmarshal           60902730 ns/op    602137348 B/op        39632 allocs/op
        gjson.Unmarshal           86691870 ns/op    105925186 B/op        39410 allocs/op
      gjson-v.Unmarshal          170791291 ns/op    105925040 B/op        39409 allocs/op
        sonic.Unmarshal          113034762 ns/op    110672918 B/op        78923 allocs/op
      sonic-v.Unmarshal          192278071 ns/op    113225413 B/op        79048 allocs/op
        codec.Unmarshal          118080016 ns/op     48866521 B/op       157623 allocs/op
          jin.Unmarshal           70748354 ns/op      2652135 B/op        78811 allocs/op
        jason.Unmarshal         1008929964 ns/op   1083156752 B/op     16167235 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal          272663070 ns/op     68120556 B/op       591309 allocs/op
     easyjson.Unmarshal          348996436 ns/op      2531165 B/op        39415 allocs/op

   * simdjson █████████████████████████████▉ 7.47
 * jsonparser █████████████████████████████▏ 7.30
        * jin █████████████████████████▋ 6.43
   * fastjson █████████████████████▌ 5.40
      * gjson ████████████████████▉ 5.25
        goccy ████████████████████▌ 5.13
        sonic ████████████████  4.02
        codec ███████████████▍ 3.85
    * gjson-v ██████████▋ 2.66
      sonic-v █████████▍ 2.37
        json2 █████████▏ 2.29
     jsoniter ████████▌ 2.14
      segment ████████▏ 2.04
       ffjson ██████▋ 1.67
     easyjson █████▏ 1.30
         json ▓▓▓▓ 1.00
           oj ███▍ 0.87
      * jason █▊ 0.45
        djson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys generically
         json.Unmarshal          848625570 ns/op    743931088 B/op     14315102 allocs/op
        json2.Unmarshal          749784609 ns/op    693785500 B/op     11948513 allocs/op
           oj.Unmarshal          860803884 ns/op    809216560 B/op     19911129 allocs/op
     fastjson.Unmarshal          220484058 ns/op     66540835 B/op      2876749 allocs/op
     jsoniter.Unmarshal          732422110 ns/op    820368916 B/op     19201911 allocs/op
   jsonparser.Unmarshal          507701320 ns/op    105248856 B/op      5437894 allocs/op
        goccy.Unmarshal          686071000 ns/op    879399736 B/op     16801591 allocs/op
      segment.Unmarshal         2918416592 ns/op   6770751984 B/op     19978755 allocs/op
     simdjson.Unmarshal          172174840 ns/op    688366901 B/op      2758583 allocs/op
        gjson.Unmarshal          428770523 ns/op    196083440 B/op      2994781 allocs/op
      gjson-v.Unmarshal          526761869 ns/op    196083456 B/op      2994781 allocs/op
        sonic.Unmarshal          394771704 ns/op    890524736 B/op      2285577 allocs/op
      sonic-v.Unmarshal          482782363 ns/op    890388192 B/op      2285574 allocs/op
        codec.Unmarshal         1132613674 ns/op    312883248 B/op     18323501 allocs/op
          jin.Unmarshal          882158059 ns/op    563459208 B/op     18165707 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal          394995232 ns/op    770787336 B/op      7537379 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     simdjson ███████████████████▋ 4.93
     fastjson ███████████████▍ 3.85
        sonic ████████▌ 2.15
        djson ████████▌ 2.15
        gjson ███████▉ 1.98
      sonic-v ███████  1.76
   jsonparser ██████▋ 1.67
      gjson-v ██████▍ 1.61
        goccy ████▉ 1.24
     jsoniter ████▋ 1.16
        json2 ████▌ 1.13
         json ▓▓▓▓ 1.00
           oj ███▉ 0.99
          jin ███▊ 0.96
        codec ██▉ 0.75
      segment █▏ 0.29
        jason >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys into struct
         json.Unmarshal          706741396 ns/op    110917608 B/op      3113016 allocs/op
        json2.Unmarshal          397323630 ns/op    129775901 B/op      1801323 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          225360383 ns/op     97217966 B/op      2955391 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal          144421506 ns/op    106193920 B/op        39603 allocs/op
      segment.Unmarshal          567098452 ns/op     68879064 B/op      2797773 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal          143026339 ns/op    133237347 B/op       157637 allocs/op
      sonic-v.Unmarshal          225443365 ns/op    133641777 B/op       157653 allocs/op
        codec.Unmarshal          301658934 ns/op    102057804 B/op      2718955 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal          400354927 ns/op    217858725 B/op      4689268 allocs/op
     easyjson.Unmarshal          257488724 ns/op    170488064 B/op      3822311 allocs/op

        sonic ███████████████████▊ 4.94
        goccy ███████████████████▌ 4.89
     jsoniter ████████████▌ 3.14
      sonic-v ████████████▌ 3.13
     easyjson ██████████▉ 2.74
        codec █████████▎ 2.34
        json2 ███████  1.78
       ffjson ███████  1.77
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
         json.Unmarshal        21217815905 ns/op    116288992 B/op      2105576 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal         3441546606 ns/op     14463336 B/op       218594 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal         2142844107 ns/op     24696416 B/op       947400 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal         1450563109 ns/op   2147515736 B/op       217342 allocs/op
        gjson.Unmarshal         2992754190 ns/op   5740687488 B/op       210561 allocs/op
      gjson-v.Unmarshal         6811110929 ns/op   5740687488 B/op       210561 allocs/op
        sonic.Unmarshal          739672573 ns/op     15224100 B/op       421120 allocs/op
      sonic-v.Unmarshal         4322040800 ns/op     15226472 B/op       421121 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal         2731479262 ns/op     14210896 B/op       421119 allocs/op
        jason.Unmarshal        38870630635 ns/op  51080100064 B/op    711850242 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ██████████████████████████████████████████████████████████████████████████████████████████████████████████████████▋ 28.69
     simdjson ██████████████████████████████████████████████████████████▌ 14.63
   jsonparser ███████████████████████████████████████▌ 9.90
          jin ███████████████████████████████  7.77
        gjson ████████████████████████████▎ 7.09
     fastjson ████████████████████████▋ 6.17
      sonic-v ███████████████████▋ 4.91
      gjson-v ████████████▍ 3.12
         json ▓▓▓▓ 1.00
        jason ██▏ 0.55
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
         json.Unmarshal        20862153997 ns/op    116289104 B/op      2105577 allocs/op
        json2.Unmarshal         9162406339 ns/op        21080 B/op           72 allocs/op
           oj >>> not supported <<<
     fastjson.Unmarshal         3442735121 ns/op     14463336 B/op       218594 allocs/op
     jsoniter.Unmarshal         9317896870 ns/op   4650048304 B/op    210767863 allocs/op
   jsonparser.Unmarshal         2132705485 ns/op     24696400 B/op       947400 allocs/op
        goccy.Unmarshal         3175665704 ns/op   5742218272 B/op       211847 allocs/op
      segment.Unmarshal         8394652736 ns/op     13537216 B/op       210563 allocs/op
     simdjson.Unmarshal         1418401169 ns/op   2184291664 B/op       217133 allocs/op
        gjson.Unmarshal         2976675050 ns/op   5740688160 B/op       210568 allocs/op
      gjson-v.Unmarshal         6751341632 ns/op   5740687968 B/op       210566 allocs/op
        sonic.Unmarshal         4331436435 ns/op   5758939400 B/op       421562 allocs/op
      sonic-v.Unmarshal         7958631049 ns/op   5771179768 B/op       422148 allocs/op
        codec.Unmarshal         4149915149 ns/op    261152232 B/op       842234 allocs/op
          jin.Unmarshal         2776872933 ns/op     14210896 B/op       421119 allocs/op
        jason.Unmarshal        38747301211 ns/op  51079922256 B/op    711849634 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal        10432928073 ns/op    362372448 B/op      3158494 allocs/op
     easyjson.Unmarshal        17569367117 ns/op     13542776 B/op       210581 allocs/op

   * simdjson ██████████████████████████████████████████████████████████▊ 14.71
 * jsonparser ███████████████████████████████████████▏ 9.78
        * jin ██████████████████████████████  7.51
      * gjson ████████████████████████████  7.01
        goccy ██████████████████████████▎ 6.57
   * fastjson ████████████████████████▏ 6.06
        codec ████████████████████  5.03
        sonic ███████████████████▎ 4.82
    * gjson-v ████████████▎ 3.09
      sonic-v ██████████▍ 2.62
      segment █████████▉ 2.49
        json2 █████████  2.28
     jsoniter ████████▉ 2.24
       ffjson ███████▉ 2.00
     easyjson ████▋ 1.19
         json ▓▓▓▓ 1.00
      * jason ██▏ 0.54
           oj >>> not supported <<<
        djson >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys generically
         json.Unmarshal        35968701547 ns/op  37172843008 B/op    701322391 allocs/op
        json2.Unmarshal        30215736748 ns/op  35262286688 B/op    579264232 allocs/op
           oj.Unmarshal        34945565922 ns/op  40944082552 B/op    988125211 allocs/op
     fastjson.Unmarshal         9641348170 ns/op   3342950024 B/op    142344580 allocs/op
     jsoniter.Unmarshal        31032520089 ns/op  41731368816 B/op    960519577 allocs/op
   jsonparser.Unmarshal        24217703281 ns/op   5294193144 B/op    273513557 allocs/op
        goccy.Unmarshal        28451286344 ns/op  44770144360 B/op    833148854 allocs/op
      segment.Unmarshal       119561382671 ns/op 341651510856 B/op    995514453 allocs/op
     simdjson.Unmarshal         6767400896 ns/op  27003703296 B/op    137924604 allocs/op
        gjson.Unmarshal        20744953501 ns/op  10406633200 B/op    146758242 allocs/op
      gjson-v.Unmarshal        25426085705 ns/op  10406634880 B/op    146758261 allocs/op
        sonic.Unmarshal        18782534412 ns/op  44980382472 B/op    111177917 allocs/op
      sonic-v.Unmarshal        22469516483 ns/op  44978448952 B/op    111178021 allocs/op
        codec.Unmarshal        54913471757 ns/op  13510684224 B/op    905186237 allocs/op
          jin.Unmarshal        43656820470 ns/op  28265592920 B/op    910027464 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal        18946867413 ns/op  39791239008 B/op    375169203 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     simdjson █████████████████████▎ 5.31
     fastjson ██████████████▉ 3.73
        sonic ███████▋ 1.92
        djson ███████▌ 1.90
        gjson ██████▉ 1.73
      sonic-v ██████▍ 1.60
   jsonparser █████▉ 1.49
      gjson-v █████▋ 1.41
        goccy █████  1.26
        json2 ████▊ 1.19
     jsoniter ████▋ 1.16
           oj ████  1.03
         json ▓▓▓▓ 1.00
          jin ███▎ 0.82
        codec ██▌ 0.66
      segment █▏ 0.30
        jason >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys into struct
         json.Unmarshal        32584727620 ns/op   4443558704 B/op    137915233 allocs/op
        json2.Unmarshal        19600269768 ns/op   6895390848 B/op     94347011 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal        10325293698 ns/op   5003848008 B/op    150338413 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         6431995408 ns/op   5754816672 B/op       223081 allocs/op
      segment.Unmarshal        25861007502 ns/op   3491109688 B/op    141915747 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal         6410811773 ns/op   7136374504 B/op      4633093 allocs/op
      sonic-v.Unmarshal        10050084437 ns/op   7153086056 B/op      4633847 allocs/op
        codec.Unmarshal        13619618877 ns/op   3062219280 B/op    133914579 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal        17334408090 ns/op  10026592856 B/op    215402659 allocs/op
     easyjson.Unmarshal        12116979105 ns/op   8947234160 B/op    198556641 allocs/op

        sonic ████████████████████▎ 5.08
        goccy ████████████████████▎ 5.07
      sonic-v ████████████▉ 3.24
     jsoniter ████████████▌ 3.16
     easyjson ██████████▊ 2.69
        codec █████████▌ 2.39
       ffjson ███████▌ 1.88
        json2 ██████▋ 1.66
      segment █████  1.26
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
         json.Marshal                 1004 ns/op          688 B/op           17 allocs/op
        json2.Marshal                 1083 ns/op          424 B/op           10 allocs/op
           oj.Builder                  890 ns/op         1128 B/op           17 allocs/op
     fastjson.Marshal                  574 ns/op          679 B/op            6 allocs/op
     jsoniter.Marshal                  523 ns/op          432 B/op            7 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                  567 ns/op          416 B/op            2 allocs/op
      segment.Marshal                  575 ns/op          176 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson.Marshal                 1643 ns/op         3024 B/op           27 allocs/op
      gjson-v >>> not supported <<<
        sonic.Marshal                  705 ns/op         1803 B/op           14 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                  608 ns/op          988 B/op            0 allocs/op
          jin.Marshal                 1506 ns/op         2544 B/op           56 allocs/op
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Marshal                  148 ns/op          144 B/op            2 allocs/op
     easyjson.Marshal                  113 ns/op          288 B/op            3 allocs/op

   * easyjson ███████████████████████████████████▌ 8.88
     * ffjson ███████████████████████████▏ 6.78
     jsoniter ███████▋ 1.92
        goccy ███████  1.77
     fastjson ██████▉ 1.75
      segment ██████▉ 1.75
        codec ██████▌ 1.65
        sonic █████▋ 1.42
           oj ████▌ 1.13
         json ▓▓▓▓ 1.00
        json2 ███▋ 0.93
          jin ██▋ 0.67
        gjson ██▍ 0.61
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
      gjson-v >>> not supported <<<
   jsonparser >>> not supported <<<

Validate []byte
         json.Validate                6449 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj.Validate                2221 ns/op            0 B/op            0 allocs/op
     fastjson.Validate                2622 ns/op            0 B/op            0 allocs/op
     jsoniter.Validate                4531 ns/op         2184 B/op          100 allocs/op
   jsonparser >>> not supported <<<
        goccy.Validate               16892 ns/op        27588 B/op          463 allocs/op
      segment.Validate                4387 ns/op            0 B/op            0 allocs/op
     simdjson >>> not supported <<<
        gjson.Validate                2344 ns/op            0 B/op            0 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                2086 ns/op            0 B/op            0 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ████████████▎ 3.09
           oj ███████████▌ 2.90
        gjson ███████████  2.75
     fastjson █████████▊ 2.46
      segment █████▉ 1.47
     jsoniter █████▋ 1.42
         json ▓▓▓▓ 1.00
        goccy █▌ 0.38
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
         json.Validate                6593 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Validate                3111 ns/op         4096 B/op            1 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Validate                2593 ns/op         4096 B/op            1 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                2415 ns/op         4098 B/op            1 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ██████████▉ 2.73
        gjson ██████████▏ 2.54
     fastjson ████████▍ 2.12
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
Unmarshal single record (2kb), read few keys generically
         json.Unmarshal              27593 ns/op        12488 B/op           21 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal               8280 ns/op        12034 B/op            7 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal               8140 ns/op        12208 B/op           17 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal               9418 ns/op        16368 B/op            9 allocs/op
      gjson-v.Unmarshal              12878 ns/op        16368 B/op            9 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin.Unmarshal               7975 ns/op        12344 B/op           10 allocs/op
        jason.Unmarshal              51373 ns/op        55997 B/op          495 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

          jin █████████████▊ 3.46
   jsonparser █████████████▌ 3.39
     fastjson █████████████▎ 3.33
        gjson ███████████▋ 2.93
      gjson-v ████████▌ 2.14
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
         json.Unmarshal              27624 ns/op        12488 B/op           21 allocs/op
        json2.Unmarshal               8930 ns/op          128 B/op            6 allocs/op
           oj >>> not supported <<<
     fastjson.Unmarshal               8378 ns/op        12080 B/op           10 allocs/op
     jsoniter.Unmarshal              12294 ns/op        14063 B/op           84 allocs/op
   jsonparser.Unmarshal               8340 ns/op        12208 B/op           17 allocs/op
        goccy.Unmarshal               9019 ns/op        16171 B/op            7 allocs/op
      segment.Unmarshal               9609 ns/op        12064 B/op           12 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal               9598 ns/op        16400 B/op           10 allocs/op
      gjson-v.Unmarshal              13110 ns/op        16400 B/op           10 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal              10584 ns/op        13240 B/op           15 allocs/op
          jin.Unmarshal               8226 ns/op        12344 B/op           10 allocs/op
        jason.Unmarshal              52616 ns/op        55997 B/op          495 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal              22038 ns/op        14206 B/op           36 allocs/op
     easyjson.Unmarshal              21021 ns/op        12326 B/op           16 allocs/op

        * jin █████████████▍ 3.36
 * jsonparser █████████████▏ 3.31
   * fastjson █████████████▏ 3.30
        json2 ████████████▎ 3.09
        goccy ████████████▎ 3.06
      * gjson ███████████▌ 2.88
      segment ███████████▍ 2.87
        codec ██████████▍ 2.61
     jsoniter ████████▉ 2.25
    * gjson-v ████████▍ 2.11
     easyjson █████▎ 1.31
       ffjson █████  1.25
         json ▓▓▓▓ 1.00
      * jason ██  0.53
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
        djson >>> not supported <<<
        sonic >>> not supported <<<
           oj >>> not supported <<<

Unmarshal single record (2kb), read all keys generically
         json.Unmarshal              36101 ns/op        29778 B/op          340 allocs/op
        json2.Unmarshal              23528 ns/op        15546 B/op          295 allocs/op
           oj.Unmarshal              23418 ns/op        19299 B/op          465 allocs/op
     fastjson.Unmarshal              13344 ns/op        13608 B/op           73 allocs/op
     jsoniter.Unmarshal              27732 ns/op        31696 B/op          457 allocs/op
   jsonparser.Unmarshal              20386 ns/op        14528 B/op          135 allocs/op
        goccy.Unmarshal              25625 ns/op        34556 B/op          398 allocs/op
      segment.Unmarshal              64595 ns/op       173242 B/op          475 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal              18818 ns/op        18336 B/op           76 allocs/op
      gjson-v.Unmarshal              22306 ns/op        18336 B/op           76 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal              37099 ns/op        19456 B/op          435 allocs/op
          jin.Unmarshal              28490 ns/op        25344 B/op          435 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal              20665 ns/op        32169 B/op          183 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     fastjson ██████████▊ 2.71
        gjson ███████▋ 1.92
   jsonparser ███████  1.77
        djson ██████▉ 1.75
      gjson-v ██████▍ 1.62
           oj ██████▏ 1.54
        json2 ██████▏ 1.53
        goccy █████▋ 1.41
     jsoniter █████▏ 1.30
          jin █████  1.27
         json ▓▓▓▓ 1.00
        codec ███▉ 0.97
      segment ██▏ 0.56
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
        jason >>> not supported <<<
        sonic >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal single record (2kb), read all keys into struct
         json.Unmarshal              32981 ns/op        14608 B/op           80 allocs/op
        json2.Unmarshal              14301 ns/op         3106 B/op           35 allocs/op
           oj.Unmarshal              26809 ns/op         9340 B/op          453 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal              13259 ns/op        14392 B/op           77 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal              11181 ns/op        16289 B/op            7 allocs/op
      segment.Unmarshal              15313 ns/op        13680 B/op           73 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal              15513 ns/op        14536 B/op           72 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal              18798 ns/op        16820 B/op          113 allocs/op
     easyjson.Unmarshal              16317 ns/op        15888 B/op           94 allocs/op

        goccy ███████████▊ 2.95
     jsoniter █████████▉ 2.49
        json2 █████████▏ 2.31
      segment ████████▌ 2.15
        codec ████████▌ 2.13
     easyjson ████████  2.02
       ffjson ███████  1.75
           oj ████▉ 1.23
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
         json.Unmarshal          664894208 ns/op     21756032 B/op       394052 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal          129191500 ns/op      2537837 B/op        39514 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal           89898131 ns/op      4614337 B/op       177229 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal          124166166 ns/op    105925557 B/op        39414 allocs/op
      gjson-v.Unmarshal          239335091 ns/op    105925216 B/op        39411 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin.Unmarshal           98072979 ns/op      2652130 B/op        78811 allocs/op
        jason.Unmarshal         1343083834 ns/op   1082925056 B/op     16167184 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

   jsonparser █████████████████████████████▌ 7.40
          jin ███████████████████████████  6.78
        gjson █████████████████████▍ 5.35
     fastjson ████████████████████▌ 5.15
      gjson-v ███████████  2.78
         json ▓▓▓▓ 1.00
        jason █▉ 0.50
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
         json.Unmarshal          664172041 ns/op     21755760 B/op       394052 allocs/op
        json2.Unmarshal          255247364 ns/op        16410 B/op           67 allocs/op
           oj.Unmarshal          676079104 ns/op    246953896 B/op     15880757 allocs/op
     fastjson.Unmarshal          127984198 ns/op      2537825 B/op        39513 allocs/op
     jsoniter.Unmarshal          286331823 ns/op     89443200 B/op      3979982 allocs/op
   jsonparser.Unmarshal           88943012 ns/op      4614339 B/op       177229 allocs/op
        goccy.Unmarshal          127639843 ns/op    106045457 B/op        39744 allocs/op
      segment.Unmarshal          185692229 ns/op      2526037 B/op        39406 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal          124225893 ns/op    105925706 B/op        39416 allocs/op
      gjson-v.Unmarshal          237990391 ns/op    105925331 B/op        39412 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal          173319819 ns/op     48866474 B/op       157623 allocs/op
          jin.Unmarshal           98280399 ns/op      2652145 B/op        78811 allocs/op
        jason.Unmarshal         1346599667 ns/op   1082885904 B/op     16167057 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal          392091208 ns/op     67900330 B/op       591344 allocs/op
     easyjson.Unmarshal          535095187 ns/op      2527752 B/op        39416 allocs/op

 * jsonparser █████████████████████████████▊ 7.47
        * jin ███████████████████████████  6.76
      * gjson █████████████████████▍ 5.35
        goccy ████████████████████▊ 5.20
   * fastjson ████████████████████▊ 5.19
        codec ███████████████▎ 3.83
      segment ██████████████▎ 3.58
    * gjson-v ███████████▏ 2.79
        json2 ██████████▍ 2.60
     jsoniter █████████▎ 2.32
       ffjson ██████▊ 1.69
     easyjson ████▉ 1.24
         json ▓▓▓▓ 1.00
           oj ███▉ 0.98
      * jason █▉ 0.49
     simdjson >>> not supported <<<
        djson >>> not supported <<<
      sonic-v >>> not supported <<<
        sonic >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys generically
         json.Unmarshal         1151385209 ns/op    743878576 B/op     14314939 allocs/op
        json2.Unmarshal          946216771 ns/op    693414512 B/op     11939686 allocs/op
           oj.Unmarshal         1169068583 ns/op    809287432 B/op     19911234 allocs/op
     fastjson.Unmarshal          366834472 ns/op     66551472 B/op      2876854 allocs/op
     jsoniter.Unmarshal          988414270 ns/op    819532340 B/op     19202085 allocs/op
   jsonparser.Unmarshal          643740292 ns/op    105469680 B/op      5437892 allocs/op
        goccy.Unmarshal          893365020 ns/op    877667784 B/op     16802529 allocs/op
      segment.Unmarshal         3024869083 ns/op   6770606720 B/op     19978959 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal          627679292 ns/op    196084440 B/op      2994787 allocs/op
      gjson-v.Unmarshal          744356500 ns/op    196084424 B/op      2994787 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal         1400056875 ns/op    313137696 B/op     18323507 allocs/op
          jin.Unmarshal         1114297417 ns/op    563338952 B/op     18165724 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal          646424062 ns/op    770805720 B/op      7537464 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     fastjson ████████████▌ 3.14
        gjson ███████▎ 1.83
   jsonparser ███████▏ 1.79
        djson ███████  1.78
      gjson-v ██████▏ 1.55
        goccy █████▏ 1.29
        json2 ████▊ 1.22
     jsoniter ████▋ 1.16
          jin ████▏ 1.03
         json ▓▓▓▓ 1.00
           oj ███▉ 0.98
        codec ███▎ 0.82
      segment █▌ 0.38
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
        jason >>> not supported <<<
        sonic >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys into struct
         json.Unmarshal          936409687 ns/op    111143384 B/op      3113021 allocs/op
        json2.Unmarshal          531664208 ns/op    129790196 B/op      1792489 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          302310635 ns/op     97329280 B/op      2955470 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal          210571950 ns/op    107169392 B/op        43110 allocs/op
      segment.Unmarshal          409084263 ns/op     69133181 B/op      2797769 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal          396267569 ns/op    102290640 B/op      2718959 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal          555748271 ns/op    218501800 B/op      4690357 allocs/op
     easyjson.Unmarshal          425918889 ns/op    170812645 B/op      3822719 allocs/op

        goccy █████████████████▊ 4.45
     jsoniter ████████████▍ 3.10
        codec █████████▍ 2.36
      segment █████████▏ 2.29
     easyjson ████████▊ 2.20
        json2 ███████  1.76
       ffjson ██████▋ 1.68
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
         json.Unmarshal        31301040958 ns/op    116289696 B/op      2105580 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal         5292408125 ns/op     14548256 B/op       218830 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal         2914637334 ns/op     24696800 B/op       947485 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal         4801432291 ns/op   5740740480 B/op       211113 allocs/op
      gjson-v.Unmarshal        10462319709 ns/op   5740724640 B/op       210948 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin.Unmarshal         3653586417 ns/op     14210960 B/op       421119 allocs/op
        jason.Unmarshal        63338545375 ns/op  51093030944 B/op    711854828 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

   jsonparser ██████████████████████████████████████████▉ 10.74
          jin ██████████████████████████████████▎ 8.57
        gjson ██████████████████████████  6.52
     fastjson ███████████████████████▋ 5.91
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
         json.Unmarshal        31317242417 ns/op    116289552 B/op      2105577 allocs/op
        json2.Unmarshal        12152636500 ns/op        17624 B/op           72 allocs/op
           oj >>> not supported <<<
     fastjson.Unmarshal         5469278000 ns/op     14548256 B/op       218830 allocs/op
     jsoniter.Unmarshal        14016159167 ns/op   4644582472 B/op    210772098 allocs/op
   jsonparser.Unmarshal         3011843625 ns/op     24696720 B/op       947485 allocs/op
        goccy.Unmarshal         5389332875 ns/op   5748884664 B/op       231974 allocs/op
      segment.Unmarshal         7908669250 ns/op     13537216 B/op       210563 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal         5010077666 ns/op   5740739520 B/op       211103 allocs/op
      gjson-v.Unmarshal        10811854167 ns/op   5740731456 B/op       211019 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal         6703905625 ns/op    261153192 B/op       842244 allocs/op
          jin.Unmarshal         3705682666 ns/op     14211104 B/op       421120 allocs/op
        jason.Unmarshal        64744879750 ns/op  51093577088 B/op    711856862 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal        16835513291 ns/op    362928112 B/op      3160071 allocs/op
     easyjson.Unmarshal        27184591250 ns/op     13548512 B/op       210629 allocs/op

 * jsonparser █████████████████████████████████████████▌ 10.40
        * jin █████████████████████████████████▊ 8.45
      * gjson █████████████████████████  6.25
        goccy ███████████████████████▏ 5.81
   * fastjson ██████████████████████▉ 5.73
        codec ██████████████████▋ 4.67
      segment ███████████████▊ 3.96
    * gjson-v ███████████▌ 2.90
        json2 ██████████▎ 2.58
     jsoniter ████████▉ 2.23
       ffjson ███████▍ 1.86
     easyjson ████▌ 1.15
         json ▓▓▓▓ 1.00
      * jason █▉ 0.48
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
        djson >>> not supported <<<
        sonic >>> not supported <<<
           oj >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys generically
         json.Unmarshal        59311769750 ns/op  37196911920 B/op    701327218 allocs/op
        json2.Unmarshal        47479322666 ns/op  35261685712 B/op    579151551 allocs/op
           oj.Unmarshal        63907148584 ns/op  40969089560 B/op    988133051 allocs/op
     fastjson.Unmarshal        18188804041 ns/op   3343061760 B/op    142345014 allocs/op
     jsoniter.Unmarshal        52841262000 ns/op  41776010648 B/op    960570620 allocs/op
   jsonparser.Unmarshal        33000124292 ns/op   5277731400 B/op    273513719 allocs/op
        goccy.Unmarshal        49823845416 ns/op  44836555536 B/op    833466654 allocs/op
      segment.Unmarshal       149980700625 ns/op 341675772024 B/op    995538666 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal        31562406958 ns/op  10406695968 B/op    146758665 allocs/op
      gjson-v.Unmarshal        37402392459 ns/op  10406686592 B/op    146758590 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal        69678185833 ns/op  13524891216 B/op    905187755 allocs/op
          jin.Unmarshal        57280385250 ns/op  28244873848 B/op    910029136 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal        34548678125 ns/op  39791614160 B/op    375172653 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     fastjson █████████████  3.26
        gjson ███████▌ 1.88
   jsonparser ███████▏ 1.80
        djson ██████▊ 1.72
      gjson-v ██████▎ 1.59
        json2 ████▉ 1.25
        goccy ████▊ 1.19
     jsoniter ████▍ 1.12
          jin ████▏ 1.04
         json ▓▓▓▓ 1.00
           oj ███▋ 0.93
        codec ███▍ 0.85
      segment █▌ 0.40
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
        jason >>> not supported <<<
        sonic >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys into struct
         json.Unmarshal        45092887875 ns/op   4426671968 B/op    137915351 allocs/op
        json2.Unmarshal        27475470625 ns/op   6898842040 B/op     94203776 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal        14504947750 ns/op   4989314664 B/op    150343123 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         9322221333 ns/op   5821877552 B/op       439323 allocs/op
      segment.Unmarshal        19876647041 ns/op   3474107224 B/op    141915831 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal        18446697125 ns/op   3045082832 B/op    133914672 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Unmarshal        24959078250 ns/op  10030816656 B/op    215458688 allocs/op
     easyjson.Unmarshal        21510108417 ns/op   8935612424 B/op    198580351 allocs/op

        goccy ███████████████████▎ 4.84
     jsoniter ████████████▍ 3.11
        codec █████████▊ 2.44
      segment █████████  2.27
     easyjson ████████▍ 2.10
       ffjson ███████▏ 1.81
        json2 ██████▌ 1.64
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
         json.Marshal                 1348 ns/op          688 B/op           17 allocs/op
        json2.Marshal                 1313 ns/op          424 B/op           10 allocs/op
           oj.Builder                 1485 ns/op         1129 B/op           17 allocs/op
     fastjson.Marshal                  880 ns/op          679 B/op            6 allocs/op
     jsoniter.Marshal                  721 ns/op          432 B/op            7 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                  845 ns/op          416 B/op            2 allocs/op
      segment.Marshal                  769 ns/op          176 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson.Marshal                 2721 ns/op         3026 B/op           27 allocs/op
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Marshal                  759 ns/op         1003 B/op            0 allocs/op
          jin.Marshal                 2305 ns/op         2546 B/op           56 allocs/op
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson.Marshal                  212 ns/op          144 B/op            2 allocs/op
     easyjson.Marshal                  189 ns/op          288 B/op            3 allocs/op

     easyjson ████████████████████████████▌ 7.13
       ffjson █████████████████████████▍ 6.36
     jsoniter ███████▍ 1.87
        codec ███████  1.78
      segment ███████  1.75
        goccy ██████▍ 1.60
     fastjson ██████▏ 1.53
        json2 ████  1.03
         json ▓▓▓▓ 1.00
           oj ███▋ 0.91
          jin ██▎ 0.58
        gjson █▉ 0.50
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
        sonic >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
      gjson-v >>> not supported <<<
   jsonparser >>> not supported <<<

Validate []byte
         json.Validate               10517 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj.Validate                3868 ns/op            0 B/op            0 allocs/op
     fastjson.Validate                4949 ns/op            0 B/op            0 allocs/op
     jsoniter.Validate                6944 ns/op         2185 B/op          100 allocs/op
   jsonparser >>> not supported <<<
        goccy.Validate               28429 ns/op        27616 B/op          463 allocs/op
      segment.Validate                3625 ns/op            0 B/op            0 allocs/op
     simdjson >>> not supported <<<
        gjson.Validate                3584 ns/op            0 B/op            0 allocs/op
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        gjson ███████████▋ 2.93
      segment ███████████▌ 2.90
           oj ██████████▉ 2.72
     fastjson ████████▌ 2.13
     jsoniter ██████  1.51
         json ▓▓▓▓ 1.00
        goccy █▍ 0.37
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
         json.Validate               10559 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Validate                5617 ns/op         4096 B/op            1 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Validate                4268 ns/op         4096 B/op            1 allocs/op
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        gjson █████████▉ 2.47
     fastjson ███████▌ 1.88
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
