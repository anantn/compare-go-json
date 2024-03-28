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
		"validate-bytes":            {name: "Validate", fun: jsoniterValidate},
		"unmarshal-single-few-keys": {name: "Unmarshal", fun: jsoniterFile1},
		"unmarshal-single-all-keys": {name: "Unmarshal", fun: jsoniterFile1All},
		"unmarshal-small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsoniterFileMany(b, openSmallLogFile())
		}},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsoniterFileManyAll(b, openSmallLogFile())
		}},
		"unmarshal-large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsoniterFileMany(b, openLargeLogFile())
		}},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jsoniterFileManyAll(b, openLargeLogFile())
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

func jsoniterFile1(b *testing.B) {
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
		}
		if err := checkPatientStruct(p); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func jsoniterFile1All(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var p Patient
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		if err := jsoni.Unmarshal(j, &p); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func jsoniterFileMany(b *testing.B, f *os.File) {
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

func jsoniterFileManyAll(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var data interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			if err := jsoni.Unmarshal(buf.Bytes(), &data); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}
