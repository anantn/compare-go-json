# Compare Go JSON

Fork of [ohler55/compare-go-json](https://github.com/ohler55/compare-go-json),
modified to focus on JSON file parsing with dynamic (unknown) schema.

# Results

## x86_64 (Linux)

### Unmarshaling
```
Unmarshal single record (2kb), read few keys generically
        sonic ████████████████████▋ 5.19
     fastjson ██████████████▉ 3.73
   jsonparser ██████████████▍ 3.61
      sonic-v ████████████▋ 3.18
        gjson ████████████▏ 3.04
      gjson-v ████████▏ 2.04
         json ▓▓▓▓ 1.00
Unmarshal single record (2kb), read few keys into struct
   * fastjson ██████████████▊ 3.72
 * jsonparser ██████████████▍ 3.62
        sonic █████████████▎ 3.34
        goccy █████████████▎ 3.32
      * gjson ████████████  3.02
        json2 ██████████▊ 2.69
        codec ██████████▎ 2.58
      sonic-v █████████▌ 2.40
     jsoniter █████████▏ 2.30
    * gjson-v ████████▎ 2.07
      segment ███████  1.75
     easyjson █████▍ 1.35
       ffjson ████▋ 1.18
         json ▓▓▓▓ 1.00
Unmarshal single record (2kb), read all keys generically
     fastjson ████████████▎ 3.08
        djson ███████▋ 1.93
        sonic ███████▌ 1.89
        gjson ██████▉ 1.74
   jsonparser ██████▊ 1.70
      sonic-v ██████▍ 1.62
      gjson-v █████▊ 1.46
        goccy █████▌ 1.40
     jsoniter █████▎ 1.33
        json2 █████▏ 1.30
     simdjson █████▏ 1.29
         json ▓▓▓▓ 1.00
        codec ███▎ 0.83
      segment █▋ 0.41
Unmarshal single record (2kb), read all keys into struct
        sonic ███████████████▌ 3.89
        goccy ██████████████  3.52
      sonic-v ███████████  2.77
     jsoniter ██████████▊ 2.71
     easyjson █████████▎ 2.31
        codec ████████▊ 2.22
        json2 ████████▊ 2.21
       ffjson ███████▌ 1.91
      segment █████▏ 1.31
         json ▓▓▓▓ 1.00

Unmarshal many records (2kb each) from small file (100MB), read few keys generically
        sonic ████████████████████████████████████████████████████▍ 13.12
     simdjson ███████████████████████████████▏ 7.78
   jsonparser ████████████████████████████▋ 7.17
     fastjson ██████████████████████  5.53
        gjson ████████████████████▊ 5.19
      sonic-v ████████████████▏ 4.05
      gjson-v ██████████▋ 2.68
         json ▓▓▓▓ 1.00
Unmarshal many records (2kb each) from small file (100MB), read few keys into struct
   * simdjson ███████████████████████████████▉ 7.99
 * jsonparser █████████████████████████████  7.27
   * fastjson ██████████████████████▋ 5.66
        goccy █████████████████████▋ 5.43
      * gjson █████████████████████▏ 5.29
        sonic █████████████████▏ 4.30
        codec ████████████████▏ 4.03
    * gjson-v ██████████▋ 2.68
      sonic-v ██████████  2.51
        json2 █████████▎ 2.34
     jsoniter █████████  2.25
      segment ████████▍ 2.12
       ffjson ██████▊ 1.71
     easyjson █████▎ 1.32
         json ▓▓▓▓ 1.00
Unmarshal many records (2kb each) from small file (100MB), read all keys generically
     simdjson ██████████████████▎ 4.58
     fastjson ██████████████▋ 3.68
        djson ███████▉ 1.98
        sonic ███████▊ 1.95
        gjson ██████▋ 1.67
      sonic-v ██████▍ 1.60
   jsonparser ██████▎ 1.57
      gjson-v █████▌ 1.40
        goccy █████▎ 1.32
     jsoniter ████▊ 1.20
        json2 ████▋ 1.16
         json ▓▓▓▓ 1.00
        codec ██▊ 0.71
      segment █▎ 0.32
Unmarshal many records (2kb each) from small file (100MB), read all keys into struct
        goccy ███████████████████▋ 4.91
        sonic ██████████████████▋ 4.66
     jsoniter ████████████▌ 3.14
      sonic-v ████████████▏ 3.06
     easyjson ██████████▌ 2.65
        codec █████████▏ 2.28
       ffjson ███████  1.77
        json2 ██████▉ 1.75
      segment ████▊ 1.22
         json ▓▓▓▓ 1.00

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys generically
        sonic ████████████████████████████████████████████████████████████████████████████████████████████████████████████████▌ 28.14
     simdjson █████████████████████████████████████████████████████████████████▉ 16.49
   jsonparser ███████████████████████████████████████▍ 9.85
        gjson █████████████████████████▊ 6.46
     fastjson █████████████████████████▏ 6.30
      sonic-v ███████████████████▏ 4.79
      gjson-v ███████████▌ 2.91
         json ▓▓▓▓ 1.00
Unmarshal many records (25kb each) from semi-large file (5GB), read few keys into struct
   * simdjson █████████████████████████████████████████████████████████████████████████████▍ 19.37
 * jsonparser ███████████████████████████████████████▍ 9.87
      * gjson █████████████████████████▋ 6.41
   * fastjson █████████████████████████▏ 6.31
        goccy ████████████████████████  6.00
        codec ███████████████████▉ 4.98
        sonic ███████████████████  4.76
    * gjson-v ███████████▍ 2.86
      sonic-v ██████████▌ 2.63
      segment █████████▊ 2.46
        json2 █████████▏ 2.29
     jsoniter █████████  2.25
       ffjson ███████▊ 1.95
     easyjson ████▊ 1.21
         json ▓▓▓▓ 1.00
Unmarshal many records from (25kb each) semi-large file (5GB), read all keys generically
     simdjson █████████████████████▍ 5.36
     fastjson ███████████████▊ 3.95
        djson ███████▉ 1.99
        sonic ███████▋ 1.93
      sonic-v ██████▍ 1.61
        gjson ██████▎ 1.58
   jsonparser █████▉ 1.49
      gjson-v █████▎ 1.32
        goccy █████▏ 1.30
        json2 ████▉ 1.22
     jsoniter ████▊ 1.20
         json ▓▓▓▓ 1.00
        codec ██▋ 0.68
      segment █▏ 0.31
Unmarshal many records from (25kb each) semi-large file (5GB), read all keys into struct
        goccy █████████████████████▎ 5.34
        sonic ████████████████████▏ 5.05
      sonic-v ████████████▌ 3.14
     jsoniter ████████████▍ 3.10
     easyjson ██████████▌ 2.65
        codec █████████▎ 2.32
       ffjson ███████▍ 1.86
        json2 ██████▍ 1.61
      segment ████▉ 1.25
         json ▓▓▓▓ 1.00

Unmarshal CanadaGeometry generically
        sonic █████████▋ 2.43
   * simdjson ███████  1.77
 * jsonparser ██████▋ 1.68
      * djson ██████▌ 1.65
        json2 ████▉ 1.22
        goccy ████▍ 1.12
   * fastjson ████▏ 1.06
         json ▓▓▓▓ 1.00
      * gjson ███▎ 0.84
        codec ███▎ 0.83
     jsoniter ███  0.77
      segment ██▍ 0.60
Unmarshal CanadaGeometry into struct
        sonic ████████████████████████▍ 6.11
      segment ██████████▍ 2.60
        goccy ██████████▏ 2.53
     easyjson ████████▍ 2.11
        codec ██████▍ 1.60
        json2 █████▊ 1.46
     jsoniter ████▋ 1.18
         json ▓▓▓▓ 1.00
       ffjson ███▏ 0.80

Unmarshal CITMCatalog generically
   * simdjson ███████████████████▉ 4.97
        sonic ████████████▍ 3.10
      * djson █████████▉ 2.49
        json2 ████████▋ 2.17
   * fastjson ████████▎ 2.08
     jsoniter ███████▋ 1.91
        goccy ██████▉ 1.74
      * gjson ████▉ 1.23
 * jsonparser ████▌ 1.14
         json ▓▓▓▓ 1.00
        codec ███▉ 0.99
      segment █▋ 0.43
Unmarshal CITMCatalog into struct
        sonic ████████████████████████████████████▉ 9.24
        goccy ██████████████████████████▍ 6.61
     jsoniter █████████████████████▉ 5.49
        codec ██████████████  3.51
     easyjson ████████████▉ 3.22
        json2 ████████████▋ 3.16
      segment ████████▊ 2.21
         json ▓▓▓▓ 1.00

Unmarshal SyntheaFHIR generically
   * simdjson ███████████████▌ 3.90
        sonic ████████████▋ 3.16
      * djson █████████▊ 2.45
   * fastjson ███████▊ 1.97
        json2 ███████▎ 1.83
        goccy ██████▋ 1.68
     jsoniter ██████▌ 1.63
      * gjson ████▌ 1.15
         json ▓▓▓▓ 1.00
 * jsonparser ███▊ 0.96
        codec ███▍ 0.85
      segment ▉ 0.22
Unmarshal SyntheaFHIR into struct
        sonic ███████████████████████████  6.76
        goccy █████████████████████▋ 5.42
     jsoniter ███████████████▉ 3.98
        codec ████████████▎ 3.09
     easyjson ████████████▏ 3.04
        json2 ██████████▏ 2.56
      segment ██████  1.50
         json ▓▓▓▓ 1.00
       ffjson ██▎ 0.58

Unmarshal TwitterStatus generically
   * simdjson █████████████████▍ 4.37
   * fastjson ████████████▋ 3.18
        sonic ███████████▉ 2.98
      * djson ████████▌ 2.15
        goccy ██████▊ 1.72
      * gjson ██████▋ 1.66
        json2 ██████▍ 1.60
     jsoniter ██████▎ 1.58
 * jsonparser █████▊ 1.46
         json ▓▓▓▓ 1.00
        codec ███▋ 0.93
      segment █▏ 0.30
Unmarshal TwitterStatus into struct
        sonic ███████████████████████████████▍ 7.86
        goccy █████████████████████▏ 5.31
     easyjson ██████████████████▎ 4.58
     jsoniter ███████████████▊ 3.95
        codec █████████████▎ 3.32
        json2 ██████████▉ 2.73
       ffjson ██████▋ 1.66
      segment █████▍ 1.34
         json ▓▓▓▓ 1.00

Unmarshal GolangSource generically
        sonic ██████████▍ 2.60
      * djson ██████████▏ 2.54
   * simdjson ████████▋ 2.18
        json2 █████▋ 1.42
     jsoniter ████▊ 1.20
   * fastjson ████▊ 1.20
        goccy ████▋ 1.17
         json ▓▓▓▓ 1.00
        codec ██▍ 0.61
 * jsonparser █▌ 0.39
      * gjson █▍ 0.36
      segment █▏ 0.31
Unmarshal GolangSource into struct
        sonic █████████████████████████▋ 6.41
        goccy ██████████████████████▋ 5.68
      segment █████████████████████▊ 5.44
     jsoniter ████████████▌ 3.14
     easyjson ████████████  3.01
        codec █████████▏ 2.31
        json2 ███████▊ 1.96
       ffjson ███████▍ 1.87
         json ▓▓▓▓ 1.00

Unmarshal StringUnicode generically
   * fastjson █████████████████████████████████████████▊ 10.46
        sonic ████████████████████████████████▍ 8.09
 * jsonparser ██████████████████████████▋ 6.68
      * gjson ███████████████████████▉ 5.97
        goccy ███████████████▌ 3.89
     jsoniter ███████████▋ 2.93
        codec ██████████▍ 2.62
        json2 ████████▋ 2.16
   * simdjson ███████  1.78
      * djson ██████▋ 1.67
         json ▓▓▓▓ 1.00
      segment ██▏ 0.54
Unmarshal StringUnicode into struct
        sonic █████████████████████████████████████████████████████████████████████████▉ 18.49
     easyjson ██████████████████████████████████████████████████████▏ 13.53
        goccy ████████████████████████▊ 6.20
       ffjson █████████████████████▌ 5.41
        codec █████████████████████▏ 5.30
     jsoniter █████████████████  4.26
        json2 ██████████▏ 2.55
         json ▓▓▓▓ 1.00
      segment ███▋ 0.92
```

### Marshaling
```
Marshal custom data through an object builder
     fastjson ███████▋ 1.91
        sonic █████▊ 1.45
         json ▓▓▓▓ 1.00
        gjson ██▌ 0.63

Marshal CanadaGeometry generically
     jsoniter █████████▉ 2.50
        sonic █████████▉ 2.49
      segment █████▊ 1.46
        json2 █████▎ 1.32
        goccy █████▏ 1.29
        codec ████▉ 1.25
         json ▓▓▓▓ 1.00
Marshal CanadaGeometry from struct
     jsoniter ████████████████▎ 4.07
        sonic █████████▍ 2.35
     easyjson ████▌ 1.14
      segment ████▌ 1.13
        goccy ████▍ 1.11
        json2 ████  1.00
         json ▓▓▓▓ 1.00
        codec ███▍ 0.85
       ffjson ██  0.52

Marshal CITMCatalog generically
        sonic █████████████▍ 3.36
        json2 █████████▊ 2.44
      segment █████████▍ 2.35
        codec ███████▎ 1.82
        goccy ███████  1.78
     jsoniter ██████▋ 1.66
         json ▓▓▓▓ 1.00
Marshal CITMCatalog from struct
        sonic ██████████▉ 2.72
        goccy ███████▌ 1.88
      segment ██████▏ 1.54
     easyjson █████▌ 1.40
     jsoniter █████▏ 1.29
         json ▓▓▓▓ 1.00
       ffjson ██▊ 0.69
        json2 ██▋ 0.66
        codec ██▌ 0.63

Marshal SyntheaFHIR generically
        sonic ███████████▊ 2.96
        json2 █████████▍ 2.36
        codec █████████▏ 2.29
      segment ████████▋ 2.19
     jsoniter ██████▊ 1.72
        goccy █████▊ 1.44
         json ▓▓▓▓ 1.00
Marshal SyntheaFHIR from struct
        sonic ██████████▍ 2.62
      segment ██████████  2.53
     easyjson ████████▎ 2.07
        goccy ███████▊ 1.94
     jsoniter █████▉ 1.48
       ffjson █████▎ 1.34
        codec ████▍ 1.09
        json2 ████  1.02
         json ▓▓▓▓ 1.00

Marshal TwitterStatus generically
        sonic ████████████████▍ 4.11
        json2 ██████████▊ 2.70
     jsoniter ██████████▌ 2.64
        codec █████████▌ 2.38
      segment ███████  1.76
        goccy █████▎ 1.32
         json ▓▓▓▓ 1.00
Marshal TwitterStatus from struct
        sonic ██████████████▊ 3.71
        goccy ██████  1.50
      segment █████▊ 1.46
     easyjson ████▋ 1.17
         json ▓▓▓▓ 1.00
     jsoniter ███▊ 0.95
       ffjson ███  0.78
        json2 ██▉ 0.73
        codec ██▏ 0.55

Marshal GolangSource generically
        sonic ██████████████████  4.52
     jsoniter ██████████▎ 2.57
        json2 ██████████▎ 2.56
      segment ████████▎ 2.09
        codec ███████▍ 1.87
        goccy █████▌ 1.39
         json ▓▓▓▓ 1.00
Marshal GolangSource from struct
        sonic ███████████▋ 2.91
      segment ██████▏ 1.56
        goccy ██████▏ 1.53
     jsoniter █████▋ 1.43
     easyjson █████  1.28
         json ▓▓▓▓ 1.00
        json2 ██▉ 0.74
        codec ██▎ 0.57
       ffjson ██  0.52

Marshal StringUnicode generically
        sonic ██████████████████████████████████  8.52
     jsoniter █████████▏ 2.29
      segment ████▉ 1.22
        goccy ████▊ 1.22
        codec ████▌ 1.15
        json2 ████▎ 1.07
         json ▓▓▓▓ 1.00
Marshal StringUnicode from struct
        sonic ███████████████████████████████▉ 7.97
     jsoniter ██████▍ 1.62
     easyjson ████▋ 1.17
        goccy ████▎ 1.08
        json2 ████▎ 1.08
      segment ████  1.02
         json ▓▓▓▓ 1.00
       ffjson ███▊ 0.95
        codec ███  0.77
```

### Validation
```
Validate []byte
        sonic █████████████▏ 3.31
        gjson ██████████▉ 2.73
     fastjson █████████▋ 2.43
      segment █████▉ 1.48
     jsoniter █████▊ 1.46
         json ▓▓▓▓ 1.00
        goccy █▌ 0.40

Validate string
        sonic ██████████▋ 2.67
        gjson █████████▍ 2.36
     fastjson ████████▎ 2.06
         json ▓▓▓▓ 1.00
```

# Raw Benchmark Data

Outputs are also available as CSV in the repository.

## x86_64 (Linux)
```
Unmarshal CanadaGeometry generically
         json.Unmarshal            2313566 ns/op      1339377 B/op        52215 allocs/op
        json2.Unmarshal            1897350 ns/op      1011586 B/op        37898 allocs/op
     fastjson.Unmarshal            2178808 ns/op      8725365 B/op        30278 allocs/op
     jsoniter.Unmarshal            3019851 ns/op      1961666 B/op        84971 allocs/op
   jsonparser.Unmarshal            1375731 ns/op          152 B/op           12 allocs/op
        goccy.Unmarshal            2061001 ns/op      1470284 B/op        59189 allocs/op
      segment.Unmarshal            3849684 ns/op      2119755 B/op        51680 allocs/op
     simdjson.Unmarshal            1308406 ns/op       987184 B/op        14324 allocs/op
        gjson.Unmarshal            2749807 ns/op       393072 B/op        14314 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal             953804 ns/op      2817017 B/op        29654 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            2772804 ns/op      1765858 B/op        37408 allocs/op
        djson.Unmarshal            1402795 ns/op      1289906 B/op        37891 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic █████████▋ 2.43
   * simdjson ███████  1.77
 * jsonparser ██████▋ 1.68
      * djson ██████▌ 1.65
        json2 ████▉ 1.22
        goccy ████▍ 1.12
   * fastjson ████▏ 1.06
         json ▓▓▓▓ 1.00
      * gjson ███▎ 0.84
        codec ███▎ 0.83
     jsoniter ███  0.77
      segment ██▍ 0.60
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal CanadaGeometry into struct
         json.Unmarshal            1990968 ns/op          472 B/op           13 allocs/op
        json2.Unmarshal            1368228 ns/op       348002 B/op         1626 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            1683099 ns/op       598821 B/op        25122 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal             786267 ns/op       278713 B/op            1 allocs/op
      segment.Unmarshal             766918 ns/op           48 B/op            4 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal             325839 ns/op       279058 B/op            2 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            1242610 ns/op         1224 B/op            7 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal            2492215 ns/op       536102 B/op           25 allocs/op
     easyjson.Unmarshal             943813 ns/op       330368 B/op          753 allocs/op

        sonic ████████████████████████▍ 6.11
      segment ██████████▍ 2.60
        goccy ██████████▏ 2.53
     easyjson ████████▍ 2.11
        codec ██████▍ 1.60
        json2 █████▊ 1.46
     jsoniter ████▋ 1.18
         json ▓▓▓▓ 1.00
       ffjson ███▏ 0.80
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal CanadaGeometry generically
         json.Marshal              1178992 ns/op       280216 B/op           21 allocs/op
        json2.Marshal               893043 ns/op       279678 B/op            7 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               471628 ns/op       172859 B/op           13 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal               913525 ns/op       280849 B/op            3 allocs/op
      segment.Marshal               806160 ns/op       279433 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               473072 ns/op       279149 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal               944891 ns/op      1186328 B/op           22 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     jsoniter █████████▉ 2.50
        sonic █████████▉ 2.49
      segment █████▊ 1.46
        json2 █████▎ 1.32
        goccy █████▏ 1.29
        codec ████▉ 1.25
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

Marshal CanadaGeometry from struct
         json.Marshal               814405 ns/op       279299 B/op            1 allocs/op
        json2.Marshal               814395 ns/op       279396 B/op            1 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               200106 ns/op       172185 B/op            2 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal               731504 ns/op       279311 B/op            1 allocs/op
      segment.Marshal               720733 ns/op       279372 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               347097 ns/op       279256 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal               955285 ns/op      1186328 B/op           22 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal              1567515 ns/op      1060643 B/op        14391 allocs/op
     easyjson.Marshal               714184 ns/op       280437 B/op           20 allocs/op

     jsoniter ████████████████▎ 4.07
        sonic █████████▍ 2.35
     easyjson ████▌ 1.14
      segment ████▌ 1.13
        goccy ████▍ 1.11
        json2 ████  1.00
         json ▓▓▓▓ 1.00
        codec ███▍ 0.85
       ffjson ██  0.52
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal CITMCatalog generically
         json.Unmarshal            7891036 ns/op      5123772 B/op        95373 allocs/op
        json2.Unmarshal            3640346 ns/op      4958955 B/op        80513 allocs/op
     fastjson.Unmarshal            3794960 ns/op     19991461 B/op        45701 allocs/op
     jsoniter.Unmarshal            4123283 ns/op      5573448 B/op       118756 allocs/op
   jsonparser.Unmarshal            6928457 ns/op       145320 B/op        15863 allocs/op
        goccy.Unmarshal            4535191 ns/op      7416581 B/op       123577 allocs/op
      segment.Unmarshal           18408851 ns/op     55526618 B/op       123714 allocs/op
     simdjson.Unmarshal            1587640 ns/op      2854121 B/op        15139 allocs/op
        gjson.Unmarshal            6440498 ns/op      1855522 B/op        15131 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal            2542329 ns/op      8645717 B/op        58177 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            7985000 ns/op      6195089 B/op       159476 allocs/op
        djson.Unmarshal            3163451 ns/op      6437632 B/op        54480 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

   * simdjson ███████████████████▉ 4.97
        sonic ████████████▍ 3.10
      * djson █████████▉ 2.49
        json2 ████████▋ 2.17
   * fastjson ████████▎ 2.08
     jsoniter ███████▋ 1.91
        goccy ██████▉ 1.74
      * gjson ████▉ 1.23
 * jsonparser ████▌ 1.14
         json ▓▓▓▓ 1.00
        codec ███▉ 0.99
      segment █▋ 0.43
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal CITMCatalog into struct
         json.Unmarshal            7235083 ns/op       253352 B/op        10950 allocs/op
        json2.Unmarshal            2286036 ns/op       888615 B/op         6295 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            1317186 ns/op        72338 B/op         2535 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal            1094338 ns/op      1782809 B/op         2226 allocs/op
      segment.Unmarshal            3267854 ns/op        52384 B/op         1354 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal             782852 ns/op      1731588 B/op          111 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            2063138 ns/op        47880 B/op         1036 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal       >>> ffjson error: (*errors.errorString)cannot unmarshal tok:string into Go value for int64 offset=40 line=3 char=19 <<<
     easyjson.Unmarshal            2247358 ns/op       896720 B/op         5705 allocs/op

        sonic ████████████████████████████████████▉ 9.24
        goccy ██████████████████████████▍ 6.61
     jsoniter █████████████████████▉ 5.49
        codec ██████████████  3.51
     easyjson ████████████▉ 3.22
        json2 ████████████▋ 3.16
      segment ████████▊ 2.21
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
      sonic-v >>> not supported <<<
     simdjson >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> ffjson error: (*errors.errorString)cannot unmarshal tok:string into Go value for int64 offset=40 line=3 char=19 <<<
     fastjson >>> not supported <<<

Marshal CITMCatalog generically
         json.Marshal              3735122 ns/op      2388150 B/op        62674 allocs/op
        json2.Marshal              1529819 ns/op       508227 B/op           16 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              2243359 ns/op      1912889 B/op        32812 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              2099306 ns/op       652344 B/op          323 allocs/op
      segment.Marshal              1589223 ns/op       511621 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              1110816 ns/op       509290 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              2057812 ns/op      2538009 B/op           25 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic █████████████▍ 3.36
        json2 █████████▊ 2.44
      segment █████████▍ 2.35
        codec ███████▎ 1.82
        goccy ███████  1.78
     jsoniter ██████▋ 1.66
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
         json.Marshal               564679 ns/op       556059 B/op          890 allocs/op
        json2.Marshal               852915 ns/op       512995 B/op          133 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               436962 ns/op       509160 B/op           22 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal               300540 ns/op       521291 B/op            3 allocs/op
      segment.Marshal               366994 ns/op       547964 B/op          616 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               207592 ns/op       508415 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal               891764 ns/op      2538009 B/op           25 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal               818540 ns/op       903518 B/op        14180 allocs/op
     easyjson.Marshal               403581 ns/op       524344 B/op         1125 allocs/op

        sonic ██████████▉ 2.72
        goccy ███████▌ 1.88
      segment ██████▏ 1.54
     easyjson █████▌ 1.40
     jsoniter █████▏ 1.29
         json ▓▓▓▓ 1.00
       ffjson ██▊ 0.69
        json2 ██▋ 0.66
        codec ██▌ 0.63
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal SyntheaFHIR generically
         json.Unmarshal            9965366 ns/op      7430179 B/op       134383 allocs/op
        json2.Unmarshal            5442546 ns/op      6938850 B/op       115831 allocs/op
     fastjson.Unmarshal            5068888 ns/op     26046547 B/op        66098 allocs/op
     jsoniter.Unmarshal            6115106 ns/op      8197420 B/op       181426 allocs/op
   jsonparser.Unmarshal           10377395 ns/op      1108792 B/op        49604 allocs/op
        goccy.Unmarshal            5917595 ns/op      9840037 B/op       159460 allocs/op
      segment.Unmarshal           44705283 ns/op     77571246 B/op       193338 allocs/op
     simdjson.Unmarshal            2555805 ns/op      6172199 B/op        26181 allocs/op
        gjson.Unmarshal            8631546 ns/op      2432544 B/op        26821 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal            3155057 ns/op      9488219 B/op        67614 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal           11708820 ns/op      9669992 B/op       247797 allocs/op
        djson.Unmarshal            4061064 ns/op      8374303 B/op        68393 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

   * simdjson ███████████████▌ 3.90
        sonic ████████████▋ 3.16
      * djson █████████▊ 2.45
   * fastjson ███████▊ 1.97
        json2 ███████▎ 1.83
        goccy ██████▋ 1.68
     jsoniter ██████▌ 1.63
      * gjson ████▌ 1.15
         json ▓▓▓▓ 1.00
 * jsonparser ███▊ 0.96
        codec ███▍ 0.85
      segment ▉ 0.22
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal SyntheaFHIR into struct
         json.Unmarshal            9466104 ns/op       625394 B/op        21883 allocs/op
        json2.Unmarshal            3699256 ns/op      3103377 B/op        14550 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            2380812 ns/op       825759 B/op        29439 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal            1746506 ns/op      2157955 B/op         2508 allocs/op
      segment.Unmarshal            6307028 ns/op       623760 B/op        21832 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal            1400355 ns/op      2017624 B/op           10 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            3062872 ns/op       623400 B/op        21826 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal           16231142 ns/op      4334447 B/op        21904 allocs/op
     easyjson.Unmarshal            3113983 ns/op      3531634 B/op        29214 allocs/op

        sonic ███████████████████████████  6.76
        goccy █████████████████████▋ 5.42
     jsoniter ███████████████▉ 3.98
        codec ████████████▎ 3.09
     easyjson ████████████▏ 3.04
        json2 ██████████▏ 2.56
      segment ██████  1.50
         json ▓▓▓▓ 1.00
       ffjson ██▎ 0.58
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal SyntheaFHIR generically
         json.Marshal              5866989 ns/op      4138929 B/op        95282 allocs/op
        json2.Marshal              2490215 ns/op      1162067 B/op            8 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              3418891 ns/op      3091033 B/op        45292 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              4066664 ns/op      2501465 B/op         6476 allocs/op
      segment.Marshal              2683452 ns/op      1162623 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              1980233 ns/op      1151987 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              2565726 ns/op      6642203 B/op           29 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ███████████▊ 2.96
        json2 █████████▍ 2.36
        codec █████████▏ 2.29
      segment ████████▋ 2.19
     jsoniter ██████▊ 1.72
        goccy █████▊ 1.44
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
         json.Marshal              5442506 ns/op      3985935 B/op        18009 allocs/op
        json2.Marshal              5344803 ns/op      3141699 B/op         2569 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              3668637 ns/op      4023484 B/op        17963 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              2802921 ns/op      3945524 B/op        17940 allocs/op
      segment.Marshal              2151809 ns/op      3335969 B/op         5338 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              2075850 ns/op      3981986 B/op        17939 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              4974450 ns/op     16792092 B/op           33 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal              4074707 ns/op      9497582 B/op        45596 allocs/op
     easyjson.Marshal              2635408 ns/op      3979747 B/op        15520 allocs/op

        sonic ██████████▍ 2.62
      segment ██████████  2.53
     easyjson ████████▎ 2.07
        goccy ███████▊ 1.94
     jsoniter █████▉ 1.48
       ffjson █████▎ 1.34
        codec ████▍ 1.09
        json2 ████  1.02
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal TwitterStatus generically
         json.Unmarshal            3248842 ns/op      2152389 B/op        31262 allocs/op
        json2.Unmarshal            2032587 ns/op      2074213 B/op        27236 allocs/op
     fastjson.Unmarshal            1020751 ns/op      5341211 B/op        11130 allocs/op
     jsoniter.Unmarshal            2059132 ns/op      2427239 B/op        45041 allocs/op
   jsonparser.Unmarshal            2231782 ns/op       300312 B/op        10184 allocs/op
        goccy.Unmarshal            1893955 ns/op      2756016 B/op        39695 allocs/op
      segment.Unmarshal           10770461 ns/op      7060650 B/op        39237 allocs/op
     simdjson.Unmarshal             743946 ns/op      1950210 B/op         5732 allocs/op
        gjson.Unmarshal            1958590 ns/op       822360 B/op         6884 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal            1091362 ns/op      2606743 B/op        12137 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            3511145 ns/op      2577333 B/op        60987 allocs/op
        djson.Unmarshal            1510992 ns/op      2470377 B/op        12678 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

   * simdjson █████████████████▍ 4.37
   * fastjson ████████████▋ 3.18
        sonic ███████████▉ 2.98
      * djson ████████▌ 2.15
        goccy ██████▊ 1.72
      * gjson ██████▋ 1.66
        json2 ██████▍ 1.60
     jsoniter ██████▎ 1.58
 * jsonparser █████▊ 1.46
         json ▓▓▓▓ 1.00
        codec ███▋ 0.93
      segment █▏ 0.30
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal TwitterStatus into struct
         json.Unmarshal            2972638 ns/op       304611 B/op         6486 allocs/op
        json2.Unmarshal            1087872 ns/op       306194 B/op         2972 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal             752441 ns/op       339699 B/op         6293 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal             560189 ns/op       647187 B/op          238 allocs/op
      segment.Unmarshal            2210263 ns/op       369712 B/op         5341 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal             378313 ns/op       694537 B/op          401 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal             895679 ns/op       215944 B/op         4530 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal            1786008 ns/op      1052357 B/op         9494 allocs/op
     easyjson.Unmarshal             648800 ns/op       324336 B/op         4907 allocs/op

        sonic ███████████████████████████████▍ 7.86
        goccy █████████████████████▏ 5.31
     easyjson ██████████████████▎ 4.58
     jsoniter ███████████████▊ 3.95
        codec █████████████▎ 3.32
        json2 ██████████▉ 2.73
       ffjson ██████▋ 1.66
      segment █████▍ 1.34
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal TwitterStatus generically
         json.Marshal              2290499 ns/op      1490170 B/op        27955 allocs/op
        json2.Marshal               849899 ns/op       469085 B/op            7 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               867351 ns/op       630770 B/op         3793 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              1741621 ns/op      1050923 B/op          543 allocs/op
      segment.Marshal              1302591 ns/op       475205 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               557744 ns/op       468452 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal               960904 ns/op      2538008 B/op           25 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ████████████████▍ 4.11
        json2 ██████████▊ 2.70
     jsoniter ██████████▌ 2.64
        codec █████████▌ 2.38
      segment ███████  1.76
        goccy █████▎ 1.32
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
         json.Marshal               425300 ns/op       512890 B/op          141 allocs/op
        json2.Marshal               578987 ns/op       503226 B/op           59 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal               446949 ns/op       504424 B/op           52 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal               283066 ns/op       509655 B/op            6 allocs/op
      segment.Marshal               291742 ns/op       512436 B/op          121 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               114569 ns/op       500170 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal               776190 ns/op      2540384 B/op           29 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal               545616 ns/op       799315 B/op         6117 allocs/op
     easyjson.Marshal               362153 ns/op       519988 B/op          953 allocs/op

        sonic ██████████████▊ 3.71
        goccy ██████  1.50
      segment █████▊ 1.46
     easyjson ████▋ 1.17
         json ▓▓▓▓ 1.00
     jsoniter ███▊ 0.95
       ffjson ███  0.78
        json2 ██▉ 0.73
        codec ██▏ 0.55
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal GolangSource generically
         json.Unmarshal           12907638 ns/op      7916353 B/op       271264 allocs/op
        json2.Unmarshal            9082002 ns/op      6990199 B/op       218839 allocs/op
     fastjson.Unmarshal           10765831 ns/op     50767009 B/op       130330 allocs/op
     jsoniter.Unmarshal           10748110 ns/op      9233841 B/op       346658 allocs/op
   jsonparser.Unmarshal           32725343 ns/op       925272 B/op        64023 allocs/op
        goccy.Unmarshal           11028293 ns/op     10873829 B/op       320840 allocs/op
      segment.Unmarshal           41535574 ns/op     67095381 B/op       334022 allocs/op
     simdjson.Unmarshal            5933364 ns/op     12653663 B/op        63942 allocs/op
        gjson.Unmarshal           35367577 ns/op      2658512 B/op        76819 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal            4957385 ns/op     10998793 B/op       128226 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal           21306783 ns/op     10502697 B/op       436539 allocs/op
        djson.Unmarshal            5081338 ns/op      7822418 B/op       117489 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ██████████▍ 2.60
      * djson ██████████▏ 2.54
   * simdjson ████████▋ 2.18
        json2 █████▋ 1.42
     jsoniter ████▊ 1.20
   * fastjson ████▊ 1.20
        goccy ████▋ 1.17
         json ▓▓▓▓ 1.00
        codec ██▍ 0.61
 * jsonparser █▌ 0.39
      * gjson █▍ 0.36
      segment █▏ 0.31
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal GolangSource into struct
         json.Unmarshal           12830100 ns/op       704424 B/op        24927 allocs/op
        json2.Unmarshal            6533352 ns/op      3236444 B/op        13941 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal            4089587 ns/op       994092 B/op        37077 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal            2257459 ns/op      1944071 B/op            1 allocs/op
      segment.Unmarshal            2358913 ns/op       411672 B/op        12806 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal            2002070 ns/op      1942226 B/op            2 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal            5562520 ns/op       412848 B/op        12809 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal            6858951 ns/op      6276805 B/op        52765 allocs/op
     easyjson.Unmarshal            4266860 ns/op      4282425 B/op        27846 allocs/op

        sonic █████████████████████████▋ 6.41
        goccy ██████████████████████▋ 5.68
      segment █████████████████████▊ 5.44
     jsoniter ████████████▌ 3.14
     easyjson ████████████  3.01
        codec █████████▏ 2.31
        json2 ███████▊ 1.96
       ffjson ███████▍ 1.87
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal GolangSource generically
         json.Marshal             12336126 ns/op      8546376 B/op       192096 allocs/op
        json2.Marshal              4812988 ns/op      1987843 B/op            7 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              4799212 ns/op      3477614 B/op        38422 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              8900485 ns/op      1942023 B/op            2 allocs/op
      segment.Marshal              5896386 ns/op      1995725 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal              2731629 ns/op      1957562 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              6614227 ns/op     10590745 B/op           31 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ██████████████████  4.52
     jsoniter ██████████▎ 2.57
        json2 ██████████▎ 2.56
      segment ████████▎ 2.09
        codec ███████▍ 1.87
        goccy █████▌ 1.39
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

Marshal GolangSource from struct
         json.Marshal              2336490 ns/op      1941587 B/op            1 allocs/op
        json2.Marshal              3157060 ns/op      1971013 B/op            1 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal              1636906 ns/op      1794133 B/op            2 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal              1525652 ns/op      1957054 B/op            1 allocs/op
      segment.Marshal              1500916 ns/op      1941585 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal               803785 ns/op      1946026 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal              4073694 ns/op     10590745 B/op           31 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal              4535635 ns/op      6490605 B/op        77031 allocs/op
     easyjson.Marshal              1824893 ns/op      1996780 B/op           73 allocs/op

        sonic ███████████▋ 2.91
      segment ██████▏ 1.56
        goccy ██████▏ 1.53
     jsoniter █████▋ 1.43
     easyjson █████  1.28
         json ▓▓▓▓ 1.00
        json2 ██▉ 0.74
        codec ██▎ 0.57
       ffjson ██  0.52
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal StringUnicode generically
         json.Unmarshal              58678 ns/op        30307 B/op          194 allocs/op
        json2.Unmarshal              27133 ns/op        27555 B/op          179 allocs/op
     fastjson.Unmarshal               5609 ns/op        35208 B/op           76 allocs/op
     jsoniter.Unmarshal              20007 ns/op        31958 B/op          262 allocs/op
   jsonparser.Unmarshal               8784 ns/op        18624 B/op          121 allocs/op
        goccy.Unmarshal              15077 ns/op        31322 B/op          191 allocs/op
      segment.Unmarshal             108277 ns/op        43274 B/op          307 allocs/op
     simdjson.Unmarshal              32971 ns/op       194760 B/op           75 allocs/op
        gjson.Unmarshal               9827 ns/op        20272 B/op           66 allocs/op
      gjson-v >>> not supported <<<
        sonic.Unmarshal               7249 ns/op        29864 B/op           72 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal              22431 ns/op        32852 B/op          315 allocs/op
        djson.Unmarshal              35221 ns/op        47483 B/op          133 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

   * fastjson █████████████████████████████████████████▊ 10.46
        sonic ████████████████████████████████▍ 8.09
 * jsonparser ██████████████████████████▋ 6.68
      * gjson ███████████████████████▉ 5.97
        goccy ███████████████▌ 3.89
     jsoniter ███████████▋ 2.93
        codec ██████████▍ 2.62
        json2 ████████▋ 2.16
   * simdjson ███████  1.78
      * djson ██████▋ 1.67
         json ▓▓▓▓ 1.00
      segment ██▏ 0.54
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal StringUnicode into struct
         json.Unmarshal              55368 ns/op        18296 B/op           66 allocs/op
        json2.Unmarshal              21731 ns/op        15473 B/op           49 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal              12982 ns/op        18912 B/op           72 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               8936 ns/op        18432 B/op            1 allocs/op
      segment.Unmarshal              60478 ns/op        35280 B/op          122 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               2995 ns/op        18912 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Unmarshal              10441 ns/op        18824 B/op           63 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal              10243 ns/op        18449 B/op           66 allocs/op
     easyjson.Unmarshal               4091 ns/op        17776 B/op           61 allocs/op

        sonic █████████████████████████████████████████████████████████████████████████▉ 18.49
     easyjson ██████████████████████████████████████████████████████▏ 13.53
        goccy ████████████████████████▊ 6.20
       ffjson █████████████████████▌ 5.41
        codec █████████████████████▏ 5.30
     jsoniter █████████████████  4.26
        json2 ██████████▏ 2.55
         json ▓▓▓▓ 1.00
      segment ███▋ 0.92
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal StringUnicode generically
         json.Marshal                27344 ns/op        23044 B/op          122 allocs/op
        json2.Marshal                25530 ns/op        19562 B/op           65 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal                11932 ns/op        18561 B/op            4 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                22475 ns/op        18440 B/op            1 allocs/op
      segment.Marshal                22356 ns/op        18435 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal                 3211 ns/op        18522 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                23788 ns/op        84504 B/op           13 allocs/op
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ██████████████████████████████████  8.52
     jsoniter █████████▏ 2.29
      segment ████▉ 1.22
        goccy ████▊ 1.22
        codec ████▌ 1.15
        json2 ████▎ 1.07
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

Marshal StringUnicode from struct
         json.Marshal                17843 ns/op        18433 B/op            1 allocs/op
        json2.Marshal                16584 ns/op        18436 B/op            1 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Marshal                10992 ns/op        18442 B/op            2 allocs/op
   jsonparser >>> not supported <<<
        goccy.Marshal                16549 ns/op        18434 B/op            1 allocs/op
      segment.Marshal                17488 ns/op        18434 B/op            1 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Marshal                 2238 ns/op        18532 B/op            4 allocs/op
      sonic-v >>> not supported <<<
        codec.Marshal                23117 ns/op        84504 B/op           13 allocs/op
        djson >>> not supported <<<
       ffjson.Marshal                18692 ns/op        50615 B/op           67 allocs/op
     easyjson.Marshal                15313 ns/op        19274 B/op           11 allocs/op

        sonic ███████████████████████████████▉ 7.97
     jsoniter ██████▍ 1.62
     easyjson ████▋ 1.17
        goccy ████▎ 1.08
        json2 ████▎ 1.08
      segment ████  1.02
         json ▓▓▓▓ 1.00
       ffjson ███▊ 0.95
        codec ███  0.77
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
      sonic-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal single record (2kb), read few keys generically
         json.Unmarshal              17372 ns/op        12488 B/op           21 allocs/op
        json2 >>> not supported <<<
     fastjson.Unmarshal               4656 ns/op        12034 B/op            7 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal               4817 ns/op        12208 B/op           17 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Unmarshal               5721 ns/op        16368 B/op            9 allocs/op
      gjson-v.Unmarshal               8514 ns/op        16368 B/op            9 allocs/op
        sonic.Unmarshal               3350 ns/op         4851 B/op           12 allocs/op
      sonic-v.Unmarshal               5456 ns/op         4851 B/op           12 allocs/op
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ████████████████████▋ 5.19
     fastjson ██████████████▉ 3.73
   jsonparser ██████████████▍ 3.61
      sonic-v ████████████▋ 3.18
        gjson ████████████▏ 3.04
      gjson-v ████████▏ 2.04
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

Unmarshal single record (2kb), read few keys into struct
         json.Unmarshal              17356 ns/op        12488 B/op           21 allocs/op
        json2.Unmarshal               6441 ns/op          128 B/op            6 allocs/op
     fastjson.Unmarshal               4669 ns/op        12080 B/op           10 allocs/op
     jsoniter.Unmarshal               7543 ns/op        14056 B/op           84 allocs/op
   jsonparser.Unmarshal               4801 ns/op        12208 B/op           17 allocs/op
        goccy.Unmarshal               5228 ns/op        16131 B/op            7 allocs/op
      segment.Unmarshal               9898 ns/op        12064 B/op           12 allocs/op
     simdjson >>> not supported <<<
        gjson.Unmarshal               5746 ns/op        16400 B/op           10 allocs/op
      gjson-v.Unmarshal               8379 ns/op        16400 B/op           10 allocs/op
        sonic.Unmarshal               5202 ns/op         8645 B/op           11 allocs/op
      sonic-v.Unmarshal               7232 ns/op         8648 B/op           11 allocs/op
        codec.Unmarshal               6734 ns/op        13240 B/op           15 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal              14711 ns/op        14185 B/op           36 allocs/op
     easyjson.Unmarshal              12842 ns/op        12320 B/op           16 allocs/op

   * fastjson ██████████████▊ 3.72
 * jsonparser ██████████████▍ 3.62
        sonic █████████████▎ 3.34
        goccy █████████████▎ 3.32
      * gjson ████████████  3.02
        json2 ██████████▊ 2.69
        codec ██████████▎ 2.58
      sonic-v █████████▌ 2.40
     jsoniter █████████▏ 2.30
    * gjson-v ████████▎ 2.07
      segment ███████  1.75
     easyjson █████▍ 1.35
       ffjson ████▋ 1.18
         json ▓▓▓▓ 1.00
     simdjson >>> not supported <<<
        djson >>> not supported <<<

Unmarshal single record (2kb), read all keys generically
         json.Unmarshal              22999 ns/op        29776 B/op          340 allocs/op
        json2.Unmarshal              17723 ns/op        15538 B/op          295 allocs/op
     fastjson.Unmarshal               7467 ns/op        13608 B/op           73 allocs/op
     jsoniter.Unmarshal              17312 ns/op        31681 B/op          457 allocs/op
   jsonparser.Unmarshal              13514 ns/op        14528 B/op          135 allocs/op
        goccy.Unmarshal              16372 ns/op        34514 B/op          398 allocs/op
      segment.Unmarshal              55446 ns/op       173240 B/op          475 allocs/op
     simdjson.Unmarshal              17764 ns/op       132224 B/op           82 allocs/op
        gjson.Unmarshal              13218 ns/op        18336 B/op           76 allocs/op
      gjson-v.Unmarshal              15768 ns/op        18336 B/op           76 allocs/op
        sonic.Unmarshal              12175 ns/op        27019 B/op           59 allocs/op
      sonic-v.Unmarshal              14212 ns/op        27016 B/op           59 allocs/op
        codec.Unmarshal              27610 ns/op        19456 B/op          435 allocs/op
        djson.Unmarshal              11895 ns/op        32168 B/op          183 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     fastjson ████████████▎ 3.08
        djson ███████▋ 1.93
        sonic ███████▌ 1.89
        gjson ██████▉ 1.74
   jsonparser ██████▊ 1.70
      sonic-v ██████▍ 1.62
      gjson-v █████▊ 1.46
        goccy █████▌ 1.40
     jsoniter █████▎ 1.33
        json2 █████▏ 1.30
     simdjson █████▏ 1.29
         json ▓▓▓▓ 1.00
        codec ███▎ 0.83
      segment █▋ 0.41
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal single record (2kb), read all keys into struct
         json.Unmarshal              21803 ns/op        14608 B/op           80 allocs/op
        json2.Unmarshal               9868 ns/op         3104 B/op           35 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal               8058 ns/op        14384 B/op           77 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal               6200 ns/op        16139 B/op            7 allocs/op
      segment.Unmarshal              16704 ns/op        13680 B/op           73 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal               5607 ns/op         9225 B/op            9 allocs/op
      sonic-v.Unmarshal               7866 ns/op         9241 B/op            9 allocs/op
        codec.Unmarshal               9827 ns/op        14536 B/op           72 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal              11445 ns/op        16786 B/op          113 allocs/op
     easyjson.Unmarshal               9428 ns/op        15888 B/op           94 allocs/op

        sonic ███████████████▌ 3.89
        goccy ██████████████  3.52
      sonic-v ███████████  2.77
     jsoniter ██████████▊ 2.71
     easyjson █████████▎ 2.31
        codec ████████▊ 2.22
        json2 ████████▊ 2.21
       ffjson ███████▌ 1.91
      segment █████▏ 1.31
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys generically
         json.Unmarshal          457149273 ns/op     21755680 B/op       394051 allocs/op
        json2 >>> not supported <<<
     fastjson.Unmarshal           82693731 ns/op      2533390 B/op        39470 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal           63766530 ns/op      4614567 B/op       177366 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal           58730447 ns/op    616503039 B/op        39630 allocs/op
        gjson.Unmarshal           88080214 ns/op    105924750 B/op        39406 allocs/op
      gjson-v.Unmarshal          170750143 ns/op    105924736 B/op        39406 allocs/op
        sonic.Unmarshal           34856392 ns/op      2845134 B/op        78811 allocs/op
      sonic-v.Unmarshal          113004549 ns/op      2850905 B/op        78811 allocs/op
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ████████████████████████████████████████████████████▍ 13.12
     simdjson ███████████████████████████████▏ 7.78
   jsonparser ████████████████████████████▋ 7.17
     fastjson ██████████████████████  5.53
        gjson ████████████████████▊ 5.19
      sonic-v ████████████████▏ 4.05
      gjson-v ██████████▋ 2.68
         json ▓▓▓▓ 1.00
        goccy >>> not supported <<<
      segment >>> not supported <<<
     jsoniter >>> not supported <<<
        json2 >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read few keys into struct
         json.Unmarshal          464504322 ns/op     21755680 B/op       394051 allocs/op
        json2.Unmarshal          198501845 ns/op        16804 B/op           64 allocs/op
     fastjson.Unmarshal           82113304 ns/op      2533390 B/op        39470 allocs/op
     jsoniter.Unmarshal          206180191 ns/op     89011819 B/op      3979908 allocs/op
   jsonparser.Unmarshal           63898075 ns/op      4614567 B/op       177366 allocs/op
        goccy.Unmarshal           85598977 ns/op    105934654 B/op        39416 allocs/op
      segment.Unmarshal          219354768 ns/op      2526041 B/op        39406 allocs/op
     simdjson.Unmarshal           58170529 ns/op    618706512 B/op        39632 allocs/op
        gjson.Unmarshal           87880124 ns/op    105924736 B/op        39406 allocs/op
      gjson-v.Unmarshal          173216449 ns/op    105924736 B/op        39406 allocs/op
        sonic.Unmarshal          108008396 ns/op    107318369 B/op        78814 allocs/op
      sonic-v.Unmarshal          185358706 ns/op    107479506 B/op        78820 allocs/op
        codec.Unmarshal          115180196 ns/op     48866308 B/op       157621 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal          271035501 ns/op     67791938 B/op       591082 allocs/op
     easyjson.Unmarshal          352490840 ns/op      2527725 B/op        39408 allocs/op

   * simdjson ███████████████████████████████▉ 7.99
 * jsonparser █████████████████████████████  7.27
   * fastjson ██████████████████████▋ 5.66
        goccy █████████████████████▋ 5.43
      * gjson █████████████████████▏ 5.29
        sonic █████████████████▏ 4.30
        codec ████████████████▏ 4.03
    * gjson-v ██████████▋ 2.68
      sonic-v ██████████  2.51
        json2 █████████▎ 2.34
     jsoniter █████████  2.25
      segment ████████▍ 2.12
       ffjson ██████▊ 1.71
     easyjson █████▎ 1.32
         json ▓▓▓▓ 1.00
        djson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys generically
         json.Unmarshal          789427684 ns/op    743939032 B/op     14315107 allocs/op
        json2.Unmarshal          682499681 ns/op    692662224 B/op     11948028 allocs/op
     fastjson.Unmarshal          214632983 ns/op     66540387 B/op      2876746 allocs/op
     jsoniter.Unmarshal          658843931 ns/op    819015708 B/op     19201248 allocs/op
   jsonparser.Unmarshal          502049971 ns/op    105248248 B/op      5437891 allocs/op
        goccy.Unmarshal          598121816 ns/op    876474116 B/op     16797675 allocs/op
      segment.Unmarshal         2443864243 ns/op   6770681504 B/op     19978344 allocs/op
     simdjson.Unmarshal          172550089 ns/op    688363872 B/op      2758581 allocs/op
        gjson.Unmarshal          472736940 ns/op    196083397 B/op      2994781 allocs/op
      gjson-v.Unmarshal          565747862 ns/op    196083384 B/op      2994781 allocs/op
        sonic.Unmarshal          404674651 ns/op    889791477 B/op      2285518 allocs/op
      sonic-v.Unmarshal          493170713 ns/op    889791306 B/op      2285516 allocs/op
        codec.Unmarshal         1118838890 ns/op    312883200 B/op     18323501 allocs/op
        djson.Unmarshal          397839021 ns/op    770793634 B/op      7537398 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     simdjson ██████████████████▎ 4.58
     fastjson ██████████████▋ 3.68
        djson ███████▉ 1.98
        sonic ███████▊ 1.95
        gjson ██████▋ 1.67
      sonic-v ██████▍ 1.60
   jsonparser ██████▎ 1.57
      gjson-v █████▌ 1.40
        goccy █████▎ 1.32
     jsoniter ████▊ 1.20
        json2 ████▋ 1.16
         json ▓▓▓▓ 1.00
        codec ██▊ 0.71
      segment █▎ 0.32
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records (2kb each) from small file (100MB), read all keys into struct
         json.Unmarshal          695996051 ns/op    110920464 B/op      3113019 allocs/op
        json2.Unmarshal          397821774 ns/op    129761162 B/op      1801316 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal          221325792 ns/op     97208075 B/op      2955386 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal          141608558 ns/op    105994028 B/op        39510 allocs/op
      segment.Unmarshal          572132836 ns/op     68879048 B/op      2797773 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal          149290887 ns/op    133177630 B/op       157629 allocs/op
      sonic-v.Unmarshal          227702612 ns/op    133209627 B/op       157635 allocs/op
        codec.Unmarshal          304764603 ns/op    102057728 B/op      2718954 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal          392576936 ns/op    217788229 B/op      4689218 allocs/op
     easyjson.Unmarshal          262267480 ns/op    170469662 B/op      3822300 allocs/op

        goccy ███████████████████▋ 4.91
        sonic ██████████████████▋ 4.66
     jsoniter ████████████▌ 3.14
      sonic-v ████████████▏ 3.06
     easyjson ██████████▌ 2.65
        codec █████████▏ 2.28
       ffjson ███████  1.77
        json2 ██████▉ 1.75
      segment ████▊ 1.22
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys generically
         json.Unmarshal        21267548015 ns/op    116288976 B/op      2105576 allocs/op
        json2 >>> not supported <<<
     fastjson.Unmarshal         3373883258 ns/op     14463336 B/op       218594 allocs/op
     jsoniter >>> not supported <<<
   jsonparser.Unmarshal         2159109729 ns/op     24696400 B/op       947400 allocs/op
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson.Unmarshal         1289848634 ns/op   2241936488 B/op       217325 allocs/op
        gjson.Unmarshal         3293954689 ns/op   5740687488 B/op       210561 allocs/op
      gjson-v.Unmarshal         7318231333 ns/op   5740687488 B/op       210561 allocs/op
        sonic.Unmarshal          755655105 ns/op     15265148 B/op       421122 allocs/op
      sonic-v.Unmarshal         4443380892 ns/op     15308568 B/op       421125 allocs/op
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ████████████████████████████████████████████████████████████████████████████████████████████████████████████████▌ 28.14
     simdjson █████████████████████████████████████████████████████████████████▉ 16.49
   jsonparser ███████████████████████████████████████▍ 9.85
        gjson █████████████████████████▊ 6.46
     fastjson █████████████████████████▏ 6.30
      sonic-v ███████████████████▏ 4.79
      gjson-v ███████████▌ 2.91
         json ▓▓▓▓ 1.00
     jsoniter >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
        json2 >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records (25kb each) from semi-large file (5GB), read few keys into struct
         json.Unmarshal        21276357190 ns/op    116288976 B/op      2105576 allocs/op
        json2.Unmarshal         9298279044 ns/op        21080 B/op           72 allocs/op
     fastjson.Unmarshal         3372183988 ns/op     14463336 B/op       218594 allocs/op
     jsoniter.Unmarshal         9455296540 ns/op   4649577344 B/op    210767649 allocs/op
   jsonparser.Unmarshal         2154893747 ns/op     24696416 B/op       947400 allocs/op
        goccy.Unmarshal         3545887029 ns/op   5741075040 B/op       210976 allocs/op
      segment.Unmarshal         8643461359 ns/op     13537216 B/op       210563 allocs/op
     simdjson.Unmarshal         1098571623 ns/op   2307488688 B/op       217231 allocs/op
        gjson.Unmarshal         3318988918 ns/op   5740687488 B/op       210561 allocs/op
      gjson-v.Unmarshal         7430612961 ns/op   5740687584 B/op       210562 allocs/op
        sonic.Unmarshal         4468286845 ns/op   5751556888 B/op       421248 allocs/op
      sonic-v.Unmarshal         8102550815 ns/op   5754911336 B/op       421407 allocs/op
        codec.Unmarshal         4275839405 ns/op    261152232 B/op       842234 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal        10907335076 ns/op    362265808 B/op      3158405 allocs/op
     easyjson.Unmarshal        17643095216 ns/op     13542760 B/op       210581 allocs/op

   * simdjson █████████████████████████████████████████████████████████████████████████████▍ 19.37
 * jsonparser ███████████████████████████████████████▍ 9.87
      * gjson █████████████████████████▋ 6.41
   * fastjson █████████████████████████▏ 6.31
        goccy ████████████████████████  6.00
        codec ███████████████████▉ 4.98
        sonic ███████████████████  4.76
    * gjson-v ███████████▍ 2.86
      sonic-v ██████████▌ 2.63
      segment █████████▊ 2.46
        json2 █████████▏ 2.29
     jsoniter █████████  2.25
       ffjson ███████▊ 1.95
     easyjson ████▊ 1.21
         json ▓▓▓▓ 1.00
        djson >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys generically
         json.Unmarshal        37506408939 ns/op  37172602160 B/op    701321529 allocs/op
        json2.Unmarshal        30705226340 ns/op  35258784568 B/op    579262980 allocs/op
     fastjson.Unmarshal         9503306708 ns/op   3342948648 B/op    142344570 allocs/op
     jsoniter.Unmarshal        31320264596 ns/op  41727388424 B/op    960518493 allocs/op
   jsonparser.Unmarshal        25139443156 ns/op   5294191192 B/op    273513548 allocs/op
        goccy.Unmarshal        28906211008 ns/op  44760233512 B/op    833136573 allocs/op
      segment.Unmarshal       122365186819 ns/op 341651377480 B/op    995513888 allocs/op
     simdjson.Unmarshal         7002746027 ns/op  26783921352 B/op    137924608 allocs/op
        gjson.Unmarshal        23757379977 ns/op  10406631168 B/op    146758233 allocs/op
      gjson-v.Unmarshal        28379567945 ns/op  10406631360 B/op    146758233 allocs/op
        sonic.Unmarshal        19415687668 ns/op  44950758448 B/op    111175291 allocs/op
      sonic-v.Unmarshal        23297311824 ns/op  44950270240 B/op    111175314 allocs/op
        codec.Unmarshal        55108309611 ns/op  13510676256 B/op    905186195 allocs/op
        djson.Unmarshal        18822211425 ns/op  39791537168 B/op    375170177 allocs/op
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     simdjson █████████████████████▍ 5.36
     fastjson ███████████████▊ 3.95
        djson ███████▉ 1.99
        sonic ███████▋ 1.93
      sonic-v ██████▍ 1.61
        gjson ██████▎ 1.58
   jsonparser █████▉ 1.49
      gjson-v █████▎ 1.32
        goccy █████▏ 1.30
        json2 ████▉ 1.22
     jsoniter ████▊ 1.20
         json ▓▓▓▓ 1.00
        codec ██▋ 0.68
      segment █▏ 0.31
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Unmarshal many records from (25kb each) semi-large file (5GB), read all keys into struct
         json.Unmarshal        31963101500 ns/op   4443556656 B/op    137915221 allocs/op
        json2.Unmarshal        19828042589 ns/op   6894675936 B/op     94346705 allocs/op
     fastjson >>> not supported <<<
     jsoniter.Unmarshal        10317144573 ns/op   5003339888 B/op    150338171 allocs/op
   jsonparser >>> not supported <<<
        goccy.Unmarshal         5986563172 ns/op   5744419304 B/op       215817 allocs/op
      segment.Unmarshal        25640867757 ns/op   3491108072 B/op    141915738 allocs/op
     simdjson >>> not supported <<<
        gjson >>> not supported <<<
      gjson-v >>> not supported <<<
        sonic.Unmarshal         6323792290 ns/op   7132254280 B/op      4632759 allocs/op
      sonic-v.Unmarshal        10170375596 ns/op   7138581688 B/op      4632982 allocs/op
        codec.Unmarshal        13796191443 ns/op   3062218368 B/op    133914575 allocs/op
        djson >>> not supported <<<
       ffjson.Unmarshal        17176614134 ns/op  10023485040 B/op    215400775 allocs/op
     easyjson.Unmarshal        12083082689 ns/op   8946305624 B/op    198555826 allocs/op

        goccy █████████████████████▎ 5.34
        sonic ████████████████████▏ 5.05
      sonic-v ████████████▌ 3.14
     jsoniter ████████████▍ 3.10
     easyjson ██████████▌ 2.65
        codec █████████▎ 2.32
       ffjson ███████▍ 1.86
        json2 ██████▍ 1.61
      segment ████▉ 1.25
         json ▓▓▓▓ 1.00
   jsonparser >>> not supported <<<
      gjson-v >>> not supported <<<
        gjson >>> not supported <<<
        djson >>> not supported <<<
     simdjson >>> not supported <<<
     fastjson >>> not supported <<<

Marshal custom data through an object builder
         json.Marshal                 1034 ns/op          688 B/op           17 allocs/op
        json2 >>> not supported <<<
     fastjson.Marshal                  542 ns/op          679 B/op            6 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Marshal                 1653 ns/op         3024 B/op           27 allocs/op
      gjson-v >>> not supported <<<
        sonic.Marshal                  715 ns/op         1803 B/op           14 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

     fastjson ███████▋ 1.91
        sonic █████▊ 1.45
         json ▓▓▓▓ 1.00
        gjson ██▌ 0.63
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

Validate []byte
         json.Validate                6541 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
     fastjson.Validate                2697 ns/op            0 B/op            0 allocs/op
     jsoniter.Validate                4490 ns/op         2184 B/op          100 allocs/op
   jsonparser >>> not supported <<<
        goccy.Validate               16214 ns/op        27585 B/op          463 allocs/op
      segment.Validate                4417 ns/op            0 B/op            0 allocs/op
     simdjson >>> not supported <<<
        gjson.Validate                2394 ns/op            0 B/op            0 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                1979 ns/op            0 B/op            0 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic █████████████▏ 3.31
        gjson ██████████▉ 2.73
     fastjson █████████▋ 2.43
      segment █████▉ 1.48
     jsoniter █████▊ 1.46
         json ▓▓▓▓ 1.00
        goccy █▌ 0.40
   jsonparser >>> not supported <<<
     simdjson >>> not supported <<<
      gjson-v >>> not supported <<<
        json2 >>> not supported <<<
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

Validate string
         json.Validate                6583 ns/op            0 B/op            0 allocs/op
        json2 >>> not supported <<<
     fastjson.Validate                3191 ns/op         4096 B/op            1 allocs/op
     jsoniter >>> not supported <<<
   jsonparser >>> not supported <<<
        goccy >>> not supported <<<
      segment >>> not supported <<<
     simdjson >>> not supported <<<
        gjson.Validate                2794 ns/op         4096 B/op            1 allocs/op
      gjson-v >>> not supported <<<
        sonic.Validate                2461 ns/op         4097 B/op            1 allocs/op
      sonic-v >>> not supported <<<
        codec >>> not supported <<<
        djson >>> not supported <<<
       ffjson >>> not supported <<<
     easyjson >>> not supported <<<

        sonic ██████████▋ 2.67
        gjson █████████▍ 2.36
     fastjson ████████▎ 2.06
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
