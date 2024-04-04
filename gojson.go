// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"testing"
)

var jsonPkg = pkg{
	name: "json",
	calls: map[string]*call{
		"validate-string":           {name: "Validate", fun: goValidate},
		"validate-bytes":            {name: "Validate", fun: goValidate},
		"unmarshal-single-few-keys": {name: "Unmarshal", fun: goFile1},
		"unmarshal-single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			goFile1All(b, false)
		}},
		"unmarshal-small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			goFileMany(b, openSmallLogFile())
		}},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			goFileManyAll(b, openSmallLogFile(), false)
		}},
		"unmarshal-large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			goFileMany(b, openLargeLogFile())
		}},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			goFileManyAll(b, openLargeLogFile(), false)
		}},
		"marshal-builder": {name: "Marshal", fun: goMarshalBuilder},
	},
}

func goValidate(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if !json.Valid(sample) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
	}
}

func goMarshalBuilder(b *testing.B) {
	var data interface{}
	err := json.Unmarshal([]byte(getSampleLog()), &data)
	if err != nil {
		benchErr = err
		b.Fail()
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, benchErr = json.Marshal(data); benchErr != nil {
			b.Fail()
		}
	}
}

func goFile1(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	// Allow partial reading into struct just to get baseline
	var p PartialPatient
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		if err := json.Unmarshal(j, &p); err != nil {
			benchErr = err
			b.Fail()
		}
		if err = checkPatientStruct(p); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func goFile1All(b *testing.B, useStruct bool) {
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
			if err := json.Unmarshal(j, &p); err != nil {
				benchErr = err
				b.Fail()
			}
		} else {
			if err := json.Unmarshal(j, &pi); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func goFileMany(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	// Allow partial reading into struct just to get baseline
	var l PartialLog
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if err := json.Unmarshal(scanner.Bytes(), &l); err != nil {
				benchErr = err
				b.Fail()
			} else if err = checkLog(l.What, l.Where[0].Line); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func goFileManyAll(b *testing.B, f *os.File, useStruct bool) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l FullLog
	var li interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if useStruct {
				if err := json.Unmarshal(scanner.Bytes(), &l); err != nil {
					benchErr = err
					b.Fail()
				}
			} else {
				if err := json.Unmarshal(scanner.Bytes(), &li); err != nil {
					benchErr = err
					b.Fail()
				}
			}
		}
	}
}
