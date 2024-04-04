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
	name: "json2",
	calls: map[string]*call{
		"unmarshal-single-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			go2File1(b, false)
		}},
		"unmarshal-single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			go2File1All(b, false)
		}},
		"unmarshal-small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			go2FileMany(b, openSmallLogFile(), false)
		}},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			go2FileManyAll(b, openSmallLogFile(), false)
		}},
		"unmarshal-large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			go2FileMany(b, openLargeLogFile(), false)
		}},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			go2FileManyAll(b, openLargeLogFile(), false)
		}},
		"marshal-builder": {name: "Marshal", fun: go2MarshalBuilder},
	},
}

func go2MarshalBuilder(b *testing.B) {
	var data interface{}
	err := json2.Unmarshal([]byte(getSampleLog()), &data)
	if err != nil {
		benchErr = err
		b.Fail()
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, benchErr = json2.Marshal(data); benchErr != nil {
			b.Fail()
		}
	}
}

func go2File1(b *testing.B, useStruct bool) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var p PartialPatient
	var pi interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		if useStruct {
			if err := json2.UnmarshalRead(f, &p); err != nil && err != io.EOF {
				benchErr = err
				b.Fail()
			}
			if err = checkPatientStruct(p); err != nil {
				benchErr = err
				b.Fail()
			}
		} else {
			if err := json2.UnmarshalRead(f, &pi); err != nil && err != io.EOF {
				benchErr = err
				b.Fail()
			}
			if err = checkPatientInterface(pi); err != nil {
				benchErr = err
				b.Fail()
			}
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

func go2FileMany(b *testing.B, f *os.File, useStruct bool) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l PartialLog
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
				} else if err = checkLog(l.What, l.Where[0].Line); err != nil {
					benchErr = err
					b.Fail()
				}
			} else {
				if err := json2.UnmarshalDecode(dec, &li); err == io.EOF {
					break
				} else if err != nil {
					benchErr = err
					b.Fail()
				} else if err = checkLogInterface(li); err != nil {
					benchErr = err
					b.Fail()
				}
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
