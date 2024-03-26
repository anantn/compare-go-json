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
     json.Unmarshal          21334 ns/op        17744 B/op          334 allocs/op
       oj.Parse              10203 ns/op         5691 B/op          364 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal          15303 ns/op        19681 B/op          451 allocs/op
 simdjson.Parse              28532 ns/op       140634 B/op          361 allocs/op
    gjson.ParseBytes         14697 ns/op        20040 B/op          175 allocs/op
    sonic.Unmarshal           9498 ns/op        23627 B/op          168 allocs/op

    sonic ███████████████▋ 2.25
       oj ██████████████▋ 2.09
    gjson ██████████▏ 1.45
 jsoniter █████████▊ 1.39
     json ▓▓▓▓▓▓▓ 1.00
 simdjson █████▏ 0.75
 fastjson >>> not supported <<<

Validate string/[]byte
     json.Valid               6610 ns/op            0 B/op            0 allocs/op
       oj.Validate            2202 ns/op            0 B/op            0 allocs/op
 fastjson.Valid               2537 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               4620 ns/op         2187 B/op          100 allocs/op
 simdjson.Validate           17432 ns/op       118633 B/op           11 allocs/op
    gjson.Validate            2367 ns/op            0 B/op            0 allocs/op
    sonic.Valid               2085 ns/op            0 B/op            0 allocs/op

    sonic ██████████████████████▏ 3.17
       oj █████████████████████  3.00
    gjson ███████████████████▌ 2.79
 fastjson ██████████████████▏ 2.61
 jsoniter ██████████  1.43
     json ▓▓▓▓▓▓▓ 1.00
 simdjson ██▋ 0.38

Iterate tokens in a string/[]byte
     json.Decode             40162 ns/op        22568 B/op         1175 allocs/op
       oj.Tokenize            4697 ns/op         1976 B/op          156 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             16142 ns/op        20360 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Decode             10245 ns/op        24077 B/op          170 allocs/op

       oj ███████████████████████████████████████████████████████████▊ 8.55
    sonic ███████████████████████████▍ 3.92
 jsoniter █████████████████▍ 2.49
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Unmarshal string/[]byte to a struct
     json.Unmarshal          20030 ns/op         2576 B/op           74 allocs/op
       oj.Unmarshal          23397 ns/op         9347 B/op          456 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal           6910 ns/op         3141 B/op          170 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Unmarshal           4724 ns/op         4973 B/op           11 allocs/op

    sonic █████████████████████████████▋ 4.24
 jsoniter ████████████████████▎ 2.90
     json ▓▓▓▓▓▓▓ 1.00
       oj █████▉ 0.86
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal simple types to string/[]byte
     json.Marshal            13717 ns/op         9873 B/op          216 allocs/op
       oj.JSON                4216 ns/op            0 B/op            0 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             7723 ns/op         7052 B/op           94 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             4311 ns/op         2807 B/op            4 allocs/op

       oj ██████████████████████▊ 3.25
    sonic ██████████████████████▎ 3.18
 jsoniter ████████████▍ 1.78
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal a struct to string/[]byte
     json.Marshal             3971 ns/op         3462 B/op            1 allocs/op
       oj.Marshal             4791 ns/op         5178 B/op           45 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             3603 ns/op         3471 B/op            2 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             1533 ns/op         3270 B/op            4 allocs/op

    sonic ██████████████████▏ 2.59
 jsoniter ███████▋ 1.10
     json ▓▓▓▓▓▓▓ 1.00
       oj █████▊ 0.83
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Read from single JSON file
     json.Decode             27107 ns/op        32385 B/op          342 allocs/op
       oj.ParseReader        11824 ns/op         9787 B/op          365 allocs/op
 fastjson.Decode              5251 ns/op        12032 B/op            6 allocs/op
 jsoniter.Decode             18923 ns/op        20328 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Decode             10708 ns/op        25038 B/op          170 allocs/op

 fastjson ████████████████████████████████████▏ 5.16
    sonic █████████████████▋ 2.53
       oj ████████████████  2.29
 jsoniter ██████████  1.43
     json ▓▓▓▓▓▓▓ 1.00
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Read multiple JSON in a small log file (100MB)
     json.Decode         765954999 ns/op   1188215152 B/op     14810435 allocs/op
       oj.ParseReader    407592089 ns/op    847037629 B/op     15402845 allocs/op
 fastjson.Decode         118953298 ns/op    626876960 B/op           52 allocs/op
 jsoniter.Decode         597037107 ns/op   1264548104 B/op     19390810 allocs/op
 simdjson.ParseReader    339655157 ns/op   1388718410 B/op     13625793 allocs/op
    gjson >>> not supported <<<
    sonic.Decode         422145893 ns/op    773577032 B/op      8887215 allocs/op

 fastjson █████████████████████████████████████████████  6.44
 simdjson ███████████████▊ 2.26
       oj █████████████▏ 1.88
    sonic ████████████▋ 1.81
 jsoniter ████████▉ 1.28
     json ▓▓▓▓▓▓▓ 1.00
    gjson >>> not supported <<<

Read multiple JSON in a semi large log file (5GB)
     json.Decode       41226370950 ns/op  28649573792 B/op    740521095 allocs/op
       oj.ParseReader  19729721853 ns/op  11590687232 B/op    770141124 allocs/op
 fastjson.Decode        4369099671 ns/op         8048 B/op           30 allocs/op
 jsoniter.Decode       33725463937 ns/op  32465556088 B/op    969506712 allocs/op
 simdjson.ParseReader  14247777399 ns/op  33742812832 B/op    681285493 allocs/op
    gjson >>> not supported <<<
    sonic.Decode       15874736309 ns/op  37207074248 B/op    444312002 allocs/op

 fastjson ██████████████████████████████████████████████████████████████████  9.44
 simdjson ████████████████████▎ 2.89
    sonic ██████████████████▏ 2.60
       oj ██████████████▋ 2.09
 jsoniter ████████▌ 1.22
     json ▓▓▓▓▓▓▓ 1.00
    gjson >>> not supported <<<

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
