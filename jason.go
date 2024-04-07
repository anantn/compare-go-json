// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/antonholmquist/jason"
)

var jasonPkg = pkg{
	name: "jason",
	calls: map[string]*call{
		"single-few-keys": {name: "Unmarshal", fun: jasonFile1Few},
		"small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jasonFileManyFew(b, smallTestFile())
		}},
		"large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jasonFileManyFew(b, largeTestFile())
		}},
	},
}

func jasonFile1Few(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)
		obj, err := jason.NewObjectFromBytes(j)
		if err != nil {
			benchErr = err
			b.Fail()
		}
		idarr, _ := obj.GetObjectArray("identifier")
		coarr, _ := idarr[0].GetObjectArray("type", "coding")
		strval, _ := coarr[0].GetString("code")
		names, _ := obj.GetObjectArray("name")
		narr, _ := names[2].GetStringArray("given")
		boolval, _ := obj.GetBoolean("deceasedBoolean")
		err = checkPatient(strval, len(narr), boolval)
		if err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func jasonFileManyFew(b *testing.B, f testfile) {
	defer func() { _ = f.handle.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.handle.Seek(0, 0)
		buf := bufio.NewScanner(f.handle)
		recordCount := 0
		for buf.Scan() {
			obj, err := jason.NewObjectFromBytes(buf.Bytes())
			if err != nil {
				benchErr = err
				b.Fail()
			}
			whatval, _ := obj.GetString("what")
			wherearr, _ := obj.GetObjectArray("where")
			whereval, _ := wherearr[0].GetInt64("line")
			err = checkLog(whatval, int(whereval))
			if err != nil {
				benchErr = err
				b.Fail()
			}
			recordCount++
		}
		if recordCount != f.numRecords {
			benchErr = fmt.Errorf("expected %d records, got %d", f.numRecords, recordCount)
			b.Fail()
		}
	}
}
