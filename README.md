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
         json.Unmarshal            2930463 ns/op      1339378 B/op        52215 allocs/op
        json2.Unmarshal            2301422 ns/op      1011496 B/op        37898 allocs/op
     fastjson.Unmarshal            2667064 ns/op      8725359 B/op        30278 allocs/op
     jsoniter.Unmarshal            3608686 ns/op      1961488 B/op        84971 allocs/op
   jsonparser.Unmarshal            2210165 ns/op          152 B/op           12 allocs/op
        goccy.Unmarshal            2496667 ns/op      1471004 B/op        59189 allocs/op
      segment.Unmarshal            4786675 ns/op      2119754 B/op        51680 allocs/op
     simdjson.Unmarshal       >>> Unsupported platform <<<
        gjson.Unmarshal            3679093 ns/op       393072 B/op        14314 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal            3131722 ns/op      2943315 B/op        52232 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            3357473 ns/op      1765857 B/op        37408 allocs/op
        djson.Unmarshal            1791333 ns/op      1289904 B/op        37891 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

      * djson ██████▌ 1.64
 * jsonparser █████▎ 1.33
        json2 █████  1.27
        goccy ████▋ 1.17
   * fastjson ████▍ 1.10
         json ▓▓▓▓ 1.00
        sonic ███▋ 0.94
        codec ███▍ 0.87
     jsoniter ███▏ 0.81
      * gjson ███▏ 0.80
      segment ██▍ 0.61
      gjson-v >>> not supported <<<
     simdjson >>> Unsupported platform <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal CanadaGeometry into struct
         json.Unmarshal            2426403 ns/op          472 B/op           13 allocs/op
        json2.Unmarshal            1543791 ns/op       347966 B/op         1626 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            1892427 ns/op       598761 B/op        25122 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal             958384 ns/op       278610 B/op            1 allocs/op
      segment.Unmarshal             984800 ns/op           48 B/op            4 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal            2602558 ns/op      1604400 B/op           30 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            1674874 ns/op         1224 B/op            7 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal            3316605 ns/op       537713 B/op           25 allocs/op
     easyjson.Unmarshal            1192676 ns/op       330368 B/op          753 allocs/op

        goccy ██████████▏ 2.53
      segment █████████▊ 2.46
     easyjson ████████▏ 2.03
        json2 ██████▎ 1.57
        codec █████▊ 1.45
     jsoniter █████▏ 1.28
         json ▓▓▓▓ 1.00
        sonic ███▋ 0.93
       ffjson ██▉ 0.73
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal CanadaGeometry generically
         json.Marshal              1584502 ns/op       280522 B/op           21 allocs/op
        json2.Marshal              1244573 ns/op       279982 B/op            7 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               619693 ns/op       172911 B/op           13 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              1513295 ns/op       278791 B/op            3 allocs/op
      segment.Marshal              1182506 ns/op       279723 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              1587658 ns/op       280570 B/op           22 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              1288753 ns/op      1186328 B/op           22 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     jsoniter ██████████▏ 2.56
      segment █████▎ 1.34
        json2 █████  1.27
        codec ████▉ 1.23
        goccy ████▏ 1.05
         json ▓▓▓▓ 1.00
        sonic ███▉ 1.00
   jsonparser >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
     fastjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Marshal CanadaGeometry from struct
         json.Marshal              1146507 ns/op       279552 B/op            1 allocs/op
        json2.Marshal              1174065 ns/op       278539 B/op            1 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               269668 ns/op       172198 B/op            2 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              1110036 ns/op       279642 B/op            1 allocs/op
      segment.Marshal              1024117 ns/op       279560 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              1150302 ns/op       279597 B/op            2 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              1235353 ns/op      1186329 B/op           22 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal              2105245 ns/op      1063342 B/op        14391 allocs/op
     easyjson.Marshal              1009099 ns/op       280579 B/op           20 allocs/op

     jsoniter █████████████████  4.25
     easyjson ████▌ 1.14
      segment ████▍ 1.12
        goccy ████▏ 1.03
         json ▓▓▓▓ 1.00
        sonic ███▉ 1.00
        json2 ███▉ 0.98
        codec ███▋ 0.93
       ffjson ██▏ 0.54
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal CITMCatalog generically
         json.Unmarshal           11646679 ns/op      5123805 B/op        95373 allocs/op
        json2.Unmarshal            4576091 ns/op      4958431 B/op        80513 allocs/op
     fastjson.Unmarshal            4580593 ns/op     19991467 B/op        45701 allocs/op
     jsoniter.Unmarshal            5609850 ns/op      5572869 B/op       118756 allocs/op
   jsonparser.Unmarshal           11463707 ns/op       145320 B/op        15863 allocs/op
        goccy.Unmarshal            5458185 ns/op      7416016 B/op       123577 allocs/op
      segment.Unmarshal           19768006 ns/op     55526553 B/op       123714 allocs/op
     simdjson.Unmarshal       >>> Unsupported platform <<<
        gjson.Unmarshal            7936037 ns/op      1855521 B/op        15131 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal           12885489 ns/op     12773365 B/op        95391 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            9411090 ns/op      6195134 B/op       159476 allocs/op
        djson.Unmarshal            4865584 ns/op      6437655 B/op        54480 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        json2 ██████████▏ 2.55
   * fastjson ██████████▏ 2.54
      * djson █████████▌ 2.39
        goccy ████████▌ 2.13
     jsoniter ████████▎ 2.08
      * gjson █████▊ 1.47
        codec ████▉ 1.24
 * jsonparser ████  1.02
         json ▓▓▓▓ 1.00
        sonic ███▌ 0.90
      segment ██▎ 0.59
     simdjson >>> Unsupported platform <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal CITMCatalog into struct
         json.Unmarshal           10519151 ns/op       253352 B/op        10950 allocs/op
        json2.Unmarshal            2691569 ns/op       888447 B/op         6293 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            1963466 ns/op        72332 B/op         2535 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal            1423802 ns/op      1781233 B/op         2226 allocs/op
      segment.Unmarshal            1635864 ns/op        52384 B/op         1354 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal           11822809 ns/op      7902978 B/op        10969 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            2409894 ns/op        47880 B/op         1036 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal       >>> ffjson error: (*errors.errorString)cannot unmarshal tok:string into Go value for int64 offset=40 line=3 char=19 <<<
     easyjson.Unmarshal            2796218 ns/op       896599 B/op         5704 allocs/op

        goccy █████████████████████████████▌ 7.39
      segment █████████████████████████▋ 6.43
     jsoniter █████████████████████▍ 5.36
        codec █████████████████▍ 4.36
        json2 ███████████████▋ 3.91
     easyjson ███████████████  3.76
         json ▓▓▓▓ 1.00
        sonic ███▌ 0.89
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> ffjson error: (*errors.errorString)cannot unmarshal tok:string into Go value for int64 offset=40 line=3 char=19 <<<
     fastjson >>> not supported <<<

