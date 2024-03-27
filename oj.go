// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/ohler55/ojg"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
)

var ojPkg = pkg{
	name: "oj",
	calls: map[string]*call{
		"parse":            {name: "Parse", fun: ojParse},
		"validate":         {name: "Validate", fun: ojValidate},
		"decode":           {name: "Tokenize", fun: ojTokenize},
		"unmarshal-struct": {name: "Unmarshal", fun: ojUnmarshalPatient},
		"marshal":          {name: "JSON", fun: ojJSON},
		"marshal-struct":   {name: "Marshal", fun: ojMarshalPatient},
		"file1":            {name: "ParseReader", fun: ojFile1},
		"small-file":       {name: "ParseReader", fun: ojFileManySmallLoad},
		"large-file":       {name: "ParseReader", fun: ojFileManyLarge},
	},
}

func ojParse(b *testing.B) {
	sample, _ := ioutil.ReadFile(filename)
	b.ResetTimer()
	p := &oj.Parser{Reuse: true}
	for n := 0; n < b.N; n++ {
		if _, benchErr = p.Parse(sample); benchErr != nil {
			b.Fail()
		}
	}
}

func ojValidate(b *testing.B) {
	sample, _ := ioutil.ReadFile(filename)
	b.ResetTimer()
	var v oj.Validator
	for n := 0; n < b.N; n++ {
		if benchErr = v.Validate(sample); benchErr != nil {
			b.Fail()
		}
	}
}

func ojUnmarshalPatient(b *testing.B) {
	sample, _ := ioutil.ReadFile(filename)
	p := oj.Parser{Reuse: true}
	b.ResetTimer()
	var out Patient
	for n := 0; n < b.N; n++ {
		if err := p.Unmarshal(sample, &out); err != nil {
			log.Fatal(err)
		}
	}
}

func ojTokenize(b *testing.B) {
	sample, _ := ioutil.ReadFile(filename)
	b.ResetTimer()
	h := oj.ZeroHandler{}
	t := oj.Tokenizer{}
	for n := 0; n < b.N; n++ {
		if err := t.Parse(sample, &h); err != nil {
			log.Fatal(err)
		}
	}
}

func ojJSON(b *testing.B) {
	data := loadSample()
	wr := oj.Writer{Options: ojg.Options{OmitNil: true}}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = wr.MustJSON(data)
	}
}

func ojMarshalPatient(b *testing.B) {
	sample, _ := ioutil.ReadFile(filename)
	var patient Patient
	if err := oj.Unmarshal(sample, &patient); err != nil {
		log.Fatal(err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		//_ = wr.MustJSON(&patient)
		if _, err := oj.Marshal(&patient); err != nil {
			log.Fatal(err)
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
		j, _ := ioutil.ReadAll(f)
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

func ojFileManySmallChan(b *testing.B) {
	f := openSmallLogFile()
	defer func() { _ = f.Close() }()

	rc := make(chan interface{}, 1000)
	ready := make(chan bool)
	go func() {
		ready <- true
		for {
			if v := <-rc; v == nil {
				break
			}
		}
	}()
	<-ready
	b.ResetTimer()

	var p oj.Parser
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		if _, benchErr = p.ParseReader(f, rc); benchErr != nil {
			b.Fail()
		}
	}
	rc <- nil
}

func ojCb(_ interface{}) bool {
	return false
}

func ojFileManySmallReader(b *testing.B) {
	f := openSmallLogFile()
	defer func() { _ = f.Close() }()
	b.ResetTimer()
	p := &oj.Parser{Reuse: true}
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		if _, benchErr = p.ParseReader(f, ojCb); benchErr != nil {
			b.Fail()
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
