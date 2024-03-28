// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
)

var ojPkg = pkg{
	name: "oj",
	calls: map[string]*call{
		"validate-bytes":                {name: "Validate", fun: ojValidate},
		"unmarshal-single-few-keys":     {name: "Unmarshal", fun: ojFile1},
		"unmarshal-single-all-keys":     {name: "Unmarshal", fun: ojUnmarshalPatient},
		"unmarshal-small-file-few-keys": {name: "Unmarshal", fun: ojFileManySmallLoad},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: ojFileManySmallLoad},
		"unmarshal-large-file-few-keys": {name: "Unmarshal", fun: ojFileManyLarge},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: ojFileManyLarge},
		"marshal-builder":               {name: "Builder", fun: ojMarshalBuilder},
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

func ojUnmarshalPatient(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	p := oj.Parser{Reuse: true}
	b.ResetTimer()
	var out Patient
	for n := 0; n < b.N; n++ {
		if err := p.Unmarshal(sample, &out); err != nil {
			log.Fatal(err)
		}
	}
}

func ojMarshalBuilder(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var bu oj.Builder
		_ = bu.Object()
		_ = bu.Value(time.Now().UnixNano(), "when")
		_ = bu.Value("Just some fake log entry for a generated log file.", "what")
		_ = bu.Array("where")
		_ = bu.Object()
		_ = bu.Value("example.go", "file")
		_ = bu.Value(123, "line")
		bu.Pop()
		bu.Pop()
		_ = bu.Value("benchmark-application", "who")
		_ = bu.Value("INFO", "level")
		bu.PopAll()
		entry := bu.Result()
		if err := checkMarshalCustom(oj.JSON(entry)); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func ojFile1(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	strpath := jp.MustParseString("$.identifier[0].type.coding[0].code")
	arrpath := jp.MustParseString("$.name[2].given")
	boolpath := jp.MustParseString("$.deceasedBoolean")

	b.ResetTimer()
	p := &oj.Parser{Reuse: true}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		var result any
		if result, benchErr = p.Parse(j); benchErr != nil {
			b.Fail()
		} else {
			strtest := strpath.Get(result)[0].(string)
			arrtest := arrpath.Get(result)[0].([]interface{})
			booltest := boolpath.Get(result)[0].(bool)
			if err = checkPatient(strtest, len(arrtest), booltest); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func ojFileManySmallLoad(b *testing.B) {
	ojCheckFileValues(b, openSmallLogFile())
}

func ojFileManyLarge(b *testing.B) {
	ojCheckFileValues(b, openLargeLogFile())
}

func ojCheckFileValues(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()

	whatpath := jp.MustParseString("$.what")
	wherepath := jp.MustParseString("$.where[0].line")

	b.ResetTimer()
	p := &oj.Parser{Reuse: true}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		if _, err := p.ParseReader(f, func(result any) {
			whatval := whatpath.Get(result)[0].(string)
			whereval := wherepath.Get(result)[0].(int64)
			if err := checkLog(whatval, int(whereval)); err != nil {
				benchErr = err
				b.Fail()
			}
		}); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}
