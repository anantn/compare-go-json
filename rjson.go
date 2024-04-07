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

	"github.com/willabides/rjson"
)

var rjsonPkg = pkg{
	name: "rjson",
	calls: map[string]*call{
		"validate-bytes":  {name: "Validate", fun: rjsonValidateBytes},
		"single-all-keys": {name: "Unmarshal", fun: rjsonFile1All},
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			rjsonFileManyAll(b, smallTestFile())
		}},
		"large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			rjsonFileManyAll(b, largeTestFile())
		}},
	},
}

func rjsonValidateBytes(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if !rjson.Valid(sample, rjsonBuffer) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
	}
}

func rjsonFile1All(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)

		rjsonVisitCount = 0
		rjsonValueHolder = nil
		err := rjsonIterate(j)
		if err != nil {
			benchErr = err
			b.Fail()
		}
		// Includes root object
		if rjsonVisitCount != singleNumChildren+1 {
			benchErr = fmt.Errorf("expected %d visits, got %d", singleNumChildren, rjsonVisitCount)
			b.Fail()
		}
		if rjsonValueHolder == nil {
			benchErr = fmt.Errorf("expected a value, got nil")
			b.Fail()
		}
	}
}

func rjsonFileManyAll(b *testing.B, f testfile) {
	defer func() { _ = f.handle.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.handle.Seek(0, 0)
		buf := bufio.NewScanner(f.handle)
		recordCount := 0
		for buf.Scan() {
			rjsonVisitCount = 0
			rjsonValueHolder = nil
			err := rjsonIterate(buf.Bytes())
			if err != nil {
				benchErr = err
				b.Fail()
			}
			// Includes root object
			if rjsonVisitCount != f.numChildren+1 {
				benchErr = fmt.Errorf("expected %d visits, got %d", f.numChildren, rjsonVisitCount)
				b.Fail()
			}
			if rjsonValueHolder == nil {
				benchErr = fmt.Errorf("expected a value, got nil")
				b.Fail()
			}
			recordCount++
		}
		if recordCount != f.numRecords {
			benchErr = fmt.Errorf("expected %d records, got %d", f.numRecords, recordCount)
			b.Fail()
		}
	}
}

var rjsonVisitCount int
var rjsonValueHolder interface{}
var rjsonBuffer = &rjson.Buffer{}

func rjsonIterate(v []byte) (err error) {
	t, p, err := rjson.NextTokenType(v)
	fmt.Println()
	fmt.Println(t)
	if err == io.EOF {
		err = nil
		return
	}
	fmt.Println(err)
	if err != nil {
		return
	}

	switch t {
	case rjson.ObjectStartType:
		rjsonVisitCount++
	case rjson.ArrayStartType:
		rjsonVisitCount++
	case rjson.StringType:
		rjsonVisitCount++
		rjsonValueHolder, p, err = rjson.ReadStringBytes(v, nil)
	case rjson.NumberType:
		rjsonVisitCount++
		rjsonValueHolder, p, err = rjson.ReadInt(v)
	case rjson.TrueType:
		fallthrough
	case rjson.FalseType:
		rjsonVisitCount++
		rjsonValueHolder, p, err = rjson.ReadBool(v)
	default:
		p, err = rjson.SkipValue(v, rjsonBuffer)
	}

	if err != nil {
		return
	}
	fmt.Println(err)
	fmt.Println(p)
	return rjsonIterate(v[p:])
}