Marshal CITMCatalog generically
         json.Marshal              4532937 ns/op      2384653 B/op        62674 allocs/op
        json2.Marshal              1877228 ns/op       512222 B/op           16 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              2862956 ns/op      1914032 B/op        32812 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              2922042 ns/op       662552 B/op          324 allocs/op
      segment.Marshal              1933875 ns/op       512282 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              4521702 ns/op      2384701 B/op        62675 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              2556736 ns/op      2538009 B/op           25 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        json2 █████████▋ 2.41
      segment █████████▍ 2.34
        codec ███████  1.77
     jsoniter ██████▎ 1.58
        goccy ██████▏ 1.55
        sonic ████  1.00
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
     fastjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Marshal CITMCatalog from struct
         json.Marshal               720344 ns/op       556026 B/op          890 allocs/op
        json2.Marshal              1039261 ns/op       513247 B/op          133 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               669491 ns/op       509127 B/op           22 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal               423630 ns/op       520911 B/op            3 allocs/op
      segment.Marshal               472943 ns/op       547998 B/op          616 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               726565 ns/op       556729 B/op          891 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              1182128 ns/op      2538009 B/op           25 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal              1086874 ns/op       904572 B/op        14180 allocs/op
     easyjson.Marshal               469986 ns/op       524049 B/op         1125 allocs/op

        goccy ██████▊ 1.70
     easyjson ██████▏ 1.53
      segment ██████  1.52
     jsoniter ████▎ 1.08
         json ▓▓▓▓ 1.00
        sonic ███▉ 0.99
        json2 ██▊ 0.69
       ffjson ██▋ 0.66
        codec ██▍ 0.61
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal SyntheaFHIR generically
         json.Unmarshal           14470526 ns/op      7430181 B/op       134383 allocs/op
        json2.Unmarshal            6868322 ns/op      6938485 B/op       115832 allocs/op
     fastjson.Unmarshal            6635612 ns/op     26046540 B/op        66098 allocs/op
     jsoniter.Unmarshal            8277775 ns/op      8196986 B/op       181427 allocs/op
   jsonparser.Unmarshal           13981398 ns/op      1108792 B/op        49604 allocs/op
        goccy.Unmarshal            7806458 ns/op      9838332 B/op       159459 allocs/op
      segment.Unmarshal           27550259 ns/op     77571258 B/op       193338 allocs/op
     simdjson.Unmarshal       >>> Unsupported platform <<<
        gjson.Unmarshal           10574975 ns/op      2432546 B/op        26821 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal           15831131 ns/op     15653921 B/op       134405 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal           14738470 ns/op      9669739 B/op       247796 allocs/op
        djson.Unmarshal            6175669 ns/op      8374215 B/op        68393 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

      * djson █████████▎ 2.34
   * fastjson ████████▋ 2.18
        json2 ████████▍ 2.11
        goccy ███████▍ 1.85
     jsoniter ██████▉ 1.75
      * gjson █████▍ 1.37
 * jsonparser ████▏ 1.03
         json ▓▓▓▓ 1.00
        codec ███▉ 0.98
        sonic ███▋ 0.91
      segment ██  0.53
     simdjson >>> Unsupported platform <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal SyntheaFHIR into struct
         json.Unmarshal           13413185 ns/op       625416 B/op        21883 allocs/op
        json2.Unmarshal            4730022 ns/op      3103233 B/op        14550 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            3118002 ns/op       825711 B/op        29439 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal            2269750 ns/op      2152029 B/op         2511 allocs/op
      segment.Unmarshal            2792717 ns/op       623760 B/op        21832 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal           15000078 ns/op      8848580 B/op        21903 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            3738040 ns/op       623400 B/op        21826 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal           22567281 ns/op      4359573 B/op        21906 allocs/op
     easyjson.Unmarshal            4186515 ns/op      3531504 B/op        29214 allocs/op

        goccy ███████████████████████▋ 5.91
      segment ███████████████████▏ 4.80
     jsoniter █████████████████▏ 4.30
        codec ██████████████▎ 3.59
     easyjson ████████████▊ 3.20
        json2 ███████████▎ 2.84
         json ▓▓▓▓ 1.00
        sonic ███▌ 0.89
       ffjson ██▍ 0.59
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal SyntheaFHIR generically
         json.Marshal              7548477 ns/op      4144461 B/op        95282 allocs/op
        json2.Marshal              3022831 ns/op      1164118 B/op            8 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              4376788 ns/op      3095758 B/op        45292 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              5295014 ns/op      2511041 B/op         6469 allocs/op
      segment.Marshal              3265443 ns/op      1146954 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              7463682 ns/op      4117961 B/op        95283 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              3194906 ns/op      6642202 B/op           29 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        json2 █████████▉ 2.50
        codec █████████▍ 2.36
      segment █████████▏ 2.31
     jsoniter ██████▉ 1.72
        goccy █████▋ 1.43
        sonic ████  1.01
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
     fastjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Marshal SyntheaFHIR from struct
         json.Marshal              7052081 ns/op      3996463 B/op        18009 allocs/op
        json2.Marshal              6697487 ns/op      3236480 B/op         2569 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              5277215 ns/op      4045662 B/op        17963 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              3506082 ns/op      3994903 B/op        17940 allocs/op
      segment.Marshal              2653399 ns/op      3309913 B/op         5338 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              6908881 ns/op      3995646 B/op        18010 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              6509511 ns/op     16792093 B/op           33 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal              4109917 ns/op      9518386 B/op        45597 allocs/op
     easyjson.Marshal              2994799 ns/op      3979085 B/op        15520 allocs/op

      segment ██████████▋ 2.66
     easyjson █████████▍ 2.35
        goccy ████████  2.01
       ffjson ██████▊ 1.72
     jsoniter █████▎ 1.34
        codec ████▎ 1.08
        json2 ████▏ 1.05
        sonic ████  1.02
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal TwitterStatus generically
         json.Unmarshal            4486665 ns/op      2152006 B/op        31260 allocs/op
        json2.Unmarshal            2547578 ns/op      2074477 B/op        27237 allocs/op
     fastjson.Unmarshal            1437217 ns/op      5341210 B/op        11130 allocs/op
     jsoniter.Unmarshal            2744495 ns/op      2427716 B/op        45043 allocs/op
   jsonparser.Unmarshal            2956080 ns/op       300312 B/op        10184 allocs/op
        goccy.Unmarshal            2435438 ns/op      2755862 B/op        39695 allocs/op
      segment.Unmarshal            5876273 ns/op      7060651 B/op        39237 allocs/op
     simdjson.Unmarshal       >>> Unsupported platform <<<
        gjson.Unmarshal            2464272 ns/op       822360 B/op         6884 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal            4940797 ns/op      5525348 B/op        31279 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            4308912 ns/op      2576802 B/op        60985 allocs/op
        djson.Unmarshal            2137382 ns/op      2470473 B/op        12678 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

   * fastjson ████████████▍ 3.12
      * djson ████████▍ 2.10
        goccy ███████▎ 1.84
      * gjson ███████▎ 1.82
        json2 ███████  1.76
     jsoniter ██████▌ 1.63
 * jsonparser ██████  1.52
        codec ████▏ 1.04
         json ▓▓▓▓ 1.00
        sonic ███▋ 0.91
      segment ███  0.76
     simdjson >>> Unsupported platform <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal TwitterStatus into struct
         json.Unmarshal            4090848 ns/op       304612 B/op         6486 allocs/op
        json2.Unmarshal            1473400 ns/op       306185 B/op         2972 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            1080897 ns/op       339687 B/op         6293 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal             783980 ns/op       647067 B/op          238 allocs/op
      segment.Unmarshal            1251024 ns/op       369712 B/op         5341 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal            4543519 ns/op      3678144 B/op         6505 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            1167933 ns/op       215944 B/op         4530 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal            2346693 ns/op      1052197 B/op         9494 allocs/op
     easyjson.Unmarshal             958591 ns/op       324336 B/op         4907 allocs/op

        goccy ████████████████████▊ 5.22
     easyjson █████████████████  4.27
     jsoniter ███████████████▏ 3.78
        codec ██████████████  3.50
      segment █████████████  3.27
        json2 ███████████  2.78
       ffjson ██████▉ 1.74
         json ▓▓▓▓ 1.00
        sonic ███▌ 0.90
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal TwitterStatus generically
         json.Marshal              2841541 ns/op      1490577 B/op        27955 allocs/op
        json2.Marshal              1021909 ns/op       469332 B/op            7 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              1102065 ns/op       631136 B/op         3793 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              2223888 ns/op      1047892 B/op          539 allocs/op
      segment.Marshal              1543387 ns/op       478470 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              2829854 ns/op      1482417 B/op        27956 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              1180194 ns/op      2538008 B/op           25 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        json2 ███████████  2.78
     jsoniter ██████████▎ 2.58
        codec █████████▋ 2.41
      segment ███████▎ 1.84
        goccy █████  1.28
        sonic ████  1.00
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
     fastjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Marshal TwitterStatus from struct
         json.Marshal               559419 ns/op       513502 B/op          141 allocs/op
        json2.Marshal               744292 ns/op       503509 B/op           59 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               590737 ns/op       503457 B/op           52 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal               383253 ns/op       509718 B/op            6 allocs/op
      segment.Marshal               407545 ns/op       512636 B/op          121 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               564524 ns/op       505979 B/op          142 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              1026333 ns/op      2540385 B/op           29 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal               645549 ns/op       798804 B/op         6117 allocs/op
     easyjson.Marshal               443796 ns/op       519861 B/op          953 allocs/op

        goccy █████▊ 1.46
      segment █████▍ 1.37
     easyjson █████  1.26
         json ▓▓▓▓ 1.00
        sonic ███▉ 0.99
     jsoniter ███▊ 0.95
       ffjson ███▍ 0.87
        json2 ███  0.75
        codec ██▏ 0.55
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal GolangSource generically
         json.Unmarshal           16314898 ns/op      7916355 B/op       271264 allocs/op
        json2.Unmarshal           10801435 ns/op      6989969 B/op       218839 allocs/op
     fastjson.Unmarshal           11501837 ns/op     50767001 B/op       130330 allocs/op
     jsoniter.Unmarshal           13231298 ns/op      9233487 B/op       346658 allocs/op
   jsonparser.Unmarshal           41288053 ns/op       925272 B/op        64023 allocs/op
        goccy.Unmarshal           13147988 ns/op     10872736 B/op       320840 allocs/op
      segment.Unmarshal           56850888 ns/op     67095377 B/op       334022 allocs/op
     simdjson.Unmarshal       >>> Unsupported platform <<<
        gjson.Unmarshal           47130026 ns/op      2658512 B/op        76819 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal           17606715 ns/op     15992867 B/op       271286 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal           25003260 ns/op     10502703 B/op       436539 allocs/op
        djson.Unmarshal            6485176 ns/op      7822419 B/op       117489 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

      * djson ██████████  2.52
        json2 ██████  1.51
   * fastjson █████▋ 1.42
        goccy ████▉ 1.24
     jsoniter ████▉ 1.23
         json ▓▓▓▓ 1.00
        sonic ███▋ 0.93
        codec ██▌ 0.65
 * jsonparser █▌ 0.40
      * gjson █▍ 0.35
      segment █▏ 0.29
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> Unsupported platform <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal GolangSource into struct
         json.Unmarshal           15821851 ns/op       704424 B/op        24927 allocs/op
        json2.Unmarshal            7198056 ns/op      3236335 B/op        13941 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            4885754 ns/op       994058 B/op        37077 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal            2967276 ns/op      1944557 B/op            1 allocs/op
      segment.Unmarshal            3012914 ns/op       411672 B/op        12806 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal           16964857 ns/op      8780936 B/op        24949 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            7033787 ns/op       412848 B/op        12809 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal            9416814 ns/op      6276466 B/op        52765 allocs/op
     easyjson.Unmarshal            5951104 ns/op      4282425 B/op        27846 allocs/op

        goccy █████████████████████▎ 5.33
      segment █████████████████████  5.25
     jsoniter ████████████▉ 3.24
     easyjson ██████████▋ 2.66
        codec ████████▉ 2.25
        json2 ████████▊ 2.20
       ffjson ██████▋ 1.68
         json ▓▓▓▓ 1.00
        sonic ███▋ 0.93
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal GolangSource generically
         json.Marshal             14703650 ns/op      8550228 B/op       192096 allocs/op
        json2.Marshal              6193179 ns/op      1996903 B/op            7 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              6634270 ns/op      3492228 B/op        38422 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal             10771525 ns/op      2465881 B/op            3 allocs/op
      segment.Marshal              7195143 ns/op      2006504 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal             14761086 ns/op      8555167 B/op       192097 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              9004832 ns/op     10590746 B/op           31 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        json2 █████████▍ 2.37
     jsoniter ████████▊ 2.22
      segment ████████▏ 2.04
        codec ██████▌ 1.63
        goccy █████▍ 1.37
         json ▓▓▓▓ 1.00
        sonic ███▉ 1.00
   jsonparser >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
     fastjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Marshal GolangSource from struct
         json.Marshal              3196429 ns/op      1952776 B/op            1 allocs/op
        json2.Marshal              3927913 ns/op      1977323 B/op            1 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              2486846 ns/op      1794080 B/op            2 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              2104843 ns/op      1960323 B/op            1 allocs/op
      segment.Marshal              2072638 ns/op      1959967 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              3213125 ns/op      1941578 B/op            2 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              5269763 ns/op     10590746 B/op           31 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal              5956170 ns/op      6511834 B/op        77032 allocs/op
     easyjson.Marshal              2287628 ns/op      1996592 B/op           73 allocs/op

      segment ██████▏ 1.54
        goccy ██████  1.52
     easyjson █████▌ 1.40
     jsoniter █████▏ 1.29
         json ▓▓▓▓ 1.00
        sonic ███▉ 0.99
        json2 ███▎ 0.81
        codec ██▍ 0.61
       ffjson ██▏ 0.54
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal StringUnicode generically
         json.Unmarshal              83452 ns/op        30305 B/op          194 allocs/op
        json2.Unmarshal              36929 ns/op        27555 B/op          179 allocs/op
     fastjson.Unmarshal               8089 ns/op        35208 B/op           76 allocs/op
     jsoniter.Unmarshal              28289 ns/op        31958 B/op          262 allocs/op
   jsonparser.Unmarshal              15876 ns/op        18624 B/op          121 allocs/op
        goccy.Unmarshal              22883 ns/op        31325 B/op          191 allocs/op
      segment.Unmarshal              93842 ns/op        43274 B/op          307 allocs/op
     simdjson.Unmarshal       >>> Unsupported platform <<<
        gjson.Unmarshal              13333 ns/op        20272 B/op           66 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal              95644 ns/op       130895 B/op          204 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal              30158 ns/op        32851 B/op          315 allocs/op
        djson.Unmarshal              48834 ns/op        47482 B/op          133 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

   * fastjson █████████████████████████████████████████▎ 10.32
      * gjson █████████████████████████  6.26
 * jsonparser █████████████████████  5.26
        goccy ██████████████▌ 3.65
     jsoniter ███████████▊ 2.95
        codec ███████████  2.77
        json2 █████████  2.26
      * djson ██████▊ 1.71
         json ▓▓▓▓ 1.00
      segment ███▌ 0.89
        sonic ███▍ 0.87
      sonic-v >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> Unsupported platform <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal StringUnicode into struct
         json.Unmarshal              80167 ns/op        18296 B/op           66 allocs/op
        json2.Unmarshal              30189 ns/op        15472 B/op           49 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal              19503 ns/op        18912 B/op           72 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal              14478 ns/op        18432 B/op            1 allocs/op
      segment.Unmarshal              56805 ns/op        35280 B/op          122 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal              91927 ns/op       118880 B/op           76 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal              14934 ns/op        18824 B/op           63 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal              18061 ns/op        18448 B/op           66 allocs/op
     easyjson.Unmarshal               5538 ns/op        17776 B/op           61 allocs/op

     easyjson █████████████████████████████████████████████████████████▉ 14.48
        goccy ██████████████████████▏ 5.54
        codec █████████████████████▍ 5.37
       ffjson █████████████████▊ 4.44
     jsoniter ████████████████▍ 4.11
        json2 ██████████▌ 2.66
      segment █████▋ 1.41
         json ▓▓▓▓ 1.00
        sonic ███▍ 0.87
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal StringUnicode generically
         json.Marshal                37675 ns/op        23042 B/op          122 allocs/op
        json2.Marshal                35726 ns/op        19562 B/op           65 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal                18405 ns/op        18561 B/op            4 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                37059 ns/op        18440 B/op            1 allocs/op
      segment.Marshal                34612 ns/op        18434 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal                37944 ns/op        23090 B/op          123 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                33886 ns/op        84504 B/op           13 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     jsoniter ████████▏ 2.05
        codec ████▍ 1.11
      segment ████▎ 1.09
        json2 ████▏ 1.05
        goccy ████  1.02
         json ▓▓▓▓ 1.00
        sonic ███▉ 0.99
   jsonparser >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
     simdjson >>> not supported <<<
      sonic-v >>> not supported <<<
     fastjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Marshal StringUnicode from struct
         json.Marshal                26880 ns/op        18433 B/op            1 allocs/op
        json2.Marshal                24038 ns/op        18433 B/op            1 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal                16771 ns/op        18441 B/op            2 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                29103 ns/op        18434 B/op            1 allocs/op
      segment.Marshal                29583 ns/op        18432 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal                27654 ns/op        18481 B/op            2 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                33456 ns/op        84504 B/op           13 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal                26557 ns/op        50612 B/op           67 allocs/op
     easyjson.Marshal                23517 ns/op        19275 B/op           11 allocs/op

     jsoniter ██████▍ 1.60
     easyjson ████▌ 1.14
        json2 ████▍ 1.12
       ffjson ████  1.01
         json ▓▓▓▓ 1.00
        sonic ███▉ 0.97
        goccy ███▋ 0.92
      segment ███▋ 0.91
        codec ███▏ 0.80
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal single record (2kb), read few keys generically
         json.Unmarshal              29084 ns/op        12488 B/op           21 allocs/op
        json2 >>> not supported <<<
     fastjson.Unmarshal               8587 ns/op        12034 B/op            7 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal               8308 ns/op        12208 B/op           17 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal               9452 ns/op        16368 B/op            9 allocs/op
      gjson-v.Unmarshal              12953 ns/op        16368 B/op            9 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

   jsonparser ██████████████  3.50
     fastjson █████████████▌ 3.39
        gjson ████████████▎ 3.08
      gjson-v ████████▉ 2.25
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

