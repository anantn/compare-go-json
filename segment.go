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
		"validate-bytes":            {name: "Validate", fun: segmentValidate},
		"unmarshal-single-few-keys": {name: "Unmarshal", fun: segmentFile1},
		"unmarshal-single-all-keys": {name: "Unmarshal", fun: segmentFile1All},
		"unmarshal-small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			segmentFileMany(b, openSmallLogFile())
		}},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			segmentFileManyAll(b, openSmallLogFile())
		}},
		"unmarshal-large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			segmentFileMany(b, openLargeLogFile())
		}},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			segmentFileManyAll(b, openLargeLogFile())
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

func segmentFile1(b *testing.B) {
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

func segmentFile1All(b *testing.B) {
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
		if err := segment.Unmarshal(j, &p); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func segmentFileMany(b *testing.B, f *os.File) {
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
			}
			if err := checkLog(l.What, l.Where[0].Line); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func segmentFileManyAll(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var data interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			if err := segment.Unmarshal(buf.Bytes(), &data); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}
