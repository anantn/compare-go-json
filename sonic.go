// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"bytes"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
)

var sonicPkg = pkg{
	name: "sonic",
	calls: map[string]*call{
		"parse":            {name: "Unmarshal", fun: sonicUnmarshal},
		"validate":         {name: "Valid", fun: sonicValid},
		"decode":           {name: "Decode", fun: sonicDecode},
		"unmarshal-struct": {name: "Unmarshal", fun: sonicUnmarshalPatient},
		"marshal":          {name: "Marshal", fun: sonicMarshal},
		"marshal-struct":   {name: "Marshal", fun: sonicMarshalPatient},
		"file1":            {name: "Decode", fun: sonicFile1},
		"small-file":       {name: "Decode", fun: sonicFileManySmall},
		"large-file":       {name: "Decode", fun: sonicFileManyLarge},
	},
}

func sonicUnmarshal(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()

	var result interface{}
	for n := 0; n < b.N; n++ {
		if benchErr = sonic.Unmarshal(sample, &result); benchErr != nil {
			b.Fail()
		}
	}
}

func sonicValid(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if !sonic.Valid(sample) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
	}
}

func sonicDecode(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()

	var data interface{}
	for n := 0; n < b.N; n++ {
		dec := sonic.ConfigDefault.NewDecoder(bytes.NewReader(sample))
		if err := dec.Decode(&data); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func sonicUnmarshalPatient(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()

	var patient Patient
	for n := 0; n < b.N; n++ {
		if benchErr = sonic.Unmarshal(sample, &patient); benchErr != nil {
			b.Fail()
		}
	}
}

func sonicMarshal(b *testing.B) {
	data := loadSample()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, benchErr = sonic.Marshal(data); benchErr != nil {
			b.Fail()
		}
	}
}

func sonicMarshalPatient(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	var patient Patient
	if err := sonic.Unmarshal(sample, &patient); err != nil {
		log.Fatal(err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, benchErr = sonic.Marshal(&patient); benchErr != nil {
			b.Fail()
		}
	}
}

func sonicFile1(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := os.ReadFile(filename)
		if root, err := sonic.Get(j); err != nil {
			benchErr = err
			b.Fail()
		} else {
			strtest, err := root.Get("identifier").Index(0).Get("type").Get("coding").Index(0).Get("code").String()
			if err != nil {
				benchErr = err
				b.Fail()
			}
			arrtest, err := root.Get("name").Index(2).Get("given").Array()
			if err != nil {
				benchErr = err
				b.Fail()
			}
			booltest, err := root.Get("deceasedBoolean").Bool()
			if err != nil {
				benchErr = err
				b.Fail()
			}
			if err = checkPatient(strtest, len(arrtest), booltest); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func sonicFileManySmall(b *testing.B) {
	f := openSmallLogFile()
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if root, err := sonic.GetFromString(scanner.Text()); err != nil {
				benchErr = err
				b.Fail()
			} else {
				sonicCheckFileValues(b, &root)
			}
		}
	}
}

func sonicFileManyLarge(b *testing.B) {
	f := openLargeLogFile()
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var data interface{}
			if err := sonic.UnmarshalString(scanner.Text(), &data); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func sonicCheckFileValues(b *testing.B, root *ast.Node) {
	whatval, err := root.Get("what").String()
	if err != nil {
		benchErr = err
		b.Fail()
	}
	whereval, err := root.Get("where").Index(0).Get("line").Int64()
	if err != nil {
		benchErr = err
		b.Fail()
	}
	if err = checkLog(whatval, int(whereval)); err != nil {
		benchErr = err
		b.Fail()
	}
}