Unmarshal single record (2kb), read few keys into struct
         json.Unmarshal              28368 ns/op        12488 B/op           21 allocs/op
        json2.Unmarshal               9113 ns/op          128 B/op            6 allocs/op
     fastjson.Unmarshal               8564 ns/op        12080 B/op           10 allocs/op
     jsoniter.Unmarshal              12644 ns/op        14056 B/op           84 allocs/op
   jsonparser.Unmarshal               8566 ns/op        12208 B/op           17 allocs/op
        goccy.Unmarshal               8826 ns/op        16129 B/op            7 allocs/op
      segment.Unmarshal               9772 ns/op        12064 B/op           12 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal               9473 ns/op        16400 B/op           10 allocs/op
      gjson-v.Unmarshal              13082 ns/op        16400 B/op           10 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal              11068 ns/op        13240 B/op           15 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal              23048 ns/op        14184 B/op           36 allocs/op
     easyjson.Unmarshal              21598 ns/op        12320 B/op           16 allocs/op

   * fastjson █████████████▏ 3.31
 * jsonparser █████████████▏ 3.31
        goccy ████████████▊ 3.21
        json2 ████████████▍ 3.11
      * gjson ███████████▉ 2.99
      segment ███████████▌ 2.90
        codec ██████████▎ 2.56
     jsoniter ████████▉ 2.24
    * gjson-v ████████▋ 2.17
     easyjson █████▎ 1.31
       ffjson ████▉ 1.23
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        djson >>> not supported <<<

