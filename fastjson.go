// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/valyala/fastjson"
)

var fastjsonPkg = pkg{
	name: "fastjson",
	calls: map[string]*call{
		"validate":       {name: "Valid", fun: fastjsonValidate},
		"marshal-custom": {name: "Marshal", fun: fastjsonMarshalCustom},
		"file1":          {name: "Decode", fun: fastjsonFile1},
		"small-file":     {name: "Decode", fun: fastjsonFileManySmall},
		"large-file":     {name: "Decode", fun: fastjsonFileManyLarge},
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

func fastjsonMarshalCustom(b *testing.B) {
	//{"when":1711509483695365000,"what":"Just some fake log entry for a generated log file.","where":[{"file":"example.go","line":123}],"who":"benchmark-application","level":"INFO"}
	var a fastjson.Arena
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		o := a.NewObject()
		o.Set("when", a.NewNumberFloat64(float64(time.Now().UnixNano())))
		o.Set("what", a.NewString("Just some fake log entry for a generated log file."))
		where := a.NewObject()
		where.Set("file", a.NewString("example.go"))
		where.Set("line", a.NewNumberInt(123))
		arr := a.NewArray()
		arr.SetArrayItem(0, where)
		o.Set("where", arr)
		o.Set("who", a.NewString("benchmark-application"))
		o.Set("level", a.NewString("INFO"))
		if err := checkMarshalCustom(string(o.MarshalTo(nil))); err != nil {
			benchErr = err
			b.Fail()
		}
		a.Reset()
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
		if val, err := p.ParseBytes(j); err != nil {
			benchErr = err
			b.Fail()
		} else {
			if val.Type() != fastjson.TypeObject {
				benchErr = fmt.Errorf("expected object, got %s", val.Type())
				b.Fail()
			}
			strtest := val.GetStringBytes("identifier", "0", "type", "coding", "0", "code")
			arrtest := val.GetArray("name", "2", "given")
			booltest := val.GetBool("deceasedBoolean")
			err = checkPatient(string(strtest), len(arrtest), booltest)
			if err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}
}

func fastjsonFileManySmall(b *testing.B) {
	fastjsonCheckFileValues(b, openSmallLogFile())
}

func fastjsonFileManyLarge(b *testing.B) {
	fastjsonCheckFileValues(b, openLargeLogFile())
}

func fastjsonCheckFileValues(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()

	var sc fastjson.Scanner
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		buf := bufio.NewScanner(f)
		for buf.Scan() {
			sc.InitBytes(buf.Bytes())
			for sc.Next() {
				val := sc.Value()
				if err := sc.Error(); err != nil {
					benchErr = err
					b.Fail()
				}
				if val.Type() != fastjson.TypeObject {
					benchErr = fmt.Errorf("expected object, got %s", val.Type())
					b.Fail()
				}
				whatval := val.GetStringBytes("what")
				whereval := val.GetInt("where", "0", "line")
				err := checkLog(string(whatval), whereval)
				if err != nil {
					benchErr = err
					b.Fail()
				}
			}
		}
	}
}
