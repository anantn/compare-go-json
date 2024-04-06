// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"testing"

	segment "github.com/segmentio/encoding/json"
)

var segmentPkg = pkg{
	name: "segment",
	calls: map[string]*call{
		"validate-bytes": {name: "Validate", fun: segmentValidate},
		"single-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			segmentFile1Few(b)
		}},
		"single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			segmentFile1All(b, false)
		}},
		"single-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			segmentFile1All(b, true)
		}},
		"small-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			segmentFileManyFew(b, openSmallLogFile())
		}},
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			segmentFileManyAll(b, openSmallLogFile(), false)
		}},
		"small-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			segmentFileManyAll(b, openSmallLogFile(), true)
		}},
		"large-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			segmentFileManyFew(b, openLargeLogFile())
		}},
		"large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			segmentFileManyAll(b, openLargeLogFile(), false)
		}},
		"large-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			segmentFileManyAll(b, openLargeLogFile(), true)
		}},
		"marshal-builder": {name: "Marshal", fun: segmentMarshalBuilder},
	},
}

func segmentValidate(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if !segment.Valid(sample) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
	}
}

func segmentMarshalBuilder(b *testing.B) {
	var data interface{}
	err := segment.Unmarshal([]byte(getSampleLog()), &data)
	if err != nil {
		benchErr = err
		b.Fail()
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, benchErr = segment.Marshal(data); benchErr != nil {
			b.Fail()
		}
	}
}

func segmentFile1Few(b *testing.B) {
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
		if err := segment.Unmarshal(j, &p); err != nil {
			benchErr = err
			b.Fail()
		}
		if err := checkPatientStruct(p); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func segmentFile1All(b *testing.B, useStruct bool) {
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
			if err := segment.Unmarshal(j, &p); err != nil {
				benchErr = err
				b.Fail()
			}
		} else {
			if err := segment.Unmarshal(j, &pi); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func segmentFileManyFew(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l PartialLog
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			if err := segment.Unmarshal(buf.Bytes(), &l); err != nil {
				benchErr = err
				b.Fail()
			} else if err := checkLogStruct(l); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func segmentFileManyAll(b *testing.B, f *os.File, useStruct bool) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l FullLog
	var li interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			if useStruct {
				if err := segment.Unmarshal(buf.Bytes(), &l); err != nil {
					benchErr = err
					b.Fail()
				}
			} else {
				if err := segment.Unmarshal(buf.Bytes(), &li); err != nil {
					benchErr = err
					b.Fail()
				}
			}
		}
	}
}
