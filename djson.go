// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/a8m/djson"
)

var djsonPkg = pkg{
	name: "djson",
	calls: map[string]*call{
		"single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			djsonFile1All(b)
		}},
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			djsonFileManyAll(b, openSmallLogFile())
		}},
		"large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			djsonFileManyAll(b, openLargeLogFile())
		}},
	},
}

func djsonFile1All(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var pi map[string]interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		dec := djson.NewDecoder(j)
		dec.AllocString()
		if pi, err = dec.DecodeObject(); err != nil {
			benchErr = err
			b.Fail()
		}
		if len(pi) != singleNumRootKeys {
			benchErr = fmt.Errorf("expected %d root keys, got %d", singleNumRootKeys, len(pi))
			b.Fail()
		}
	}
}

func djsonFileManyAll(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()
	b.ResetTimer()

	var err error
	var li map[string]interface{}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			dec := djson.NewDecoder(scanner.Bytes())
			dec.AllocString()
			if li, err = dec.DecodeObject(); err != nil {
				benchErr = err
				b.Fail()
			}
			if len(li) != logNumRootKeys {
				benchErr = fmt.Errorf("expected %d root keys, got %d", logNumRootKeys, len(li))
				b.Fail()
			}
		}
	}
}
