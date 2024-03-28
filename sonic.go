// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
)

var sonicShouldValidate bool

var sonicPkg = pkg{
	name: "sonic",
	calls: map[string]*call{
		"validate-bytes":  {name: "Validate", fun: sonicValidate},
		"validate-string": {name: "Validate", fun: sonicValidateString},
		"unmarshal-single-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicShouldValidate = false
			sonicFile1(b)
		}},
		"unmarshal-single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicShouldValidate = false
			sonicFile1All(b)
		}},
		"unmarshal-small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicShouldValidate = false
			sonicFileMany(b, openSmallLogFile(), smallLogFileLen)
		}},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicShouldValidate = false
			sonicFileManyAll(b, openSmallLogFile(), smallLogFileLen)
		}},
		"unmarshal-large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicShouldValidate = false
			sonicFileMany(b, openLargeLogFile(), largeLogFileLen)
		}},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicShouldValidate = false
			sonicFileManyAll(b, openLargeLogFile(), largeLogFileLen)
		}},
		"marshal-builder": {name: "Marshal", fun: sonicMarshalBuilder},
	},
}

var sonicValidatePkg = pkg{
	name: "sonic-v",
	calls: map[string]*call{
		"unmarshal-single-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicShouldValidate = true
			sonicFile1(b)
		}},
		"unmarshal-single-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicShouldValidate = true
			sonicFile1All(b)
		}},
		"unmarshal-small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicShouldValidate = true
			sonicFileMany(b, openSmallLogFile(), smallLogFileLen)
		}},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicShouldValidate = true
			sonicFileManyAll(b, openSmallLogFile(), smallLogFileLen)
		}},
		"unmarshal-large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicShouldValidate = true
			sonicFileMany(b, openLargeLogFile(), largeLogFileLen)
		}},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicShouldValidate = true
			sonicFileManyAll(b, openLargeLogFile(), largeLogFileLen)
		}},
	},
}

func sonicValidate(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if !sonic.Valid(sample) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
	}
}

func sonicValidateString(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if !sonic.ValidString(string(sample)) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
	}
}

func sonicMarshalBuilder(b *testing.B) {
	// {"when":1711509483695365000,"what":"Just some fake log entry for a generated log file.","where":[{"file":"example.go","line":123}],"who":"benchmark-application","level":"INFO"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dst := ast.NewObject(nil)
		dst.Set("when", ast.NewNumber(strconv.FormatInt(time.Now().UnixNano(), 10)))
		dst.Set("what", ast.NewString("Just some fake log entry for a generated log file."))
		where := ast.NewObject([]ast.Pair{
			{Key: "file", Value: ast.NewString("example.go")},
			{Key: "line", Value: ast.NewNumber(strconv.Itoa(123))},
		})
		dst.Set("where", ast.NewArray([]ast.Node{where}))
		dst.Set("who", ast.NewString("benchmark-application"))
		dst.Set("level", ast.NewString("INFO"))
		output, err := dst.MarshalJSON()
		if err != nil {
			benchErr = err
			b.Fail()
		}
		if benchErr = checkMarshalCustom(string(output)); benchErr != nil {
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
		if sonicShouldValidate && !sonic.Valid(j) {
			benchErr = errors.New("json not valid")
			b.Fail()
		}
		strnode, _ := sonic.Get(j, "identifier", 0, "type", "coding", 0, "code")
		strval, _ := strnode.String()
		arrnode, _ := sonic.Get(j, "name", 2, "given")
		arrval, _ := arrnode.Array()
		boolnode, _ := sonic.Get(j, "deceasedBoolean")
		boolval, _ := boolnode.Bool()
		if err := checkPatient(strval, len(arrval), boolval); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func sonicFile1All(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		j, _ := os.ReadFile(filename)
		if sonicShouldValidate && !sonic.Valid(j) {
			benchErr = errors.New("json not valid")
			b.Fail()
		}
		if root, err := sonic.Get(j); err != nil {
			benchErr = err
			b.Fail()
		} else {
			root.ForEach(sonicVisitChildren)
			if sonicVisitCount != singleNumChildren {
				benchErr = fmt.Errorf("expected %d children, got %d",
					singleNumChildren, sonicVisitCount)
				b.Fail()
			}
			if sonicValueHolder == nil {
				benchErr = fmt.Errorf("expected a value, got nil")
				b.Fail()
			}
			sonicVisitCount = 0
			sonicValueHolder = nil
		}
	}
}

// Store values to avoid compiler optimizations
var sonicVisitCount int
var sonicValueHolder *ast.Node

func sonicVisitChildren(path ast.Sequence, node *ast.Node) bool {
	sonicValueHolder = node
	sonicVisitCount++
	if node.Type() == ast.V_OBJECT || node.Type() == ast.V_ARRAY {
		node.ForEach(sonicVisitChildren)
	}
	return true
}

func sonicFileMany(b *testing.B, f *os.File, count int) {
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		sonicRecordCount := 0
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			record := scanner.Bytes()
			if sonicShouldValidate && !sonic.Valid(record) {
				benchErr = errors.New("json not valid")
				b.Fail()
			}
			whatnode, _ := sonic.Get(record, "what")
			whatval, _ := whatnode.String()
			wherenode, _ := sonic.Get(record, "where", 0, "line")
			whereval, _ := wherenode.Int64()
			if err := checkLog(whatval, int(whereval)); err != nil {
				benchErr = err
				b.Fail()
			}
			sonicRecordCount++
		}
		if sonicRecordCount != count {
			benchErr = fmt.Errorf("expected %d records, got %d", count, sonicRecordCount)
			b.Fail()
		}
	}
}

func sonicFileManyAll(b *testing.B, f *os.File, count int) {
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		sonicRecordCount := 0
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			record := scanner.Bytes()
			if sonicShouldValidate && !sonic.Valid(record) {
				benchErr = errors.New("json not valid")
				b.Fail()
			}
			if root, err := sonic.Get(record); err != nil {
				benchErr = err
				b.Fail()
			} else {
				sonicVisitCount = 0
				sonicValueHolder = nil
				root.ForEach(sonicVisitChildren)
				if sonicVisitCount != logNumChildren {
					benchErr = fmt.Errorf("expected %d children, got %d",
						logNumChildren, sonicVisitCount)
					b.Fail()
				}
				if sonicValueHolder == nil {
					benchErr = fmt.Errorf("expected a value, got nil")
					b.Fail()
				}
			}
			sonicRecordCount++
		}
		if sonicRecordCount != count {
			benchErr = fmt.Errorf("expected %d records, got %d", count, sonicRecordCount)
			b.Fail()
		}
	}
}
