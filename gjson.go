// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/tidwall/gjson"
)

var gjsonPkg = pkg{
	name: "gjson",
	calls: map[string]*call{
		"parse":      {name: "ParseBytes", fun: gjsonParse},
		"validate":   {name: "Validate", fun: gjsonValid},
		"file1":      {name: "Decode", fun: gjsonFile1},
		"small-file": {name: "Decode", fun: gjsonFileManySmall},
		"large-file": {name: "Decode", fun: gjsonFileManyLarge},
	},
}

func gjsonParse(b *testing.B) {
	sample, _ := ioutil.ReadFile(filename)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if !gjson.ValidBytes(sample) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
		_ = gjson.ParseBytes(sample).Value()
	}
}

func gjsonValid(b *testing.B) {
	sample, _ := ioutil.ReadFile(filename)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if !gjson.ValidBytes(sample) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
	}
}

func gjsonFile1(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		v := string(j)
		if !gjson.Valid(v) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
		result := gjson.Parse(v)
		if !result.IsObject() {
			benchErr = errors.New("expected object")
			b.Fail()
		} else {
			strtest := result.Get("identifier.0.type.coding.0.code").String()
			arrtest := result.Get("name.2.given").Array()
			booltest := result.Get("deceasedBoolean").Bool()
			err = checkPatient(strtest, len(arrtest), booltest)
			if err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func gjsonFileManySmall(b *testing.B) {
	gjsonCheckFileValues(b, openSmallLogFile())
}

func gjsonFileManyLarge(b *testing.B) {
	gjsonCheckFileValues(b, openLargeLogFile())
}

func gjsonCheckFileValues(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			v := scanner.Text()
			if !gjson.Valid(v) {
				benchErr = errors.New("JSON not valid")
				b.Fail()
			}
			if result := gjson.Parse(v); !result.IsObject() {
				benchErr = errors.New("expected object")
				b.Fail()
			} else {
				if !result.IsObject() {
					benchErr = errors.New("expected object")
					b.Fail()
				} else {
					whatval := result.Get("what").String()
					whereval := result.Get("where.0.line").Int()
					err := checkLog(whatval, int(whereval))
					if err != nil {
						benchErr = err
						b.Fail()
					}
				}
			}
		}
	}
}
