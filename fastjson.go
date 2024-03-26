// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/valyala/fastjson"
)

var fastjsonPkg = pkg{
	name: "fastjson",
	calls: map[string]*call{
		"validate":   {name: "Valid", fun: fastjsonValidate},
		"file1":      {name: "Decode", fun: fastjsonFile1},
		"small-file": {name: "Decode", fun: fastjsonFileManySmall},
		"large-file": {name: "Decode", fun: fastjsonFileManyLarge},
	},
}

func fastjsonValidate(b *testing.B) {
	sample, _ := ioutil.ReadFile(filename)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if benchErr = fastjson.ValidateBytes(sample); benchErr != nil {
			b.Fail()
		}
	}
}

func fastjsonFile1(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	var p fastjson.Parser
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		if _, err := p.ParseBytes(j); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func fastjsonFileManySmall(b *testing.B) {
	f := openSmallLogFile()
	defer func() { _ = f.Close() }()

	var sc fastjson.Scanner
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := ioutil.ReadAll(f)
		sc.InitBytes(j)
		for sc.Next() {
			if sc.Value().Exists() {
				continue
			}
			if err := sc.Error(); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func fastjsonFileManyLarge(b *testing.B) {
	f := openLargeLogFile()
	defer func() { _ = f.Close() }()

	var sc fastjson.Scanner
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			sc.InitBytes(buf.Bytes())
			for sc.Next() {
				if sc.Value().Exists() {
					continue
				}
				if err := sc.Error(); err != nil {
					benchErr = err
					b.Fail()
				}
			}
		}
	}
}
