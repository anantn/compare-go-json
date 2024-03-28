// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"errors"
	"log"
	"os"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

var jsoniterPkg = pkg{
	name: "jsoniter",
	calls: map[string]*call{
		"validate-bytes":                {name: "Valid", fun: jsoniterValid},
		"unmarshal-single-few-keys":     {name: "Unmarshal", fun: jsoniterFile1},
		"unmarshal-single-all-keys":     {name: "Unmarshal", fun: jsoniterUnmarshalPatient},
		"unmarshal-small-file-few-keys": {name: "Unmarshal", fun: jsoniterFileManySmall},
		"unmarshal-small-file-all-keys": {name: "Unmarshal", fun: jsoniterFileManySmall},
		"unmarshal-large-file-few-keys": {name: "Unmarshal", fun: jsoniterFileManyLarge},
		"unmarshal-large-file-all-keys": {name: "Unmarshal", fun: jsoniterFileManyLarge},
		"marshal-builder":               {name: "Marshal", fun: jsoniterMarshalBuilder},
	},
}

func jsoniterValid(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if !jsoniter.Valid(sample) {
			benchErr = errors.New("JSON not valid")
			b.Fail()
		}
	}
}

func jsoniterUnmarshalPatient(b *testing.B) {
	sample, _ := os.ReadFile(filename)
	b.ResetTimer()

	var patient Patient
	for n := 0; n < b.N; n++ {
		if benchErr = jsoniter.Unmarshal(sample, &patient); benchErr != nil {
			b.Fail()
		}
	}
}

func jsoniterMarshalBuilder(b *testing.B) {
	var data interface{}
	err := jsoniter.UnmarshalFromString(`{"when":1711509483695365000,"what":"Just some fake log entry for a generated log file.","where":[{"file":"example.go","line":123}],"who":"benchmark-application","level":"INFO"}`, &data)
	if err != nil {
		benchErr = err
		b.Fail()
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if _, benchErr = jsoniter.Marshal(data); benchErr != nil {
			b.Fail()
		}
	}
}

func jsoniterFile1(b *testing.B) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read %s. %s\n", filename, err)
	}
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		dec := jsoniter.NewDecoder(f)
		var data interface{}
		if err := dec.Decode(&data); err != nil {
			benchErr = err
			b.Fail()
		}
		obj := data.(map[string]interface{})
		strtest := obj["identifier"].([]interface{})[0].(map[string]interface{})["type"].(map[string]interface{})["coding"].([]interface{})[0].(map[string]interface{})["code"].(string)
		arrtest := obj["name"].([]interface{})[2].(map[string]interface{})["given"].([]interface{})
		booltest := obj["deceasedBoolean"].(bool)
		if err := checkPatient(strtest, len(arrtest), booltest); err != nil {
			benchErr = err
			b.Fail()
		}
	}
}

func jsoniterFileManySmall(b *testing.B) {
	jsoniterCheckFileValues(b, openSmallLogFile())
}

func jsoniterFileManyLarge(b *testing.B) {
	jsoniterCheckFileValues(b, openLargeLogFile())
}

func jsoniterCheckFileValues(b *testing.B, f *os.File) {
	defer func() { _ = f.Close() }()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = f.Seek(0, 0)
		dec := jsoniter.NewDecoder(f)
		for {
			var data interface{}
			if !dec.More() {
				break
			}
			if err := dec.Decode(&data); err != nil {
				benchErr = err
				b.Fail()
			}
			whatval := data.(map[string]interface{})["what"].(string)
			whereval := data.(map[string]interface{})["where"].([]interface{})[0].(map[string]interface{})["line"].(float64)
			if err := checkLog(whatval, int(whereval)); err != nil {
				benchErr = err
				b.Fail()
			}
		}
	}

}
