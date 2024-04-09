package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	json2 "github.com/go-json-experiment/json"
	"tailscale.com/util/must"
)

type jsonbenchMark struct {
	name                  string
	new                   func() any
	data                  []byte
	fromStruct            any
	fromInterface         any
	compactedStructlen    int
	compactedInterfacelen int
}

var jsonbench = []jsonbenchMark{
	{"CanadaGeometry", func() any { return new(canadaRoot) }, mustRead("data/canada_geometry.json.gz"), new(canadaRoot), nil, 0, 0},
	{"CITMCatalog", func() any { return new(citmRoot) }, mustRead("data/citm_catalog.json.gz"), new(citmRoot), nil, 0, 0},
	{"SyntheaFHIR", func() any { return new(syntheaRoot) }, mustRead("data/synthea_fhir.json.gz"), new(syntheaRoot), nil, 0, 0},
	{"TwitterStatus", func() any { return new(twitterRoot) }, mustRead("data/twitter_status.json.gz"), new(twitterRoot), nil, 0, 0},
	{"GolangSource", func() any { return new(golangRoot) }, mustRead("data/golang_source.json.gz"), new(golangRoot), nil, 0, 0},
	{"StringUnicode", func() any { return new(stringRoot) }, mustRead("data/string_unicode.json.gz"), new(stringRoot), nil, 0, 0},
}

func jsonbenchSuites(pkgs []*pkg) []*suite {
	var suites []*suite
	for i, b := range jsonbench {
		suites = append(suites, &suite{
			fun:   b.name + "-unmarshal",
			title: "Unmarshal " + b.name + " generically",
			ref:   "json",
		})
		suites = append(suites, &suite{
			fun:   b.name + "-unmarshal-struct",
			title: "Unmarshal " + b.name + " into struct",
			ref:   "json",
		})
		suites = append(suites, &suite{
			fun:   b.name + "-marshal",
			title: "Marshal " + b.name + " generically",
			ref:   "json",
		})
		suites = append(suites, &suite{
			fun:   b.name + "-marshal-struct",
			title: "Marshal " + b.name + " from struct",
			ref:   "json",
		})

		// Store data for marshaling
		must.Do(json.Unmarshal(b.data, b.fromStruct))
		var dst interface{}
		must.Do(json.Unmarshal(b.data, &dst))
		b.fromInterface = dst

		// Figure out marshaled lengths
		var compactedStruct, compactedInterface bytes.Buffer
		marshaledStruct, err := json2.Marshal(b.fromStruct)
		if err != nil {
			log.Fatalln(err)
		}
		marshaledInterface, err := json2.Marshal(b.fromInterface)
		if err != nil {
			log.Fatalln(err)
		}
		must.Do(json.Compact(&compactedStruct, marshaledStruct))
		must.Do(json.Compact(&compactedInterface, marshaledInterface))

		// Account for some shortening due to floating point representation
		b.compactedStructlen = compactedStruct.Len() - compactedStruct.Len()/10
		b.compactedInterfacelen = compactedInterface.Len() - compactedInterface.Len()/10
		jsonbench[i] = b
	}

	for _, pkg := range pkgs {
		if pkg.unmarshal != nil {
			for _, bench := range jsonbench {
				if pkg.name != "ffjson" && pkg.name != "easyjson" {
					pkg.calls[bench.name+"-unmarshal"] = &call{
						name:   "Unmarshal",
						fun:    unmarshalFunc(pkg, bench, nil),
						caveat: !pkg.canUnmarshalStruct,
					}
				}
				if pkg.canUnmarshalStruct {
					pkg.calls[bench.name+"-unmarshal-struct"] = &call{
						name:   "Unmarshal",
						fun:    unmarshalFunc(pkg, bench, bench.new()),
						caveat: !pkg.canUnmarshalStruct,
					}
				}
			}
		}
		if pkg.marshal != nil {
			for _, bench := range jsonbench {
				if pkg.name != "ffjson" && pkg.name != "easyjson" {
					pkg.calls[bench.name+"-marshal"] = &call{
						name:   "Marshal",
						fun:    marshalFunc(pkg, bench, bench.fromInterface, false),
						caveat: !pkg.canMarshalStruct,
					}
				}
				if pkg.canMarshalStruct {
					pkg.calls[bench.name+"-marshal-struct"] = &call{
						name:   "Marshal",
						fun:    marshalFunc(pkg, bench, bench.fromStruct, true),
						caveat: !pkg.canMarshalStruct,
					}
				}
			}
		}
	}
	return suites
}

func mustRead(path string) []byte {
	b := must.Get(os.ReadFile(path))
	zr := must.Get(gzip.NewReader(bytes.NewReader(b)))
	b = must.Get(io.ReadAll(zr))
	return b
}

func unmarshalFunc(pkg *pkg, bench jsonbenchMark, to any) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if to != nil {
				if err := pkg.unmarshal(bench.data, to); err != nil {
					benchErr = err
					b.Fail()
				}
			} else {
				var dst interface{}
				if err := pkg.unmarshal(bench.data, &dst); err != nil {
					benchErr = err
					b.Fail()
				}
			}
		}
	}
}

func marshalFunc(pkg *pkg, bench jsonbenchMark, from any, fromStruct bool) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if output, err := pkg.marshal(from); err != nil {
				benchErr = err
				b.Fail()
			} else {
				var cmp = bench.compactedInterfacelen
				if fromStruct {
					cmp = bench.compactedStructlen
				}
				// Ignore for jsoniter
				if len(output) < cmp && pkg.name != "jsoniter" {
					benchErr = fmt.Errorf("output is too short: %d < %d", len(output), cmp)
					b.Fail()
				}
			}
		}
	}
}
