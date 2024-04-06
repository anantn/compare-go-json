# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Validate []byte
         json.Validate                6543 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj.Validate                2205 ns/op            0 B/op            0 allocs/op
     fastjson.Validate                2645 ns/op            0 B/op            0 allocs/op
     jsoniter.Validate                4654 ns/op         2188 B/op          100 allocs/op
   jsonparser >>> not supported <<<
        goccy.Validate               19846 ns/op        27653 B/op          463 allocs/op
      segment.Validate                4313 ns/op            0 B/op            0 allocs/op
     simdjson >>> not supported <<<
        gjson.Validate                2430 ns/op            0 B/op            0 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                1988 ns/op            0 B/op            0 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<

        sonic █████████████▏ 3.29
           oj ███████████▊ 2.97
        gjson ██████████▊ 2.69
     fastjson █████████▉ 2.47
      segment ██████  1.52
     jsoniter █████▌ 1.41
         json ▓▓▓▓ 1.00
        goccy █▎ 0.33
   jsonparser >>> not supported <<<
     simdjson >>> not supported <<<
      gjson-v >>> not supported <<<
        json2 >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<

Validate string
         json.Validate                6323 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Validate                3376 ns/op         4096 B/op            1 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Validate                3032 ns/op         4096 B/op            1 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                2640 ns/op         4139 B/op            1 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<

        sonic █████████▌ 2.40
        gjson ████████▎ 2.09
     fastjson ███████▍ 1.87
         json ▓▓▓▓ 1.00
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
           oj >>> not supported <<<
      gjson-v >>> not supported <<<
        json2 >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<

Unmarshal single record (2kb), read few keys generically
         json.Unmarshal              18312 ns/op        12488 B/op           21 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal               5298 ns/op        12034 B/op            7 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal               5375 ns/op        12147 B/op           13 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal               6343 ns/op        16368 B/op            9 allocs/op
      gjson-v.Unmarshal               9123 ns/op        16368 B/op            9 allocs/op
        sonic.Unmarshal               3539 ns/op         4966 B/op           12 allocs/op
      sonic-v.Unmarshal               5624 ns/op         4928 B/op           12 allocs/op
        codec >>> not supported <<<

        sonic ████████████████████▋ 5.17
     fastjson █████████████▊ 3.46
   jsonparser █████████████▋ 3.41
      sonic-v █████████████  3.26
        gjson ███████████▌ 2.89
      gjson-v ████████  2.01
         json ▓▓▓▓ 1.00
     jsoniter >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
           oj >>> not supported <<<
        json2 >>> not supported <<<
        codec >>> not supported <<<

Unmarshal single record (2kb), read few keys into struct
         json.Unmarshal              18300 ns/op        12488 B/op           21 allocs/op
        json2.Unmarshal               6215 ns/op          128 B/op            6 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal               8616 ns/op        14087 B/op           84 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               6123 ns/op        16299 B/op            7 allocs/op
      segment.Unmarshal              10566 ns/op        12064 B/op           12 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               5647 ns/op         9061 B/op           11 allocs/op
      sonic-v.Unmarshal               7850 ns/op         9330 B/op           11 allocs/op
        codec.Unmarshal               7697 ns/op        13240 B/op           15 allocs/op

        sonic ████████████▉ 3.24
        goccy ███████████▉ 2.99
        json2 ███████████▊ 2.94
        codec █████████▌ 2.38
      sonic-v █████████▎ 2.33
     jsoniter ████████▍ 2.12
      segment ██████▉ 1.73
         json ▓▓▓▓ 1.00
     fastjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
           oj >>> not supported <<<

Unmarshal single record (2kb), read all keys generically
         json.Unmarshal              26424 ns/op        29776 B/op          340 allocs/op
        json2.Unmarshal              19571 ns/op        15572 B/op          295 allocs/op
           oj.Unmarshal              18371 ns/op        19300 B/op          465 allocs/op
     fastjson.Unmarshal               8555 ns/op        13608 B/op           73 allocs/op
     jsoniter.Unmarshal              20026 ns/op        31753 B/op          457 allocs/op
   jsonparser.Unmarshal              14545 ns/op        14528 B/op          135 allocs/op
        goccy.Unmarshal              19273 ns/op        34668 B/op          398 allocs/op
      segment.Unmarshal              64973 ns/op       173241 B/op          475 allocs/op
     simdjson.Unmarshal              22892 ns/op       132224 B/op           82 allocs/op
        gjson.Unmarshal              14497 ns/op        18336 B/op           76 allocs/op
      gjson-v.Unmarshal              17088 ns/op        18336 B/op           76 allocs/op
        sonic.Unmarshal              14026 ns/op        27722 B/op           59 allocs/op
      sonic-v.Unmarshal              16207 ns/op        27727 B/op           59 allocs/op
        codec.Unmarshal              29725 ns/op        19456 B/op          435 allocs/op

     fastjson ████████████▎ 3.09
        sonic ███████▌ 1.88
        gjson ███████▎ 1.82
   jsonparser ███████▎ 1.82
      sonic-v ██████▌ 1.63
      gjson-v ██████▏ 1.55
           oj █████▊ 1.44
        goccy █████▍ 1.37
        json2 █████▍ 1.35
     jsoniter █████▎ 1.32
     simdjson ████▌ 1.15
         json ▓▓▓▓ 1.00
        codec ███▌ 0.89
      segment █▋ 0.41

