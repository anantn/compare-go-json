# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Validate []byte
         json.Validate                6319 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj.Validate                2303 ns/op            0 B/op            0 allocs/op
     fastjson.Validate                2658 ns/op            0 B/op            0 allocs/op
     jsoniter.Validate                4667 ns/op         2188 B/op          100 allocs/op
   jsonparser >>> not supported <<<
        goccy.Validate               20154 ns/op        27654 B/op          463 allocs/op
      segment.Validate                4441 ns/op            0 B/op            0 allocs/op
     simdjson >>> not supported <<<
        gjson.Validate                2207 ns/op            0 B/op            0 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                2058 ns/op            0 B/op            0 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

        sonic ████████████▎ 3.07
        gjson ███████████▍ 2.86
           oj ██████████▉ 2.74
     fastjson █████████▌ 2.38
      segment █████▋ 1.42
     jsoniter █████▍ 1.35
         json ▓▓▓▓ 1.00
        goccy █▎ 0.31
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        json2 >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

Validate string
         json.Validate                6368 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Validate                3233 ns/op         4096 B/op            1 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Validate                2760 ns/op         4096 B/op            1 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                2587 ns/op         4140 B/op            1 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

        sonic █████████▊ 2.46
        gjson █████████▏ 2.31
     fastjson ███████▉ 1.97
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
        djson >>> not supported <<<

Unmarshal single record (2kb), read few keys generically
         json.Unmarshal              18253 ns/op        12488 B/op           21 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal               5429 ns/op        12034 B/op            7 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal               5277 ns/op        12147 B/op           13 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal               6107 ns/op        16368 B/op            9 allocs/op
      gjson-v.Unmarshal               8648 ns/op        16368 B/op            9 allocs/op
        sonic.Unmarshal               3578 ns/op         4966 B/op           12 allocs/op
      sonic-v.Unmarshal               5674 ns/op         4937 B/op           12 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal               5384 ns/op        12040 B/op            8 allocs/op
        jason.Unmarshal              38516 ns/op        55999 B/op          495 allocs/op
        djson >>> not supported <<<

        sonic ████████████████████▍ 5.10
   jsonparser █████████████▊ 3.46
          jin █████████████▌ 3.39
     fastjson █████████████▍ 3.36
      sonic-v ████████████▊ 3.22
        gjson ███████████▉ 2.99
      gjson-v ████████▍ 2.11
         json ▓▓▓▓ 1.00
        jason █▉ 0.47
     jsoniter >>> not supported <<<
     simdjson >>> not supported <<<
      segment >>> not supported <<<
        goccy >>> not supported <<<
        codec >>> not supported <<<
           oj >>> not supported <<<
        json2 >>> not supported <<<
        djson >>> not supported <<<

