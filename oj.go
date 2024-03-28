// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ohler55/ojg/oj"
)

var ojPkg = pkg{
	name: "oj",
	calls: map[string]*call{
		"validate-bytes":            {name: "Validate", fun: ojValidate},
		"unmarshal-single-all-keys": {name: "Unmarshal", fun: ojFile1All},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			ojFileManyAll(b, openSmallLogFile())
		}},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			ojFileManyAll(b, openLargeLogFile())
		}},
		"marshal-builder": {name: "Builder", fun: ojMarshalBuilder},
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

func ojFile1All(b *testing.B) {
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

var ojAllData any

func ojFileManyAll(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	p := &oj.Parser{Reuse: true}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			if _, err := p.Parse(buf.Bytes(), func(result any) {
				ojAllData = result
			}); err != nil {
				benchErr = err
				b.Fail()
			}
		}
		if ojAllData == nil {
			benchErr = errors.New("expected a value, got nil")
			b.Fail()
		}
	}
}
