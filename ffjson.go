// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"testing"

	"github.com/pquerna/ffjson/ffjson"
)

var ffjsonPkg = pkg{
	name:               "ffjson",
	canUnmarshalStruct: true,
	unmarshal: func(data []byte, v interface{}) error {
		return ffjson.Unmarshal(data, v)
	},
	canMarshalStruct: true,
	marshal: func(v interface{}) ([]byte, error) {
		return ffjson.Marshal(v)
	},
	calls: map[string]*call{
		"single-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			ffjsonFile1Few(b)
		}},
		"single-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			ffjsonFile1All(b)
		}},
		"small-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			ffjsonFileManyFew(b, openSmallLogFile())
		}},
		"small-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			ffjsonFileManyAll(b, openSmallLogFile())
		}},
		"large-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			ffjsonFileManyFew(b, openLargeLogFile())
		}},
		"large-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			ffjsonFileManyAll(b, openLargeLogFile())
		}},
	},
}

func ffjsonFile1Few(b *testing.B) {
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
		if err := ffjson.Unmarshal(j, &p); err != nil {
			benchErr = err
			b.Fail()
		}
		if err := checkPatientStruct(p); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func ffjsonFile1All(b *testing.B) {
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
		if err := ffjson.Unmarshal(j, &p); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func ffjsonFileManyFew(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l PartialLog
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			if err := ffjson.Unmarshal(buf.Bytes(), &l); err != nil {
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

func ffjsonFileManyAll(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l FullLog
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			if err := ffjson.Unmarshal(buf.Bytes(), &l); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}
