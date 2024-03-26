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
     json.Unmarshal          21385 ns/op        17745 B/op          334 allocs/op
       oj.Parse              10201 ns/op         5691 B/op          364 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal          15520 ns/op        19681 B/op          451 allocs/op
 simdjson.Parse              29773 ns/op       140635 B/op          361 allocs/op
    gjson.ParseBytes         14910 ns/op        20040 B/op          175 allocs/op
    sonic.Unmarshal           9444 ns/op        23637 B/op          168 allocs/op

    sonic ███████████████▊ 2.26
       oj ██████████████▋ 2.10
    gjson ██████████  1.43
 jsoniter █████████▋ 1.38
     json ▓▓▓▓▓▓▓ 1.00
 simdjson █████  0.72
 fastjson >>> not supported <<<

Validate string/[]byte
     json.Valid               6412 ns/op            0 B/op            0 allocs/op
       oj.Validate            2184 ns/op            0 B/op            0 allocs/op
 fastjson.Validate            2631 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               4627 ns/op         2187 B/op          100 allocs/op
 simdjson.Validate           16502 ns/op       118633 B/op           11 allocs/op
    gjson.Validate            2387 ns/op            0 B/op            0 allocs/op
    sonic.Valid               2007 ns/op            0 B/op            0 allocs/op

    sonic ██████████████████████▎ 3.19
       oj ████████████████████▌ 2.94
    gjson ██████████████████▊ 2.69
 fastjson █████████████████  2.44
 jsoniter █████████▋ 1.39
     json ▓▓▓▓▓▓▓ 1.00
 simdjson ██▋ 0.39

Iterate tokens in a string/[]byte
     json.Decode             40300 ns/op        22568 B/op         1175 allocs/op
       oj.Tokenize            4633 ns/op         1976 B/op          156 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             16174 ns/op        20360 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Decode             10268 ns/op        24046 B/op          170 allocs/op

       oj ████████████████████████████████████████████████████████████▉ 8.70
    sonic ███████████████████████████▍ 3.92
 jsoniter █████████████████▍ 2.49
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Unmarshal string/[]byte to a struct
     json.Unmarshal          19897 ns/op         2576 B/op           74 allocs/op
       oj.Unmarshal          23042 ns/op         9351 B/op          456 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal           6806 ns/op         3141 B/op          170 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Unmarshal           4739 ns/op         4972 B/op           11 allocs/op

    sonic █████████████████████████████▍ 4.20
 jsoniter ████████████████████▍ 2.92
     json ▓▓▓▓▓▓▓ 1.00
       oj ██████  0.86
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal simple types to string/[]byte
     json.Marshal            14001 ns/op         9873 B/op          216 allocs/op
       oj.JSON                3958 ns/op            0 B/op            0 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             7624 ns/op         7052 B/op           94 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             4261 ns/op         2831 B/op            4 allocs/op

       oj ████████████████████████▊ 3.54
    sonic ███████████████████████  3.29
 jsoniter ████████████▊ 1.84
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal a struct to string/[]byte
     json.Marshal             3979 ns/op         3462 B/op            1 allocs/op
       oj.Marshal             4558 ns/op         5178 B/op           45 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             3686 ns/op         3471 B/op            2 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             1517 ns/op         3277 B/op            4 allocs/op

    sonic ██████████████████▎ 2.62
 jsoniter ███████▌ 1.08
     json ▓▓▓▓▓▓▓ 1.00
       oj ██████  0.87
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Read from single JSON file
     json.Decode             27421 ns/op        32385 B/op          342 allocs/op
       oj.ParseReader        11840 ns/op         9787 B/op          365 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             18643 ns/op        20328 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Decode             10777 ns/op        25002 B/op          170 allocs/op

    sonic █████████████████▊ 2.54
       oj ████████████████▏ 2.32
 jsoniter ██████████▎ 1.47
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Read multiple JSON in a small log file (100MB)
     json.Decode         771646616 ns/op   1188215304 B/op     14810437 allocs/op
       oj.ParseReader    412583700 ns/op    847037656 B/op     15402845 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode         569737993 ns/op   1264548280 B/op     19390812 allocs/op
 simdjson.ParseReader    331845565 ns/op   1388716292 B/op     13625790 allocs/op
    gjson >>> not supported <<<
    sonic.Decode         422505635 ns/op    773041789 B/op      8887212 allocs/op

 simdjson ████████████████▎ 2.33
       oj █████████████  1.87
    sonic ████████████▊ 1.83
 jsoniter █████████▍ 1.35
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
    gjson >>> not supported <<<

Read multiple JSON in a semi large log file (5GB)
     json.Decode       40871276155 ns/op  28649556880 B/op    740520823 allocs/op
       oj.ParseReader  20064434804 ns/op  11590685184 B/op    770141071 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode       33060294357 ns/op  32465562600 B/op    969506714 allocs/op
 simdjson.ParseReader  14157286042 ns/op  33700804872 B/op    681285401 allocs/op
    gjson >>> not supported <<<
    sonic.Decode       16126523198 ns/op  37209310600 B/op    444312021 allocs/op

 simdjson ████████████████████▏ 2.89
    sonic █████████████████▋ 2.53
       oj ██████████████▎ 2.04
 jsoniter ████████▋ 1.24
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
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
