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
