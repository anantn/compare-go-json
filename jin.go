// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ecoshub/jin"
)

var jinPkg = pkg{
	name: "jin",
	calls: map[string]*call{
		"single-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jinFile1Few(b, false)
		}},
		"single-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			jinFile1Few(b, true)
		}, caveat: true},
		"single-all-keys": {name: "Unmarshal", fun: jinFile1All},
		"small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jinFileManyFew(b, smallTestFile(), false)
		}},
		"small-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			jinFileManyFew(b, smallTestFile(), true)
		}, caveat: true},
		"small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jinFileManyAll(b, smallTestFile())
		}},
		"large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jinFileManyFew(b, largeTestFile(), false)
		}},
		"large-file-few-keys-struct": {name: "Unmarshal", fun: func(b *testing.B) {
			jinFileManyFew(b, largeTestFile(), true)
		}, caveat: true},
		"large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			jinFileManyAll(b, largeTestFile())
		}},
		"marshal-builder": {name: "Marshal", fun: jinMarshalBuilder},
	},
}

func jinMarshalBuilder(b *testing.B) {
	// {"when":1711509483695365000,"what":"Just some fake log entry for a generated log file.","where":[{"file":"example.go","line":123}],"who":"benchmark-application","level":"INFO"}
	rootSchema := jin.MakeScheme("when", "what", "where", "who", "level")
	whereSchema := jin.MakeScheme("file", "line")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		output := rootSchema.MakeJson(
			time.Now().UnixNano(),
			"Just some fake log entry for a generated log file.",
			// Can't set array directly, this seems sub-optimal
			string(jin.MakeArray(string(whereSchema.MakeJson("example.go", 123)))),
			"benchmark-application",
			"INFO")
		if benchErr = checkMarshalCustom(string(output)); benchErr != nil {
			b.Fail()
		}
	}
}

func jinFile1Few(b *testing.B, useStruct bool) {
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
		obj := jin.New(j)
		strval, _ := obj.GetString("identifier", "0", "type", "coding", "0", "code")
		arrval, _ := obj.GetStringArray("name", "2", "given")
		boolval, _ := obj.GetBool("deceasedBoolean")

		if useStruct {
			p.Identifier[0].Type.Coding[0].Code = strval
			p.Name[0].Given = arrval
			p.DeceasedBoolean = boolval
			err = checkPatientStruct(p)
		} else {
			err = checkPatient(strval, len(arrval), boolval)
		}
		if err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func jinFile1All(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := io.ReadAll(f)

		jinVisitCount = 0
		jinValueHolder = nil
		err := jinIterate(j)
		if err != nil {
			benchErr = err
			b.Fail()
		}
		// Includes root object
		if jinVisitCount != singleNumChildren+1 {
			benchErr = fmt.Errorf("expected %d visits, got %d", singleNumChildren, jinVisitCount)
			b.Fail()
		}
		if jinValueHolder == nil {
			benchErr = fmt.Errorf("expected a value, got nil")
			b.Fail()
		}
	}
}

func jinFileManyFew(b *testing.B, f testfile, useStruct bool) {
	defer func() { _ = f.handle.Close() }()

	b.ResetTimer()
	var err error
	var l = getEmptyPartialLog()
	for n := 0; n < b.N; n++ {
		_, _ = f.handle.Seek(0, 0)
		buf := bufio.NewScanner(f.handle)
		recordCount := 0
		for buf.Scan() {
			obj := jin.New(buf.Bytes())
			whatval, _ := obj.GetString("what")
			whereval, _ := obj.GetInt("where", "0", "line")

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

func jinFileManyAll(b *testing.B, f testfile) {
	defer func() { _ = f.handle.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.handle.Seek(0, 0)
		buf := bufio.NewScanner(f.handle)
		recordCount := 0
		for buf.Scan() {
			jinVisitCount = 0
			jinValueHolder = nil
			err := jinIterate(buf.Bytes())
			if err != nil {
				benchErr = fmt.Errorf("expected a value, got nil")
				b.Fail()
			}
			// Includes root object
			if jinVisitCount != f.numChildren+1 {
				benchErr = fmt.Errorf("expected %d visits, got %d", f.numChildren, jinVisitCount)
				b.Fail()
			}
			if jinValueHolder == nil {
				benchErr = fmt.Errorf("expected a value, got nil")
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

var jinVisitCount int
var jinValueHolder interface{}

func jinIterate(v []byte) error {
	jinVisitCount++
	t, err := jin.GetType(v)
	if err != nil {
		return err
	}

	switch t {
	case jin.TypeObject:
		err = jin.IterateKeyValue(v, func(key, value []byte) (bool, error) {
			err = jinIterate(value)
			return err == nil, err
		})
		if err != nil {
			return err
		}
	case jin.TypeArray:
		err = jin.IterateArray(v, func(value []byte) (bool, error) {
			err = jinIterate(value)
			return err == nil, err
		})
		if err != nil {
			return err
		}
	case jin.TypeString:
		jinValueHolder, err = jin.GetString(v)
		if err != nil {
			return err
		}
	case jin.TypeNumber:
		jinValueHolder, err = jin.GetInt(v)
		if err != nil {
			return err
		}
	case jin.TypeBoolean:
		jinValueHolder, err = jin.GetBool(v)
		if err != nil {
			return err
		}
	}
	return nil
}
