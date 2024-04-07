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
		"single-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jasonFile1Few(b, false)
		}},
		"single-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			jasonFile1Few(b, true)
		}, caveat: true},
		"small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jasonFileManyFew(b, smallTestFile(), false)
		}},
		"small-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			jasonFileManyFew(b, smallTestFile(), true)
		}, caveat: true},
		"large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jasonFileManyFew(b, largeTestFile(), false)
		}},
		"large-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			jasonFileManyFew(b, largeTestFile(), true)
		}, caveat: true},
	},
}

func jasonFile1Few(b *testing.B, useStruct bool) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	var p = getEmptyPartialPatient()
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

		if useStruct {
			p.Identifier[0].Type.Coding[0].Code = strval
			p.Name[0].Given = narr
			p.DeceasedBoolean = boolval
			err = checkPatientStruct(p)
		} else {
			err = checkPatient(strval, len(narr), boolval)
		}
		if err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func jasonFileManyFew(b *testing.B, f testfile, useStruct bool) {
	defer func() { _ = f.handle.Close() }()

	b.ResetTimer()
	var l = getEmptyPartialLog()
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

			if useStruct {
				l.What = whatval
				l.Where[0].Line = int(whereval)
				err = checkLogStruct(l)
			} else {
				err = checkLog(whatval, int(whereval))
			}
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
