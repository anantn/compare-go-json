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

var sonicPkg = pkg{
	name: "sonic",
	calls: map[string]*call{
		"validate-bytes":            {name: "Valid", fun: sonicValid},
		"validate-string":           {name: "Valid", fun: sonicValidString},
		"unmarshal-single-few-keys": {name: "Unmarshal", fun: sonicFile1},
		"unmarshal-single-all-keys": {name: "Unmarshal", fun: sonicFile1All},
		"unmarshal-small-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicFileMany(b, openSmallLogFile(), smallLogFileLen)
		}},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicFileManyAll(b, openSmallLogFile(), smallLogFileLen)
		}},
		"unmarshal-large-file-few-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicFileMany(b, openLargeLogFile(), largeLogFileLen)
		}},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: func(b *testing.B) {
			sonicFileMany(b, openLargeLogFile(), largeLogFileLen)
		}},
		"marshal-builder": {name: "Marshal", fun: sonicMarshalBuilder},
	},
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

func sonicValidString(b *testing.B) {
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
	//{"when":1711509483695365000,"what":"Just some fake log entry for a generated log file.","where":[{"file":"example.go","line":123}],"who":"benchmark-application","level":"INFO"}
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

func sonicFile1All(b *testing.B) {
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
			if root, err := sonic.GetFromString(scanner.Text()); err != nil {
				benchErr = err
				b.Fail()
			} else {
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
			if root, err := sonic.GetFromString(scanner.Text()); err != nil {
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