Unmarshal single record (2kb), read few keys into struct
         json.Unmarshal              18229 ns/op        12488 B/op           21 allocs/op
        json2.Unmarshal               6034 ns/op          128 B/op            6 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal               8599 ns/op        14087 B/op           84 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               6223 ns/op        16296 B/op            7 allocs/op
      segment.Unmarshal              10408 ns/op        12064 B/op           12 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               5696 ns/op         9036 B/op           11 allocs/op
      sonic-v.Unmarshal               7721 ns/op         9319 B/op           11 allocs/op
        codec.Unmarshal               7649 ns/op        13240 B/op           15 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

        sonic ████████████▊ 3.20
        json2 ████████████  3.02
        goccy ███████████▋ 2.93
        codec █████████▌ 2.38
      sonic-v █████████▍ 2.36
     jsoniter ████████▍ 2.12
      segment ███████  1.75
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
         json.Unmarshal              25842 ns/op        29776 B/op          340 allocs/op
        json2.Unmarshal              19618 ns/op        15571 B/op          295 allocs/op
           oj.Unmarshal              18530 ns/op        19301 B/op          465 allocs/op
     fastjson.Unmarshal               8515 ns/op        13608 B/op           73 allocs/op
     jsoniter.Unmarshal              19996 ns/op        31753 B/op          457 allocs/op
   jsonparser.Unmarshal              14380 ns/op        14528 B/op          135 allocs/op
        goccy.Unmarshal              19519 ns/op        34668 B/op          398 allocs/op
      segment.Unmarshal              64734 ns/op       173241 B/op          475 allocs/op
     simdjson.Unmarshal              22999 ns/op       132224 B/op           82 allocs/op
        gjson.Unmarshal              12771 ns/op        18336 B/op           76 allocs/op
      gjson-v.Unmarshal              15370 ns/op        18336 B/op           76 allocs/op
        sonic.Unmarshal              14180 ns/op        27701 B/op           59 allocs/op
      sonic-v.Unmarshal              16269 ns/op        27713 B/op           59 allocs/op
        codec.Unmarshal              29904 ns/op        19456 B/op          435 allocs/op
          jin.Unmarshal              23108 ns/op        25344 B/op          435 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal              14408 ns/op        32169 B/op          183 allocs/op

     fastjson ████████████▏ 3.03
        gjson ████████  2.02
        sonic ███████▎ 1.82
   jsonparser ███████▏ 1.80
        djson ███████▏ 1.79
      gjson-v ██████▋ 1.68
      sonic-v ██████▎ 1.59
           oj █████▌ 1.39
        goccy █████▎ 1.32
        json2 █████▎ 1.32
     jsoniter █████▏ 1.29
     simdjson ████▍ 1.12
          jin ████▍ 1.12
         json ▓▓▓▓ 1.00
        codec ███▍ 0.86
      segment █▌ 0.40
        jason >>> not supported <<<