Unmarshal single record (2kb), read all keys into struct
         json.Unmarshal              23195 ns/op        14608 B/op           80 allocs/op
        json2.Unmarshal              10072 ns/op         3111 B/op           35 allocs/op
           oj.Unmarshal              21523 ns/op         9340 B/op          453 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal               8877 ns/op        14413 B/op           77 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               6939 ns/op        16595 B/op            7 allocs/op
      segment.Unmarshal              17149 ns/op        13680 B/op           73 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               6239 ns/op         9579 B/op            9 allocs/op
      sonic-v.Unmarshal               8363 ns/op         9850 B/op            9 allocs/op
        codec.Unmarshal              10218 ns/op        14536 B/op           72 allocs/op

        sonic ██████████████▊ 3.72
        goccy █████████████▎ 3.34
      sonic-v ███████████  2.77
     jsoniter ██████████▍ 2.61
        json2 █████████▏ 2.30
        codec █████████  2.27
      segment █████▍ 1.35
           oj ████▎ 1.08
         json ▓▓▓▓ 1.00
     fastjson >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys generically
         json.Unmarshal          449103938 ns/op     21755877 B/op       394053 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal           79866503 ns/op      2533397 B/op        39470 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal           60665430 ns/op      4614608 B/op       177366 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal           58612256 ns/op    616507998 B/op        39638 allocs/op
        gjson.Unmarshal           90530344 ns/op    105925432 B/op        39413 allocs/op
      gjson-v.Unmarshal          179496065 ns/op    105925024 B/op        39409 allocs/op
        sonic.Unmarshal           33150877 ns/op      2888247 B/op        78814 allocs/op
      sonic-v.Unmarshal          107114343 ns/op      2874844 B/op        78814 allocs/op
        codec >>> not supported <<<

        sonic ██████████████████████████████████████████████████████▏ 13.55
     simdjson ██████████████████████████████▋ 7.66
   jsonparser █████████████████████████████▌ 7.40
     fastjson ██████████████████████▍ 5.62
        gjson ███████████████████▊ 4.96
      sonic-v ████████████████▊ 4.19
      gjson-v ██████████  2.50
         json ▓▓▓▓ 1.00
     jsoniter >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
           oj >>> not supported <<<
        json2 >>> not supported <<<
        codec >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys into struct
         json.Unmarshal          448665406 ns/op     21755744 B/op       394051 allocs/op
        json2.Unmarshal          193996560 ns/op        16804 B/op           64 allocs/op
           oj.Unmarshal          512668503 ns/op    246889584 B/op     15880758 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          209958229 ns/op     89155668 B/op      3979976 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal           92138756 ns/op    106273400 B/op        39728 allocs/op
      segment.Unmarshal          212521119 ns/op      2526041 B/op        39406 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal          112309089 ns/op    110180698 B/op        78920 allocs/op
      sonic-v.Unmarshal          184754258 ns/op    111464132 B/op        79024 allocs/op
        codec.Unmarshal          125313249 ns/op     48866468 B/op       157622 allocs/op

        goccy ███████████████████▍ 4.87
        sonic ███████████████▉ 3.99
        codec ██████████████▎ 3.58
      sonic-v █████████▋ 2.43
        json2 █████████▎ 2.31
     jsoniter ████████▌ 2.14
      segment ████████▍ 2.11
         json ▓▓▓▓ 1.00
           oj ███▌ 0.88
     fastjson >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys generically
         json.Unmarshal          851884989 ns/op    743961696 B/op     14315204 allocs/op
        json2.Unmarshal          740929570 ns/op    693843044 B/op     11948565 allocs/op
           oj.Unmarshal          847104598 ns/op    809242168 B/op     19911215 allocs/op
     fastjson.Unmarshal          212182792 ns/op     66540736 B/op      2876748 allocs/op
     jsoniter.Unmarshal          716033299 ns/op    820425120 B/op     19201955 allocs/op
   jsonparser.Unmarshal          499016452 ns/op    105248738 B/op      5437892 allocs/op
        goccy.Unmarshal          671647300 ns/op    879498980 B/op     16801563 allocs/op
      segment.Unmarshal         2756717977 ns/op   6770756288 B/op     19978800 allocs/op
     simdjson.Unmarshal          169658748 ns/op    688367742 B/op      2758584 allocs/op
        gjson.Unmarshal          450645516 ns/op    196083461 B/op      2994781 allocs/op
      gjson-v.Unmarshal          547742327 ns/op    196083432 B/op      2994781 allocs/op
        sonic.Unmarshal          380532327 ns/op    890333429 B/op      2285571 allocs/op
      sonic-v.Unmarshal          465294827 ns/op    890335085 B/op      2285572 allocs/op
        codec.Unmarshal         1085979981 ns/op    312883280 B/op     18323501 allocs/op

     simdjson ████████████████████  5.02
     fastjson ████████████████  4.01
        sonic ████████▉ 2.24
        gjson ███████▌ 1.89
      sonic-v ███████▎ 1.83
   jsonparser ██████▊ 1.71
      gjson-v ██████▏ 1.56
        goccy █████  1.27
     jsoniter ████▊ 1.19
        json2 ████▌ 1.15
           oj ████  1.01
         json ▓▓▓▓ 1.00
        codec ███▏ 0.78
      segment █▏ 0.31

