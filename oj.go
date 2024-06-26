// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/ohler55/ojg/oj"
)

var ojPkg = pkg{
	name:               "oj",
	canUnmarshalStruct: true,
	unmarshal: func(data []byte, v interface{}) error {
		return oj.Unmarshal(data, v)
	},
	canMarshalStruct: true,
	marshal: func(v interface{}) ([]byte, error) {
		return []byte(oj.JSON(v)), nil
	},
	calls: map[string]*call{
		"validate-bytes": {name: "Validate", fun: ojValidate},
		// oj does not seem to work with anonymous structs like in PartialPatient
		//	"single-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
		//		ojFile1Few(b)
		//	}},
		//
		// oj also has issues unmarhaling into a struct with int64/float64 fields?
		// disabling these benchmarks for now
		//
		// "small-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
		//		ojFileManyAll(b, openSmallLogFile(), true)
		// }},
		// "large-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
		//		ojFileManyFew(b, openLargeLogFile())
		// }},
		// "large-file-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
		//		ojFileManyAll(b, openLargeLogFile(), true)
		// }},
		"single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			ojFile1All(b, false)
		}},
		"single-all-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			ojFile1All(b, true)
		}},
		"small-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			ojFileManyFew(b, openSmallLogFile())
		}},
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			ojFileManyAll(b, openSmallLogFile(), false)
		}},
		"large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			ojFileManyAll(b, openLargeLogFile(), false)
		}},
	},
}

func ojValidate(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()
	var v oj.Validator
	for n := 0; n < b.N; n++ {
		if benchErr = v.Validate(sample); benchErr != nil {
			b.Fail()
		}
	}
}

func ojFile1Few(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	ojp := oj.Parser{Reuse: true}
	b.ResetTimer()

	var p PartialPatient
	for n := 0; n < b.N; n++ {
		if err := ojp.Unmarshal(sample, &p); err != nil {
			benchErr = err
			b.Fail()
		} else if err := checkPatientStruct(p); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func ojFile1All(b *testing.B, useStruct bool) {
	sample, _ := os.ReadFile(filename)
	ojp := oj.Parser{Reuse: true}
	b.ResetTimer()

	var p Patient
	var pi interface{}
	for n := 0; n < b.N; n++ {
		if useStruct {
			if err := ojp.Unmarshal(sample, &p); err != nil {
				benchErr = err
				b.Fail()
			}
		} else {
			if err := ojp.Unmarshal(sample, &pi); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func ojFileManyFew(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	ojp := &oj.Parser{Reuse: true}
	b.ResetTimer()

	var l PartialLog
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			if err := ojp.Unmarshal(buf.Bytes(), &l); err != nil {
				benchErr = err
				b.Fail()
			} else if err := checkLogStruct(l); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func ojFileManyAll(b *testing.B, f *os.File, useStruct bool) {
	defer func() { _ = f.Close() }()
	ojp := &oj.Parser{Reuse: true}
	b.ResetTimer()

	var l FullLog
	var li interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			if useStruct {
				if err := ojp.Unmarshal(buf.Bytes(), &l); err != nil {
					benchErr = err
					b.Fail()
				}
			} else {
				if err := ojp.Unmarshal(buf.Bytes(), &li); err != nil {
					benchErr = err
					b.Fail()
				}
			}
		}
	}
}
