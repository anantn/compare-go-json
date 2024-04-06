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
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			simdjsonFileManyAll(b, smallTestFile())
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

func simdjsonFileManyAll(b *testing.B, t testfile) {
	defer func() { _ = t.handle.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		// simdjson closes the chan when done parsing so a new one has to be
		// created on each parse.
		reuse := make(chan *simdjson.ParsedJson, 1000)
		res := make(chan simdjson.Stream, 1000)

		_, _ = t.handle.Seek(0, 0)
		simdjson.ParseNDStream(t.handle, res, reuse)

		for got := range res {
			if got.Error != nil {
				if got.Error == io.EOF {
					break
				}
				benchErr = got.Error
				b.Fail()
			}
			err := got.Value.ForEach(func(i simdjson.Iter) error {
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
