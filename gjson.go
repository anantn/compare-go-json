// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var gjsonShouldValidate bool

var gjsonPkg = pkg{
	name: "gjson",
	calls: map[string]*call{
		"validate-bytes":  {name: "Validate", fun: gjsonValidate},
		"validate-string": {name: "Validate", fun: gjsonValidateString},
		"unmarshal-single-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = false
			gjsonFile1(b)
		}},
		"unmarshal-single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = false
			gjsonFile1All(b)
		}},
		"unmarshal-small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = false
			gjsonFileMany(b, smallTestFile())
		}},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = false
			gjsonFileManyAll(b, smallTestFile())
		}},
		"unmarshal-large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = false
			gjsonFileMany(b, largeTestFile())
		}},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = false
			gjsonFileManyAll(b, largeTestFile())
		}},
		"marshal-builder": {name: "Marshal", fun: gjsonMarshalBuilder},
	},
}

var gjsonValidatePkg = pkg{
	name: "gjson-v",
	calls: map[string]*call{
		"unmarshal-single-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = true
			gjsonFile1(b)
		}},
		"unmarshal-single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = true
			gjsonFile1All(b)
		}},
		"unmarshal-small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = true
			gjsonFileMany(b, smallTestFile())
		}},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = true
			gjsonFileManyAll(b, smallTestFile())
		}},
		"unmarshal-large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = true
			gjsonFileMany(b, largeTestFile())
		}},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = true
			gjsonFileManyAll(b, largeTestFile())
		}},
	},
}

func gjsonValidate(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if !gjson.ValidBytes(sample) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
	}
}

func gjsonValidateString(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if !gjson.Valid(string(sample)) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
	}
}

func gjsonMarshalBuilder(b *testing.B) {
	// {"when":1711509483695365000,"what":"Just some fake log entry for a generated log file.","where":[{"file":"example.go","line":123}],"who":"benchmark-application","level":"INFO"}
	b.ResetTimer()
	var output = make([]byte, 0, 256)
	for n := 0; n < b.N; n++ {
		v, _ := sjson.SetBytes(output, "when", time.Now().UnixNano())
		v, _ = sjson.SetBytes(v, "what", "Just some fake log entry for a generated log file.")
		v, _ = sjson.SetBytes(v, "where.0.file", "example.go")
		v, _ = sjson.SetBytes(v, "where.0.line", 123)
		v, _ = sjson.SetBytes(v, "who", "benchmark-application")
		v, _ = sjson.SetBytes(v, "level", "INFO")
		if err := checkMarshalCustom(string(v)); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func gjsonFile1(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		if gjsonShouldValidate && !gjson.ValidBytes(j) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
		result := gjson.ParseBytes(j)
		if !result.IsObject() {
			benchErr = errors.New("expected object")
			b.Fail()
		} else {
			strtest := result.Get("identifier.0.type.coding.0.code").String()
			arrtest := result.Get("name.2.given").Array()
			booltest := result.Get("deceasedBoolean").Bool()
			err = checkPatient(strtest, len(arrtest), booltest)
			if err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func gjsonFile1All(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		if gjsonShouldValidate && !gjson.ValidBytes(j) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
		result := gjson.ParseBytes(j)
		if !result.IsObject() {
			benchErr = errors.New("expected object")
			b.Fail()
		} else {
			result.ForEach(gjsonVisitChildren)
			if gjsonVisitCount != singleNumChildren {
				benchErr = fmt.Errorf("expected %d children, got %d",
					singleNumChildren, gjsonVisitCount)
				b.Fail()
			}
			if gjsonValueHolder == nil {
				benchErr = fmt.Errorf("expected a value, got nil")
				b.Fail()
			}
			gjsonVisitCount = 0
			gjsonValueHolder = nil
		}
	}
}

// Store values to avoid compiler optimizations
var gjsonVisitCount int
var gjsonValueHolder interface{}

func gjsonVisitChildren(k, v gjson.Result) bool {
	switch v.Type {
	case gjson.String:
		gjsonValueHolder = v.Str
	case gjson.Number:
		gjsonValueHolder = v.Num
	case gjson.True:
		gjsonValueHolder = true
	case gjson.False:
		gjsonValueHolder = false
	}
	gjsonVisitCount++
	if v.IsObject() || v.IsArray() {
		v.ForEach(gjsonVisitChildren)
	}
	return true
}

func gjsonFileMany(b *testing.B, f testfile) {
	gjsonCheckFileValues(b, f.handle, f.numRecords, func(result gjson.Result) {
		whatval := result.Get("what").String()
		whereval := result.Get("where.0.line").Int()
		err := checkLog(whatval, int(whereval))
		if err != nil {
			benchErr = err
			b.Fail()
		}
	})
}

func gjsonFileManyAll(b *testing.B, f testfile) {
	gjsonCheckFileValues(b, f.handle, f.numRecords, func(result gjson.Result) {
		result.ForEach(gjsonVisitChildren)
		if gjsonVisitCount != f.numChildren {
			benchErr = fmt.Errorf("expected %d children, got %d",
				f.numChildren, gjsonVisitCount)
			b.Fail()
		}
		if gjsonValueHolder == nil {
			benchErr = fmt.Errorf("expected a value, got nil")
			b.Fail()
		}
		gjsonVisitCount = 0
		gjsonValueHolder = nil
	})
}

func gjsonCheckFileValues(b *testing.B, f *os.File, count int, cb func(gjson.Result)) {
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		gjsonRecordCount := 0
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			j := scanner.Bytes()
			if gjsonShouldValidate && !gjson.ValidBytes(j) {
				benchErr = errors.New("JSON not valid")
				b.Fail()
			}
			if result := gjson.ParseBytes(j); !result.IsObject() {
				benchErr = errors.New("expected object")
				b.Fail()
			} else {
				if !result.IsObject() {
					benchErr = errors.New("expected object")
					b.Fail()
				} else {
					cb(result)
				}
			}
			gjsonRecordCount++
		}
		if gjsonRecordCount != count {
			benchErr = fmt.Errorf("expected %d records, got %d", count, gjsonRecordCount)
			b.Fail()
		}
		gjsonRecordCount = 0
	}
}
