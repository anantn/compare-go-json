# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Benchmarks

## x86_64 (Linux)
```
Validate []byte
         json.Validate                6580 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj.Validate                2304 ns/op            0 B/op            0 allocs/op
     fastjson.Validate                2644 ns/op            0 B/op            0 allocs/op
     jsoniter.Validate                4697 ns/op         2188 B/op          100 allocs/op
   jsonparser >>> not supported <<<
        goccy.Validate               19532 ns/op        27656 B/op          463 allocs/op
      segment.Validate                4455 ns/op            0 B/op            0 allocs/op
     simdjson >>> not supported <<<
        gjson.Validate                2367 ns/op            0 B/op            0 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                1985 ns/op            0 B/op            0 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<

        sonic █████████████▎ 3.31
           oj ███████████▍ 2.86
        gjson ███████████  2.78
     fastjson █████████▉ 2.49
      segment █████▉ 1.48
     jsoniter █████▌ 1.40
         json ▓▓▓▓ 1.00
        goccy █▎ 0.34
   jsonparser >>> not supported <<<
     simdjson >>> not supported <<<
      gjson-v >>> not supported <<<
        json2 >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<

Validate string
         json.Validate                6290 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Validate                3347 ns/op         4096 B/op            1 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Validate                3013 ns/op         4096 B/op            1 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                2646 ns/op         4138 B/op            1 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
          jin >>> not supported <<<

        sonic █████████▌ 2.38
        gjson ████████▎ 2.09
     fastjson ███████▌ 1.88
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
          jin >>> not supported <<<

Unmarshal single record (2kb), read few keys generically
         json.Unmarshal              18250 ns/op        12488 B/op           21 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal               5296 ns/op        12034 B/op            7 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal               5192 ns/op        12147 B/op           13 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal               6345 ns/op        16368 B/op            9 allocs/op
      gjson-v.Unmarshal               9117 ns/op        16368 B/op            9 allocs/op
        sonic.Unmarshal               3508 ns/op         4960 B/op           12 allocs/op
      sonic-v.Unmarshal               5689 ns/op         4943 B/op           12 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal               5379 ns/op        12040 B/op            8 allocs/op

        sonic ████████████████████▊ 5.20
   jsonparser ██████████████  3.52
     fastjson █████████████▊ 3.45
          jin █████████████▌ 3.39
      sonic-v ████████████▊ 3.21
        gjson ███████████▌ 2.88
      gjson-v ████████  2.00
         json ▓▓▓▓ 1.00
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
     jsoniter >>> not supported <<<
           oj >>> not supported <<<
        codec >>> not supported <<<
        json2 >>> not supported <<<

Unmarshal single record (2kb), read few keys into struct
         json.Unmarshal              18195 ns/op        12488 B/op           21 allocs/op
        json2.Unmarshal               6221 ns/op          128 B/op            6 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal               8646 ns/op        14087 B/op           84 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               6071 ns/op        16299 B/op            7 allocs/op
      segment.Unmarshal              10533 ns/op        12064 B/op           12 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               5728 ns/op         9037 B/op           11 allocs/op
      sonic-v.Unmarshal               7956 ns/op         9335 B/op           11 allocs/op
        codec.Unmarshal               7807 ns/op        13240 B/op           15 allocs/op
          jin >>> not supported <<<

        sonic ████████████▋ 3.18
        goccy ███████████▉ 3.00
        json2 ███████████▋ 2.92
        codec █████████▎ 2.33
      sonic-v █████████▏ 2.29
     jsoniter ████████▍ 2.10
      segment ██████▉ 1.73
         json ▓▓▓▓ 1.00
     fastjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
           oj >>> not supported <<<
          jin >>> not supported <<<

Unmarshal single record (2kb), read all keys generically
         json.Unmarshal              26724 ns/op        29776 B/op          340 allocs/op
        json2.Unmarshal              19537 ns/op        15572 B/op          295 allocs/op
           oj.Unmarshal              18888 ns/op        19299 B/op          465 allocs/op
     fastjson.Unmarshal               8331 ns/op        13608 B/op           73 allocs/op
     jsoniter.Unmarshal              20153 ns/op        31753 B/op          457 allocs/op
   jsonparser.Unmarshal              14257 ns/op        14528 B/op          135 allocs/op
        goccy.Unmarshal              19001 ns/op        34669 B/op          398 allocs/op
      segment.Unmarshal              65257 ns/op       173242 B/op          475 allocs/op
     simdjson.Unmarshal              23167 ns/op       132224 B/op           82 allocs/op
        gjson.Unmarshal              13865 ns/op        18336 B/op           76 allocs/op
      gjson-v.Unmarshal              16761 ns/op        18336 B/op           76 allocs/op
        sonic.Unmarshal              14114 ns/op        27720 B/op           59 allocs/op
      sonic-v.Unmarshal              16266 ns/op        27730 B/op           59 allocs/op
        codec.Unmarshal              29816 ns/op        19456 B/op          435 allocs/op
          jin.Unmarshal              22634 ns/op        25344 B/op          435 allocs/op

     fastjson ████████████▊ 3.21
        gjson ███████▋ 1.93
        sonic ███████▌ 1.89
   jsonparser ███████▍ 1.87
      sonic-v ██████▌ 1.64
      gjson-v ██████▍ 1.59
           oj █████▋ 1.41
        goccy █████▋ 1.41
        json2 █████▍ 1.37
     jsoniter █████▎ 1.33
          jin ████▋ 1.18
     simdjson ████▌ 1.15
         json ▓▓▓▓ 1.00
        codec ███▌ 0.90
      segment █▋ 0.41

Unmarshal single record (2kb), read all keys into struct
         json.Unmarshal              23026 ns/op        14608 B/op           80 allocs/op
        json2.Unmarshal               9991 ns/op         3110 B/op           35 allocs/op
           oj.Unmarshal              21542 ns/op         9340 B/op          453 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal               8813 ns/op        14413 B/op           77 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               6925 ns/op        16595 B/op            7 allocs/op
      segment.Unmarshal              17218 ns/op        13680 B/op           73 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               6158 ns/op         9615 B/op            9 allocs/op
      sonic-v.Unmarshal               8451 ns/op         9885 B/op            9 allocs/op
        codec.Unmarshal              10210 ns/op        14536 B/op           72 allocs/op
          jin >>> not supported <<<

        sonic ██████████████▉ 3.74
        goccy █████████████▎ 3.33
      sonic-v ██████████▉ 2.72
     jsoniter ██████████▍ 2.61
        json2 █████████▏ 2.30
        codec █████████  2.26
      segment █████▎ 1.34
           oj ████▎ 1.07
         json ▓▓▓▓ 1.00
     fastjson >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
          jin >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys generically
         json.Unmarshal          450784943 ns/op     21755813 B/op       394052 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal           79440578 ns/op      2533390 B/op        39470 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal           56506741 ns/op      4614631 B/op       177366 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal           56912758 ns/op    616926756 B/op        39632 allocs/op
        gjson.Unmarshal           92148776 ns/op    105925064 B/op        39409 allocs/op
      gjson-v.Unmarshal          181301368 ns/op    105925040 B/op        39409 allocs/op
        sonic.Unmarshal           33737302 ns/op      2886978 B/op        78814 allocs/op
      sonic-v.Unmarshal          107391693 ns/op      2858433 B/op        78813 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal           69259925 ns/op      2652133 B/op        78811 allocs/op

        sonic █████████████████████████████████████████████████████▍ 13.36
   jsonparser ███████████████████████████████▉ 7.98
     simdjson ███████████████████████████████▋ 7.92
          jin ██████████████████████████  6.51
     fastjson ██████████████████████▋ 5.67
        gjson ███████████████████▌ 4.89
      sonic-v ████████████████▊ 4.20
      gjson-v █████████▉ 2.49
         json ▓▓▓▓ 1.00
        goccy >>> not supported <<<
      segment >>> not supported <<<
     jsoniter >>> not supported <<<
           oj >>> not supported <<<
        codec >>> not supported <<<
        json2 >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys into struct
         json.Unmarshal          460579898 ns/op     21755765 B/op       394051 allocs/op
        json2.Unmarshal          201043927 ns/op        16804 B/op           64 allocs/op
           oj.Unmarshal          531329654 ns/op    246888432 B/op     15880754 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          213045821 ns/op     89157707 B/op      3979976 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal           93824257 ns/op    106274879 B/op        39731 allocs/op
      segment.Unmarshal          213760844 ns/op      2526041 B/op        39406 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal          112409780 ns/op    110170158 B/op        78922 allocs/op
      sonic-v.Unmarshal          191832126 ns/op    112471853 B/op        79034 allocs/op
        codec.Unmarshal          128750539 ns/op     48866442 B/op       157622 allocs/op
          jin >>> not supported <<<

        goccy ███████████████████▋ 4.91
        sonic ████████████████▍ 4.10
        codec ██████████████▎ 3.58
      sonic-v █████████▌ 2.40
        json2 █████████▏ 2.29
     jsoniter ████████▋ 2.16
      segment ████████▌ 2.15
         json ▓▓▓▓ 1.00
           oj ███▍ 0.87
     fastjson >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
          jin >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys generically
         json.Unmarshal          856373244 ns/op    743937104 B/op     14315119 allocs/op
        json2.Unmarshal          753594016 ns/op    693884292 B/op     11948687 allocs/op
           oj.Unmarshal          856633076 ns/op    809238840 B/op     19911204 allocs/op
     fastjson.Unmarshal          216846761 ns/op     66540694 B/op      2876748 allocs/op
     jsoniter.Unmarshal          730482165 ns/op    820437352 B/op     19201982 allocs/op
   jsonparser.Unmarshal          501289724 ns/op    105248952 B/op      5437894 allocs/op
        goccy.Unmarshal          668574918 ns/op    879523336 B/op     16801739 allocs/op
      segment.Unmarshal         2885315072 ns/op   6770756112 B/op     19978805 allocs/op
     simdjson.Unmarshal          164058608 ns/op    688367336 B/op      2758584 allocs/op
        gjson.Unmarshal          453662931 ns/op    196083461 B/op      2994781 allocs/op
      gjson-v.Unmarshal          549178954 ns/op    196083408 B/op      2994781 allocs/op
        sonic.Unmarshal          385484686 ns/op    890360826 B/op      2285573 allocs/op
      sonic-v.Unmarshal          464874163 ns/op    890101845 B/op      2285571 allocs/op
        codec.Unmarshal         1087452577 ns/op    312883584 B/op     18323502 allocs/op
          jin.Unmarshal          850547806 ns/op    563459112 B/op     18165707 allocs/op

     simdjson ████████████████████▉ 5.22
     fastjson ███████████████▊ 3.95
        sonic ████████▉ 2.22
        gjson ███████▌ 1.89
      sonic-v ███████▎ 1.84
   jsonparser ██████▊ 1.71
      gjson-v ██████▏ 1.56
        goccy █████  1.28
     jsoniter ████▋ 1.17
        json2 ████▌ 1.14
          jin ████  1.01
         json ▓▓▓▓ 1.00
           oj ███▉ 1.00
        codec ███▏ 0.79
      segment █▏ 0.30

Unmarshal many records (2kb each) from small file (100MB), read all keys into struct
         json.Unmarshal          652592543 ns/op    110917600 B/op      3113016 allocs/op
        json2.Unmarshal          381838449 ns/op    129775773 B/op      1801322 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          205965819 ns/op     97217944 B/op      2955391 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal          134923656 ns/op    106193562 B/op        39601 allocs/op
      segment.Unmarshal          543089164 ns/op     68879056 B/op      2797773 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal          137985501 ns/op    133240526 B/op       157636 allocs/op
      sonic-v.Unmarshal          217199076 ns/op    133592680 B/op       157654 allocs/op
        codec.Unmarshal          272106616 ns/op    102057748 B/op      2718954 allocs/op
          jin >>> not supported <<<

        goccy ███████████████████▎ 4.84
        sonic ██████████████████▉ 4.73
     jsoniter ████████████▋ 3.17
      sonic-v ████████████  3.00
        codec █████████▌ 2.40
        json2 ██████▊ 1.71
      segment ████▊ 1.20
         json ▓▓▓▓ 1.00
     fastjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
           oj >>> not supported <<<
          jin >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys generically
         json.Unmarshal        20810649606 ns/op    116288976 B/op      2105576 allocs/op
        json2 >>> not supported <<<
           oj >>> not supported <<<
     fastjson.Unmarshal         3241623680 ns/op     14463336 B/op       218594 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal         1813855450 ns/op     24696400 B/op       947400 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal         1122838841 ns/op   2187001720 B/op       217358 allocs/op
        gjson.Unmarshal         3167609979 ns/op   5740687776 B/op       210564 allocs/op
      gjson-v.Unmarshal         7083463910 ns/op   5740688256 B/op       210569 allocs/op
        sonic.Unmarshal          751081029 ns/op     15285532 B/op       421122 allocs/op
      sonic-v.Unmarshal         4257614278 ns/op     15308584 B/op       421125 allocs/op
        codec >>> not supported <<<
          jin.Unmarshal         2755638701 ns/op     14210880 B/op       421118 allocs/op

        sonic ██████████████████████████████████████████████████████████████████████████████████████████████████████████████▊ 27.71
     simdjson ██████████████████████████████████████████████████████████████████████████▏ 18.53
   jsonparser █████████████████████████████████████████████▉ 11.47
          jin ██████████████████████████████▏ 7.55
        gjson ██████████████████████████▎ 6.57
     fastjson █████████████████████████▋ 6.42
      sonic-v ███████████████████▌ 4.89
      gjson-v ███████████▊ 2.94
         json ▓▓▓▓ 1.00
        goccy >>> not supported <<<
      segment >>> not supported <<<
     jsoniter >>> not supported <<<
           oj >>> not supported <<<
        codec >>> not supported <<<
        json2 >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys into struct
         json.Unmarshal        20358071841 ns/op    116289040 B/op      2105576 allocs/op
        json2.Unmarshal         9215644604 ns/op        21080 B/op           72 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal         8905840910 ns/op   4650048336 B/op    210767861 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         3357652530 ns/op   5742213528 B/op       211702 allocs/op
      segment.Unmarshal         8499006680 ns/op     13537216 B/op       210563 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal         4305244681 ns/op   5758513928 B/op       421563 allocs/op
      sonic-v.Unmarshal         7821605533 ns/op   5769501536 B/op       422125 allocs/op
        codec.Unmarshal         4452677240 ns/op    261152328 B/op       842235 allocs/op
          jin >>> not supported <<<

        goccy ████████████████████████▎ 6.06
        sonic ██████████████████▉ 4.73
        codec ██████████████████▎ 4.57
      sonic-v ██████████▍ 2.60
      segment █████████▌ 2.40
     jsoniter █████████▏ 2.29
        json2 ████████▊ 2.21
         json ▓▓▓▓ 1.00
     fastjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
           oj >>> not supported <<<
          jin >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys generically
         json.Unmarshal        35159821787 ns/op  37172826320 B/op    701322312 allocs/op
        json2.Unmarshal        29785922674 ns/op  35262017872 B/op    579263257 allocs/op
           oj.Unmarshal        34075060218 ns/op  40944011608 B/op    988124957 allocs/op
     fastjson.Unmarshal         9314468548 ns/op   3342949752 B/op    142344578 allocs/op
     jsoniter.Unmarshal        30120070478 ns/op  41731708808 B/op    960520700 allocs/op
   jsonparser.Unmarshal        23199132243 ns/op   5294193208 B/op    273513560 allocs/op
        goccy.Unmarshal        27050837651 ns/op  44769802320 B/op    833147019 allocs/op
      segment.Unmarshal       113503845534 ns/op 341651494840 B/op    995514358 allocs/op
     simdjson.Unmarshal         6506581399 ns/op  27178832992 B/op    137924713 allocs/op
        gjson.Unmarshal        22575525453 ns/op  10406633456 B/op    146758244 allocs/op
      gjson-v.Unmarshal        27532559609 ns/op  10406634352 B/op    146758257 allocs/op
        sonic.Unmarshal        18304008000 ns/op  44973473296 B/op    111177614 allocs/op
      sonic-v.Unmarshal        22056433793 ns/op  44972960760 B/op    111177667 allocs/op
        codec.Unmarshal        53401177661 ns/op  13510679728 B/op    905186207 allocs/op
          jin.Unmarshal        42102001211 ns/op  28265587304 B/op    910027415 allocs/op

     simdjson █████████████████████▌ 5.40
     fastjson ███████████████  3.77
        sonic ███████▋ 1.92
      sonic-v ██████▍ 1.59
        gjson ██████▏ 1.56
   jsonparser ██████  1.52
        goccy █████▏ 1.30
      gjson-v █████  1.28
        json2 ████▋ 1.18
     jsoniter ████▋ 1.17
           oj ████▏ 1.03
         json ▓▓▓▓ 1.00
          jin ███▎ 0.84
        codec ██▋ 0.66
      segment █▏ 0.31

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys into struct
         json.Unmarshal        30586544335 ns/op   4443558064 B/op    137915231 allocs/op
        json2.Unmarshal        19234114784 ns/op   6895331104 B/op     94346981 allocs/op
           oj >>> not supported <<<
     fastjson >>> not supported <<<
     jsoniter.Unmarshal         9591990021 ns/op   5003803968 B/op    150338401 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         5800360172 ns/op   5754075992 B/op       222075 allocs/op
      segment.Unmarshal        24880385936 ns/op   3491109768 B/op    141915746 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal         6151476021 ns/op   7138552344 B/op      4633089 allocs/op
      sonic-v.Unmarshal         9856322257 ns/op   7151305096 B/op      4633758 allocs/op
        codec.Unmarshal        12667776116 ns/op   3062219376 B/op    133914581 allocs/op
          jin >>> not supported <<<

        goccy █████████████████████  5.27
        sonic ███████████████████▉ 4.97
     jsoniter ████████████▊ 3.19
      sonic-v ████████████▍ 3.10
        codec █████████▋ 2.41
        json2 ██████▎ 1.59
      segment ████▉ 1.23
         json ▓▓▓▓ 1.00
     fastjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
   jsonparser >>> not supported <<<
           oj >>> not supported <<<
          jin >>> not supported <<<

Marshal custom data through an object builder
         json.Marshal                  953 ns/op          688 B/op           17 allocs/op
        json2.Marshal                 1045 ns/op          424 B/op           10 allocs/op
           oj.Builder                  865 ns/op         1128 B/op           17 allocs/op
     fastjson.Marshal                  534 ns/op          679 B/op            6 allocs/op
     jsoniter.Marshal                  486 ns/op          432 B/op            7 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                  535 ns/op          416 B/op            2 allocs/op
      segment.Marshal                  555 ns/op          176 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson.Marshal                 1618 ns/op         3024 B/op           27 allocs/op
      gjson-v >>> not supported <<<
        sonic.Marshal                  683 ns/op         1803 B/op           14 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                  587 ns/op          957 B/op            0 allocs/op
          jin.Marshal                 1469 ns/op         2544 B/op           56 allocs/op

     jsoniter ███████▊ 1.96
     fastjson ███████▏ 1.78
        goccy ███████▏ 1.78
      segment ██████▊ 1.72
        codec ██████▍ 1.62
        sonic █████▌ 1.40
           oj ████▍ 1.10
         json ▓▓▓▓ 1.00
        json2 ███▋ 0.91
          jin ██▌ 0.65
        gjson ██▎ 0.59
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
