// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bytes"
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
		"unmarshal-single-all-keys": {name: "Unmarshal", fun: goFile1All},
		"unmarshal-small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			goFileMany(b, openSmallLogFile())
		}},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			goFileManyAll(b, openSmallLogFile())
		}},
		"unmarshal-large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			goFileMany(b, openLargeLogFile())
		}},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			goFileManyAll(b, openLargeLogFile())
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

	var p PartialPatient
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		dec := json.NewDecoder(f)
		if err := dec.Decode(&p); err != nil && err != io.EOF {
			benchErr = err
			b.Fail()
		}
		if err = checkPatient(
			p.Identifier[0].Type.Coding[0].Code,
			len(p.Name[0].Given),
			p.DeceasedBoolean); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func goFile1All(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var p Patient
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		dec := json.NewDecoder(f)
		if err := dec.Decode(&p); err != nil && err != io.EOF {
			benchErr = err
			b.Fail()
		}
	}
}

func goFileMany(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l PartialLog
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		dec := json.NewDecoder(bytes.NewReader(j))
		for {
			if err := dec.Decode(&l); err == io.EOF {
				break
			} else if err != nil {
				benchErr = err
				b.Fail()
			} else if err = checkLog(l.What, l.Where[0].Line); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func goFileManyAll(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var data interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		dec := json.NewDecoder(bytes.NewReader(j))
		for {
			if err := dec.Decode(&data); err == io.EOF {
				break
			} else if err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}
