// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"testing"

	"github.com/ugorji/go/codec"
)

var jsonHandle codec.Handle = new(codec.JsonHandle)

var codecPkg = pkg{
	name:             "codec",
	canMarshalStruct: true,
	marshal: func(v interface{}) ([]byte, error) {
		var output []byte
		encoder := codec.NewEncoderBytes(&output, jsonHandle)
		if err := encoder.Encode(v); err != nil {
			return nil, err
		}
		return output, nil
	},
	canUnmarshalStruct: true,
	unmarshal: func(data []byte, v interface{}) error {
		decoder := codec.NewDecoderBytes(data, jsonHandle)
		return decoder.Decode(v)
	},
	calls: map[string]*call{
		"single-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			codecFile1Few(b)
		}},
		"single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			codecFile1All(b, false)
		}},
		"single-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			codecFile1All(b, true)
		}},
		"small-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			codecFileManyFew(b, openSmallLogFile())
		}},
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			codecFileManyAll(b, openSmallLogFile(), false)
		}},
		"small-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			codecFileManyAll(b, openSmallLogFile(), true)
		}},
		"large-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			codecFileManyFew(b, openLargeLogFile())
		}},
		"large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			codecFileManyAll(b, openLargeLogFile(), false)
		}},
		"large-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			codecFileManyAll(b, openLargeLogFile(), true)
		}},
	},
}

func codecFile1Few(b *testing.B) {
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
		var decoder = codec.NewDecoderBytes(j, jsonHandle)
		if err := decoder.Decode(&p); err != nil {
			benchErr = err
			b.Fail()
		}
		if err := checkPatientStruct(p); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func codecFile1All(b *testing.B, useStruct bool) {
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
		var decoder = codec.NewDecoderBytes(j, jsonHandle)
		if useStruct {
			if err := decoder.Decode(&p); err != nil {
				benchErr = err
				b.Fail()
			}
		} else {
			if err := decoder.Decode(&pi); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func codecFileManyFew(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l PartialLog
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			var decoder = codec.NewDecoderBytes(buf.Bytes(), jsonHandle)
			if err := decoder.Decode(&l); err != nil {
				benchErr = err
				b.Fail()
			}
			if err := checkLogStruct(l); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func codecFileManyAll(b *testing.B, f *os.File, useStruct bool) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var l FullLog
	var li interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			var decoder = codec.NewDecoderBytes(buf.Bytes(), jsonHandle)
			if useStruct {
				if err := decoder.Decode(&l); err != nil {
					benchErr = err
					b.Fail()
				}
			} else {
				if err := decoder.Decode(&li); err != nil {
					benchErr = err
					b.Fail()
				}
			}
		}
	}
}
