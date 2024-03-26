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
       oj.Parse              10121 ns/op         5691 B/op          364 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal          15507 ns/op        19681 B/op          451 allocs/op
 simdjson.Parse              28308 ns/op       140634 B/op          361 allocs/op
    gjson.ParseBytes         14874 ns/op        20039 B/op          175 allocs/op

       oj ██████████████▊ 2.11
    gjson ██████████  1.43
 jsoniter █████████▋ 1.38
     json ▓▓▓▓▓▓▓ 1.00
 simdjson █████▎ 0.75
 fastjson >>> not supported <<<

Validate string/[]byte
     json.Valid               6539 ns/op            0 B/op            0 allocs/op
       oj.Validate            2217 ns/op            0 B/op            0 allocs/op
 fastjson.Validate            2671 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               4644 ns/op         2187 B/op          100 allocs/op
 simdjson.Validate           17492 ns/op       118632 B/op           11 allocs/op
    gjson.Validate            2462 ns/op            0 B/op            0 allocs/op

       oj ████████████████████▋ 2.95
    gjson ██████████████████▌ 2.66
 fastjson █████████████████▏ 2.45
 jsoniter █████████▊ 1.41
     json ▓▓▓▓▓▓▓ 1.00
 simdjson ██▌ 0.37

Iterate tokens in a string/[]byte
     json.Decode             40580 ns/op        22568 B/op         1175 allocs/op
       oj.Tokenize            4674 ns/op         1976 B/op          156 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             16230 ns/op        20360 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

       oj ████████████████████████████████████████████████████████████▊ 8.68
 jsoniter █████████████████▌ 2.50
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Unmarshal string/[]byte to a struct
     json.Unmarshal          20405 ns/op         2576 B/op           74 allocs/op
       oj.Unmarshal          23222 ns/op         9352 B/op          456 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal           6875 ns/op         3141 B/op          170 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

 jsoniter ████████████████████▊ 2.97
     json ▓▓▓▓▓▓▓ 1.00
       oj ██████▏ 0.88
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal simple types to string/[]byte
     json.Marshal            14043 ns/op         9872 B/op          216 allocs/op
       oj.JSON                4025 ns/op            0 B/op            0 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             7595 ns/op         7051 B/op           94 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

       oj ████████████████████████▍ 3.49
 jsoniter ████████████▉ 1.85
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal a struct to string/[]byte
     json.Marshal             4005 ns/op         3461 B/op            1 allocs/op
       oj.Marshal             4606 ns/op         5177 B/op           45 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             3658 ns/op         3470 B/op            2 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

 jsoniter ███████▋ 1.09
     json ▓▓▓▓▓▓▓ 1.00
       oj ██████  0.87
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Read from single JSON file
     json.Decode             27564 ns/op        32385 B/op          342 allocs/op
       oj.ParseReader        11689 ns/op         9787 B/op          365 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             18899 ns/op        20328 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

       oj ████████████████▌ 2.36
 jsoniter ██████████▏ 1.46
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Read multiple JSON in a small log file (100MB)
     json.Decode         760783831 ns/op   1188215304 B/op     14810437 allocs/op
       oj.ParseReader    413655548 ns/op    847037661 B/op     15402845 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode         585038751 ns/op   1264548400 B/op     19390813 allocs/op
 simdjson.ParseReader    341302998 ns/op   1388713836 B/op     13625783 allocs/op
    gjson >>> not supported <<<

 simdjson ███████████████▌ 2.23
       oj ████████████▊ 1.84
 jsoniter █████████  1.30
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
    gjson >>> not supported <<<

Read multiple JSON in a semi large log file (5GB)
     json.Decode       40372080538 ns/op  28649543792 B/op    740520816 allocs/op
       oj.ParseReader  19963998777 ns/op  11590677568 B/op    770141065 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode       32718174898 ns/op  32465516824 B/op    969506396 allocs/op
 simdjson.ParseReader >>> out of memory <<<
    gjson >>> not supported <<<

       oj ██████████████▏ 2.02
 jsoniter ████████▋ 1.23
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> out of memory <<<
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
