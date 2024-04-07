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
		"single-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = false
			gjsonFile1Few(b, false)
		}},
		"single-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = false
			gjsonFile1Few(b, true)
		}, caveat: true},
		"single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = false
			gjsonFile1All(b)
		}},
		"small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = false
			gjsonFileManyFew(b, smallTestFile(), false)
		}},
		"small-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = false
			gjsonFileManyFew(b, smallTestFile(), true)
		}, caveat: true},
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = false
			gjsonFileManyAll(b, smallTestFile())
		}},
		"large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = false
			gjsonFileManyFew(b, largeTestFile(), false)
		}},
		"large-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = false
			gjsonFileManyFew(b, largeTestFile(), true)
		}, caveat: true},
		"large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = false
			gjsonFileManyAll(b, largeTestFile())
		}},
		"marshal-builder": {name: "Marshal", fun: gjsonMarshalBuilder},
	},
}

var gjsonValidatePkg = pkg{
	name: "gjson-v",
	calls: map[string]*call{
		"single-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = true
			gjsonFile1Few(b, false)
		}},
		"single-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = true
			gjsonFile1Few(b, true)
		}, caveat: true},
		"single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = true
			gjsonFile1All(b)
		}},
		"small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = true
			gjsonFileManyFew(b, smallTestFile(), false)
		}},
		"small-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = true
			gjsonFileManyFew(b, smallTestFile(), true)
		}, caveat: true},
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = true
			gjsonFileManyAll(b, smallTestFile())
		}},
		"large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = true
			gjsonFileManyFew(b, largeTestFile(), false)
		}},
		"large-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			gjsonShouldValidate = true
			gjsonFileManyFew(b, largeTestFile(), true)
		}, caveat: true},
		"large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
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

func gjsonFile1Few(b *testing.B, useStruct bool) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	var p = getEmptyPartialPatient()
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

			if useStruct {
				p.Identifier[0].Type.Coding[0].Code = strtest
				p.Name[0].Given = make([]string, len(arrtest))
				for i, v := range arrtest {
					p.Name[0].Given[i] = v.String()
				}
				p.DeceasedBoolean = booltest
				err = checkPatientStruct(p)
			} else {
				err = checkPatient(strtest, len(arrtest), booltest)
			}
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

func gjsonFileManyFew(b *testing.B, f testfile, useStruct bool) {
	var err error
	var l = getEmptyPartialLog()
	gjsonCheckFileValues(b, f.handle, f.numRecords, func(result gjson.Result) {
		whatval := result.Get("what").String()
		whereval := result.Get("where.0.line").Int()

		if useStruct {
			l.What = whatval
			l.Where[0].Line = int(whereval)
			err = checkLogStruct(l)
		} else {
			err = checkLog(whatval, int(whereval))
		}
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
