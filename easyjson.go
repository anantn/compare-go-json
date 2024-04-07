// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"testing"

	"github.com/mailru/easyjson"
)

var easyjsonPkg = pkg{
	name: "easyjson",
	calls: map[string]*call{
		"single-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			easyjsonFile1Few(b)
		}},
		"single-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			easyjsonFile1All(b)
		}},
		"small-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			easyjsonFileManyFew(b, openSmallLogFile())
		}},
		"small-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			easyjsonFileManyAll(b, openSmallLogFile())
		}},
		"large-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			easyjsonFileManyFew(b, openLargeLogFile())
		}},
		"large-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			easyjsonFileManyAll(b, openLargeLogFile())
		}},
		"marshal-builder": {name: "Marshal", fun: easyjsonMarshalBuilder, caveat: true},
	},
}

func easyjsonMarshalBuilder(b *testing.B) {
	var data PartialLog
	err := easyjson.Unmarshal([]byte(getSampleLog()), &data)
	if err != nil {
		benchErr = err
		b.Fail()
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, benchErr = easyjson.Marshal(data); benchErr != nil {
			b.Fail()
		}
	}
}

func easyjsonFile1Few(b *testing.B) {
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
		if err := easyjson.Unmarshal(j, &p); err != nil {
			benchErr = err
			b.Fail()
		}
		if err := checkPatientStruct(p); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func easyjsonFile1All(b *testing.B) {
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
		if err := easyjson.Unmarshal(j, &p); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func easyjsonFileManyFew(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l PartialLog
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			if err := easyjson.Unmarshal(buf.Bytes(), &l); err != nil {
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

func easyjsonFileManyAll(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l FullLog
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			if err := easyjson.Unmarshal(buf.Bytes(), &l); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}
