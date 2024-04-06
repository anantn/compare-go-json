// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

var jsoni = jsoniter.ConfigFastest
var jsoniterPkg = pkg{
	name: "jsoniter",
	calls: map[string]*call{
		"validate-bytes": {name: "Validate", fun: jsoniterValidate},
		"single-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			jsoniterFile1Few(b)
		}},
		"single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsoniterFile1All(b, false)
		}},
		"single-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			jsoniterFile1All(b, true)
		}},
		"small-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			jsoniterFileManyFew(b, openSmallLogFile())
		}},
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsoniterFileManyAll(b, openSmallLogFile(), false)
		}},
		"small-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			jsoniterFileManyAll(b, openSmallLogFile(), true)
		}},
		"large-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			jsoniterFileManyFew(b, openLargeLogFile())
		}},
		"large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsoniterFileManyAll(b, openLargeLogFile(), false)
		}},
		"large-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			jsoniterFileManyAll(b, openLargeLogFile(), true)
		}},
		"marshal-builder": {name: "Marshal", fun: jsoniterMarshalBuilder},
	},
}

func jsoniterValidate(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if !jsoniter.Valid(sample) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
	}
}

func jsoniterMarshalBuilder(b *testing.B) {
	var data interface{}
	err := jsoni.UnmarshalFromString(getSampleLog(), &data)
	if err != nil {
		benchErr = err
		b.Fail()
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, benchErr = jsoni.Marshal(data); benchErr != nil {
			b.Fail()
		}
	}
}

func jsoniterFile1Few(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var p PartialPatient
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		if err := jsoni.Unmarshal(j, &p); err != nil {
			benchErr = err
			b.Fail()
		} else if err := checkPatientStruct(p); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func jsoniterFile1All(b *testing.B, useStruct bool) {
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
		j, _ := io.ReadAll(f)
		if useStruct {
			if err := jsoni.Unmarshal(j, &p); err != nil {
				benchErr = err
				b.Fail()
			}
		} else {
			if err := jsoni.Unmarshal(j, &pi); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func jsoniterFileManyFew(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l PartialLog
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			if err := jsoni.Unmarshal(buf.Bytes(), &l); err != nil {
				benchErr = err
				b.Fail()
			}
			if err := checkLog(l.What, l.Where[0].Line); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func jsoniterFileManyAll(b *testing.B, f *os.File, useStruct bool) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l FullLog
	var li interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			if useStruct {
				if err := jsoni.Unmarshal(buf.Bytes(), &l); err != nil {
					benchErr = err
					b.Fail()
				}
			} else {
				if err := jsoni.Unmarshal(buf.Bytes(), &li); err != nil {
					benchErr = err
					b.Fail()
				}
			}
		}
	}
}
