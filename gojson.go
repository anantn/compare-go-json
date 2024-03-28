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
		"validate-string":               {name: "Valid", fun: goValidate},
		"validate-bytes":                {name: "Valid", fun: goValidate},
		"unmarshal-single-few-keys":     {name: "Unmarshal", fun: goFile1},
		"unmarshal-single-all-keys":     {name: "Unmarshal", fun: goFile1},
		"unmarshal-small-file-few-keys": {name: "Unmarshal", fun: goFileManySmall},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: goFileManySmall},
		"unmarshal-large-file-few-keys": {name: "Unmarshal", fun: goFileManyLarge},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: goFileManyLarge},
		"marshal-builder":               {name: "Marshal", fun: goMarshalBuilder},
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
	err := json.Unmarshal([]byte(`{"when":1711509483695365000,"what":"Just some fake log entry for a generated log file.","where":[{"file":"example.go","line":123}],"who":"benchmark-application","level":"INFO"}`), &data)
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
	var data interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		dec := json.NewDecoder(f)
		if err := dec.Decode(&data); err != nil && err != io.EOF {
			benchErr = err
			b.Fail()
		}
	}
}

func goFileManySmall(b *testing.B) {
	f := openSmallLogFile()
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

func goFileManyLarge(b *testing.B) {
	f := openLargeLogFile()
	defer func() { _ = f.Close() }()
	b.ResetTimer()
	var data interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		dec := json.NewDecoder(f)
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