Unmarshal many records (2kb each) from small file (100MB), read all keys into struct
         json.Unmarshal          651950448 ns/op    110917592 B/op      3113016 allocs/op
        json2.Unmarshal          380806820 ns/op    129775810 B/op      1801322 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          216124071 ns/op     97217828 B/op      2955390 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal          132777424 ns/op    106193571 B/op        39607 allocs/op
      segment.Unmarshal          553206541 ns/op     68879048 B/op      2797773 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal          144347826 ns/op    133222069 B/op       157636 allocs/op
      sonic-v.Unmarshal          224150407 ns/op    133490958 B/op       157651 allocs/op
        codec.Unmarshal          270345501 ns/op    102057768 B/op      2718954 allocs/op

        goccy ███████████████████▋ 4.91
        sonic ██████████████████  4.52
     jsoniter ████████████  3.02
      sonic-v ███████████▋ 2.91
        codec █████████▋ 2.41
        json2 ██████▊ 1.71
      segment ████▋ 1.18
         json ▓▓▓▓ 1.00
     fastjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
           oj >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys generically
         json.Unmarshal        20380892328 ns/op    116289008 B/op      2105576 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal         3244260320 ns/op     14463336 B/op       218594 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal         2062854099 ns/op     24696400 B/op       947400 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal         1393992740 ns/op   2315407432 B/op       217378 allocs/op
        gjson.Unmarshal         3142537931 ns/op   5740688448 B/op       210571 allocs/op
      gjson-v.Unmarshal         7097862972 ns/op   5740688736 B/op       210574 allocs/op
        sonic.Unmarshal          759498419 ns/op     15285444 B/op       421121 allocs/op
      sonic-v.Unmarshal         4411219799 ns/op     15308584 B/op       421125 allocs/op
        codec >>> not supported <<<

        sonic ███████████████████████████████████████████████████████████████████████████████████████████████████████████▎ 26.83
     simdjson ██████████████████████████████████████████████████████████▍ 14.62
   jsonparser ███████████████████████████████████████▌ 9.88
        gjson █████████████████████████▉ 6.49
     fastjson █████████████████████████▏ 6.28
      sonic-v ██████████████████▍ 4.62
      gjson-v ███████████▍ 2.87
         json ▓▓▓▓ 1.00
     jsoniter >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
           oj >>> not supported <<<
        json2 >>> not supported <<<
        codec >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys into struct
         json.Unmarshal        20567898765 ns/op    116289040 B/op      2105576 allocs/op
        json2.Unmarshal         8980506104 ns/op        21080 B/op           72 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal         9076040337 ns/op   4650048736 B/op    210767867 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         3302601715 ns/op   5742212408 B/op       211652 allocs/op
      segment.Unmarshal         8317731549 ns/op     13537216 B/op       210563 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal         4289252745 ns/op   5757342216 B/op       421549 allocs/op
      sonic-v.Unmarshal         7640563540 ns/op   5768272144 B/op       422112 allocs/op
        codec.Unmarshal         4396254464 ns/op    261152232 B/op       842234 allocs/op

        goccy ████████████████████████▉ 6.23
        sonic ███████████████████▏ 4.80
        codec ██████████████████▋ 4.68
      sonic-v ██████████▊ 2.69
      segment █████████▉ 2.47
        json2 █████████▏ 2.29
     jsoniter █████████  2.27
         json ▓▓▓▓ 1.00
     fastjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
           oj >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys generically
         json.Unmarshal        34962855059 ns/op  37173008704 B/op    701322962 allocs/op
        json2.Unmarshal        29650099638 ns/op  35262080352 B/op    579263479 allocs/op
           oj.Unmarshal        34153935033 ns/op  40944559400 B/op    988126864 allocs/op
     fastjson.Unmarshal         9278227196 ns/op   3342950024 B/op    142344580 allocs/op
     jsoniter.Unmarshal        29890821458 ns/op  41731430304 B/op    960519724 allocs/op
   jsonparser.Unmarshal        24100998320 ns/op   5294193400 B/op    273513561 allocs/op
        goccy.Unmarshal        27000850295 ns/op  44769959200 B/op    833147616 allocs/op
      segment.Unmarshal       114554113358 ns/op 341651496920 B/op    995514401 allocs/op
     simdjson.Unmarshal         6841674868 ns/op  27340493480 B/op    137924749 allocs/op
        gjson.Unmarshal        23633304238 ns/op  10406633376 B/op    146758246 allocs/op
      gjson-v.Unmarshal        27390030095 ns/op  10406635040 B/op    146758263 allocs/op
        sonic.Unmarshal        18436672408 ns/op  44980265136 B/op    111177924 allocs/op
      sonic-v.Unmarshal        22206080614 ns/op  44976860352 B/op    111177926 allocs/op
        codec.Unmarshal        53293578918 ns/op  13510683824 B/op    905186239 allocs/op

     simdjson ████████████████████▍ 5.11
     fastjson ███████████████  3.77
        sonic ███████▌ 1.90
      sonic-v ██████▎ 1.57
        gjson █████▉ 1.48
   jsonparser █████▊ 1.45
        goccy █████▏ 1.29
      gjson-v █████  1.28
        json2 ████▋ 1.18
     jsoniter ████▋ 1.17
           oj ████  1.02
         json ▓▓▓▓ 1.00
        codec ██▌ 0.66
      segment █▏ 0.31

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys into struct
         json.Unmarshal        30521041775 ns/op   4443558896 B/op    137915237 allocs/op
        json2.Unmarshal        19180908958 ns/op   6895391376 B/op     94347017 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal        10104027688 ns/op   5003847208 B/op    150338411 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         5876180206 ns/op   5755023024 B/op       224266 allocs/op
      segment.Unmarshal        25068226928 ns/op   3491110024 B/op    141915750 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal         6323706252 ns/op   7137118600 B/op      4633085 allocs/op
      sonic-v.Unmarshal        10049675296 ns/op   7153684552 B/op      4633850 allocs/op
        codec.Unmarshal        12642388404 ns/op   3062219376 B/op    133914580 allocs/op

        goccy ████████████████████▊ 5.19
        sonic ███████████████████▎ 4.83
      sonic-v ████████████▏ 3.04
     jsoniter ████████████  3.02
        codec █████████▋ 2.41
        json2 ██████▎ 1.59
      segment ████▊ 1.22
         json ▓▓▓▓ 1.00
     fastjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
           oj >>> not supported <<<

Marshal custom data through an object builder
         json.Marshal                  930 ns/op          688 B/op           17 allocs/op
        json2.Marshal                 1033 ns/op          424 B/op           10 allocs/op
           oj.Builder                  855 ns/op         1128 B/op           17 allocs/op
     fastjson.Marshal                  531 ns/op          679 B/op            6 allocs/op
     jsoniter.Marshal                  503 ns/op          432 B/op            7 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                  537 ns/op          416 B/op            2 allocs/op
      segment.Marshal                  573 ns/op          176 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson.Marshal                 1604 ns/op         3024 B/op           27 allocs/op
      gjson-v >>> not supported <<<
        sonic.Marshal                  688 ns/op         1803 B/op           14 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                  593 ns/op          990 B/op            0 allocs/op

     jsoniter ███████▍ 1.85
     fastjson ███████  1.75
        goccy ██████▉ 1.73
      segment ██████▍ 1.62
        codec ██████▎ 1.57
        sonic █████▍ 1.35
           oj ████▎ 1.09
         json ▓▓▓▓ 1.00
        json2 ███▌ 0.90
        gjson ██▎ 0.58
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
