# Compare Go JSON

Not all JSON tools cover the same features which make it difficult to
select a set of tools for a project. Here is an attempt to compare
features and benchmarks for a few of the JSON tools for Go.

## Features

| Feature                         | [go/json](https://golang.org/pkg/encoding/json/) | [fastjson](https://github.com/valyala/fastjson) | [jsoniter](https://github.com/json-iterator/go) | [OjG](https://github.com/ohler55/ojg) | [simdjson](https://github.com/minio/simdjson-go) | [gjson](https://github.com/tidwall/gjson)
| ------------------------------- | ------------------ | ------------------ | ------------------ | ------------------ | ------------------ | ------------------ |
| Parse []byte to simple go types | :white_check_mark: | :x:                | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark::boom: |
| Validate                        | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| Parse - io.Reader (large file)  | :white_check_mark: | :x:                | :white_check_mark: | :white_check_mark: | :x:                | :x:                |
| Parse from file                 | :white_check_mark: | :x:                | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x:                |
| Parse to structs                | :white_check_mark: | :x:                | :white_check_mark: | :white_check_mark: | :x:                | :x:                |
| Parse to interface types        | :x:                | :x:                | :x:                | :white_check_mark: | :x:                | :x:                |
| Multiple JSON file/stream       | :white_check_mark: | :x:                | :white_check_mark: | :white_check_mark: | :x:                | :x:                |
| ndjson (newline separated)      | :white_check_mark: | :x:                | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x:                |
| Marshal/Write                   | :white_check_mark: | :x:                | :white_check_mark: | :white_check_mark: | :x:                | :x:                |
| JSON Builder                    | :x:                | :x:                | :x:                | :white_check_mark: | :x:                | :x:                |
| JSONPath                        | :x:                | :x:                | :x:                | :white_check_mark: | :x:                | :x::small_blue_diamond: |
| Data type converters            | :x:                | :x:                | :x:                | :white_check_mark: | :x:                | :x:                |
| Simple Encoding Notation        | :x:                | :x:                | :x:                | :white_check_mark: | :x:                | :x:                |
| Test coverage                   | --                 | 93%                | 21%                | 100%               | 57.4%              | 91.5%              |

 :boom: _gjson does not validate while parsing (try a number of 1.2e3e4) although it does catch that error in validation._

 :small_blue_diamond: _gjson has an alternative search feature_

[_Details of each feature listed are at the bottom of the page_](#Feature-Explanations)

# Benchmarks

```
Parse string/[]byte to simple go types ([]interface{}, int64, string, etc)
     json.Unmarshal          21466 ns/op        17744 B/op          334 allocs/op
       oj.Parse              10307 ns/op         5691 B/op          364 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal          15518 ns/op        19681 B/op          451 allocs/op
 simdjson.Parse              29127 ns/op       140634 B/op          361 allocs/op
    gjson.ParseBytes         14806 ns/op        20040 B/op          175 allocs/op
    sonic.Unmarshal           9568 ns/op        23664 B/op          168 allocs/op

    sonic ███████████████▋ 2.24
       oj ██████████████▌ 2.08
    gjson ██████████▏ 1.45
 jsoniter █████████▋ 1.38
     json ▓▓▓▓▓▓▓ 1.00
 simdjson █████▏ 0.74
 fastjson >>> not supported <<<

Validate string/[]byte
     json.Valid               6379 ns/op            0 B/op            0 allocs/op
       oj.Validate            2262 ns/op            0 B/op            0 allocs/op
 fastjson.Valid               2567 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               4622 ns/op         2187 B/op          100 allocs/op
 simdjson.Validate           19669 ns/op       118632 B/op           11 allocs/op
    gjson.Validate            2386 ns/op            0 B/op            0 allocs/op
    sonic.Valid               1986 ns/op            0 B/op            0 allocs/op

    sonic ██████████████████████▍ 3.21
       oj ███████████████████▋ 2.82
    gjson ██████████████████▋ 2.67
 fastjson █████████████████▍ 2.49
 jsoniter █████████▋ 1.38
     json ▓▓▓▓▓▓▓ 1.00
 simdjson ██▎ 0.32

Iterate tokens in a string/[]byte
     json.Decode             39695 ns/op        22568 B/op         1175 allocs/op
       oj.Tokenize            4647 ns/op         1976 B/op          156 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             16388 ns/op        20360 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Decode             10400 ns/op        24098 B/op          170 allocs/op

       oj ███████████████████████████████████████████████████████████▊ 8.54
    sonic ██████████████████████████▋ 3.82
 jsoniter ████████████████▉ 2.42
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Unmarshal string/[]byte to a struct
     json.Unmarshal          20120 ns/op         2576 B/op           74 allocs/op
       oj.Unmarshal          23657 ns/op         9352 B/op          456 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal           6884 ns/op         3141 B/op          170 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Unmarshal           4646 ns/op         4965 B/op           11 allocs/op

    sonic ██████████████████████████████▎ 4.33
 jsoniter ████████████████████▍ 2.92
     json ▓▓▓▓▓▓▓ 1.00
       oj █████▉ 0.85
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal simple types to string/[]byte
     json.Marshal            13565 ns/op         9873 B/op          216 allocs/op
       oj.JSON                4162 ns/op            0 B/op            0 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             7461 ns/op         7052 B/op           94 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             4255 ns/op         2796 B/op            4 allocs/op

       oj ██████████████████████▊ 3.26
    sonic ██████████████████████▎ 3.19
 jsoniter ████████████▋ 1.82
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal a struct to string/[]byte
     json.Marshal             3825 ns/op         3462 B/op            1 allocs/op
       oj.Marshal             4842 ns/op         5178 B/op           45 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             3586 ns/op         3471 B/op            2 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             1541 ns/op         3310 B/op            4 allocs/op

    sonic █████████████████▍ 2.48
 jsoniter ███████▍ 1.07
     json ▓▓▓▓▓▓▓ 1.00
       oj █████▌ 0.79
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Read from single JSON file
     json.Decode             26635 ns/op        32383 B/op          342 allocs/op
       oj.ParseReader        15348 ns/op        21228 B/op          376 allocs/op
 fastjson.Decode              5350 ns/op        12034 B/op            7 allocs/op
 jsoniter.Decode             19016 ns/op        20328 B/op          456 allocs/op
 simdjson.Decode             21728 ns/op       130664 B/op           17 allocs/op
    gjson.Decode              6631 ns/op        16368 B/op            9 allocs/op
    sonic.Decode              7363 ns/op        12527 B/op           21 allocs/op

 fastjson ██████████████████████████████████▊ 4.98
    gjson ████████████████████████████  4.02
    sonic █████████████████████████▎ 3.62
       oj ████████████▏ 1.74
 jsoniter █████████▊ 1.40
 simdjson ████████▌ 1.23
     json ▓▓▓▓▓▓▓ 1.00

Read multiple JSON in a small log file (100MB)
     json.Decode         760619982 ns/op   1188215216 B/op     14810436 allocs/op
       oj.ParseReader    772730508 ns/op   1615707280 B/op     17772476 allocs/op
 fastjson.Decode         169581834 ns/op    758001769 B/op       592494 allocs/op
 jsoniter.Decode         615869736 ns/op   1264548416 B/op     19390813 allocs/op
 simdjson.ParseReader    339722034 ns/op   1388717630 B/op     13625793 allocs/op
    gjson.Decode          81924325 ns/op    720083398 B/op           51 allocs/op
    sonic.Decode         449701700 ns/op    804422544 B/op      2963353 allocs/op

    gjson ████████████████████████████████████████████████████████████████▉ 9.28
 fastjson ███████████████████████████████▍ 4.49
 simdjson ███████████████▋ 2.24
    sonic ███████████▊ 1.69
 jsoniter ████████▋ 1.24
     json ▓▓▓▓▓▓▓ 1.00
       oj ██████▉ 0.98

Read multiple JSON in a semi large log file (5GB)
     json.Decode       41145660800 ns/op  28649554480 B/op    740520826 allocs/op
       oj.ParseReader  38856723451 ns/op  80785308272 B/op    888625626 allocs/op
 fastjson.Decode       32093603765 ns/op  99289492656 B/op    414697316 allocs/op
 jsoniter.Decode       35022161164 ns/op  32465530648 B/op    969506510 allocs/op
 simdjson.ParseReader  14220122085 ns/op  33742788096 B/op    681285434 allocs/op
    gjson.Decode        8844119698 ns/op   5213263136 B/op     29620791 allocs/op
    sonic.Decode       16295161291 ns/op  37209417880 B/op    444312030 allocs/op

    gjson ████████████████████████████████▌ 4.65
 simdjson ████████████████████▎ 2.89
    sonic █████████████████▋ 2.53
 fastjson ████████▉ 1.28
 jsoniter ████████▏ 1.17
       oj ███████▍ 1.06
     json ▓▓▓▓▓▓▓ 1.00

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

## Feature Explanations

 - **Parse** parse a string to []byte slice in simple go types of
   `[]interface`, `map[string]interface{}`, `string`, `float64`,
   `int64`, `bool`, or `nil`. This support the use case of extracting
   data from a JSON suitable for natigating as well as handing off to
   other packages such as a database for storage.

 - **Validate** a string or []byte slice without extracting values.

 - **Read from io.Reader** indicates a source such as a socket or file
   larger than will fit into memory can be parsed.

 - **Read from file** indicates a parser can read from a file if not
   directly then using ioutils.

 - **Parse to structs** is the ability to reconstitute a struct type
   from JSON.

 - **Parse to interface types** is the ability to reconstitutes types
   even if they are included as interfaces in a containing struct or
   slice.

 - **Multiple JSON** indicates a file or stream with multiple JSON
   documents can be parsed. This is no restricted to the limited case
   of exactly one JSON element per line. Encountered in database dumps
   and load files.

 - **ndjson** is a multiple document JSON where each JSON document
   must be on exactly one line. Found in log files.

 - **Marshal/Write** is the ability of the package to marshal go types
   in JSON.

 - **JSON Builder** is the ability to create new data structures suitable for JSON encoding.

 - **[JSONPath](https://goessner.net/articles/JsonPath)** is the
   ability to navigate data using JSONPath expressions.

 - **Data type converters** tools for converting from type to simple
   data types. Basically marshalling and unmarshalling to simple types
   instead of to JSON.

 - **[Simple Encoding Notation](https://github.com/ohler55/ojg/blob/develop/sen.md)** is
   a lazy JSON format where quotes and commas are optional in most
   cases. A merge of JSON and GraphQL formats for those of us that
   don't want to be bothered with strict syntax checking.

 - **Parser Test coverage** percent unit test coverage of the parser
   package. It does not include coverage of other package in the
   offering.
