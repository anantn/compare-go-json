// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/valyala/fastjson"
)

var fastjsonPkg = pkg{
	name: "fastjson",
	calls: map[string]*call{
		"validate-bytes":  {name: "Validate", fun: fastjsonValidate},
		"validate-string": {name: "Validate", fun: fastjsonValidateString},
		"single-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			fastjsonFile1Few(b, false)
		}},
		"single-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			fastjsonFile1Few(b, true)
		}, caveat: true},
		"single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			fastjsonFile1All(b)
		}},
		"small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			fastjsonFileManyFew(b, smallTestFile(), false)
		}},
		"small-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			fastjsonFileManyFew(b, smallTestFile(), true)
		}, caveat: true},
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			fastjsonFileManyAll(b, smallTestFile())
		}},
		"large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			fastjsonFileManyFew(b, largeTestFile(), false)
		}},
		"large-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			fastjsonFileManyFew(b, largeTestFile(), true)
		}, caveat: true},
		"large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			fastjsonFileManyAll(b, largeTestFile())
		}},
		"marshal-builder": {name: "Marshal", fun: fastjsonMarshalBuilder},
	},
}

func fastjsonValidate(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if benchErr = fastjson.ValidateBytes(sample); benchErr != nil {
			b.Fail()
		}
	}
}

func fastjsonValidateString(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if benchErr = fastjson.Validate(string(sample)); benchErr != nil {
			b.Fail()
		}
	}
}

func fastjsonMarshalBuilder(b *testing.B) {
	// {"when":1711509483695365000,"what":"Just some fake log entry for a generated log file.","where":[{"file":"example.go","line":123}],"who":"benchmark-application","level":"INFO"}
	var a fastjson.Arena
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		o := a.NewObject()
		o.Set("when", a.NewNumberFloat64(float64(time.Now().UnixNano())))
		o.Set("what", a.NewString("Just some fake log entry for a generated log file."))
		where := a.NewObject()
		where.Set("file", a.NewString("example.go"))
		where.Set("line", a.NewNumberInt(123))
		arr := a.NewArray()
		arr.SetArrayItem(0, where)
		o.Set("where", arr)
		o.Set("who", a.NewString("benchmark-application"))
		o.Set("level", a.NewString("INFO"))
		if err := checkMarshalCustom(string(o.MarshalTo(nil))); err != nil {
			benchErr = err
			b.Fail()
		}
		a.Reset()
	}
}

func fastjsonFile1Few(b *testing.B, useStruct bool) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	var parser fastjson.Parser
	var p = getEmptyPartialPatient()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		if val, err := parser.ParseBytes(j); err != nil {
			benchErr = err
			b.Fail()
		} else {
			if val.Type() != fastjson.TypeObject {
				benchErr = fmt.Errorf("expected object, got %s", val.Type())
				b.Fail()
			}
			strtest := val.GetStringBytes("identifier", "0", "type", "coding", "0", "code")
			arrtest := val.GetArray("name", "2", "given")
			booltest := val.GetBool("deceasedBoolean")

			if useStruct {
				p.Identifier[0].Type.Coding[0].Code = string(strtest)
				p.Name[0].Given = make([]string, len(arrtest))
				for i, v := range arrtest {
					p.Name[0].Given[i] = string(v.GetStringBytes())
				}
				p.DeceasedBoolean = booltest
				err = checkPatientStruct(p)
			} else {
				err = checkPatient(string(strtest), len(arrtest), booltest)
			}
			if err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func fastjsonFile1All(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	var p fastjson.Parser
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		if val, err := p.ParseBytes(j); err != nil {
			benchErr = err
			b.Fail()
		} else {
			if val.Type() != fastjson.TypeObject {
				benchErr = fmt.Errorf("expected object, got %s", val.Type())
				b.Fail()
			}
			root, _ := val.Object()
			root.Visit(fastjsonVisitChildren)
			if fastjsonVisitCount != singleNumChildren {
				benchErr = fmt.Errorf("expected %d children, got %d",
					singleNumChildren, fastjsonVisitCount)
				b.Fail()
			}
			if fastjsonValueHolder == nil {
				benchErr = fmt.Errorf("expected a value, got nil")
				b.Fail()
			}
			fastjsonVisitCount = 0
			fastjsonValueHolder = nil
		}
	}
}

// Store values to avoid compiler optimizations
var fastjsonVisitCount int
var fastjsonValueHolder interface{}

func fastjsonVisitChildren(k []byte, v *fastjson.Value) {
	fastjsonVisitCount++
	switch v.Type() {
	case fastjson.TypeFalse:
		fastjsonValueHolder = false
	case fastjson.TypeTrue:
		fastjsonValueHolder = true
	case fastjson.TypeNumber:
		fastjsonValueHolder = v.GetFloat64()
	case fastjson.TypeString:
		fastjsonValueHolder = v.GetStringBytes()
	case fastjson.TypeObject:
		children, _ := v.Object()
		children.Visit(fastjsonVisitChildren)
	case fastjson.TypeArray:
		children, _ := v.Array()
		for _, child := range children {
			fastjsonVisitChildren(k, child)
		}
	}
}

func fastjsonFileManyFew(b *testing.B, f testfile, useStruct bool) {
	var err error
	var l = getEmptyPartialLog()
	fastjsonCheckFileValues(b, f.handle, f.numRecords, func(val *fastjson.Value) {
		whatval := val.GetStringBytes("what")
		whereval := val.GetInt("where", "0", "line")

		if useStruct {
			l.What = string(whatval)
			l.Where[0].Line = int(whereval)
			err = checkLogStruct(l)
		} else {
			err = checkLog(string(whatval), whereval)
		}
		if err != nil {
			benchErr = err
			b.Fail()
		}
	})
}

func fastjsonFileManyAll(b *testing.B, f testfile) {
	fastjsonCheckFileValues(b, f.handle, f.numRecords, func(val *fastjson.Value) {
		obj, _ := val.Object()
		obj.Visit(fastjsonVisitChildren)
		if fastjsonVisitCount != f.numChildren {
			benchErr = fmt.Errorf("expected %d children, got %d",
				f.numChildren, fastjsonVisitCount)
			b.Fail()
		}
		if fastjsonValueHolder == nil {
			benchErr = fmt.Errorf("expected a value, got nil")
			b.Fail()
		}
		fastjsonVisitCount = 0
		fastjsonValueHolder = nil
	})
}

func fastjsonCheckFileValues(b *testing.B, f *os.File, count int, cb func(*fastjson.Value)) {
	defer func() { _ = f.Close() }()

	var p fastjson.Parser
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		fastjsonManyRecordCount := 0
		for buf.Scan() {
			if val, err := p.ParseBytes(buf.Bytes()); err != nil {
				benchErr = err
				b.Fail()
			} else {
				if val.Type() != fastjson.TypeObject {
					benchErr = fmt.Errorf("expected object, got %s", val.Type())
					b.Fail()
				}
				fastjsonManyRecordCount++
				cb(val)
			}
		}
		if fastjsonManyRecordCount != count {
			benchErr = fmt.Errorf("expected %d records, got %d", count, fastjsonManyRecordCount)
			b.Fail()
		}
		fastjsonManyRecordCount = 0
	}
}
