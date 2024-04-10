# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Results

# Raw Benchmark Data

Outputs are also available as CSV in the repository.
[x86_64](benchmarks-x86_64.csv) and [arm64](benchmarks-arm64.csv).

## x86_64 (Linux)
```
Unmarshal CanadaGeometry generically
         json.Unmarshal            2201059 ns/op      1339377 B/op        52215 allocs/op
        json2.Unmarshal            1808184 ns/op      1011594 B/op        37898 allocs/op
     fastjson.Unmarshal            2143287 ns/op      8725359 B/op        30278 allocs/op
     jsoniter.Unmarshal            2910583 ns/op      1961671 B/op        84971 allocs/op
   jsonparser.Unmarshal            1269026 ns/op          152 B/op           12 allocs/op
        goccy.Unmarshal            1932274 ns/op      1470385 B/op        59189 allocs/op
      segment.Unmarshal            3653966 ns/op      2119753 B/op        51680 allocs/op
     simdjson.Unmarshal            1284845 ns/op       987182 B/op        14324 allocs/op
        gjson.Unmarshal            2702880 ns/op       393072 B/op        14314 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal             922450 ns/op      2816834 B/op        29654 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            2649064 ns/op      1765858 B/op        37408 allocs/op
        djson.Unmarshal            1310825 ns/op      1289905 B/op        37891 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal             951890 ns/op      1011609 B/op        37903 allocs/op

        sonic █████████▌ 2.39
       sonnet █████████▏ 2.31
 * jsonparser ██████▉ 1.73
   * simdjson ██████▊ 1.71
      * djson ██████▋ 1.68
        json2 ████▊ 1.22
        goccy ████▌ 1.14
   * fastjson ████  1.03
         json ▓▓▓▓ 1.00
        codec ███▎ 0.83
      * gjson ███▎ 0.81
     jsoniter ███  0.76
      segment ██▍ 0.60
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal CanadaGeometry into struct
         json.Unmarshal            1999023 ns/op          472 B/op           13 allocs/op
        json2.Unmarshal            1272186 ns/op       347991 B/op         1626 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            1572584 ns/op       598818 B/op        25122 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal             749830 ns/op       278736 B/op            1 allocs/op
      segment.Unmarshal             739785 ns/op           48 B/op            4 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal             310513 ns/op       278834 B/op            2 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            1268821 ns/op         1224 B/op            7 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal            2484316 ns/op       535504 B/op           25 allocs/op
     easyjson.Unmarshal             919723 ns/op       330368 B/op          753 allocs/op
       sonnet.Unmarshal             429378 ns/op       143113 B/op          967 allocs/op

        sonic █████████████████████████▊ 6.44
       sonnet ██████████████████▌ 4.66
      segment ██████████▊ 2.70
        goccy ██████████▋ 2.67
     easyjson ████████▋ 2.17
        codec ██████▎ 1.58
        json2 ██████▎ 1.57
     jsoniter █████  1.27
         json ▓▓▓▓ 1.00
       ffjson ███▏ 0.80
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal CanadaGeometry generically
         json.Marshal              1105380 ns/op       280161 B/op           21 allocs/op
        json2.Marshal               839797 ns/op       279610 B/op            7 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               460046 ns/op       172567 B/op           13 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal               870669 ns/op       280751 B/op            3 allocs/op
      segment.Marshal               783300 ns/op       279437 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               449976 ns/op       279144 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal               921019 ns/op      1186328 B/op           22 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Marshal               698024 ns/op       278946 B/op            5 allocs/op

        sonic █████████▊ 2.46
     jsoniter █████████▌ 2.40
       sonnet ██████▎ 1.58
      segment █████▋ 1.41
        json2 █████▎ 1.32
        goccy █████  1.27
        codec ████▊ 1.20
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal CanadaGeometry from struct
         json.Marshal               765114 ns/op       279277 B/op            1 allocs/op
        json2.Marshal               799925 ns/op       279392 B/op            1 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               196240 ns/op       172182 B/op            2 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal               703238 ns/op       279308 B/op            1 allocs/op
      segment.Marshal               702256 ns/op       279293 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               336709 ns/op       278743 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal               918709 ns/op      1186329 B/op           22 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal              1490419 ns/op      1058608 B/op        14391 allocs/op
     easyjson.Marshal               686210 ns/op       280299 B/op           20 allocs/op
       sonnet.Marshal               723355 ns/op       278907 B/op            5 allocs/op

     jsoniter ███████████████▌ 3.90
        sonic █████████  2.27
     easyjson ████▍ 1.11
      segment ████▎ 1.09
        goccy ████▎ 1.09
       sonnet ████▏ 1.06
         json ▓▓▓▓ 1.00
        json2 ███▊ 0.96
        codec ███▎ 0.83
       ffjson ██  0.51
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal CITMCatalog generically
         json.Unmarshal            7496686 ns/op      5123835 B/op        95373 allocs/op
        json2.Unmarshal            3389433 ns/op      4958941 B/op        80513 allocs/op
     fastjson.Unmarshal            3768654 ns/op     19991453 B/op        45701 allocs/op
     jsoniter.Unmarshal            3936442 ns/op      5573428 B/op       118756 allocs/op
   jsonparser.Unmarshal            6630522 ns/op       145320 B/op        15863 allocs/op
        goccy.Unmarshal            4256725 ns/op      7416711 B/op       123577 allocs/op
      segment.Unmarshal           17097780 ns/op     55526612 B/op       123714 allocs/op
     simdjson.Unmarshal            1582903 ns/op      2854125 B/op        15139 allocs/op
        gjson.Unmarshal            6042012 ns/op      1855520 B/op        15131 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal            2438729 ns/op      8644975 B/op        58177 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            7694167 ns/op      6195005 B/op       159475 allocs/op
        djson.Unmarshal            2985875 ns/op      6437585 B/op        54480 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal            2363520 ns/op      4968468 B/op        80980 allocs/op

   * simdjson ██████████████████▉ 4.74
       sonnet ████████████▋ 3.17
        sonic ████████████▎ 3.07
      * djson ██████████  2.51
        json2 ████████▊ 2.21
   * fastjson ███████▉ 1.99
     jsoniter ███████▌ 1.90
        goccy ███████  1.76
      * gjson ████▉ 1.24
 * jsonparser ████▌ 1.13
         json ▓▓▓▓ 1.00
        codec ███▉ 0.97
      segment █▊ 0.44
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal CITMCatalog into struct
         json.Unmarshal            6899554 ns/op       253352 B/op        10950 allocs/op
        json2.Unmarshal            2157551 ns/op       888608 B/op         6295 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            1296765 ns/op        72338 B/op         2535 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal            1071219 ns/op      1782677 B/op         2226 allocs/op
      segment.Unmarshal            3095855 ns/op        52384 B/op         1354 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal             759359 ns/op      1731648 B/op          111 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            1920141 ns/op        47880 B/op         1036 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal       >>> ffjson error: (*errors.errorString)cannot unmarshal tok:string into Go value for int64 offset=40 line=3 char=19 <<<
     easyjson.Unmarshal            2146114 ns/op       896724 B/op         5705 allocs/op
       sonnet.Unmarshal            1412789 ns/op       613507 B/op        13348 allocs/op

        sonic ████████████████████████████████████▎ 9.09
        goccy █████████████████████████▊ 6.44
     jsoniter █████████████████████▎ 5.32
       sonnet ███████████████████▌ 4.88
        codec ██████████████▎ 3.59
     easyjson ████████████▊ 3.21
        json2 ████████████▊ 3.20
      segment ████████▉ 2.23
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> ffjson error: (*errors.errorString)cannot unmarshal tok:string into Go value for int64 offset=40 line=3 char=19 <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal CITMCatalog generically
         json.Marshal              3470270 ns/op      2384809 B/op        62674 allocs/op
        json2.Marshal              1540170 ns/op       511740 B/op           16 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              2118618 ns/op      1912572 B/op        32812 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              1987244 ns/op       658881 B/op          331 allocs/op
      segment.Marshal              1482226 ns/op       508028 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              1044746 ns/op       509444 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              1964994 ns/op      2538009 B/op           25 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Marshal              1234419 ns/op       508393 B/op            5 allocs/op

        sonic █████████████▎ 3.32
       sonnet ███████████▏ 2.81
      segment █████████▎ 2.34
        json2 █████████  2.25
        codec ███████  1.77
        goccy ██████▉ 1.75
     jsoniter ██████▌ 1.64
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal CITMCatalog from struct
         json.Marshal               549952 ns/op       556600 B/op          890 allocs/op
        json2.Marshal               822414 ns/op       512881 B/op          133 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               418564 ns/op       510207 B/op           22 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal               295109 ns/op       521035 B/op            3 allocs/op
      segment.Marshal               359428 ns/op       547059 B/op          616 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               201527 ns/op       508430 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal               889686 ns/op      2538009 B/op           25 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal               807025 ns/op       904321 B/op        14180 allocs/op
     easyjson.Marshal               391360 ns/op       524115 B/op         1125 allocs/op
       sonnet.Marshal               340656 ns/op       508294 B/op            5 allocs/op

        sonic ██████████▉ 2.73
        goccy ███████▍ 1.86
       sonnet ██████▍ 1.61
      segment ██████  1.53
     easyjson █████▌ 1.41
     jsoniter █████▎ 1.31
         json ▓▓▓▓ 1.00
       ffjson ██▋ 0.68
        json2 ██▋ 0.67
        codec ██▍ 0.62
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal SyntheaFHIR generically
         json.Unmarshal            9510204 ns/op      7430310 B/op       134384 allocs/op
        json2.Unmarshal            5222519 ns/op      6939006 B/op       115832 allocs/op
     fastjson.Unmarshal            5067856 ns/op     26046546 B/op        66098 allocs/op
     jsoniter.Unmarshal            5907076 ns/op      8197336 B/op       181426 allocs/op
   jsonparser.Unmarshal           10258052 ns/op      1108793 B/op        49604 allocs/op
        goccy.Unmarshal            5543987 ns/op      9838170 B/op       159459 allocs/op
      segment.Unmarshal           43495810 ns/op     77571274 B/op       193338 allocs/op
     simdjson.Unmarshal            2530651 ns/op      6172197 B/op        26181 allocs/op
        gjson.Unmarshal            8402637 ns/op      2432544 B/op        26821 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal            3056453 ns/op      9489690 B/op        67615 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal           11316617 ns/op      9669900 B/op       247797 allocs/op
        djson.Unmarshal            3875910 ns/op      8374521 B/op        68394 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal            3817061 ns/op      7410722 B/op       133127 allocs/op

   * simdjson ███████████████  3.76
        sonic ████████████▍ 3.11
       sonnet █████████▉ 2.49
      * djson █████████▊ 2.45
   * fastjson ███████▌ 1.88
        json2 ███████▎ 1.82
        goccy ██████▊ 1.72
     jsoniter ██████▍ 1.61
      * gjson ████▌ 1.13
         json ▓▓▓▓ 1.00
 * jsonparser ███▋ 0.93
        codec ███▎ 0.84
      segment ▊ 0.22
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal SyntheaFHIR into struct
         json.Unmarshal            9258864 ns/op       625392 B/op        21883 allocs/op
        json2.Unmarshal            3566256 ns/op      3103391 B/op        14550 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            2237231 ns/op       825754 B/op        29439 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal            1671283 ns/op      2157891 B/op         2507 allocs/op
      segment.Unmarshal            6145498 ns/op       623760 B/op        21832 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal            1319841 ns/op      2017598 B/op           10 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            3030203 ns/op       623400 B/op        21826 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal           15816187 ns/op      4334228 B/op        21904 allocs/op
     easyjson.Unmarshal            2959394 ns/op      3531633 B/op        29214 allocs/op
       sonnet.Unmarshal            2453738 ns/op      3446361 B/op        33093 allocs/op

        sonic ████████████████████████████  7.02
        goccy ██████████████████████▏ 5.54
     jsoniter ████████████████▌ 4.14
       sonnet ███████████████  3.77
     easyjson ████████████▌ 3.13
        codec ████████████▏ 3.06
        json2 ██████████▍ 2.60
      segment ██████  1.51
         json ▓▓▓▓ 1.00
       ffjson ██▎ 0.59
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal SyntheaFHIR generically
         json.Marshal              5581444 ns/op      4137929 B/op        95282 allocs/op
        json2.Marshal              2417076 ns/op      1161868 B/op            8 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              3173327 ns/op      3089512 B/op        45292 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              3866950 ns/op      2495264 B/op         6474 allocs/op
      segment.Marshal              2522778 ns/op      1161746 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              2001790 ns/op      1147635 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              2498663 ns/op      6642202 B/op           29 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Marshal              2098520 ns/op      1147314 B/op            5 allocs/op

        sonic ███████████▏ 2.79
       sonnet ██████████▋ 2.66
        json2 █████████▏ 2.31
        codec ████████▉ 2.23
      segment ████████▊ 2.21
     jsoniter ███████  1.76
        goccy █████▊ 1.44
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal SyntheaFHIR from struct
         json.Marshal              5057858 ns/op      3983870 B/op        18009 allocs/op
        json2.Marshal              5154576 ns/op      3217367 B/op         2569 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              3435776 ns/op      3970837 B/op        17963 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              2625887 ns/op      3945464 B/op        17940 allocs/op
      segment.Marshal              2015756 ns/op      3272481 B/op         5338 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              1971755 ns/op      3998786 B/op        17939 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              4806434 ns/op     16792096 B/op           33 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal              3804123 ns/op      9493195 B/op        45596 allocs/op
     easyjson.Marshal              2460771 ns/op      3979943 B/op        15520 allocs/op
       sonnet.Marshal              3118662 ns/op      3944830 B/op        17940 allocs/op

        sonic ██████████▎ 2.57
      segment ██████████  2.51
     easyjson ████████▏ 2.06
        goccy ███████▋ 1.93
       sonnet ██████▍ 1.62
     jsoniter █████▉ 1.47
       ffjson █████▎ 1.33
        codec ████▏ 1.05
         json ▓▓▓▓ 1.00
        json2 ███▉ 0.98
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal TwitterStatus generically
         json.Unmarshal            3096270 ns/op      2152570 B/op        31262 allocs/op
        json2.Unmarshal            1953068 ns/op      2074857 B/op        27238 allocs/op
     fastjson.Unmarshal            1001276 ns/op      5341210 B/op        11130 allocs/op
     jsoniter.Unmarshal            1996329 ns/op      2427579 B/op        45042 allocs/op
   jsonparser.Unmarshal            2143972 ns/op       300312 B/op        10184 allocs/op
        goccy.Unmarshal            1817143 ns/op      2755798 B/op        39695 allocs/op
      segment.Unmarshal           10431377 ns/op      7060656 B/op        39237 allocs/op
     simdjson.Unmarshal             743708 ns/op      1950208 B/op         5732 allocs/op
        gjson.Unmarshal            1870382 ns/op       822360 B/op         6884 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal            1049349 ns/op      2606585 B/op        12136 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            3326482 ns/op      2576810 B/op        60985 allocs/op
        djson.Unmarshal            1438365 ns/op      2470278 B/op        12677 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal            1377480 ns/op      2085723 B/op        29439 allocs/op

   * simdjson ████████████████▋ 4.16
   * fastjson ████████████▎ 3.09
        sonic ███████████▊ 2.95
       sonnet ████████▉ 2.25
      * djson ████████▌ 2.15
        goccy ██████▊ 1.70
      * gjson ██████▌ 1.66
        json2 ██████▎ 1.59
     jsoniter ██████▏ 1.55
 * jsonparser █████▊ 1.44
         json ▓▓▓▓ 1.00
        codec ███▋ 0.93
      segment █▏ 0.30
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal TwitterStatus into struct
         json.Unmarshal            2826550 ns/op       304611 B/op         6486 allocs/op
        json2.Unmarshal            1054087 ns/op       306193 B/op         2972 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal             733289 ns/op       339699 B/op         6293 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal             527155 ns/op       647251 B/op          238 allocs/op
      segment.Unmarshal            2133152 ns/op       369712 B/op         5341 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal             376674 ns/op       694348 B/op          401 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal             856766 ns/op       215944 B/op         4530 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal            1708915 ns/op      1052320 B/op         9494 allocs/op
     easyjson.Unmarshal             616803 ns/op       324336 B/op         4907 allocs/op
       sonnet.Unmarshal             723734 ns/op       373601 B/op         6257 allocs/op

        sonic ██████████████████████████████  7.50
        goccy █████████████████████▍ 5.36
     easyjson ██████████████████▎ 4.58
       sonnet ███████████████▌ 3.91
     jsoniter ███████████████▍ 3.85
        codec █████████████▏ 3.30
        json2 ██████████▋ 2.68
       ffjson ██████▌ 1.65
      segment █████▎ 1.33
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal TwitterStatus generically
         json.Marshal              2274129 ns/op      1488060 B/op        27955 allocs/op
        json2.Marshal               821631 ns/op       469040 B/op            7 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               838312 ns/op       630831 B/op         3793 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              1688990 ns/op      1050622 B/op          551 allocs/op
      segment.Marshal              1274136 ns/op       478116 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               565156 ns/op       467159 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal               922706 ns/op      2538008 B/op           25 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Marshal              1117461 ns/op       475538 B/op            5 allocs/op

        sonic ████████████████  4.02
        json2 ███████████  2.77
     jsoniter ██████████▊ 2.71
        codec █████████▊ 2.46
       sonnet ████████▏ 2.04
      segment ███████▏ 1.78
        goccy █████▍ 1.35
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal TwitterStatus from struct
         json.Marshal               411430 ns/op       513404 B/op          141 allocs/op
        json2.Marshal               555077 ns/op       503163 B/op           59 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               427647 ns/op       503249 B/op           52 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal               273936 ns/op       509511 B/op            6 allocs/op
      segment.Marshal               287307 ns/op       512407 B/op          121 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               111841 ns/op       500238 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal               747740 ns/op      2540385 B/op           29 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal               527071 ns/op       798096 B/op         6117 allocs/op
     easyjson.Marshal               355337 ns/op       520014 B/op          953 allocs/op
       sonnet.Marshal               321534 ns/op       508298 B/op            5 allocs/op

        sonic ██████████████▋ 3.68
        goccy ██████  1.50
      segment █████▋ 1.43
       sonnet █████  1.28
     easyjson ████▋ 1.16
         json ▓▓▓▓ 1.00
     jsoniter ███▊ 0.96
       ffjson ███  0.78
        json2 ██▉ 0.74
        codec ██▏ 0.55
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal GolangSource generically
         json.Unmarshal           11968064 ns/op      7916353 B/op       271264 allocs/op
        json2.Unmarshal            8276294 ns/op      6990106 B/op       218839 allocs/op
     fastjson.Unmarshal            9467542 ns/op     50767007 B/op       130330 allocs/op
     jsoniter.Unmarshal           10051329 ns/op      9233781 B/op       346658 allocs/op
   jsonparser.Unmarshal           30813505 ns/op       925272 B/op        64023 allocs/op
        goccy.Unmarshal           10253947 ns/op     10870510 B/op       320839 allocs/op
      segment.Unmarshal           39725799 ns/op     67095366 B/op       334022 allocs/op
     simdjson.Unmarshal            5902033 ns/op     12653678 B/op        63942 allocs/op
        gjson.Unmarshal           35776431 ns/op      2658512 B/op        76819 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal            4789534 ns/op     10997414 B/op       128226 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal           20468895 ns/op     10502701 B/op       436539 allocs/op
        djson.Unmarshal            4750686 ns/op      7822418 B/op       117489 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal            5113454 ns/op      7011913 B/op       219940 allocs/op

      * djson ██████████  2.52
        sonic █████████▉ 2.50
       sonnet █████████▎ 2.34
   * simdjson ████████  2.03
        json2 █████▊ 1.45
   * fastjson █████  1.26
     jsoniter ████▊ 1.19
        goccy ████▋ 1.17
         json ▓▓▓▓ 1.00
        codec ██▎ 0.58
 * jsonparser █▌ 0.39
      * gjson █▎ 0.33
      segment █▏ 0.30
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal GolangSource into struct
         json.Unmarshal           12829904 ns/op       704424 B/op        24927 allocs/op
        json2.Unmarshal            6322344 ns/op      3236434 B/op        13941 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            4089098 ns/op       994092 B/op        37077 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal            2295718 ns/op      1944324 B/op            1 allocs/op
      segment.Unmarshal            2363810 ns/op       411672 B/op        12806 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal            1930832 ns/op      1942003 B/op            2 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            5392874 ns/op       412848 B/op        12809 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal            7011536 ns/op      6276801 B/op        52765 allocs/op
     easyjson.Unmarshal            4262419 ns/op      4282424 B/op        27846 allocs/op
       sonnet.Unmarshal            2956560 ns/op      1811651 B/op        26312 allocs/op

        sonic ██████████████████████████▌ 6.64
        goccy ██████████████████████▎ 5.59
      segment █████████████████████▋ 5.43
       sonnet █████████████████▎ 4.34
     jsoniter ████████████▌ 3.14
     easyjson ████████████  3.01
        codec █████████▌ 2.38
        json2 ████████  2.03
       ffjson ███████▎ 1.83
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal GolangSource generically
         json.Marshal             11721767 ns/op      8546871 B/op       192096 allocs/op
        json2.Marshal              4843868 ns/op      1986292 B/op            7 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              4876010 ns/op      3477985 B/op        38422 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              8589343 ns/op      2115868 B/op            2 allocs/op
      segment.Marshal              5490932 ns/op      1991387 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              2771808 ns/op      1950993 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              6485767 ns/op     10590746 B/op           31 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Marshal              4412152 ns/op      1941945 B/op            5 allocs/op

        sonic ████████████████▉ 4.23
       sonnet ██████████▋ 2.66
        json2 █████████▋ 2.42
     jsoniter █████████▌ 2.40
      segment ████████▌ 2.13
        codec ███████▏ 1.81
        goccy █████▍ 1.36
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal GolangSource from struct
         json.Marshal              2384544 ns/op      1950638 B/op            1 allocs/op
        json2.Marshal              3045885 ns/op      1971261 B/op            1 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              1616169 ns/op      1810031 B/op            2 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              1477480 ns/op      1956591 B/op            1 allocs/op
      segment.Marshal              1535164 ns/op      1956526 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               780608 ns/op      1948577 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              4031067 ns/op     10590745 B/op           31 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal              4415026 ns/op      6493051 B/op        77031 allocs/op
     easyjson.Marshal              1699971 ns/op      1996641 B/op           73 allocs/op
       sonnet.Marshal              1645396 ns/op      1941848 B/op            5 allocs/op

        sonic ████████████▏ 3.05
        goccy ██████▍ 1.61
      segment ██████▏ 1.55
     jsoniter █████▉ 1.48
       sonnet █████▊ 1.45
     easyjson █████▌ 1.40
         json ▓▓▓▓ 1.00
        json2 ███▏ 0.78
        codec ██▎ 0.59
       ffjson ██▏ 0.54
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal StringUnicode generically
         json.Unmarshal              56253 ns/op        30304 B/op          194 allocs/op
        json2.Unmarshal              26004 ns/op        27556 B/op          179 allocs/op
     fastjson.Unmarshal               5361 ns/op        35208 B/op           76 allocs/op
     jsoniter.Unmarshal              19741 ns/op        31960 B/op          262 allocs/op
   jsonparser.Unmarshal               8649 ns/op        18624 B/op          121 allocs/op
        goccy.Unmarshal              14755 ns/op        31325 B/op          191 allocs/op
      segment.Unmarshal             105580 ns/op        43274 B/op          307 allocs/op
     simdjson.Unmarshal              32476 ns/op       194760 B/op           75 allocs/op
        gjson.Unmarshal               9488 ns/op        20272 B/op           66 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal               7060 ns/op        29865 B/op           72 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal              22143 ns/op        32850 B/op          315 allocs/op
        djson.Unmarshal              33363 ns/op        47482 B/op          133 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal              25137 ns/op        30468 B/op          193 allocs/op

   * fastjson █████████████████████████████████████████▉ 10.49
        sonic ███████████████████████████████▊ 7.97
 * jsonparser ██████████████████████████  6.50
      * gjson ███████████████████████▋ 5.93
        goccy ███████████████▏ 3.81
     jsoniter ███████████▍ 2.85
        codec ██████████▏ 2.54
       sonnet ████████▉ 2.24
        json2 ████████▋ 2.16
   * simdjson ██████▉ 1.73
      * djson ██████▋ 1.69
         json ▓▓▓▓ 1.00
      segment ██▏ 0.53
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal StringUnicode into struct
         json.Unmarshal              54169 ns/op        18296 B/op           66 allocs/op
        json2.Unmarshal              20685 ns/op        15472 B/op           49 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal              12587 ns/op        18912 B/op           72 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               8605 ns/op        18432 B/op            1 allocs/op
      segment.Unmarshal              57850 ns/op        35280 B/op          122 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               2913 ns/op        18913 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal              10062 ns/op        18824 B/op           63 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal              10054 ns/op        18449 B/op           66 allocs/op
     easyjson.Unmarshal               3835 ns/op        17776 B/op           61 allocs/op
       sonnet.Unmarshal              21366 ns/op        18384 B/op           63 allocs/op

        sonic ██████████████████████████████████████████████████████████████████████████▍ 18.60
     easyjson ████████████████████████████████████████████████████████▍ 14.12
        goccy █████████████████████████▏ 6.30
       ffjson █████████████████████▌ 5.39
        codec █████████████████████▌ 5.38
     jsoniter █████████████████▏ 4.30
        json2 ██████████▍ 2.62
       sonnet ██████████▏ 2.54
         json ▓▓▓▓ 1.00
      segment ███▋ 0.94
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal StringUnicode generically
         json.Marshal                26138 ns/op        23042 B/op          122 allocs/op
        json2.Marshal                24904 ns/op        19560 B/op           65 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal                11566 ns/op        18561 B/op            4 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                22563 ns/op        18442 B/op            1 allocs/op
      segment.Marshal                22335 ns/op        18433 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal                 3109 ns/op        18519 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                22907 ns/op        84504 B/op           13 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Marshal                24016 ns/op        18776 B/op            5 allocs/op

        sonic █████████████████████████████████▋ 8.41
     jsoniter █████████  2.26
      segment ████▋ 1.17
        goccy ████▋ 1.16
        codec ████▌ 1.14
       sonnet ████▎ 1.09
        json2 ████▏ 1.05
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal StringUnicode from struct
         json.Marshal                17270 ns/op        18433 B/op            1 allocs/op
        json2.Marshal                16141 ns/op        18435 B/op            1 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal                10541 ns/op        18441 B/op            2 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                16427 ns/op        18434 B/op            1 allocs/op
      segment.Marshal                17192 ns/op        18433 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal                 2144 ns/op        18527 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                22736 ns/op        84504 B/op           13 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal                18691 ns/op        50615 B/op           67 allocs/op
     easyjson.Marshal                15109 ns/op        19274 B/op           11 allocs/op
       sonnet.Marshal                18415 ns/op        18776 B/op            5 allocs/op

        sonic ████████████████████████████████▏ 8.06
     jsoniter ██████▌ 1.64
     easyjson ████▌ 1.14
        json2 ████▎ 1.07
        goccy ████▏ 1.05
      segment ████  1.00
         json ▓▓▓▓ 1.00
       sonnet ███▊ 0.94
       ffjson ███▋ 0.92
        codec ███  0.76
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal single record (2kb), read few keys generically
         json.Unmarshal              16811 ns/op        12488 B/op           21 allocs/op
        json2 >>> not supported <<<
     fastjson.Unmarshal               4533 ns/op        12034 B/op            7 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal               4715 ns/op        12208 B/op           17 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal               5504 ns/op        16368 B/op            9 allocs/op
      gjson-v.Unmarshal               8209 ns/op        16368 B/op            9 allocs/op
        sonic.Unmarshal               3258 ns/op         4852 B/op           12 allocs/op
      sonic-v.Unmarshal               5368 ns/op         4851 B/op           12 allocs/op
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

        sonic ████████████████████▋ 5.16
     fastjson ██████████████▊ 3.71
   jsonparser ██████████████▎ 3.57
      sonic-v ████████████▌ 3.13
        gjson ████████████▏ 3.05
      gjson-v ████████▏ 2.05
         json ▓▓▓▓ 1.00
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
     jsoniter >>> not supported <<<
        json2 >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

Unmarshal single record (2kb), read few keys into struct
         json.Unmarshal              17171 ns/op        12488 B/op           21 allocs/op
        json2.Unmarshal               6285 ns/op          128 B/op            6 allocs/op
     fastjson.Unmarshal               4520 ns/op        12080 B/op           10 allocs/op
     jsoniter.Unmarshal               7381 ns/op        14056 B/op           84 allocs/op
   jsonparser.Unmarshal               4725 ns/op        12208 B/op           17 allocs/op
        goccy.Unmarshal               5075 ns/op        16131 B/op            7 allocs/op
      segment.Unmarshal               9755 ns/op        12064 B/op           12 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal               5563 ns/op        16400 B/op           10 allocs/op
      gjson-v.Unmarshal               8198 ns/op        16400 B/op           10 allocs/op
        sonic.Unmarshal               5160 ns/op         8644 B/op           11 allocs/op
      sonic-v.Unmarshal               7094 ns/op         8653 B/op           11 allocs/op
        codec.Unmarshal               6475 ns/op        13240 B/op           15 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal              14638 ns/op        14185 B/op           36 allocs/op
     easyjson.Unmarshal              12580 ns/op        12320 B/op           16 allocs/op
       sonnet.Unmarshal               7035 ns/op        12370 B/op           21 allocs/op

   * fastjson ███████████████▏ 3.80
 * jsonparser ██████████████▌ 3.63
        goccy █████████████▌ 3.38
        sonic █████████████▎ 3.33
      * gjson ████████████▎ 3.09
        json2 ██████████▉ 2.73
        codec ██████████▌ 2.65
       sonnet █████████▊ 2.44
      sonic-v █████████▋ 2.42
     jsoniter █████████▎ 2.33
    * gjson-v ████████▍ 2.09
      segment ███████  1.76
     easyjson █████▍ 1.36
       ffjson ████▋ 1.17
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        djson >>> not supported <<<

Unmarshal single record (2kb), read all keys generically
         json.Unmarshal              22109 ns/op        29776 B/op          340 allocs/op
        json2.Unmarshal              17582 ns/op        15537 B/op          295 allocs/op
     fastjson.Unmarshal               7252 ns/op        13608 B/op           73 allocs/op
     jsoniter.Unmarshal              16745 ns/op        31681 B/op          457 allocs/op
   jsonparser.Unmarshal              13157 ns/op        14528 B/op          135 allocs/op
        goccy.Unmarshal              15850 ns/op        34515 B/op          398 allocs/op
      segment.Unmarshal              53950 ns/op       173240 B/op          475 allocs/op
     simdjson.Unmarshal              17249 ns/op       132224 B/op           82 allocs/op
        gjson.Unmarshal              12612 ns/op        18336 B/op           76 allocs/op
      gjson-v.Unmarshal              15700 ns/op        18336 B/op           76 allocs/op
        sonic.Unmarshal              12044 ns/op        27012 B/op           59 allocs/op
      sonic-v.Unmarshal              14109 ns/op        27017 B/op           59 allocs/op
        codec.Unmarshal              27885 ns/op        19456 B/op          435 allocs/op
        djson.Unmarshal              11561 ns/op        32168 B/op          183 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal              12115 ns/op        30352 B/op          341 allocs/op

     fastjson ████████████▏ 3.05
        djson ███████▋ 1.91
        sonic ███████▎ 1.84
       sonnet ███████▎ 1.82
        gjson ███████  1.75
   jsonparser ██████▋ 1.68
      sonic-v ██████▎ 1.57
      gjson-v █████▋ 1.41
        goccy █████▌ 1.39
     jsoniter █████▎ 1.32
     simdjson █████▏ 1.28
        json2 █████  1.26
         json ▓▓▓▓ 1.00
        codec ███▏ 0.79
      segment █▋ 0.41
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal single record (2kb), read all keys into struct
         json.Unmarshal              22123 ns/op        14608 B/op           80 allocs/op
        json2.Unmarshal               9718 ns/op         3104 B/op           35 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal               7777 ns/op        14384 B/op           77 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               5844 ns/op        16138 B/op            7 allocs/op
      segment.Unmarshal              16147 ns/op        13680 B/op           73 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               5512 ns/op         9227 B/op            9 allocs/op
      sonic-v.Unmarshal               7630 ns/op         9234 B/op            9 allocs/op
        codec.Unmarshal               9464 ns/op        14536 B/op           72 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal              11124 ns/op        16786 B/op          113 allocs/op
     easyjson.Unmarshal               8930 ns/op        15888 B/op           94 allocs/op
       sonnet.Unmarshal               9096 ns/op        17070 B/op          112 allocs/op

        sonic ████████████████  4.01
        goccy ███████████████▏ 3.79
      sonic-v ███████████▌ 2.90
     jsoniter ███████████▍ 2.84
     easyjson █████████▉ 2.48
       sonnet █████████▋ 2.43
        codec █████████▎ 2.34
        json2 █████████  2.28
       ffjson ███████▉ 1.99
      segment █████▍ 1.37
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys generically
         json.Unmarshal          462480203 ns/op     21756058 B/op       394052 allocs/op
        json2 >>> not supported <<<
     fastjson.Unmarshal           82742770 ns/op      2533390 B/op        39470 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal           60570679 ns/op      4614566 B/op       177366 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal           55404783 ns/op    623113625 B/op        39635 allocs/op
        gjson.Unmarshal           82261617 ns/op    105924736 B/op        39406 allocs/op
      gjson-v.Unmarshal          167015961 ns/op    105924736 B/op        39406 allocs/op
        sonic.Unmarshal           33074794 ns/op      2843970 B/op        78811 allocs/op
      sonic-v.Unmarshal          107984365 ns/op      2841749 B/op        78811 allocs/op
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

        sonic ███████████████████████████████████████████████████████▉ 13.98
     simdjson █████████████████████████████████▍ 8.35
   jsonparser ██████████████████████████████▌ 7.64
        gjson ██████████████████████▍ 5.62
     fastjson ██████████████████████▎ 5.59
      sonic-v █████████████████▏ 4.28
      gjson-v ███████████  2.77
         json ▓▓▓▓ 1.00
        json2 >>> not supported <<<
      segment >>> not supported <<<
     jsoniter >>> not supported <<<
        goccy >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys into struct
         json.Unmarshal          440616257 ns/op     21755674 B/op       394051 allocs/op
        json2.Unmarshal          202403672 ns/op        16782 B/op           64 allocs/op
     fastjson.Unmarshal           79826813 ns/op      2533390 B/op        39470 allocs/op
     jsoniter.Unmarshal          199038416 ns/op     89012049 B/op      3979908 allocs/op
   jsonparser.Unmarshal           62052158 ns/op      4614566 B/op       177366 allocs/op
        goccy.Unmarshal           82650913 ns/op    105933909 B/op        39416 allocs/op
      segment.Unmarshal          213258620 ns/op      2526041 B/op        39406 allocs/op
     simdjson.Unmarshal           58100800 ns/op    616506667 B/op        39632 allocs/op
        gjson.Unmarshal           85115398 ns/op    105924744 B/op        39406 allocs/op
      gjson-v.Unmarshal          170529040 ns/op    105924752 B/op        39406 allocs/op
        sonic.Unmarshal          105638249 ns/op    107296600 B/op        78814 allocs/op
      sonic-v.Unmarshal          180855338 ns/op    107487698 B/op        78819 allocs/op
        codec.Unmarshal          111876865 ns/op     48866307 B/op       157621 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal          261120890 ns/op     67792052 B/op       591084 allocs/op
     easyjson.Unmarshal          360334074 ns/op      2527848 B/op        39411 allocs/op
       sonnet.Unmarshal          140695046 ns/op      7649930 B/op        88672 allocs/op

   * simdjson ██████████████████████████████▎ 7.58
 * jsonparser ████████████████████████████▍ 7.10
   * fastjson ██████████████████████  5.52
        goccy █████████████████████▎ 5.33
      * gjson ████████████████████▋ 5.18
        sonic ████████████████▋ 4.17
        codec ███████████████▊ 3.94
       sonnet ████████████▌ 3.13
    * gjson-v ██████████▎ 2.58
      sonic-v █████████▋ 2.44
     jsoniter ████████▊ 2.21
        json2 ████████▋ 2.18
      segment ████████▎ 2.07
       ffjson ██████▋ 1.69
     easyjson ████▉ 1.22
         json ▓▓▓▓ 1.00
        djson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys generically
         json.Unmarshal          754451642 ns/op    743935208 B/op     14315094 allocs/op
        json2.Unmarshal          652682656 ns/op    692712324 B/op     11948210 allocs/op
     fastjson.Unmarshal          206168785 ns/op     66540390 B/op      2876746 allocs/op
     jsoniter.Unmarshal          629315000 ns/op    819030712 B/op     19201308 allocs/op
   jsonparser.Unmarshal          497189904 ns/op    105248237 B/op      5437891 allocs/op
        goccy.Unmarshal          582491652 ns/op    876422088 B/op     16797521 allocs/op
      segment.Unmarshal         2365200180 ns/op   6770681184 B/op     19978341 allocs/op
     simdjson.Unmarshal          169647308 ns/op    688361966 B/op      2758577 allocs/op
        gjson.Unmarshal          452640619 ns/op    196083413 B/op      2994781 allocs/op
      gjson-v.Unmarshal          558415469 ns/op    196083392 B/op      2994781 allocs/op
        sonic.Unmarshal          386595946 ns/op    889762610 B/op      2285517 allocs/op
      sonic-v.Unmarshal          472349945 ns/op    889707816 B/op      2285514 allocs/op
        codec.Unmarshal         1073348145 ns/op    312883216 B/op     18323501 allocs/op
        djson.Unmarshal          386343714 ns/op    770795954 B/op      7537406 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal          434009834 ns/op    760643714 B/op     14245930 allocs/op

     simdjson █████████████████▊ 4.45
     fastjson ██████████████▋ 3.66
        djson ███████▊ 1.95
        sonic ███████▊ 1.95
       sonnet ██████▉ 1.74
        gjson ██████▋ 1.67
      sonic-v ██████▍ 1.60
   jsonparser ██████  1.52
      gjson-v █████▍ 1.35
        goccy █████▏ 1.30
     jsoniter ████▊ 1.20
        json2 ████▌ 1.16
         json ▓▓▓▓ 1.00
        codec ██▊ 0.70
      segment █▎ 0.32
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys into struct
         json.Unmarshal          660889293 ns/op    110920456 B/op      3113019 allocs/op
        json2.Unmarshal          383399684 ns/op    129761162 B/op      1801316 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          220380775 ns/op     97208107 B/op      2955386 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal          132661459 ns/op    105993936 B/op        39508 allocs/op
      segment.Unmarshal          553316389 ns/op     68879048 B/op      2797773 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal          141257183 ns/op    133160446 B/op       157628 allocs/op
      sonic-v.Unmarshal          217593987 ns/op    133196484 B/op       157634 allocs/op
        codec.Unmarshal          302149239 ns/op    102057732 B/op      2718954 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal          377911694 ns/op    217787789 B/op      4689215 allocs/op
     easyjson.Unmarshal          262566180 ns/op    170469542 B/op      3822296 allocs/op
       sonnet.Unmarshal          291645685 ns/op    212511236 B/op      4817275 allocs/op

        goccy ███████████████████▉ 4.98
        sonic ██████████████████▋ 4.68
      sonic-v ████████████▏ 3.04
     jsoniter ███████████▉ 3.00
     easyjson ██████████  2.52
       sonnet █████████  2.27
        codec ████████▋ 2.19
       ffjson ██████▉ 1.75
        json2 ██████▉ 1.72
      segment ████▊ 1.19
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys generically
         json.Unmarshal        20514828640 ns/op    116288976 B/op      2105576 allocs/op
        json2 >>> not supported <<<
     fastjson.Unmarshal         3322266875 ns/op     14463336 B/op       218594 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal         2057975153 ns/op     24696400 B/op       947400 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal         1128456302 ns/op   2317987480 B/op       217335 allocs/op
        gjson.Unmarshal         3183278889 ns/op   5740687584 B/op       210562 allocs/op
      gjson-v.Unmarshal         7152873085 ns/op   5740687680 B/op       210563 allocs/op
        sonic.Unmarshal          753822448 ns/op     15285540 B/op       421122 allocs/op
      sonic-v.Unmarshal         4324609816 ns/op     15308568 B/op       421125 allocs/op
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

        sonic ████████████████████████████████████████████████████████████████████████████████████████████████████████████▊ 27.21
     simdjson ████████████████████████████████████████████████████████████████████████▋ 18.18
   jsonparser ███████████████████████████████████████▊ 9.97
        gjson █████████████████████████▊ 6.44
     fastjson ████████████████████████▋ 6.17
      sonic-v ██████████████████▉ 4.74
      gjson-v ███████████▍ 2.87
         json ▓▓▓▓ 1.00
        json2 >>> not supported <<<
      segment >>> not supported <<<
     jsoniter >>> not supported <<<
        goccy >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys into struct
         json.Unmarshal        20317809829 ns/op    116288976 B/op      2105576 allocs/op
        json2.Unmarshal         9032656253 ns/op        21080 B/op           72 allocs/op
     fastjson.Unmarshal         3303466255 ns/op     14463336 B/op       218594 allocs/op
     jsoniter.Unmarshal         9052638753 ns/op   4649567072 B/op    210767642 allocs/op
   jsonparser.Unmarshal         2060464169 ns/op     24696400 B/op       947400 allocs/op
        goccy.Unmarshal         3509143094 ns/op   5741056368 B/op       210995 allocs/op
      segment.Unmarshal         8236836487 ns/op     13537216 B/op       210563 allocs/op
     simdjson.Unmarshal         1437336521 ns/op   2218407648 B/op       217173 allocs/op
        gjson.Unmarshal         3172437334 ns/op   5740687488 B/op       210561 allocs/op
      gjson-v.Unmarshal         7224891728 ns/op   5740687584 B/op       210562 allocs/op
        sonic.Unmarshal         4293992892 ns/op   5752825000 B/op       421255 allocs/op
      sonic-v.Unmarshal         7776198800 ns/op   5753842744 B/op       421385 allocs/op
        codec.Unmarshal         4256501488 ns/op    261152232 B/op       842234 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal        10682830428 ns/op    362267224 B/op      3158419 allocs/op
     easyjson.Unmarshal        16788012574 ns/op     13542952 B/op       210583 allocs/op
       sonnet.Unmarshal         5909597907 ns/op     40935856 B/op       474066 allocs/op

   * simdjson ████████████████████████████████████████████████████████▌ 14.14
 * jsonparser ███████████████████████████████████████▍ 9.86
      * gjson █████████████████████████▌ 6.40
   * fastjson ████████████████████████▌ 6.15
        goccy ███████████████████████▏ 5.79
        codec ███████████████████  4.77
        sonic ██████████████████▉ 4.73
       sonnet █████████████▊ 3.44
    * gjson-v ███████████▏ 2.81
      sonic-v ██████████▍ 2.61
      segment █████████▊ 2.47
        json2 ████████▉ 2.25
     jsoniter ████████▉ 2.24
       ffjson ███████▌ 1.90
     easyjson ████▊ 1.21
         json ▓▓▓▓ 1.00
        djson >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys generically
         json.Unmarshal        35174089142 ns/op  37173091824 B/op    701323233 allocs/op
        json2.Unmarshal        30087611531 ns/op  35258025952 B/op    579260512 allocs/op
     fastjson.Unmarshal         9550894046 ns/op   3342949112 B/op    142344575 allocs/op
     jsoniter.Unmarshal        30559321568 ns/op  41727145640 B/op    960517839 allocs/op
   jsonparser.Unmarshal        24024196128 ns/op   5294191672 B/op    273513554 allocs/op
        goccy.Unmarshal        27500553268 ns/op  44759951792 B/op    833135290 allocs/op
      segment.Unmarshal       116244224334 ns/op 341651379576 B/op    995513943 allocs/op
     simdjson.Unmarshal         6755738193 ns/op  27136803824 B/op    137924682 allocs/op
        gjson.Unmarshal        23206073684 ns/op  10406631856 B/op    146758239 allocs/op
      gjson-v.Unmarshal        27525123774 ns/op  10406631296 B/op    146758235 allocs/op
        sonic.Unmarshal        18515060510 ns/op  44947706216 B/op    111175205 allocs/op
      sonic-v.Unmarshal        22230097055 ns/op  44946642936 B/op    111175197 allocs/op
        codec.Unmarshal        52829223324 ns/op  13510675472 B/op    905186189 allocs/op
        djson.Unmarshal        18041584974 ns/op  39791204944 B/op    375169017 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal        20571237096 ns/op  35929463440 B/op    699112538 allocs/op

     simdjson ████████████████████▊ 5.21
     fastjson ██████████████▋ 3.68
        djson ███████▊ 1.95
        sonic ███████▌ 1.90
       sonnet ██████▊ 1.71
      sonic-v ██████▎ 1.58
        gjson ██████  1.52
   jsonparser █████▊ 1.46
        goccy █████  1.28
      gjson-v █████  1.28
        json2 ████▋ 1.17
     jsoniter ████▌ 1.15
         json ▓▓▓▓ 1.00
        codec ██▋ 0.67
      segment █▏ 0.30
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys into struct
         json.Unmarshal        31258995701 ns/op   4443559312 B/op    137915228 allocs/op
        json2.Unmarshal        19497175880 ns/op   6894666336 B/op     94346704 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal        10237784788 ns/op   5003334632 B/op    150338165 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         5797626425 ns/op   5744199560 B/op       215210 allocs/op
      segment.Unmarshal        25441911438 ns/op   3491108088 B/op    141915737 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal         6222656282 ns/op   7132244760 B/op      4632758 allocs/op
      sonic-v.Unmarshal         9927829760 ns/op   7138095576 B/op      4632973 allocs/op
        codec.Unmarshal        13393824473 ns/op   3062217968 B/op    133914572 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal        16833904321 ns/op  10023436984 B/op    215400711 allocs/op
     easyjson.Unmarshal        11714213495 ns/op   8946289824 B/op    198555754 allocs/op
       sonnet.Unmarshal        12764908711 ns/op   8218829088 B/op    236144932 allocs/op

        goccy █████████████████████▌ 5.39
        sonic ████████████████████  5.02
      sonic-v ████████████▌ 3.15
     jsoniter ████████████▏ 3.05
     easyjson ██████████▋ 2.67
       sonnet █████████▊ 2.45
        codec █████████▎ 2.33
       ffjson ███████▍ 1.86
        json2 ██████▍ 1.60
      segment ████▉ 1.23
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal custom data through an object builder
         json.Marshal                  970 ns/op          688 B/op           17 allocs/op
        json2 >>> not supported <<<
     fastjson.Marshal                  541 ns/op          679 B/op            6 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Marshal                 1653 ns/op         3024 B/op           27 allocs/op
      gjson-v >>> not supported <<<
        sonic.Marshal                  685 ns/op         1803 B/op           14 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

     fastjson ███████▏ 1.79
        sonic █████▋ 1.42
         json ▓▓▓▓ 1.00
        gjson ██▎ 0.59
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
     jsoniter >>> not supported <<<
      gjson-v >>> not supported <<<
        json2 >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

Validate []byte
         json.Validate                6280 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
     fastjson.Validate                2659 ns/op            0 B/op            0 allocs/op
     jsoniter.Validate                4537 ns/op         2184 B/op          100 allocs/op
   jsonparser >>> not supported <<<
        goccy.Validate               15814 ns/op        27584 B/op          463 allocs/op
      segment.Validate                4277 ns/op            0 B/op            0 allocs/op
     simdjson >>> not supported <<<
        gjson.Validate                2473 ns/op            0 B/op            0 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                2036 ns/op            0 B/op            0 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Validate                3885 ns/op            0 B/op            0 allocs/op

        sonic ████████████▎ 3.08
        gjson ██████████▏ 2.54
     fastjson █████████▍ 2.36
       sonnet ██████▍ 1.62
      segment █████▊ 1.47
     jsoniter █████▌ 1.38
         json ▓▓▓▓ 1.00
        goccy █▌ 0.40
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
        json2 >>> not supported <<<

Validate string
         json.Validate                6226 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
     fastjson.Validate                3133 ns/op         4096 B/op            1 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Validate                2803 ns/op         4096 B/op            1 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                2414 ns/op         4097 B/op            1 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

        sonic ██████████▎ 2.58
        gjson ████████▉ 2.22
     fastjson ███████▉ 1.99
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
     jsoniter >>> not supported <<<
      gjson-v >>> not supported <<<
        json2 >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

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
Unmarshal CanadaGeometry generically
         json.Unmarshal            3028476 ns/op      1339387 B/op        52215 allocs/op
        json2.Unmarshal            2316036 ns/op      1011496 B/op        37898 allocs/op
     fastjson.Unmarshal            2645426 ns/op      8725352 B/op        30278 allocs/op
     jsoniter.Unmarshal            3585789 ns/op      1961487 B/op        84971 allocs/op
   jsonparser.Unmarshal            2205988 ns/op          152 B/op           12 allocs/op
        goccy.Unmarshal            2462537 ns/op      1470356 B/op        59189 allocs/op
      segment.Unmarshal            4705931 ns/op      2119753 B/op        51680 allocs/op
     simdjson.Unmarshal       >>> Unsupported platform <<<
        gjson.Unmarshal            3670229 ns/op       393072 B/op        14314 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal            3141429 ns/op      2943315 B/op        52232 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            3303187 ns/op      1765859 B/op        37408 allocs/op
        djson.Unmarshal            1852211 ns/op      1289905 B/op        37891 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal            1357541 ns/op      1011609 B/op        37903 allocs/op

       sonnet ████████▉ 2.23
      * djson ██████▌ 1.64
 * jsonparser █████▍ 1.37
        json2 █████▏ 1.31
        goccy ████▉ 1.23
   * fastjson ████▌ 1.14
         json ▓▓▓▓ 1.00
        sonic ███▊ 0.96
        codec ███▋ 0.92
     jsoniter ███▍ 0.84
      * gjson ███▎ 0.83
      segment ██▌ 0.64
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> Unsupported platform <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal CanadaGeometry into struct
         json.Unmarshal            2547508 ns/op          472 B/op           13 allocs/op
        json2.Unmarshal            1554266 ns/op       347966 B/op         1626 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            1892327 ns/op       598762 B/op        25122 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal             962134 ns/op       278635 B/op            1 allocs/op
      segment.Unmarshal             981722 ns/op           48 B/op            4 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal            2663387 ns/op      1604401 B/op           30 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            1669278 ns/op         1224 B/op            7 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal            3418503 ns/op       536875 B/op           25 allocs/op
     easyjson.Unmarshal            1188248 ns/op       330368 B/op          753 allocs/op
       sonnet.Unmarshal             615821 ns/op       143141 B/op          967 allocs/op

       sonnet ████████████████▌ 4.14
        goccy ██████████▌ 2.65
      segment ██████████▍ 2.59
     easyjson ████████▌ 2.14
        json2 ██████▌ 1.64
        codec ██████  1.53
     jsoniter █████▍ 1.35
         json ▓▓▓▓ 1.00
        sonic ███▊ 0.96
       ffjson ██▉ 0.75
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal CanadaGeometry generically
         json.Marshal              1584742 ns/op       280519 B/op           21 allocs/op
        json2.Marshal              1246393 ns/op       279987 B/op            7 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               621609 ns/op       172918 B/op           13 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              1508664 ns/op       282039 B/op            3 allocs/op
      segment.Marshal              1172066 ns/op       279712 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              1584264 ns/op       279179 B/op           22 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              1290783 ns/op      1186328 B/op           22 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Marshal              1038781 ns/op       278898 B/op            5 allocs/op

     jsoniter ██████████▏ 2.55
       sonnet ██████  1.53
      segment █████▍ 1.35
        json2 █████  1.27
        codec ████▉ 1.23
        goccy ████▏ 1.05
        sonic ████  1.00
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal CanadaGeometry from struct
         json.Marshal              1145192 ns/op       279548 B/op            1 allocs/op
        json2.Marshal              1180193 ns/op       279710 B/op            1 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               269037 ns/op       172199 B/op            2 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              1102638 ns/op       279636 B/op            1 allocs/op
      segment.Marshal              1025058 ns/op       278539 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              1150283 ns/op       278587 B/op            2 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              1239434 ns/op      1186328 B/op           22 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal              2092942 ns/op      1059612 B/op        14391 allocs/op
     easyjson.Marshal              1014252 ns/op       280350 B/op           20 allocs/op
       sonnet.Marshal              1079498 ns/op       278883 B/op            5 allocs/op

     jsoniter █████████████████  4.26
     easyjson ████▌ 1.13
      segment ████▍ 1.12
       sonnet ████▏ 1.06
        goccy ████▏ 1.04
         json ▓▓▓▓ 1.00
        sonic ███▉ 1.00
        json2 ███▉ 0.97
        codec ███▋ 0.92
       ffjson ██▏ 0.55
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal CITMCatalog generically
         json.Unmarshal           11534838 ns/op      5123660 B/op        95372 allocs/op
        json2.Unmarshal            4533186 ns/op      4958490 B/op        80513 allocs/op
     fastjson.Unmarshal            4669186 ns/op     19991452 B/op        45701 allocs/op
     jsoniter.Unmarshal            5644672 ns/op      5573007 B/op       118756 allocs/op
   jsonparser.Unmarshal           11374455 ns/op       145320 B/op        15863 allocs/op
        goccy.Unmarshal            5528200 ns/op      7416061 B/op       123577 allocs/op
      segment.Unmarshal           20706205 ns/op     55526632 B/op       123714 allocs/op
     simdjson.Unmarshal       >>> Unsupported platform <<<
        gjson.Unmarshal            7908014 ns/op      1855520 B/op        15131 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal           12938835 ns/op     12773356 B/op        95391 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            9344647 ns/op      6195119 B/op       159476 allocs/op
        djson.Unmarshal            4915077 ns/op      6437595 B/op        54480 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal            3502266 ns/op      4968494 B/op        80980 allocs/op

       sonnet █████████████▏ 3.29
        json2 ██████████▏ 2.54
   * fastjson █████████▉ 2.47
      * djson █████████▍ 2.35
        goccy ████████▎ 2.09
     jsoniter ████████▏ 2.04
      * gjson █████▊ 1.46
        codec ████▉ 1.23
 * jsonparser ████  1.01
         json ▓▓▓▓ 1.00
        sonic ███▌ 0.89
      segment ██▏ 0.56
     simdjson >>> Unsupported platform <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal CITMCatalog into struct
         json.Unmarshal           10491377 ns/op       253352 B/op        10950 allocs/op
        json2.Unmarshal            2673335 ns/op       888435 B/op         6293 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            1947619 ns/op        72333 B/op         2535 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal            1408529 ns/op      1781218 B/op         2226 allocs/op
      segment.Unmarshal            1621568 ns/op        52384 B/op         1354 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal           11745355 ns/op      7902978 B/op        10969 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            2395920 ns/op        47880 B/op         1036 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal       >>> ffjson error: (*errors.errorString)cannot unmarshal tok:string into Go value for int64 offset=40 line=3 char=19 <<<
     easyjson.Unmarshal            2770633 ns/op       896595 B/op         5704 allocs/op
       sonnet.Unmarshal            2434426 ns/op       612941 B/op        13348 allocs/op

        goccy █████████████████████████████▊ 7.45
      segment █████████████████████████▉ 6.47
     jsoniter █████████████████████▌ 5.39
        codec █████████████████▌ 4.38
       sonnet █████████████████▏ 4.31
        json2 ███████████████▋ 3.92
     easyjson ███████████████▏ 3.79
         json ▓▓▓▓ 1.00
        sonic ███▌ 0.89
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> ffjson error: (*errors.errorString)cannot unmarshal tok:string into Go value for int64 offset=40 line=3 char=19 <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal CITMCatalog generically
         json.Marshal              4474942 ns/op      2388625 B/op        62674 allocs/op
        json2.Marshal              1880509 ns/op       512224 B/op           16 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              2835216 ns/op      1907899 B/op        32812 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              2885377 ns/op       661856 B/op          325 allocs/op
      segment.Marshal              1890465 ns/op       507960 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              4486555 ns/op      2392616 B/op        62675 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              2552594 ns/op      2538010 B/op           25 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Marshal              1553461 ns/op       508310 B/op            5 allocs/op

       sonnet ███████████▌ 2.88
        json2 █████████▌ 2.38
      segment █████████▍ 2.37
        codec ███████  1.75
     jsoniter ██████▎ 1.58
        goccy ██████▏ 1.55
         json ▓▓▓▓ 1.00
        sonic ███▉ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal CITMCatalog from struct
         json.Marshal               717661 ns/op       556675 B/op          890 allocs/op
        json2.Marshal              1044817 ns/op       513251 B/op          133 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               695943 ns/op       510652 B/op           22 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal               425700 ns/op       521605 B/op            3 allocs/op
      segment.Marshal               495866 ns/op       548047 B/op          616 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               727676 ns/op       556723 B/op          891 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              1172659 ns/op      2538009 B/op           25 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal              1093837 ns/op       909240 B/op        14181 allocs/op
     easyjson.Marshal               465827 ns/op       524039 B/op         1125 allocs/op
       sonnet.Marshal               481918 ns/op       508264 B/op            5 allocs/op

        goccy ██████▋ 1.69
     easyjson ██████▏ 1.54
       sonnet █████▉ 1.49
      segment █████▊ 1.45
     jsoniter ████  1.03
         json ▓▓▓▓ 1.00
        sonic ███▉ 0.99
        json2 ██▋ 0.69
       ffjson ██▌ 0.66
        codec ██▍ 0.61
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal SyntheaFHIR generically
         json.Unmarshal           14846661 ns/op      7430382 B/op       134384 allocs/op
        json2.Unmarshal            6973265 ns/op      6938933 B/op       115833 allocs/op
     fastjson.Unmarshal            6682266 ns/op     26046535 B/op        66098 allocs/op
     jsoniter.Unmarshal            8407300 ns/op      8196621 B/op       181425 allocs/op
   jsonparser.Unmarshal           14294803 ns/op      1108792 B/op        49604 allocs/op
        goccy.Unmarshal            7967740 ns/op      9838457 B/op       159460 allocs/op
      segment.Unmarshal           27946400 ns/op     77571247 B/op       193338 allocs/op
     simdjson.Unmarshal       >>> Unsupported platform <<<
        gjson.Unmarshal           10748612 ns/op      2432545 B/op        26821 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal           16234866 ns/op     15653669 B/op       134404 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal           14968979 ns/op      9669829 B/op       247797 allocs/op
        djson.Unmarshal            6306420 ns/op      8374270 B/op        68393 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal            5814157 ns/op      7410927 B/op       133127 allocs/op

       sonnet ██████████▏ 2.55
      * djson █████████▍ 2.35
   * fastjson ████████▉ 2.22
        json2 ████████▌ 2.13
        goccy ███████▍ 1.86
     jsoniter ███████  1.77
      * gjson █████▌ 1.38
 * jsonparser ████▏ 1.04
         json ▓▓▓▓ 1.00
        codec ███▉ 0.99
        sonic ███▋ 0.91
      segment ██▏ 0.53
     simdjson >>> Unsupported platform <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal SyntheaFHIR into struct
         json.Unmarshal           13986089 ns/op       625408 B/op        21883 allocs/op
        json2.Unmarshal            4703384 ns/op      3103203 B/op        14550 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            3263476 ns/op       825712 B/op        29439 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal            2347565 ns/op      2152099 B/op         2512 allocs/op
      segment.Unmarshal            2983134 ns/op       623760 B/op        21832 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal           15285082 ns/op      8848586 B/op        21903 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            3931695 ns/op       623400 B/op        21826 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal           23151530 ns/op      4343884 B/op        21905 allocs/op
     easyjson.Unmarshal            4231476 ns/op      3531507 B/op        29214 allocs/op
       sonnet.Unmarshal            4059131 ns/op      3442515 B/op        33100 allocs/op

        goccy ███████████████████████▊ 5.96
      segment ██████████████████▊ 4.69
     jsoniter █████████████████▏ 4.29
        codec ██████████████▏ 3.56
       sonnet █████████████▊ 3.45
     easyjson █████████████▏ 3.31
        json2 ███████████▉ 2.97
         json ▓▓▓▓ 1.00
        sonic ███▋ 0.92
       ffjson ██▍ 0.60
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal SyntheaFHIR generically
         json.Marshal              7512386 ns/op      4145156 B/op        95282 allocs/op
        json2.Marshal              3034593 ns/op      1164118 B/op            8 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              4411762 ns/op      3096034 B/op        45292 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              5311418 ns/op      2511568 B/op         6476 allocs/op
      segment.Marshal              3211111 ns/op      1165231 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              7361808 ns/op      4117958 B/op        95283 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              3150172 ns/op      6642203 B/op           29 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Marshal              2600150 ns/op      1147256 B/op            5 allocs/op

       sonnet ███████████▌ 2.89
        json2 █████████▉ 2.48
        codec █████████▌ 2.38
      segment █████████▎ 2.34
     jsoniter ██████▊ 1.70
        goccy █████▋ 1.41
        sonic ████  1.02
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal SyntheaFHIR from struct
         json.Marshal              7160941 ns/op      3999591 B/op        18009 allocs/op
        json2.Marshal              6802592 ns/op      3237565 B/op         2569 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              5317643 ns/op      3970695 B/op        17963 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              3469655 ns/op      3945283 B/op        17940 allocs/op
      segment.Marshal              2621452 ns/op      3272290 B/op         5338 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              6950456 ns/op      3995933 B/op        18010 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              6560657 ns/op     16792092 B/op           33 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal              4170733 ns/op      9523767 B/op        45597 allocs/op
     easyjson.Marshal              2980309 ns/op      3979286 B/op        15520 allocs/op
       sonnet.Marshal              4049575 ns/op      3944829 B/op        17940 allocs/op

      segment ██████████▉ 2.73
     easyjson █████████▌ 2.40
        goccy ████████▎ 2.06
       sonnet ███████  1.77
       ffjson ██████▊ 1.72
     jsoniter █████▍ 1.35
        codec ████▎ 1.09
        json2 ████▏ 1.05
        sonic ████  1.03
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal TwitterStatus generically
         json.Unmarshal            4452721 ns/op      2152500 B/op        31262 allocs/op
        json2.Unmarshal            2557140 ns/op      2074205 B/op        27236 allocs/op
     fastjson.Unmarshal            1431299 ns/op      5341210 B/op        11130 allocs/op
     jsoniter.Unmarshal            2731736 ns/op      2427349 B/op        45042 allocs/op
   jsonparser.Unmarshal            2961659 ns/op       300312 B/op        10184 allocs/op
        goccy.Unmarshal            2476262 ns/op      2756223 B/op        39697 allocs/op
      segment.Unmarshal            5879230 ns/op      7060651 B/op        39237 allocs/op
     simdjson.Unmarshal       >>> Unsupported platform <<<
        gjson.Unmarshal            2447384 ns/op       822360 B/op         6884 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal            4927465 ns/op      5525474 B/op        31279 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            4324698 ns/op      2576546 B/op        60984 allocs/op
        djson.Unmarshal            2152758 ns/op      2470346 B/op        12678 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal            2029763 ns/op      2085995 B/op        29440 allocs/op

   * fastjson ████████████▍ 3.11
       sonnet ████████▊ 2.19
      * djson ████████▎ 2.07
      * gjson ███████▎ 1.82
        goccy ███████▏ 1.80
        json2 ██████▉ 1.74
     jsoniter ██████▌ 1.63
 * jsonparser ██████  1.50
        codec ████  1.03
         json ▓▓▓▓ 1.00
        sonic ███▌ 0.90
      segment ███  0.76
     simdjson >>> Unsupported platform <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal TwitterStatus into struct
         json.Unmarshal            4080642 ns/op       304612 B/op         6486 allocs/op
        json2.Unmarshal            1472714 ns/op       306184 B/op         2972 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            1088199 ns/op       339686 B/op         6293 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal             773768 ns/op       646850 B/op          238 allocs/op
      segment.Unmarshal            1227859 ns/op       369712 B/op         5341 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal            4512110 ns/op      3678146 B/op         6505 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            1178833 ns/op       215944 B/op         4530 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal            2329277 ns/op      1052192 B/op         9494 allocs/op
     easyjson.Unmarshal             957630 ns/op       324336 B/op         4907 allocs/op
       sonnet.Unmarshal            1188251 ns/op       373612 B/op         6257 allocs/op

        goccy █████████████████████  5.27
     easyjson █████████████████  4.26
     jsoniter ██████████████▉ 3.75
        codec █████████████▊ 3.46
       sonnet █████████████▋ 3.43
      segment █████████████▎ 3.32
        json2 ███████████  2.77
       ffjson ███████  1.75
         json ▓▓▓▓ 1.00
        sonic ███▌ 0.90
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal TwitterStatus generically
         json.Marshal              2836205 ns/op      1490542 B/op        27955 allocs/op
        json2.Marshal              1017655 ns/op       469314 B/op            7 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              1114337 ns/op       631147 B/op         3793 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              2223944 ns/op      1043498 B/op          546 allocs/op
      segment.Marshal              1536556 ns/op       478499 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              2821800 ns/op      1482388 B/op        27956 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              1186464 ns/op      2538008 B/op           25 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Marshal              1326179 ns/op       475514 B/op            5 allocs/op

        json2 ███████████▏ 2.79
     jsoniter ██████████▏ 2.55
        codec █████████▌ 2.39
       sonnet ████████▌ 2.14
      segment ███████▍ 1.85
        goccy █████  1.28
        sonic ████  1.01
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal TwitterStatus from struct
         json.Marshal               557420 ns/op       513490 B/op          141 allocs/op
        json2.Marshal               746795 ns/op       501848 B/op           59 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               585360 ns/op       503448 B/op           52 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal               378943 ns/op       508888 B/op            6 allocs/op
      segment.Marshal               405900 ns/op       512620 B/op          121 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               560270 ns/op       505344 B/op          142 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              1016153 ns/op      2540384 B/op           29 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal               648015 ns/op       798515 B/op         6117 allocs/op
     easyjson.Marshal               440687 ns/op       519853 B/op          953 allocs/op
       sonnet.Marshal               419051 ns/op       508266 B/op            5 allocs/op

        goccy █████▉ 1.47
      segment █████▍ 1.37
       sonnet █████▎ 1.33
     easyjson █████  1.26
         json ▓▓▓▓ 1.00
        sonic ███▉ 0.99
     jsoniter ███▊ 0.95
       ffjson ███▍ 0.86
        json2 ██▉ 0.75
        codec ██▏ 0.55
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal GolangSource generically
         json.Unmarshal           16211597 ns/op      7916353 B/op       271264 allocs/op
        json2.Unmarshal           10664056 ns/op      6989954 B/op       218839 allocs/op
     fastjson.Unmarshal           11642933 ns/op     50767000 B/op       130330 allocs/op
     jsoniter.Unmarshal           13176505 ns/op      9233496 B/op       346658 allocs/op
   jsonparser.Unmarshal           41118407 ns/op       925272 B/op        64023 allocs/op
        goccy.Unmarshal           13222840 ns/op     10871715 B/op       320839 allocs/op
      segment.Unmarshal           56876995 ns/op     67095368 B/op       334022 allocs/op
     simdjson.Unmarshal       >>> Unsupported platform <<<
        gjson.Unmarshal           46907648 ns/op      2658512 B/op        76819 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal           17441768 ns/op     15992865 B/op       271286 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal           24971700 ns/op     10502698 B/op       436539 allocs/op
        djson.Unmarshal            6657840 ns/op      7822417 B/op       117489 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal            6897581 ns/op      7011914 B/op       219940 allocs/op

      * djson █████████▋ 2.43
       sonnet █████████▍ 2.35
        json2 ██████  1.52
   * fastjson █████▌ 1.39
     jsoniter ████▉ 1.23
        goccy ████▉ 1.23
         json ▓▓▓▓ 1.00
        sonic ███▋ 0.93
        codec ██▌ 0.65
 * jsonparser █▌ 0.39
      * gjson █▍ 0.35
      segment █▏ 0.29
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> Unsupported platform <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal GolangSource into struct
         json.Unmarshal           15610452 ns/op       704431 B/op        24927 allocs/op
        json2.Unmarshal            7128316 ns/op      3236283 B/op        13941 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            4885073 ns/op       994058 B/op        37077 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal            2950107 ns/op      1945795 B/op            1 allocs/op
      segment.Unmarshal            2978547 ns/op       411672 B/op        12806 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal           16919528 ns/op      8780937 B/op        24949 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            6988188 ns/op       412848 B/op        12809 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal            9374540 ns/op      6276464 B/op        52765 allocs/op
     easyjson.Unmarshal            5880207 ns/op      4282425 B/op        27846 allocs/op
       sonnet.Unmarshal            4254730 ns/op      1813819 B/op        26312 allocs/op

        goccy █████████████████████▏ 5.29
      segment ████████████████████▉ 5.24
       sonnet ██████████████▋ 3.67
     jsoniter ████████████▊ 3.20
     easyjson ██████████▌ 2.65
        codec ████████▉ 2.23
        json2 ████████▊ 2.19
       ffjson ██████▋ 1.67
         json ▓▓▓▓ 1.00
        sonic ███▋ 0.92
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal GolangSource generically
         json.Marshal             14920102 ns/op      8551541 B/op       192096 allocs/op
        json2.Marshal              6263663 ns/op      1998077 B/op            7 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              6625623 ns/op      3493225 B/op        38422 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal             10848215 ns/op      2511535 B/op            3 allocs/op
      segment.Marshal              7328991 ns/op      1941573 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal             14894631 ns/op      8556744 B/op       192097 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              9149828 ns/op     10590745 B/op           31 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Marshal              6085491 ns/op      1941901 B/op            5 allocs/op

       sonnet █████████▊ 2.45
        json2 █████████▌ 2.38
     jsoniter █████████  2.25
      segment ████████▏ 2.04
        codec ██████▌ 1.63
        goccy █████▌ 1.38
        sonic ████  1.00
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal GolangSource from struct
         json.Marshal              3206228 ns/op      1952807 B/op            1 allocs/op
        json2.Marshal              3965679 ns/op      1941527 B/op            1 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              2490902 ns/op      1794080 B/op            2 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              2107295 ns/op      1941529 B/op            1 allocs/op
      segment.Marshal              2097867 ns/op      1960426 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              3213669 ns/op      1952824 B/op            2 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              5381643 ns/op     10590744 B/op           31 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal              5983599 ns/op      6504741 B/op        77032 allocs/op
     easyjson.Marshal              2276656 ns/op      1996898 B/op           73 allocs/op
       sonnet.Marshal              2337799 ns/op      1941848 B/op            5 allocs/op

      segment ██████  1.53
        goccy ██████  1.52
     easyjson █████▋ 1.41
       sonnet █████▍ 1.37
     jsoniter █████▏ 1.29
         json ▓▓▓▓ 1.00
        sonic ███▉ 1.00
        json2 ███▏ 0.81
        codec ██▍ 0.60
       ffjson ██▏ 0.54
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal StringUnicode generically
         json.Unmarshal              83257 ns/op        30303 B/op          194 allocs/op
        json2.Unmarshal              36680 ns/op        27556 B/op          179 allocs/op
     fastjson.Unmarshal               8453 ns/op        35208 B/op           76 allocs/op
     jsoniter.Unmarshal              28210 ns/op        31956 B/op          262 allocs/op
   jsonparser.Unmarshal              16124 ns/op        18624 B/op          121 allocs/op
        goccy.Unmarshal              22979 ns/op        31320 B/op          191 allocs/op
      segment.Unmarshal              97541 ns/op        43274 B/op          307 allocs/op
     simdjson.Unmarshal       >>> Unsupported platform <<<
        gjson.Unmarshal              13531 ns/op        20272 B/op           66 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal              97158 ns/op       130906 B/op          204 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal              30041 ns/op        32851 B/op          315 allocs/op
        djson.Unmarshal              49450 ns/op        47484 B/op          133 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal              38758 ns/op        30468 B/op          193 allocs/op

   * fastjson ███████████████████████████████████████▍ 9.85
      * gjson ████████████████████████▌ 6.15
 * jsonparser ████████████████████▋ 5.16
        goccy ██████████████▍ 3.62
     jsoniter ███████████▊ 2.95
        codec ███████████  2.77
        json2 █████████  2.27
       sonnet ████████▌ 2.15
      * djson ██████▋ 1.68
         json ▓▓▓▓ 1.00
        sonic ███▍ 0.86
      segment ███▍ 0.85
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
     simdjson >>> Unsupported platform <<<

Unmarshal StringUnicode into struct
         json.Unmarshal              80579 ns/op        18296 B/op           66 allocs/op
        json2.Unmarshal              30295 ns/op        15472 B/op           49 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal              19593 ns/op        18912 B/op           72 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal              14540 ns/op        18432 B/op            1 allocs/op
      segment.Unmarshal              56874 ns/op        35280 B/op          122 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal              93255 ns/op       118880 B/op           76 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal              15037 ns/op        18824 B/op           63 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal              18049 ns/op        18448 B/op           66 allocs/op
     easyjson.Unmarshal               5593 ns/op        17776 B/op           61 allocs/op
       sonnet.Unmarshal              33812 ns/op        18384 B/op           63 allocs/op

     easyjson █████████████████████████████████████████████████████████▋ 14.41
        goccy ██████████████████████▏ 5.54
        codec █████████████████████▍ 5.36
       ffjson █████████████████▊ 4.46
     jsoniter ████████████████▍ 4.11
        json2 ██████████▋ 2.66
       sonnet █████████▌ 2.38
      segment █████▋ 1.42
         json ▓▓▓▓ 1.00
        sonic ███▍ 0.86
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal StringUnicode generically
         json.Marshal                37504 ns/op        23042 B/op          122 allocs/op
        json2.Marshal                35552 ns/op        19562 B/op           65 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal                18355 ns/op        18561 B/op            4 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                37616 ns/op        18435 B/op            1 allocs/op
      segment.Marshal                34633 ns/op        18434 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal                38006 ns/op        23094 B/op          123 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                34272 ns/op        84504 B/op           13 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Marshal                35861 ns/op        18776 B/op            5 allocs/op

     jsoniter ████████▏ 2.04
        codec ████▍ 1.09
      segment ████▎ 1.08
        json2 ████▏ 1.05
       sonnet ████▏ 1.05
         json ▓▓▓▓ 1.00
        goccy ███▉ 1.00
        sonic ███▉ 0.99
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal StringUnicode from struct
         json.Marshal                27084 ns/op        18433 B/op            1 allocs/op
        json2.Marshal                24119 ns/op        18433 B/op            1 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal                16865 ns/op        18441 B/op            2 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                29448 ns/op        18434 B/op            1 allocs/op
      segment.Marshal                29221 ns/op        18434 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal                27321 ns/op        18480 B/op            2 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                35197 ns/op        84504 B/op           13 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal                26443 ns/op        50611 B/op           67 allocs/op
     easyjson.Marshal                23158 ns/op        19272 B/op           11 allocs/op
       sonnet.Marshal                29417 ns/op        18776 B/op            5 allocs/op

     jsoniter ██████▍ 1.61
     easyjson ████▋ 1.17
        json2 ████▍ 1.12
       ffjson ████  1.02
         json ▓▓▓▓ 1.00
        sonic ███▉ 0.99
      segment ███▋ 0.93
       sonnet ███▋ 0.92
        goccy ███▋ 0.92
        codec ███  0.77
   jsonparser >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal single record (2kb), read few keys generically
         json.Unmarshal              28748 ns/op        12488 B/op           21 allocs/op
        json2 >>> not supported <<<
     fastjson.Unmarshal               8948 ns/op        12034 B/op            7 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal               8623 ns/op        12208 B/op           17 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal               9429 ns/op        16368 B/op            9 allocs/op
      gjson-v.Unmarshal              12970 ns/op        16368 B/op            9 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

   jsonparser █████████████▎ 3.33
     fastjson ████████████▊ 3.21
        gjson ████████████▏ 3.05
      gjson-v ████████▊ 2.22
         json ▓▓▓▓ 1.00
      segment >>> not supported <<<
     jsoniter >>> not supported <<<
     simdjson >>> not supported <<<
        goccy >>> not supported <<<
        json2 >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

Unmarshal single record (2kb), read few keys into struct
         json.Unmarshal              28473 ns/op        12488 B/op           21 allocs/op
        json2.Unmarshal               9196 ns/op          128 B/op            6 allocs/op
     fastjson.Unmarshal               9216 ns/op        12080 B/op           10 allocs/op
     jsoniter.Unmarshal              13176 ns/op        14056 B/op           84 allocs/op
   jsonparser.Unmarshal               8719 ns/op        12208 B/op           17 allocs/op
        goccy.Unmarshal               9400 ns/op        16129 B/op            7 allocs/op
      segment.Unmarshal              10298 ns/op        12064 B/op           12 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal               9808 ns/op        16400 B/op           10 allocs/op
      gjson-v.Unmarshal              13432 ns/op        16400 B/op           10 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal              11294 ns/op        13240 B/op           15 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal              23496 ns/op        14184 B/op           36 allocs/op
     easyjson.Unmarshal              22418 ns/op        12320 B/op           16 allocs/op
       sonnet.Unmarshal              13042 ns/op        12368 B/op           21 allocs/op

 * jsonparser █████████████  3.27
        json2 ████████████▍ 3.10
   * fastjson ████████████▎ 3.09
        goccy ████████████  3.03
      * gjson ███████████▌ 2.90
      segment ███████████  2.76
        codec ██████████  2.52
       sonnet ████████▋ 2.18
     jsoniter ████████▋ 2.16
    * gjson-v ████████▍ 2.12
     easyjson █████  1.27
       ffjson ████▊ 1.21
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        djson >>> not supported <<<

Unmarshal single record (2kb), read all keys generically
         json.Unmarshal              36906 ns/op        29775 B/op          340 allocs/op
        json2.Unmarshal              23750 ns/op        15537 B/op          295 allocs/op
     fastjson.Unmarshal              13882 ns/op        13608 B/op           73 allocs/op
     jsoniter.Unmarshal              27562 ns/op        31680 B/op          457 allocs/op
   jsonparser.Unmarshal              21049 ns/op        14528 B/op          135 allocs/op
        goccy.Unmarshal              24833 ns/op        34514 B/op          398 allocs/op
      segment.Unmarshal              55402 ns/op       173240 B/op          475 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal              18768 ns/op        18336 B/op           76 allocs/op
      gjson-v.Unmarshal              22455 ns/op        18336 B/op           76 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal              38454 ns/op        19456 B/op          435 allocs/op
        djson.Unmarshal              20405 ns/op        32168 B/op          183 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal              21216 ns/op        30351 B/op          341 allocs/op

     fastjson ██████████▋ 2.66
        gjson ███████▊ 1.97
        djson ███████▏ 1.81
   jsonparser ███████  1.75
       sonnet ██████▉ 1.74
      gjson-v ██████▌ 1.64
        json2 ██████▏ 1.55
        goccy █████▉ 1.49
     jsoniter █████▎ 1.34
         json ▓▓▓▓ 1.00
        codec ███▊ 0.96
      segment ██▋ 0.67
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
        sonic >>> not supported <<<

Unmarshal single record (2kb), read all keys into struct
         json.Unmarshal              34803 ns/op        14608 B/op           80 allocs/op
        json2.Unmarshal              14833 ns/op         3104 B/op           35 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal              13232 ns/op        14384 B/op           77 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal              10578 ns/op        16131 B/op            7 allocs/op
      segment.Unmarshal              14947 ns/op        13680 B/op           73 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal              15188 ns/op        14536 B/op           72 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal              17358 ns/op        16784 B/op          113 allocs/op
     easyjson.Unmarshal              15353 ns/op        15888 B/op           94 allocs/op
       sonnet.Unmarshal              16222 ns/op        17066 B/op          112 allocs/op

        goccy █████████████▏ 3.29
     jsoniter ██████████▌ 2.63
        json2 █████████▍ 2.35
      segment █████████▎ 2.33
        codec █████████▏ 2.29
     easyjson █████████  2.27
       sonnet ████████▌ 2.15
       ffjson ████████  2.01
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys generically
         json.Unmarshal          674884854 ns/op     21756264 B/op       394052 allocs/op
        json2 >>> not supported <<<
     fastjson.Unmarshal          127916541 ns/op      2537813 B/op        39513 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal           88929032 ns/op      4614294 B/op       177229 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal          119249074 ns/op    105924746 B/op        39406 allocs/op
      gjson-v.Unmarshal          238851908 ns/op    105924755 B/op        39406 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

   jsonparser ██████████████████████████████▎ 7.59
        gjson ██████████████████████▋ 5.66
     fastjson █████████████████████  5.28
      gjson-v ███████████▎ 2.83
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
      segment >>> not supported <<<
        json2 >>> not supported <<<
     jsoniter >>> not supported <<<
        goccy >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys into struct
         json.Unmarshal          675047166 ns/op     21755688 B/op       394052 allocs/op
        json2.Unmarshal          257752406 ns/op        16442 B/op           67 allocs/op
     fastjson.Unmarshal          128682536 ns/op      2537813 B/op        39513 allocs/op
     jsoniter.Unmarshal          288187781 ns/op     89397132 B/op      3979910 allocs/op
   jsonparser.Unmarshal           88987083 ns/op      4614292 B/op       177229 allocs/op
        goccy.Unmarshal          121972638 ns/op    105927824 B/op        39418 allocs/op
      segment.Unmarshal          185530555 ns/op      2526037 B/op        39406 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal          118294935 ns/op    105924736 B/op        39406 allocs/op
      gjson-v.Unmarshal          237914591 ns/op    105924755 B/op        39406 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal          173944902 ns/op     48866314 B/op       157621 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal          397625944 ns/op     67784197 B/op       591086 allocs/op
     easyjson.Unmarshal          566405625 ns/op      2526852 B/op        39410 allocs/op
       sonnet.Unmarshal          191543076 ns/op      7644805 B/op        88606 allocs/op

 * jsonparser ██████████████████████████████▎ 7.59
      * gjson ██████████████████████▊ 5.71
        goccy ██████████████████████▏ 5.53
   * fastjson ████████████████████▉ 5.25
        codec ███████████████▌ 3.88
      segment ██████████████▌ 3.64
       sonnet ██████████████  3.52
    * gjson-v ███████████▎ 2.84
        json2 ██████████▍ 2.62
     jsoniter █████████▎ 2.34
       ffjson ██████▊ 1.70
     easyjson ████▊ 1.19
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        djson >>> not supported <<<
      sonic-v >>> not supported <<<
        sonic >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys generically
         json.Unmarshal         1065006916 ns/op    743894832 B/op     14314982 allocs/op
        json2.Unmarshal          872200542 ns/op    693014624 B/op     11939101 allocs/op
     fastjson.Unmarshal          374939389 ns/op     66551194 B/op      2876853 allocs/op
     jsoniter.Unmarshal          915412895 ns/op    819024048 B/op     19201288 allocs/op
   jsonparser.Unmarshal          657618562 ns/op    105469184 B/op      5437891 allocs/op
        goccy.Unmarshal          834659708 ns/op    876413268 B/op     16797650 allocs/op
      segment.Unmarshal         2430035041 ns/op   6770504624 B/op     19978349 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal          623158708 ns/op    196083400 B/op      2994781 allocs/op
      gjson-v.Unmarshal          736469458 ns/op    196083504 B/op      2994782 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal         1399426791 ns/op    313135648 B/op     18323502 allocs/op
        djson.Unmarshal          588255792 ns/op    770779456 B/op      7537349 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal          636517520 ns/op    760673208 B/op     14246037 allocs/op

     fastjson ███████████▎ 2.84
        djson ███████▏ 1.81
        gjson ██████▊ 1.71
       sonnet ██████▋ 1.67
   jsonparser ██████▍ 1.62
      gjson-v █████▊ 1.45
        goccy █████  1.28
        json2 ████▉ 1.22
     jsoniter ████▋ 1.16
         json ▓▓▓▓ 1.00
        codec ███  0.76
      segment █▊ 0.44
     simdjson >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys into struct
         json.Unmarshal          940127479 ns/op    111143608 B/op      3113018 allocs/op
        json2.Unmarshal          538296770 ns/op    129718460 B/op      1792387 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          298948000 ns/op     97277368 B/op      2955390 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal          206827458 ns/op    105954774 B/op        39536 allocs/op
      segment.Unmarshal          403718638 ns/op     69132701 B/op      2797767 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal          397341472 ns/op    102289994 B/op      2718957 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal          528466958 ns/op    218001024 B/op      4689225 allocs/op
     easyjson.Unmarshal          418722250 ns/op    170697664 B/op      3822301 allocs/op
       sonnet.Unmarshal          414539097 ns/op    212704890 B/op      4817452 allocs/op

        goccy ██████████████████▏ 4.55
     jsoniter ████████████▌ 3.14
        codec █████████▍ 2.37
      segment █████████▎ 2.33
       sonnet █████████  2.27
     easyjson ████████▉ 2.25
       ffjson ███████  1.78
        json2 ██████▉ 1.75
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys generically
         json.Unmarshal        31172616209 ns/op    116288976 B/op      2105576 allocs/op
        json2 >>> not supported <<<
     fastjson.Unmarshal         5405913958 ns/op     14548160 B/op       218829 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal         3002996458 ns/op     24696576 B/op       947485 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal         4470717000 ns/op   5740688544 B/op       210572 allocs/op
      gjson-v.Unmarshal        10147630792 ns/op   5740689120 B/op       210578 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

   jsonparser █████████████████████████████████████████▌ 10.38
        gjson ███████████████████████████▉ 6.97
     fastjson ███████████████████████  5.77
      gjson-v ████████████▎ 3.07
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
      segment >>> not supported <<<
        json2 >>> not supported <<<
     jsoniter >>> not supported <<<
        goccy >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys into struct
         json.Unmarshal        31149364417 ns/op    116288976 B/op      2105576 allocs/op
        json2.Unmarshal        12681093792 ns/op        17624 B/op           72 allocs/op
     fastjson.Unmarshal         5826332750 ns/op     14548160 B/op       218829 allocs/op
     jsoniter.Unmarshal        13799118542 ns/op   4641819184 B/op    210767687 allocs/op
   jsonparser.Unmarshal         3027409833 ns/op     24696576 B/op       947485 allocs/op
        goccy.Unmarshal         5143607541 ns/op   5740848736 B/op       211049 allocs/op
      segment.Unmarshal         8206325583 ns/op     13537216 B/op       210563 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal         4477475333 ns/op   5740688928 B/op       210576 allocs/op
      gjson-v.Unmarshal        10521156042 ns/op   5740688256 B/op       210569 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal         7150501583 ns/op    261152232 B/op       842234 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal        16743516167 ns/op    362240664 B/op      3158425 allocs/op
     easyjson.Unmarshal        28045107917 ns/op     13539480 B/op       210580 allocs/op
       sonnet.Unmarshal         8328851083 ns/op     40871920 B/op       473226 allocs/op

 * jsonparser █████████████████████████████████████████▏ 10.29
      * gjson ███████████████████████████▊ 6.96
        goccy ████████████████████████▏ 6.06
   * fastjson █████████████████████▍ 5.35
        codec █████████████████▍ 4.36
      segment ███████████████▏ 3.80
       sonnet ██████████████▉ 3.74
    * gjson-v ███████████▊ 2.96
        json2 █████████▊ 2.46
     jsoniter █████████  2.26
       ffjson ███████▍ 1.86
     easyjson ████▍ 1.11
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        djson >>> not supported <<<
      sonic-v >>> not supported <<<
        sonic >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys generically
         json.Unmarshal        51681838167 ns/op  37196165376 B/op    701322680 allocs/op
        json2.Unmarshal        40896711833 ns/op  35236079368 B/op    579113253 allocs/op
     fastjson.Unmarshal        17873952875 ns/op   3343034432 B/op    142344813 allocs/op
     jsoniter.Unmarshal        44033619167 ns/op  41742177248 B/op    960517311 allocs/op
   jsonparser.Unmarshal        32952000209 ns/op   5277693816 B/op    273513559 allocs/op
        goccy.Unmarshal        39318439375 ns/op  44758904720 B/op    833139284 allocs/op
      segment.Unmarshal       122778763208 ns/op 341672344088 B/op    995514183 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal        31392917416 ns/op  10406633872 B/op    146758257 allocs/op
      gjson-v.Unmarshal        37498572291 ns/op  10406633136 B/op    146758249 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal        69650495500 ns/op  13524661600 B/op    905186211 allocs/op
        djson.Unmarshal        28334121458 ns/op  39791099552 B/op    375168664 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Unmarshal        29948491958 ns/op  35952028760 B/op    699058950 allocs/op

     fastjson ███████████▌ 2.89
        djson ███████▎ 1.82
       sonnet ██████▉ 1.73
        gjson ██████▌ 1.65
   jsonparser ██████▎ 1.57
      gjson-v █████▌ 1.38
        goccy █████▎ 1.31
        json2 █████  1.26
     jsoniter ████▋ 1.17
         json ▓▓▓▓ 1.00
        codec ██▉ 0.74
      segment █▋ 0.42
     simdjson >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys into struct
         json.Unmarshal        44744801333 ns/op   4426642256 B/op    137915227 allocs/op
        json2.Unmarshal        27063054917 ns/op   6894515848 B/op     94197508 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal        14012362708 ns/op   4986239976 B/op    150338211 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         8640527458 ns/op   5742314424 B/op       216787 allocs/op
      segment.Unmarshal        19493120459 ns/op   3474082280 B/op    141915746 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal        18452377375 ns/op   3045058464 B/op    133914577 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal        22947244625 ns/op  10005805496 B/op    215401282 allocs/op
     easyjson.Unmarshal        20246948875 ns/op   8928967416 B/op    198556048 allocs/op
       sonnet.Unmarshal        18969659791 ns/op   8199695408 B/op    236094590 allocs/op

        goccy ████████████████████▋ 5.18
     jsoniter ████████████▊ 3.19
        codec █████████▋ 2.42
       sonnet █████████▍ 2.36
      segment █████████▏ 2.30
     easyjson ████████▊ 2.21
       ffjson ███████▊ 1.95
        json2 ██████▌ 1.65
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
        djson >>> not supported <<<
        gjson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal custom data through an object builder
         json.Marshal                 1268 ns/op          688 B/op           17 allocs/op
        json2 >>> not supported <<<
     fastjson.Marshal                  823 ns/op          679 B/op            6 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Marshal                 2286 ns/op         3024 B/op           27 allocs/op
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

     fastjson ██████▏ 1.54
         json ▓▓▓▓ 1.00
        gjson ██▏ 0.55
   jsonparser >>> not supported <<<
        sonic >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        json2 >>> not supported <<<
      gjson-v >>> not supported <<<
     jsoniter >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

Validate []byte
         json.Validate               10566 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
     fastjson.Validate                4912 ns/op            0 B/op            0 allocs/op
     jsoniter.Validate                6846 ns/op         2184 B/op          100 allocs/op
   jsonparser >>> not supported <<<
        goccy.Validate               23522 ns/op        27584 B/op          463 allocs/op
      segment.Validate                3568 ns/op            0 B/op            0 allocs/op
     simdjson >>> not supported <<<
        gjson.Validate                3506 ns/op            0 B/op            0 allocs/op
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet.Validate                5957 ns/op            0 B/op            0 allocs/op

        gjson ████████████  3.01
      segment ███████████▊ 2.96
     fastjson ████████▌ 2.15
       sonnet ███████  1.77
     jsoniter ██████▏ 1.54
         json ▓▓▓▓ 1.00
        goccy █▊ 0.45
   jsonparser >>> not supported <<<
     simdjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
        json2 >>> not supported <<<

Validate string
         json.Validate               10453 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
     fastjson.Validate                5232 ns/op         4096 B/op            1 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Validate                3814 ns/op         4096 B/op            1 allocs/op
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

        gjson ██████████▉ 2.74
     fastjson ███████▉ 2.00
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
        sonic >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        json2 >>> not supported <<<
      gjson-v >>> not supported <<<
     jsoniter >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<
       sonnet >>> not supported <<<

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
