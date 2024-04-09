// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/mailru/easyjson"
)

var easyjsonPkg = pkg{
	name:               "easyjson",
	canUnmarshalStruct: true,
	unmarshal: func(data []byte, v interface{}) error {
		if i, ok := v.(easyjson.Unmarshaler); ok {
			return easyjson.Unmarshal(data, i)
		}
		return fmt.Errorf("easyjson: Unmarshal only supports easyjson.Unmarshaler")
	},
	canMarshalStruct: true,
	marshal: func(v interface{}) ([]byte, error) {
		if i, ok := v.(easyjson.Marshaler); ok {
			return easyjson.Marshal(i)
		}
		return nil, fmt.Errorf("easyjson: Marshal only supports easyjson.Marshaler")
	},
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
	},
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
