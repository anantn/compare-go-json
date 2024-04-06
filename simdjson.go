// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/minio/simdjson-go"
)

var simdjsonPkg = pkg{
	name: "simdjson",
	calls: map[string]*call{
		"single-all-keys": {name: "Unmarshal", fun: simdjsonFile1All},
		"small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			simdjsonFileManyFew(b, smallTestFile())
		}},
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			simdjsonFileManyAll(b, smallTestFile())
		}},
		"large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			simdjsonFileManyFew(b, largeTestFile())
		}},
		"large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			simdjsonFileManyAll(b, largeTestFile())
		}},
	},
}

func simdjsonFile1All(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		pj, err := simdjson.Parse(j, nil)
		if err != nil {
			benchErr = err
			b.Fail()
		}
		simdjsonVisitCount = 0
		simdjsonValueHolder = nil
		err = pj.ForEach(simdjsonVisitChildren)
		if err != nil {
			benchErr = err
			b.Fail()
		}
		if simdjsonVisitCount != singleNumChildren+1 {
			benchErr = fmt.Errorf("expected %d children but got %d", singleNumChildren, simdjsonVisitCount)
			b.Fail()
		}
		if simdjsonValueHolder == nil {
			benchErr = fmt.Errorf("value holder is nil")
			b.Fail()
		}
	}
}

// Selecting only a few fields is so painful to write, we are not doing it
// for simdjsonFile1 methods but only here.
func simdjsonFileManyFew(b *testing.B, t testfile) {
	defer func() { _ = t.handle.Close() }()

	b.ResetTimer()
	var obj = &simdjson.Object{}
	var element = &simdjson.Element{}
	for n := 0; n < b.N; n++ {
		// simdjson closes the chan when done parsing so a new one has to be
		// created on each parse.
		reuse := make(chan *simdjson.ParsedJson, 1000)
		res := make(chan simdjson.Stream, 1000)
		_, _ = t.handle.Seek(0, 0)
		simdjson.ParseNDStream(t.handle, res, reuse)

		recordCount := 0
		for got := range res {
			if got.Error != nil {
				if got.Error == io.EOF {
					break
				}
				benchErr = got.Error
				b.Fail()
			}

			var err error
			var whatval string
			var whereval int64
			err = got.Value.ForEach(func(i simdjson.Iter) error {
				recordCount++
				obj, err = i.Object(obj)
				if err != nil {
					return err
				}
				element = obj.FindKey("what", element)
				if element != nil && element.Type == simdjson.TypeString {
					whatval, err = element.Iter.String()
				}
				element = obj.FindKey("where", element)
				if element != nil && element.Type == simdjson.TypeArray {
					firstval := element.Iter.Advance()
					if firstval == simdjson.TypeObject {
						obj, err = element.Iter.Object(obj)
						element = obj.FindKey("line", element)
						if element != nil && element.Type == simdjson.TypeInt {
							whereval, err = element.Iter.Int()
						}
					}
				}
				if err = checkLog(whatval, int(whereval)); err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				benchErr = err
				b.Fail()
			}
			reuse <- got.Value
		}

		if recordCount != t.numRecords {
			benchErr = fmt.Errorf("expected %d records but got %d", t.numRecords, recordCount)
			b.Fail()
		}
	}
}

func simdjsonFileManyAll(b *testing.B, t testfile) {
	defer func() { _ = t.handle.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		reuse := make(chan *simdjson.ParsedJson, 1000)
		res := make(chan simdjson.Stream, 1000)
		_, _ = t.handle.Seek(0, 0)
		simdjson.ParseNDStream(t.handle, res, reuse)

		recordCount := 0
		for got := range res {
			if got.Error != nil {
				if got.Error == io.EOF {
					break
				}
				benchErr = got.Error
				b.Fail()
			}
			err := got.Value.ForEach(func(i simdjson.Iter) error {
				recordCount++
				simdjsonVisitCount = 0
				simdjsonValueHolder = nil
				if err := simdjsonVisitChildren(i); err != nil {
					return err
				}
				if simdjsonVisitCount != t.numChildren+1 {
					return fmt.Errorf("expected %d children but got %d", t.numChildren, simdjsonVisitCount)
				}
				if simdjsonValueHolder == nil {
					return fmt.Errorf("value holder is nil")
				}
				return nil
			})
			if err != nil {
				benchErr = err
				b.Fail()
			}
			reuse <- got.Value
		}
		if recordCount != t.numRecords {
			benchErr = fmt.Errorf("expected %d records but got %d", t.numRecords, recordCount)
			b.Fail()
		}
	}
}

// Store values to avoid compiler optimizations
var simdjsonVisitCount int
var simdjsonValueHolder interface{}
var simdjsonArr = &simdjson.Array{}
var simdjsonObj = &simdjson.Object{}

func simdjsonVisitChildren(iter simdjson.Iter) (err error) {
	simdjsonVisitCount++
	typ := iter.Type()
	switch typ {
	case simdjson.TypeArray:
		simdjsonArr, err = iter.Array(simdjsonArr)
		simdjsonArr.ForEach(func(i simdjson.Iter) {
			err = simdjsonVisitChildren(i)
		})
	case simdjson.TypeObject:
		simdjsonObj, err = iter.Object(simdjsonObj)
		simdjsonObj.ForEach(func(key []byte, i simdjson.Iter) {
			err = simdjsonVisitChildren(i)
		}, nil)
	case simdjson.TypeString:
		simdjsonValueHolder, err = iter.StringBytes()
	case simdjson.TypeInt:
		simdjsonValueHolder, err = iter.Int()
	case simdjson.TypeUint:
		simdjsonValueHolder, err = iter.Uint()
	case simdjson.TypeFloat:
		simdjsonValueHolder, err = iter.Float()
	case simdjson.TypeBool:
		simdjsonValueHolder, err = iter.Bool()
	}
	return
}
