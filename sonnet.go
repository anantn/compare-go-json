// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"testing"

	"github.com/sugawarayuuta/sonnet"
)

var sonnetPkg = pkg{
	name:               "sonnet",
	canUnmarshalStruct: true,
	unmarshal: func(data []byte, v interface{}) error {
		return sonnet.Unmarshal(data, v)
	},
	canMarshalStruct: true,
	marshal: func(v interface{}) ([]byte, error) {
		return sonnet.Marshal(v)
	},
	calls: map[string]*call{
		"validate-bytes": {name: "Validate", fun: sonnetValidate},
		"single-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			sonnetFile1Few(b)
		}},
		"single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonnetFile1All(b, false)
		}},
		"single-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			sonnetFile1All(b, true)
		}},
		"small-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			sonnetFileManyFew(b, openSmallLogFile())
		}},
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonnetFileManyAll(b, openSmallLogFile(), false)
		}},
		"small-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			sonnetFileManyAll(b, openSmallLogFile(), true)
		}},
		"large-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			sonnetFileManyFew(b, openLargeLogFile())
		}},
		"large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonnetFileManyAll(b, openLargeLogFile(), false)
		}},
		"large-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			sonnetFileManyAll(b, openLargeLogFile(), true)
		}},
	},
}

func sonnetValidate(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if !sonnet.Valid(sample) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
	}
}

func sonnetFile1Few(b *testing.B) {
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
		if err := sonnet.Unmarshal(j, &p); err != nil {
			benchErr = err
			b.Fail()
		}
		if err := checkPatientStruct(p); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func sonnetFile1All(b *testing.B, useStruct bool) {
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
			if err := sonnet.Unmarshal(j, &p); err != nil {
				benchErr = err
				b.Fail()
			}
		} else {
			if err := sonnet.Unmarshal(j, &pi); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func sonnetFileManyFew(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l PartialLog
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			if err := sonnet.Unmarshal(buf.Bytes(), &l); err != nil {
				benchErr = err
				b.Fail()
			} else if err := checkLogStruct(l); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func sonnetFileManyAll(b *testing.B, f *os.File, useStruct bool) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l FullLog
	var li interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			if useStruct {
				if err := sonnet.Unmarshal(buf.Bytes(), &l); err != nil {
					benchErr = err
					b.Fail()
				}
			} else {
				if err := sonnet.Unmarshal(buf.Bytes(), &li); err != nil {
					benchErr = err
					b.Fail()
				}
			}
		}
	}
}