Unmarshal single record (2kb), read all keys generically
         json.Unmarshal              36512 ns/op        29776 B/op          340 allocs/op
        json2.Unmarshal              23921 ns/op        15538 B/op          295 allocs/op
     fastjson.Unmarshal              13826 ns/op        13608 B/op           73 allocs/op
     jsoniter.Unmarshal              27173 ns/op        31681 B/op          457 allocs/op
   jsonparser.Unmarshal              21026 ns/op        14528 B/op          135 allocs/op
        goccy.Unmarshal              25032 ns/op        34513 B/op          398 allocs/op
      segment.Unmarshal              57080 ns/op       173240 B/op          475 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal              18790 ns/op        18336 B/op           76 allocs/op
      gjson-v.Unmarshal              22453 ns/op        18336 B/op           76 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal              39093 ns/op        19456 B/op          435 allocs/op
        djson.Unmarshal              20166 ns/op        32168 B/op          183 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     fastjson ██████████▌ 2.64
        gjson ███████▊ 1.94
        djson ███████▏ 1.81
   jsonparser ██████▉ 1.74
      gjson-v ██████▌ 1.63
        json2 ██████  1.53
        goccy █████▊ 1.46
     jsoniter █████▎ 1.34
         json ▓▓▓▓ 1.00
        codec ███▋ 0.93
      segment ██▌ 0.64
     simdjson >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal single record (2kb), read all keys into struct
         json.Unmarshal              34513 ns/op        14608 B/op           80 allocs/op
        json2.Unmarshal              14639 ns/op         3104 B/op           35 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal              13400 ns/op        14384 B/op           77 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal              10755 ns/op        16132 B/op            7 allocs/op
      segment.Unmarshal              15092 ns/op        13680 B/op           73 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal              15213 ns/op        14536 B/op           72 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal              17519 ns/op        16784 B/op          113 allocs/op
     easyjson.Unmarshal              15541 ns/op        15888 B/op           94 allocs/op

        goccy ████████████▊ 3.21
     jsoniter ██████████▎ 2.58
        json2 █████████▍ 2.36
      segment █████████▏ 2.29
        codec █████████  2.27
     easyjson ████████▉ 2.22
       ffjson ███████▉ 1.97
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys generically
         json.Unmarshal          663703062 ns/op     21755688 B/op       394052 allocs/op
        json2 >>> not supported <<<
     fastjson.Unmarshal          127384099 ns/op      2537813 B/op        39513 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal           89242968 ns/op      4614294 B/op       177229 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal          118244509 ns/op    105924736 B/op        39406 allocs/op
      gjson-v.Unmarshal          234592308 ns/op    105924793 B/op        39406 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

   jsonparser █████████████████████████████▋ 7.44
        gjson ██████████████████████▍ 5.61
     fastjson ████████████████████▊ 5.21
      gjson-v ███████████▎ 2.83
         json ▓▓▓▓ 1.00
        goccy >>> not supported <<<
      segment >>> not supported <<<
        json2 >>> not supported <<<
     jsoniter >>> not supported <<<
     simdjson >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys into struct
         json.Unmarshal          664814750 ns/op     21756840 B/op       394053 allocs/op
        json2.Unmarshal          256031698 ns/op        16410 B/op           67 allocs/op
     fastjson.Unmarshal          128269651 ns/op      2537813 B/op        39513 allocs/op
     jsoniter.Unmarshal          283401333 ns/op     89397072 B/op      3979910 allocs/op
   jsonparser.Unmarshal           90054884 ns/op      4614292 B/op       177229 allocs/op
        goccy.Unmarshal          122290402 ns/op    105928180 B/op        39420 allocs/op
      segment.Unmarshal          183718548 ns/op      2526037 B/op        39406 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal          117889824 ns/op    105924736 B/op        39406 allocs/op
      gjson-v.Unmarshal          232452141 ns/op    105924793 B/op        39406 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal          173896236 ns/op     48866314 B/op       157621 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal          396128986 ns/op     67783970 B/op       591085 allocs/op
     easyjson.Unmarshal          533720166 ns/op      2526756 B/op        39407 allocs/op

 * jsonparser █████████████████████████████▌ 7.38
      * gjson ██████████████████████▌ 5.64
        goccy █████████████████████▋ 5.44
   * fastjson ████████████████████▋ 5.18
        codec ███████████████▎ 3.82
      segment ██████████████▍ 3.62
    * gjson-v ███████████▍ 2.86
        json2 ██████████▍ 2.60
     jsoniter █████████▍ 2.35
       ffjson ██████▋ 1.68
     easyjson ████▉ 1.25
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        djson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys generically
         json.Unmarshal         1076777375 ns/op    743894784 B/op     14314981 allocs/op
        json2.Unmarshal          865747708 ns/op    692986004 B/op     11939001 allocs/op
     fastjson.Unmarshal          367342180 ns/op     66551205 B/op      2876853 allocs/op
     jsoniter.Unmarshal          910307125 ns/op    818983916 B/op     19201148 allocs/op
   jsonparser.Unmarshal          649875458 ns/op    105469168 B/op      5437891 allocs/op
        goccy.Unmarshal          816875167 ns/op    876440948 B/op     16797747 allocs/op
      segment.Unmarshal         2429726833 ns/op   6770505120 B/op     19978352 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal          623964416 ns/op    196083408 B/op      2994781 allocs/op
      gjson-v.Unmarshal          740883146 ns/op    196083424 B/op      2994781 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal         1406450417 ns/op    313135904 B/op     18323504 allocs/op
        djson.Unmarshal          580971958 ns/op    770804640 B/op      7537437 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     fastjson ███████████▋ 2.93
        djson ███████▍ 1.85
        gjson ██████▉ 1.73
   jsonparser ██████▋ 1.66
      gjson-v █████▊ 1.45
        goccy █████▎ 1.32
        json2 ████▉ 1.24
     jsoniter ████▋ 1.18
         json ▓▓▓▓ 1.00
        codec ███  0.77
      segment █▊ 0.44
     simdjson >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys into struct
         json.Unmarshal          941790958 ns/op    111143592 B/op      3113018 allocs/op
        json2.Unmarshal          533704750 ns/op    129718564 B/op      1792388 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          301008750 ns/op     97277288 B/op      2955389 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal          206479741 ns/op    105959286 B/op        39515 allocs/op
      segment.Unmarshal          403420639 ns/op     69132728 B/op      2797768 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal          392385847 ns/op    102289994 B/op      2718957 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal          532519271 ns/op    218001976 B/op      4689236 allocs/op
     easyjson.Unmarshal          412181708 ns/op    170697685 B/op      3822301 allocs/op

        goccy ██████████████████▏ 4.56
     jsoniter ████████████▌ 3.13
        codec █████████▌ 2.40
      segment █████████▎ 2.33
     easyjson █████████▏ 2.28
       ffjson ███████  1.77
        json2 ███████  1.76
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys generically
         json.Unmarshal        31151799708 ns/op    116289088 B/op      2105577 allocs/op
        json2 >>> not supported <<<
     fastjson.Unmarshal         5362334708 ns/op     14548160 B/op       218829 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal         2989742292 ns/op     24696576 B/op       947485 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal         4458367292 ns/op   5740688160 B/op       210568 allocs/op
      gjson-v.Unmarshal        10100536709 ns/op   5740688928 B/op       210576 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

   jsonparser █████████████████████████████████████████▋ 10.42
        gjson ███████████████████████████▉ 6.99
     fastjson ███████████████████████▏ 5.81
      gjson-v ████████████▎ 3.08
         json ▓▓▓▓ 1.00
        goccy >>> not supported <<<
      segment >>> not supported <<<
        json2 >>> not supported <<<
     jsoniter >>> not supported <<<
     simdjson >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys into struct
         json.Unmarshal        31182632375 ns/op    116289088 B/op      2105577 allocs/op
        json2.Unmarshal        12207852708 ns/op        17624 B/op           72 allocs/op
     fastjson.Unmarshal         5392101125 ns/op     14548160 B/op       218829 allocs/op
     jsoniter.Unmarshal        13370146542 ns/op   4641821040 B/op    210767683 allocs/op
   jsonparser.Unmarshal         2963057584 ns/op     24696576 B/op       947485 allocs/op
        goccy.Unmarshal         4746360542 ns/op   5740860880 B/op       211103 allocs/op
      segment.Unmarshal         7741170875 ns/op     13537216 B/op       210563 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal         4436551042 ns/op   5740688256 B/op       210569 allocs/op
      gjson-v.Unmarshal        10114040250 ns/op   5740687968 B/op       210566 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal         6530384792 ns/op    261152328 B/op       842235 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal        16949572083 ns/op    362241864 B/op      3158442 allocs/op
     easyjson.Unmarshal        27074810792 ns/op     13539304 B/op       210578 allocs/op

 * jsonparser ██████████████████████████████████████████  10.52
      * gjson ████████████████████████████  7.03
        goccy ██████████████████████████▎ 6.57
   * fastjson ███████████████████████▏ 5.78
        codec ███████████████████  4.78
      segment ████████████████  4.03
    * gjson-v ████████████▎ 3.08
        json2 ██████████▏ 2.55
     jsoniter █████████▎ 2.33
       ffjson ███████▎ 1.84
     easyjson ████▌ 1.15
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        djson >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys generically
         json.Unmarshal        51868342084 ns/op  37195708256 B/op    701321100 allocs/op
        json2.Unmarshal        41304437291 ns/op  35236098312 B/op    579113277 allocs/op
     fastjson.Unmarshal        17808638958 ns/op   3343034320 B/op    142344813 allocs/op
     jsoniter.Unmarshal        44444182166 ns/op  41742626960 B/op    960518791 allocs/op
   jsonparser.Unmarshal        33000001708 ns/op   5277693576 B/op    273513556 allocs/op
        goccy.Unmarshal        38858971000 ns/op  44758760712 B/op    833138811 allocs/op
      segment.Unmarshal       120991613459 ns/op 341672361064 B/op    995514314 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal        31087792833 ns/op  10406634304 B/op    146758262 allocs/op
      gjson-v.Unmarshal        37048535958 ns/op  10406633776 B/op    146758257 allocs/op
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal        69191465959 ns/op  13524660400 B/op    905186207 allocs/op
        djson.Unmarshal        27292381459 ns/op  39791201424 B/op    375169029 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     fastjson ███████████▋ 2.91
        djson ███████▌ 1.90
        gjson ██████▋ 1.67
   jsonparser ██████▎ 1.57
      gjson-v █████▌ 1.40
        goccy █████▎ 1.33
        json2 █████  1.26
     jsoniter ████▋ 1.17
         json ▓▓▓▓ 1.00
        codec ██▉ 0.75
      segment █▋ 0.43
     simdjson >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys into struct
         json.Unmarshal        44308667167 ns/op   4426643632 B/op    137915230 allocs/op
        json2.Unmarshal        26901260709 ns/op   6894523384 B/op     94197532 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal        13866489292 ns/op   4986246416 B/op    150338239 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         8546593541 ns/op   5742405160 B/op       216991 allocs/op
      segment.Unmarshal        19351603500 ns/op   3474082328 B/op    141915744 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec.Unmarshal        18246980333 ns/op   3045058688 B/op    133914577 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal        22954650041 ns/op  10005827984 B/op    215401333 allocs/op
     easyjson.Unmarshal        19775896791 ns/op   8928973888 B/op    198556044 allocs/op

        goccy ████████████████████▋ 5.18
     jsoniter ████████████▊ 3.20
        codec █████████▋ 2.43
      segment █████████▏ 2.29
     easyjson ████████▉ 2.24
       ffjson ███████▋ 1.93
        json2 ██████▌ 1.65
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal custom data through an object builder
         json.Marshal                 1297 ns/op          688 B/op           17 allocs/op
        json2 >>> not supported <<<
     fastjson.Marshal                  807 ns/op          679 B/op            6 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Marshal                 2271 ns/op         3024 B/op           27 allocs/op
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     fastjson ██████▍ 1.61
         json ▓▓▓▓ 1.00
        gjson ██▎ 0.57
   jsonparser >>> not supported <<<
     jsoniter >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        json2 >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Validate []byte
         json.Validate               10461 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
     fastjson.Validate                4941 ns/op            0 B/op            0 allocs/op
     jsoniter.Validate                6672 ns/op         2184 B/op          100 allocs/op
   jsonparser >>> not supported <<<
        goccy.Validate               22950 ns/op        27585 B/op          463 allocs/op
      segment.Validate                3611 ns/op            0 B/op            0 allocs/op
     simdjson >>> not supported <<<
        gjson.Validate                3532 ns/op            0 B/op            0 allocs/op
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        gjson ███████████▊ 2.96
      segment ███████████▌ 2.90
     fastjson ████████▍ 2.12
     jsoniter ██████▎ 1.57
         json ▓▓▓▓ 1.00
        goccy █▊ 0.46
   jsonparser >>> not supported <<<
     simdjson >>> not supported <<<
        json2 >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Validate string
         json.Validate               10478 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
     fastjson.Validate                5349 ns/op         4096 B/op            1 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Validate                3858 ns/op         4096 B/op            1 allocs/op
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        gjson ██████████▊ 2.72
     fastjson ███████▊ 1.96
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
     jsoniter >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        json2 >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
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
