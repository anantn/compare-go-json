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
     json.Unmarshal          21440 ns/op        17744 B/op          334 allocs/op
       oj.Parse              10011 ns/op         5691 B/op          364 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal          15499 ns/op        19681 B/op          451 allocs/op
 simdjson.Parse              28843 ns/op       140634 B/op          361 allocs/op
    gjson.ParseBytes         14772 ns/op        20040 B/op          175 allocs/op
    sonic.Unmarshal           9496 ns/op        23649 B/op          168 allocs/op

    sonic ███████████████▊ 2.26
       oj ██████████████▉ 2.14
    gjson ██████████▏ 1.45
 jsoniter █████████▋ 1.38
     json ▓▓▓▓▓▓▓ 1.00
 simdjson █████▏ 0.74
 fastjson >>> not supported <<<

Validate string/[]byte
     json.Valid               6565 ns/op            0 B/op            0 allocs/op
       oj.Validate            2280 ns/op            0 B/op            0 allocs/op
 fastjson.Validate            2751 ns/op            0 B/op            0 allocs/op
 jsoniter.Valid               4644 ns/op         2187 B/op          100 allocs/op
 simdjson.Validate           19955 ns/op       118632 B/op           11 allocs/op
    gjson.Validate            2371 ns/op            0 B/op            0 allocs/op
    sonic.Valid               2064 ns/op            0 B/op            0 allocs/op

    sonic ██████████████████████▎ 3.18
       oj ████████████████████▏ 2.88
    gjson ███████████████████▍ 2.77
 fastjson ████████████████▋ 2.39
 jsoniter █████████▉ 1.41
     json ▓▓▓▓▓▓▓ 1.00
 simdjson ██▎ 0.33

Iterate tokens in a string/[]byte
     json.Decode             40707 ns/op        22568 B/op         1175 allocs/op
       oj.Tokenize            4700 ns/op         1976 B/op          156 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             16137 ns/op        20361 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Decode             10378 ns/op        24056 B/op          170 allocs/op

       oj ████████████████████████████████████████████████████████████▋ 8.66
    sonic ███████████████████████████▍ 3.92
 jsoniter █████████████████▋ 2.52
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Unmarshal string/[]byte to a struct
     json.Unmarshal          20273 ns/op         2576 B/op           74 allocs/op
       oj.Unmarshal          23560 ns/op         9353 B/op          456 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Unmarshal           6950 ns/op         3141 B/op          170 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Unmarshal           4665 ns/op         4963 B/op           11 allocs/op

    sonic ██████████████████████████████▍ 4.35
 jsoniter ████████████████████▍ 2.92
     json ▓▓▓▓▓▓▓ 1.00
       oj ██████  0.86
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal simple types to string/[]byte
     json.Marshal            13876 ns/op         9873 B/op          216 allocs/op
       oj.JSON                3979 ns/op            0 B/op            0 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             7633 ns/op         7052 B/op           94 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             4277 ns/op         2828 B/op            4 allocs/op

       oj ████████████████████████▍ 3.49
    sonic ██████████████████████▋ 3.24
 jsoniter ████████████▋ 1.82
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Marshal a struct to string/[]byte
     json.Marshal             3965 ns/op         3462 B/op            1 allocs/op
       oj.Marshal             4556 ns/op         5178 B/op           45 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Marshal             3623 ns/op         3471 B/op            2 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Marshal             1542 ns/op         3307 B/op            4 allocs/op

    sonic █████████████████▉ 2.57
 jsoniter ███████▋ 1.09
     json ▓▓▓▓▓▓▓ 1.00
       oj ██████  0.87
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Read from single JSON file
     json.Decode             27063 ns/op        32384 B/op          342 allocs/op
       oj.ParseReader        11639 ns/op         9787 B/op          365 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode             18713 ns/op        20328 B/op          456 allocs/op
 simdjson >>> not supported <<<
    gjson >>> not supported <<<
    sonic.Decode             10953 ns/op        25040 B/op          170 allocs/op

    sonic █████████████████▎ 2.47
       oj ████████████████▎ 2.33
 jsoniter ██████████  1.45
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
 simdjson >>> not supported <<<
    gjson >>> not supported <<<

Read multiple JSON in a small log file (100MB)
     json.Decode         748883892 ns/op   1188215344 B/op     14810437 allocs/op
       oj.ParseReader    422272343 ns/op    847037789 B/op     15402846 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode         588487835 ns/op   1264548528 B/op     19390815 allocs/op
 simdjson.ParseReader    344629794 ns/op   1388715210 B/op     13625788 allocs/op
    gjson >>> not supported <<<
    sonic.Decode         416351367 ns/op    769350437 B/op      8887201 allocs/op

 simdjson ███████████████▏ 2.17
    sonic ████████████▌ 1.80
       oj ████████████▍ 1.77
 jsoniter ████████▉ 1.27
     json ▓▓▓▓▓▓▓ 1.00
 fastjson >>> not supported <<<
    gjson >>> not supported <<<

Read multiple JSON in a semi large log file (5GB)
     json.Decode       40776417124 ns/op  28649553216 B/op    740520835 allocs/op
       oj.ParseReader  20071660004 ns/op  11590687088 B/op    770141076 allocs/op
 fastjson >>> not supported <<<
 jsoniter.Decode       32981371904 ns/op  32465521752 B/op    969506417 allocs/op
 simdjson.ParseReader >>> out of memory <<<
    gjson >>> not supported <<<
    sonic.Decode       21112730041 ns/op  38544678992 B/op    444360111 allocs/op

       oj ██████████████▏ 2.03
    sonic █████████████▌ 1.93
 jsoniter ████████▋ 1.24
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
