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
	name: "codec",
	calls: map[string]*call{
		"unmarshal-single-few-keys": {name: "Unmarshal", fun: codecFile1},
		"unmarshal-single-all-keys": {name: "Unmarshal", fun: codecFile1All},
		"unmarshal-small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			codecFileMany(b, openSmallLogFile())
		}},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			codecFileManyAll(b, openSmallLogFile())
		}},
		"unmarshal-large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			codecFileMany(b, openLargeLogFile())
		}},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			codecFileManyAll(b, openLargeLogFile())
		}},
		"marshal-builder": {name: "Marshal", fun: codecMarshalBuilder},
	},
}

func codecMarshalBuilder(b *testing.B) {
	var data interface{}
	var decoder *codec.Decoder = codec.NewDecoderString(getSampleLog(), jsonHandle)
	err := decoder.Decode(&data)
	if err != nil {
		benchErr = err
		b.Fail()
	}

	b.ResetTimer()
	var output = make([]byte, 0, 768)
	var encoder = codec.NewEncoderBytes(&output, jsonHandle)
	for n := 0; n < b.N; n++ {
		if benchErr = encoder.Encode(data); benchErr != nil {
			b.Fail()
		}
	}
}

func codecFile1(b *testing.B) {
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

func codecFile1All(b *testing.B) {
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
		var decoder = codec.NewDecoderBytes(j, jsonHandle)
		if err := decoder.Decode(&p); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func codecFileMany(b *testing.B, f *os.File) {
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
			if err := checkLog(l.What, l.Where[0].Line); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func codecFileManyAll(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var data interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			var decoder = codec.NewDecoderBytes(buf.Bytes(), jsonHandle)
			if err := decoder.Decode(&data); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}