Unmarshal single record (2kb), read all keys into struct
         json.Unmarshal              23132 ns/op        14608 B/op           80 allocs/op
        json2.Unmarshal              10184 ns/op         3111 B/op           35 allocs/op
           oj.Unmarshal              21474 ns/op         9339 B/op          453 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal               8844 ns/op        14412 B/op           77 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               7299 ns/op        16594 B/op            7 allocs/op
      segment.Unmarshal              17286 ns/op        13680 B/op           73 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               6154 ns/op         9592 B/op            9 allocs/op
      sonic-v.Unmarshal               8353 ns/op         9849 B/op            9 allocs/op
        codec.Unmarshal              10152 ns/op        14536 B/op           72 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

        sonic ███████████████  3.76
        goccy ████████████▋ 3.17
      sonic-v ███████████  2.77
     jsoniter ██████████▍ 2.62
        codec █████████  2.28
        json2 █████████  2.27
      segment █████▎ 1.34
           oj ████▎ 1.08
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     fastjson >>> not supported <<<
   jsonparser >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys generically
         json.Unmarshal          450573915 ns/op     21755781 B/op       394052 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal           82393398 ns/op      2533417 B/op        39470 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal           57503810 ns/op      4614622 B/op       177366 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal           59188980 ns/op    612107390 B/op        39635 allocs/op
        gjson.Unmarshal           86955573 ns/op    105925348 B/op        39412 allocs/op
      gjson-v.Unmarshal          170461471 ns/op    105925104 B/op        39409 allocs/op
        sonic.Unmarshal           33814027 ns/op      2892849 B/op        78814 allocs/op
      sonic-v.Unmarshal          107136950 ns/op      2866654 B/op        78814 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal           69031295 ns/op      2652139 B/op        78811 allocs/op
        jason.Unmarshal          995813176 ns/op   1083141296 B/op     16167192 allocs/op
        djson >>> not supported <<<

        sonic █████████████████████████████████████████████████████▎ 13.33
   jsonparser ███████████████████████████████▎ 7.84
     simdjson ██████████████████████████████▍ 7.61
          jin ██████████████████████████  6.53
     fastjson █████████████████████▊ 5.47
        gjson ████████████████████▋ 5.18
      sonic-v ████████████████▊ 4.21
      gjson-v ██████████▌ 2.64
         json ▓▓▓▓ 1.00
        jason █▊ 0.45
      segment >>> not supported <<<
        goccy >>> not supported <<<
     jsoniter >>> not supported <<<
           oj >>> not supported <<<
        codec >>> not supported <<<
        json2 >>> not supported <<<
        djson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys into struct
         json.Unmarshal          447826277 ns/op     21755856 B/op       394052 allocs/op
        json2.Unmarshal          189722343 ns/op        16804 B/op           64 allocs/op
           oj.Unmarshal          520284601 ns/op    246896112 B/op     15880782 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          207919243 ns/op     89156681 B/op      3979976 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal           89358638 ns/op    106272872 B/op        39718 allocs/op
      segment.Unmarshal          218663992 ns/op      2526041 B/op        39406 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal          113049766 ns/op    110334943 B/op        78920 allocs/op
      sonic-v.Unmarshal          189297721 ns/op    112435069 B/op        79035 allocs/op
        codec.Unmarshal          126612607 ns/op     48866382 B/op       157622 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

        goccy ████████████████████  5.01
        sonic ███████████████▊ 3.96
        codec ██████████████▏ 3.54
      sonic-v █████████▍ 2.37
        json2 █████████▍ 2.36
     jsoniter ████████▌ 2.15
      segment ████████▏ 2.05
         json ▓▓▓▓ 1.00
           oj ███▍ 0.86
   jsonparser >>> not supported <<<
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     fastjson >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys generically
         json.Unmarshal          849433662 ns/op    743908888 B/op     14315026 allocs/op
        json2.Unmarshal          742199693 ns/op    693865688 B/op     11948654 allocs/op
           oj.Unmarshal          851101649 ns/op    809198168 B/op     19911062 allocs/op
     fastjson.Unmarshal          213916531 ns/op     66540681 B/op      2876748 allocs/op
     jsoniter.Unmarshal          718897752 ns/op    820427492 B/op     19201953 allocs/op
   jsonparser.Unmarshal          486967444 ns/op    105248941 B/op      5437894 allocs/op
        goccy.Unmarshal          681166895 ns/op    879551032 B/op     16801885 allocs/op
      segment.Unmarshal         2750541850 ns/op   6770760256 B/op     19978836 allocs/op
     simdjson.Unmarshal          167813249 ns/op    688369008 B/op      2758582 allocs/op
        gjson.Unmarshal          412481338 ns/op    196083413 B/op      2994781 allocs/op
      gjson-v.Unmarshal          519314171 ns/op    196083416 B/op      2994781 allocs/op
        sonic.Unmarshal          383255122 ns/op    890251445 B/op      2285568 allocs/op
      sonic-v.Unmarshal          467994058 ns/op    890239597 B/op      2285571 allocs/op
        codec.Unmarshal         1080303597 ns/op    312883248 B/op     18323501 allocs/op
          jin.Unmarshal          847140008 ns/op    563459104 B/op     18165707 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal          389260078 ns/op    770790210 B/op      7537387 allocs/op

     simdjson ████████████████████▏ 5.06
     fastjson ███████████████▉ 3.97
        sonic ████████▊ 2.22
        djson ████████▋ 2.18
        gjson ████████▏ 2.06
      sonic-v ███████▎ 1.82
   jsonparser ██████▉ 1.74
      gjson-v ██████▌ 1.64
        goccy ████▉ 1.25
     jsoniter ████▋ 1.18
        json2 ████▌ 1.14
          jin ████  1.00
         json ▓▓▓▓ 1.00
           oj ███▉ 1.00
        codec ███▏ 0.79
      segment █▏ 0.31
        jason >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys into struct
         json.Unmarshal          666517026 ns/op    110917640 B/op      3113017 allocs/op
        json2.Unmarshal          383404713 ns/op    129775789 B/op      1801322 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          211910636 ns/op     97217995 B/op      2955391 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal          144636596 ns/op    106194393 B/op        39625 allocs/op
      segment.Unmarshal          542951087 ns/op     68879064 B/op      2797773 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal          141153952 ns/op    133200526 B/op       157635 allocs/op
      sonic-v.Unmarshal          221895368 ns/op    133733617 B/op       157656 allocs/op
        codec.Unmarshal          271644014 ns/op    102057760 B/op      2718954 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

        sonic ██████████████████▉ 4.72
        goccy ██████████████████▍ 4.61
     jsoniter ████████████▌ 3.15
      sonic-v ████████████  3.00
        codec █████████▊ 2.45
        json2 ██████▉ 1.74
      segment ████▉ 1.23
         json ▓▓▓▓ 1.00
     fastjson >>> not supported <<<
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
           oj >>> not supported <<<
   jsonparser >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys generically
         json.Unmarshal        20781966067 ns/op    116289120 B/op      2105577 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal         3469462354 ns/op     14463336 B/op       218594 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal         1923834628 ns/op     24696400 B/op       947400 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal         1388723450 ns/op   2349547896 B/op       217458 allocs/op
        gjson.Unmarshal         2956191214 ns/op   5740687680 B/op       210563 allocs/op
      gjson-v.Unmarshal         6726572727 ns/op   5740688064 B/op       210567 allocs/op
        sonic.Unmarshal          770053351 ns/op     15244484 B/op       421120 allocs/op
      sonic-v.Unmarshal         4390838164 ns/op     15308392 B/op       421123 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal         2674136863 ns/op     14210896 B/op       421118 allocs/op
        jason.Unmarshal        37317770441 ns/op  51079748160 B/op    711849010 allocs/op
        djson >>> not supported <<<

        sonic ███████████████████████████████████████████████████████████████████████████████████████████████████████████▉ 26.99
     simdjson ███████████████████████████████████████████████████████████▊ 14.96
   jsonparser ███████████████████████████████████████████▏ 10.80
          jin ███████████████████████████████  7.77
        gjson ████████████████████████████  7.03
     fastjson ███████████████████████▉ 5.99
      sonic-v ██████████████████▉ 4.73
      gjson-v ████████████▎ 3.09
         json ▓▓▓▓ 1.00
        jason ██▏ 0.56
      segment >>> not supported <<<
        goccy >>> not supported <<<
     jsoniter >>> not supported <<<
           oj >>> not supported <<<
        codec >>> not supported <<<
        json2 >>> not supported <<<
        djson >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys into struct
         json.Unmarshal        20573385083 ns/op    116288992 B/op      2105576 allocs/op
        json2.Unmarshal         8975810965 ns/op        21080 B/op           72 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal         8931055865 ns/op   4650048432 B/op    210767862 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         3072960693 ns/op   5742214848 B/op       211735 allocs/op
      segment.Unmarshal         8452832471 ns/op     13537216 B/op       210563 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal         4306649981 ns/op   5755745096 B/op       421538 allocs/op
      sonic-v.Unmarshal         7717709984 ns/op   5768312856 B/op       422113 allocs/op
        codec.Unmarshal         4362483815 ns/op    261152232 B/op       842234 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

        goccy ██████████████████████████▊ 6.69
        sonic ███████████████████  4.78
        codec ██████████████████▊ 4.72
      sonic-v ██████████▋ 2.67
      segment █████████▋ 2.43
     jsoniter █████████▏ 2.30
        json2 █████████▏ 2.29
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
         json.Unmarshal        35147413704 ns/op  37173140608 B/op    701323406 allocs/op
        json2.Unmarshal        29645410052 ns/op  35262314864 B/op    579264295 allocs/op
           oj.Unmarshal        34158059579 ns/op  40944281288 B/op    988125905 allocs/op
     fastjson.Unmarshal         9390894991 ns/op   3342949928 B/op    142344579 allocs/op
     jsoniter.Unmarshal        29734387745 ns/op  41731506816 B/op    960520003 allocs/op
   jsonparser.Unmarshal        23375225282 ns/op   5294193032 B/op    273513557 allocs/op
        goccy.Unmarshal        28240626779 ns/op  44770036936 B/op    833148105 allocs/op
      segment.Unmarshal       114463310522 ns/op 341651488056 B/op    995514294 allocs/op
     simdjson.Unmarshal         6507662742 ns/op  27022855696 B/op    137924684 allocs/op
        gjson.Unmarshal        20270092100 ns/op  10406633440 B/op    146758240 allocs/op
      gjson-v.Unmarshal        24768375028 ns/op  10406634320 B/op    146758254 allocs/op
        sonic.Unmarshal        18516214686 ns/op  44978803664 B/op    111177929 allocs/op
      sonic-v.Unmarshal        22211858955 ns/op  44977925872 B/op    111177962 allocs/op
        codec.Unmarshal        53437824389 ns/op  13510680448 B/op    905186204 allocs/op
          jin.Unmarshal        42562206608 ns/op  28265590744 B/op    910027443 allocs/op
        jason >>> not supported <<<
        djson.Unmarshal        18799426007 ns/op  39791291552 B/op    375169342 allocs/op

     simdjson █████████████████████▌ 5.40
     fastjson ██████████████▉ 3.74
        sonic ███████▌ 1.90
        djson ███████▍ 1.87
        gjson ██████▉ 1.73
      sonic-v ██████▎ 1.58
   jsonparser ██████  1.50
      gjson-v █████▋ 1.42
        goccy ████▉ 1.24
        json2 ████▋ 1.19
     jsoniter ████▋ 1.18
           oj ████  1.03
         json ▓▓▓▓ 1.00
          jin ███▎ 0.83
        codec ██▋ 0.66
      segment █▏ 0.31
        jason >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys into struct
         json.Unmarshal        30595766323 ns/op   4443558512 B/op    137915232 allocs/op
        json2.Unmarshal        19395247290 ns/op   6895391504 B/op     94347019 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal         9655571901 ns/op   5003847384 B/op    150338412 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         6195779529 ns/op   5755022440 B/op       223501 allocs/op
      segment.Unmarshal        25069133978 ns/op   3491109864 B/op    141915748 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal         6142954450 ns/op   7136799496 B/op      4633087 allocs/op
      sonic-v.Unmarshal         9911038693 ns/op   7152764744 B/op      4633808 allocs/op
        codec.Unmarshal        12685270152 ns/op   3062218992 B/op    133914575 allocs/op
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

        sonic ███████████████████▉ 4.98
        goccy ███████████████████▊ 4.94
     jsoniter ████████████▋ 3.17
      sonic-v ████████████▎ 3.09
        codec █████████▋ 2.41
        json2 ██████▎ 1.58
      segment ████▉ 1.22
         json ▓▓▓▓ 1.00
     fastjson >>> not supported <<<
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
           oj >>> not supported <<<
   jsonparser >>> not supported <<<
          jin >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

