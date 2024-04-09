// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"io"
	"log"
	"os"
	"testing"

	json2 "github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

var json2Pkg = pkg{
	name:               "json2",
	canUnmarshalStruct: true,
	unmarshal: func(data []byte, v interface{}) error {
		return json2.Unmarshal(data, v)
	},
	canMarshalStruct: true,
	marshal: func(v interface{}) ([]byte, error) {
		return json2.Marshal(v)
	},
	calls: map[string]*call{
		"single-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			go2File1Few(b)
		}},
		"single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			go2File1All(b, false)
		}},
		"single-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			go2File1All(b, true)
		}},
		"small-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			go2FileManyFew(b, openSmallLogFile())
		}},
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			go2FileManyAll(b, openSmallLogFile(), false)
		}},
		"small-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			go2FileManyAll(b, openSmallLogFile(), true)
		}},
		"large-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			go2FileManyFew(b, openLargeLogFile())
		}},
		"large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			go2FileManyAll(b, openLargeLogFile(), false)
		}},
		"large-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			go2FileManyAll(b, openLargeLogFile(), true)
		}},
	},
}

func go2File1Few(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var p PartialPatient
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		if err := json2.UnmarshalRead(f, &p); err != nil && err != io.EOF {
			benchErr = err
			b.Fail()
		} else if err = checkPatientStruct(p); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func go2File1All(b *testing.B, useStruct bool) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var p Patient
	var pi interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		if useStruct {
			if err := json2.UnmarshalRead(f, &p); err != nil && err != io.EOF {
				benchErr = err
				b.Fail()
			}
		} else {
			if err := json2.UnmarshalRead(f, &pi); err != nil && err != io.EOF {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func go2FileManyFew(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l PartialLog
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		dec := jsontext.NewDecoder(f)
		for {
			if err := json2.UnmarshalDecode(dec, &l); err == io.EOF {
				break
			} else if err != nil {
				benchErr = err
				b.Fail()
			} else if err = checkLogStruct(l); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func go2FileManyAll(b *testing.B, f *os.File, useStruct bool) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l FullLog
	var li interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		dec := jsontext.NewDecoder(f)
		for {
			if useStruct {
				if err := json2.UnmarshalDecode(dec, &l); err == io.EOF {
					break
				} else if err != nil {
					benchErr = err
					b.Fail()
				}
			} else {
				if err := json2.UnmarshalDecode(dec, &li); err == io.EOF {
					break
				} else if err != nil {
					benchErr = err
					b.Fail()
				}
			}
		}
	}
}
