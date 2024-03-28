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
		"validate-bytes":            {name: "Validate", fun: fastjsonValidate},
		"validate-string":           {name: "Validate", fun: fastjsonValidateString},
		"unmarshal-single-few-keys": {name: "Unmarshal", fun: fastjsonFile1},
		"unmarshal-single-all-keys": {name: "Unmarshal", fun: fastjsonFile1All},
		"unmarshal-small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			fastjsonFileMany(b, openSmallLogFile(), smallLogFileLen)
		}},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			fastjsonFileManyAll(b, openSmallLogFile(), smallLogFileLen)
		}},
		"unmarshal-large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			fastjsonFileMany(b, openLargeLogFile(), largeLogFileLen)
		}},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			fastjsonFileManyAll(b, openLargeLogFile(), largeLogFileLen)
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

func fastjsonFile1(b *testing.B) {
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
			strtest := val.GetStringBytes("identifier", "0", "type", "coding", "0", "code")
			arrtest := val.GetArray("name", "2", "given")
			booltest := val.GetBool("deceasedBoolean")
			err = checkPatient(string(strtest), len(arrtest), booltest)
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

func fastjsonFileMany(b *testing.B, f *os.File, count int) {
	fastjsonCheckFileValues(b, f, count, func(val *fastjson.Value) {
		whatval := val.GetStringBytes("what")
		whereval := val.GetInt("where", "0", "line")
		err := checkLog(string(whatval), whereval)
		if err != nil {
			benchErr = err
			b.Fail()
		}
	})
}

func fastjsonFileManyAll(b *testing.B, f *os.File, count int) {
	fastjsonCheckFileValues(b, f, count, func(val *fastjson.Value) {
		obj, _ := val.Object()
		obj.Visit(fastjsonVisitChildren)
		if fastjsonVisitCount != logNumChildren {
			benchErr = fmt.Errorf("expected %d children, got %d",
				logNumChildren, fastjsonVisitCount)
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
