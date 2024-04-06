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
	calls: map[string]*call{
		"single-few-keys": {name: "Unmarshal", fun: jsonparserFile1Few},
		"single-all-keys": {name: "Unmarshal", fun: jsonparserFile1All},
		"small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsonparserFileManyFew(b, smallTestFile())
		}},
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsonparserFileManyAll(b, smallTestFile())
		}},
		"large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsonparserFileManyFew(b, largeTestFile())
		}},
		"large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsonparserFileManyAll(b, largeTestFile())
		}},
	},
}

func jsonparserFile1Few(b *testing.B) {
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
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		var strval string
		var arrval int
		var boolval bool
		jsonparser.EachKey(j, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
			switch idx {
			case 0:
				strval = string(value)
			case 1:
				jsonparser.ArrayEach(value, func(val []byte, vt jsonparser.ValueType, offset int, err error) {
					arrval++
				})
			case 3:
				boolval = vt == jsonparser.Boolean
			}
		}, paths...)
		err = checkPatient(strval, arrval, boolval)
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
		jsonparserVisitCount = 0
		jsonparserValueHolder = nil
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

func jsonparserFileManyFew(b *testing.B, f testfile) {
	paths := [][]string{
		{"what"},
		{"where", "[0]", "line"},
	}
	var whatval string
	var whereval int64
	jsonparserCheckFileValues(b, f.handle, f.numRecords, func(j []byte) {
		jsonparser.EachKey(j, func(idx int, value []byte, vt jsonparser.ValueType, err error) {
			switch idx {
			case 0:
				whatval = string(value)
			case 1:
				whereval, _ = jsonparser.ParseInt(value)
			}
		}, paths...)
		err := checkLog(whatval, int(whereval))
		if err != nil {
			benchErr = err
			b.Fail()
		}
	})
}

func jsonparserFileManyAll(b *testing.B, f testfile) {
	jsonparserCheckFileValues(b, f.handle, f.numRecords, func(j []byte) {
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
		jsonparserVisitCount = 0
		jsonparserValueHolder = nil
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
