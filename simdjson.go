// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/minio/simdjson-go"
)

var simdjsonPkg = pkg{
	name: "simdjson",
	calls: map[string]*call{
		"unmarshal-single-all-keys":     {name: "Unmarshal", fun: simdjsonFile1},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: simdjsonFileManySmall},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: simdjsonFileManyLarge},
	},
}

func simdjsonFile1(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		var pj simdjson.ParsedJson
		if _, err := simdjson.Parse(j, &pj); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func simdjsonFileManySmall(b *testing.B) {
	f := openSmallLogFile()
	defer func() { _ = f.Close() }()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		// simdjson closes the chan when done parsing so a new one has to be
		// created on each parse.
		done := make(chan bool)
		rc := make(chan simdjson.Stream, 1000)
		go func() {
			cnt := 0
			for {
				v := <-rc
				cnt++
				if v.Error != nil {
					if v.Error != io.EOF {
						benchErr = v.Error
						b.Fail()
					}
					break
				}
				if benchErr = simdjsonExtract(v.Value); benchErr != nil {
					b.Fail()
				}
			}
			done <- true
		}()
		_, _ = f.Seek(0, 0)
		simdjson.ParseNDStream(f, rc, nil)
		<-done
	}
}

func simdjsonFileManyLarge(b *testing.B) {
	// simdjson will keep reading from the file until the rc stream blocks.
	f := openLargeLogFile()
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		// simdjson closes the chan when done parsing so a new one has to be
		// created on each parse.
		done := make(chan bool)
		rc := make(chan simdjson.Stream)
		reuse := make(chan *simdjson.ParsedJson)
		go func() {
			cnt := 0
			for {
				v := <-rc
				cnt++
				if v.Error != nil {
					if v.Error != io.EOF {
						benchErr = v.Error
						b.Fail()
					}
					break
				}
				if benchErr = simdjsonExtract(v.Value); benchErr != nil {
					b.Fail()
				} else {
					go func() {
						reuse <- v.Value
					}()
				}
			}
			done <- true
		}()
		_, _ = f.Seek(0, 0)
		simdjson.ParseNDStream(f, rc, reuse)
		<-done
	}
}

func simdjsonExtract(pj *simdjson.ParsedJson) (err error) {
	tmp := &simdjson.Iter{}

	iter := pj.Iter()
	for {
		typ := iter.Advance()
		switch typ {
		case simdjson.TypeRoot:
			if typ, tmp, err = iter.Root(tmp); err != nil {
				return
			}
			switch typ {
			case simdjson.TypeArray:
				ary := &simdjson.Array{}
				if ary, err = tmp.Array(ary); err != nil {
					return
				}
				if _, err = ary.Interface(); err != nil {
					return
				}
			case simdjson.TypeObject:
				obj := &simdjson.Object{}
				if obj, err = tmp.Object(obj); err != nil {
					return
				}
				var m map[string]interface{}
				if _, err = obj.Map(m); err != nil {
					return
				}
			}
		default:
			return
		}
	}
}
