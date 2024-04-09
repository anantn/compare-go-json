// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/buger/jsonparser"
)

var jsonparserPkg = pkg{
	name: "jsonparser",
	unmarshal: func(data []byte, v interface{}) error {
		jsonparserVisitCount = 0
		jsonparserValueHolder = nil
		jsonparser.ObjectEach(data, jsonparserVisitChildren)
		return nil
	},
	calls: map[string]*call{
		"single-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsonparserFile1Few(b, false)
		}},
		"single-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			jsonparserFile1Few(b, true)
		}, caveat: true},
		"single-all-keys": {name: "Unmarshal", fun: jsonparserFile1All},
		"small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsonparserFileManyFew(b, smallTestFile(), false)
		}},
		"small-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			jsonparserFileManyFew(b, smallTestFile(), true)
		}, caveat: true},
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsonparserFileManyAll(b, smallTestFile())
		}},
		"large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsonparserFileManyFew(b, largeTestFile(), false)
		}},
		"large-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			jsonparserFileManyFew(b, largeTestFile(), true)
		}, caveat: true},
		"large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsonparserFileManyAll(b, largeTestFile())
		}},
	},
}

func jsonparserFile1Few(b *testing.B, useStruct bool) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	paths := [][]string{
		{"identifier", "[0]", "type", "coding", "[0]", "code"},
		{"name", "[2]", "given"},
		{"deceasedBoolean"},
	}
	b.ResetTimer()
	var p = getEmptyPartialPatient()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		var strval string
		var arrval []string
		var boolval bool
		jsonparser.EachKey(j, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
			switch idx {
			case 0:
				strval = string(value)
			case 1:
				jsonparser.ArrayEach(value, func(val []byte, vt jsonparser.ValueType, offset int, err error) {
					arrval = append(arrval, string(val))
				})
			case 3:
				boolval = vt == jsonparser.Boolean
			}
		}, paths...)

		if useStruct {
			p.Identifier[0].Type.Coding[0].Code = strval
			p.Name[0].Given = arrval
			p.DeceasedBoolean = boolval
			err = checkPatientStruct(p)
		} else {
			err = checkPatient(strval, len(arrval), boolval)
		}
		if err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func jsonparserFile1All(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		jsonparserVisitCount = 0
		jsonparserValueHolder = nil
		jsonparser.ObjectEach(j, jsonparserVisitChildren)
		if jsonparserVisitCount != singleNumChildren {
			benchErr = fmt.Errorf("expected %d children, got %d",
				singleNumChildren, jsonparserVisitCount)
			b.Fail()
		}
		if jsonparserValueHolder == nil {
			benchErr = fmt.Errorf("expected a value, got nil")
			b.Fail()
		}
	}
}

// Store values to avoid compiler optimizations
var jsonparserVisitCount int
var jsonparserValueHolder interface{}

func jsonparserVisitChildren(key []byte, value []byte, vt jsonparser.ValueType, offset int) error {
	jsonparserVisitCount++
	switch vt {
	case jsonparser.String:
		jsonparserValueHolder = string(value)
	case jsonparser.Number:
		jsonparserValueHolder, _ = jsonparser.ParseInt(value)
	case jsonparser.Boolean:
		jsonparserValueHolder, _ = jsonparser.ParseBoolean(value)
	case jsonparser.Object:
		jsonparser.ObjectEach(value, jsonparserVisitChildren)
	case jsonparser.Array:
		jsonparser.ArrayEach(value, func(val []byte, vt jsonparser.ValueType, offset int, err error) {
			jsonparserVisitChildren(nil, val, vt, offset)
		})
	}
	return nil
}

func jsonparserFileManyFew(b *testing.B, f testfile, useStruct bool) {
	paths := [][]string{
		{"what"},
		{"where", "[0]", "line"},
	}
	var err error
	var whatval string
	var whereval int64
	var l = getEmptyPartialLog()
	jsonparserCheckFileValues(b, f.handle, f.numRecords, func(j []byte) {
		jsonparser.EachKey(j, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
			switch idx {
			case 0:
				whatval = string(value)
			case 1:
				whereval, _ = jsonparser.ParseInt(value)
			}
		}, paths...)

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

func jsonparserFileManyAll(b *testing.B, f testfile) {
	jsonparserCheckFileValues(b, f.handle, f.numRecords, func(j []byte) {
		jsonparserVisitCount = 0
		jsonparserValueHolder = nil
		jsonparser.ObjectEach(j, jsonparserVisitChildren)
		if jsonparserVisitCount != f.numChildren {
			benchErr = fmt.Errorf("expected %d children, got %d",
				f.numChildren, jsonparserVisitCount)
			b.Fail()
		}
		if jsonparserValueHolder == nil {
			benchErr = fmt.Errorf("expected a value, got nil")
			b.Fail()
		}
	})
}

func jsonparserCheckFileValues(b *testing.B, f *os.File, count int, cb func([]byte)) {
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		jsonparserManyRecordCount := 0
		for buf.Scan() {
			cb(buf.Bytes())
			jsonparserManyRecordCount++
		}
		if jsonparserManyRecordCount != count {
			benchErr = fmt.Errorf("expected %d records, got %d", count, jsonparserManyRecordCount)
			b.Fail()
		}
		jsonparserManyRecordCount = 0
	}
}