Marshal custom data through an object builder
         json.Marshal                  965 ns/op          688 B/op           17 allocs/op
        json2.Marshal                 1026 ns/op          424 B/op           10 allocs/op
           oj.Builder                  843 ns/op         1128 B/op           17 allocs/op
     fastjson.Marshal                  557 ns/op          679 B/op            6 allocs/op
     jsoniter.Marshal                  482 ns/op          432 B/op            7 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                  537 ns/op          416 B/op            2 allocs/op
      segment.Marshal                  544 ns/op          176 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson.Marshal                 1567 ns/op         3024 B/op           27 allocs/op
      gjson-v >>> not supported <<<
        sonic.Marshal                  690 ns/op         1803 B/op           14 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                  596 ns/op          942 B/op            0 allocs/op
          jin.Marshal                 1465 ns/op         2544 B/op           56 allocs/op
        jason >>> not supported <<<
        djson >>> not supported <<<

     jsoniter ████████  2.00
        goccy ███████▏ 1.80
      segment ███████  1.77
     fastjson ██████▉ 1.73
        codec ██████▍ 1.62
        sonic █████▌ 1.40
           oj ████▌ 1.14
         json ▓▓▓▓ 1.00
        json2 ███▊ 0.94
          jin ██▋ 0.66
        gjson ██▍ 0.62
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
   jsonparser >>> not supported <<<
        jason >>> not supported <<<
        djson >>> not supported <<<

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
