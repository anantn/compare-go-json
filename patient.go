// Copyright (c) 2021, Peter Ohler, All rights reserved.

package main

import (
	"fmt"
	"strings"
)

// Patient is a struct used for Marshal and Unmarshal benchmarks.
type Patient struct {
	ResourceType         string           `json:"resourceType"`
	ID                   string           `json:"id"`
	Text                 Text             `json:"text"`
	Identifier           []*Identifier    `json:"identifier"`
	Active               bool             `json:"active"`
	Name                 []*Name          `json:"name"`
	Telecom              []*Telecom       `json:"telecom"`
	Gender               string           `json:"gender"`
	BirthDate            string           `json:"birthDate"`
	XBirthDate           X                `json:"_birthDate"`
	DeceasedBoolean      bool             `json:"deceasedBoolean"`
	Address              []*Address       `json:"address"`
	Contact              []*Contact       `json:"contact"`
	Communication        []*Communication `json:"communication"`
	ManagingOrganization Ref              `json:"managingOrganization"`
	Meta                 Meta             `json:"meta"`
}

// Text is a struct used for Marshal and Unmarshal benchmarks.
type Text struct {
	Status string `json:"status"`
	Div    string `json:"div"`
}

// Name is a struct used for Marshal and Unmarshal benchmarks.
type Name struct {
	Given   []string `json:"given"`
	Family  string   `json:"family"`
	XFamily X        `json:"_family"`
	Use     string   `json:"use"`
	Period  Period   `json:"period"`
}

// Ref is a struct used for Marshal and Unmarshal benchmarks.
type Ref struct {
	Reference string `json:"reference"`
	Display   string `json:"display"`
}

// Identifier is a struct used for Marshal and Unmarshal benchmarks.
type Identifier struct {
	Use      string `json:"use"`
	Type     CC     `json:"type"`
	System   string `json:"system"`
	Value    string `json:"value"`
	Period   Period `json:"period"`
	Assigner Ref    `json:"assigner"`
}

// CC is a struct used for Marshal and Unmarshal benchmarks.
type CC struct {
	Coding []*Tag `json:"coding"`
	Text   string `json:"text"`
}

// Period is a struct used for Marshal and Unmarshal benchmarks.
type Period struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// Meta is a struct used for Marshal and Unmarshal benchmarks.
type Meta struct {
	Tag []*Tag `json:"tag"`
}

// Tag is a struct used for Marshal and Unmarshal benchmarks.
type Tag struct {
	System string `json:"system"`
	Code   string `json:"code"`
}

// X is a struct used for Marshal and Unmarshal benchmarks.
type X struct {
	Extension []Extension `json:"extension"`
}

// Extension is a struct used for Marshal and Unmarshal benchmarks.
type Extension struct {
	URL           string `json:"url"`
	ValueDateTime string `json:"valueDateTime"`
}

// Address is a struct used for Marshal and Unmarshal benchmarks.
type Address struct {
	Use        string   `json:"use"`
	Type       string   `json:"type"`
	Text       string   `json:"text"`
	Line       []string `json:"line"`
	City       string   `json:"city"`
	District   string   `json:"district"`
	State      string   `json:"state"`
	PostalCode string   `json:"postalCode"`
	Country    string   `json:"country"`
	Period     Period   `json:"period"`
}

// Telecom is a struct used for Marshal and Unmarshal benchmarks.
type Telecom struct {
	Use    string `json:"use"`
	System string `json:"system"`
	Value  string `json:"value"`
	Rank   int    `json:"rank"`
	Period Period `json:"period"`
}

// Contact is a struct used for Marshal and Unmarshal benchmarks.
type Contact struct {
	Relationship []*CC      `json:"relationship"`
	Name         Name       `json:"name"`
	Telecom      []*Telecom `json:"telecom"`
	Address      Address    `json:"address"`
	Gender       string     `json:"gender"`
	Period       Period     `json:"period"`
}

// Communication is a struct used for Marshal and Unmarshal benchmarks.
type Communication struct {
	Language  CC   `json:"language"`
	Preferred bool `json:"preferred"`
}

type PartialPatient struct {
	Identifier []struct {
		Type struct {
			Coding []struct {
				Code string `json:"code"`
			} `json:"coding"`
		} `json:"type"`
	} `json:"identifier"`
	Name []struct {
		Given []string `json:"given"`
	} `json:"name"`
	DeceasedBoolean bool `json:"deceasedBoolean"`
}

func checkPatient(strval string, arrval int, boolval bool) error {
	if strval != "MR" {
		return fmt.Errorf("expected MR, got %s", strval)
	}
	if arrval != 2 {
		return fmt.Errorf("expected len(2), got len(%d)", arrval)
	}
	if boolval {
		return fmt.Errorf("expected false, got true")
	}
	return nil
}

func checkPatientStruct(p PartialPatient) error {
	if len(p.Identifier) < 1 {
		return fmt.Errorf("expected 1+ identifier, got %d", len(p.Identifier))
	}
	if len(p.Identifier[0].Type.Coding) < 1 {
		return fmt.Errorf("expected 1+ coding, got %d", len(p.Identifier[0].Type.Coding))
	}
	if len(p.Name) < 1 {
		return fmt.Errorf("expected 1+ name, got %d", len(p.Name))
	}
	return checkPatient(
		p.Identifier[0].Type.Coding[0].Code,
		len(p.Name[0].Given),
		p.DeceasedBoolean)
}

type PartialLog struct {
	What  string `json:"what"`
	Where []struct {
		Line int `json:"line"`
	} `json:"where"`
}

type FullLog struct {
	What  string `json:"what"`
	Where []struct {
		File string `json:"file"`
		Line int    `json:"line"`
	}
	Who      string    `json:"who"`
	Level    string    `json:"level"`
	When     int64     `json:"when"`
	Patients []Patient `json:"patients"`
}

func checkLog(what string, where int) error {
	if what != "Just some fake log entry for a generated log file." {
		return fmt.Errorf("expected log entry got '%s'", what)
	}
	if where != 123 {
		return fmt.Errorf("expected 123 got %d", where)
	}
	return nil
}

func checkLogStruct(l PartialLog) error {
	if len(l.Where) < 1 {
		return fmt.Errorf("expected 1+ where, got %d", len(l.Where))
	}
	return checkLog(l.What, l.Where[0].Line)
}

func checkMarshalCustom(data string) error {
	// Account for minification/spacing differences
	if len(data) < 170 || len(data) > 180 {
		return fmt.Errorf("expected 170-180 bytes, got %d", len(data))
	}
	for _, c := range []string{"when", "what", "where", "file", "line", "who", "level"} {
		if !strings.Contains(data, c) {
			return fmt.Errorf("expected %s", c)
		}
	}
	return nil
}

func getSamplePatient() string {
	return `{"resourceType":"Patient","id":"example","text":{"status":"generated","div":"<div xmlns=\"http://www.w3.org/1999/xhtml\">\n\t\t\t<table>\n\t\t\t\t<tbody>\n\t\t\t\t\t<tr>\n\t\t\t\t\t\t<td>Name</td>\n\t\t\t\t\t\t<td>Peter James \n              <b>Chalmers</b> (&quot;Jim&quot;)\n            </td>\n\t\t\t\t\t</tr>\n\t\t\t\t\t<tr>\n\t\t\t\t\t\t<td>Address</td>\n\t\t\t\t\t\t<td>534 Erewhon, Pleasantville, Vic, 3999</td>\n\t\t\t\t\t</tr>\n\t\t\t\t\t<tr>\n\t\t\t\t\t\t<td>Contacts</td>\n\t\t\t\t\t\t<td>Home: unknown. Work: (03) 5555 6473</td>\n\t\t\t\t\t</tr>\n\t\t\t\t\t<tr>\n\t\t\t\t\t\t<td>Id</td>\n\t\t\t\t\t\t<td>MRN: 12345 (Acme Healthcare)</td>\n\t\t\t\t\t</tr>\n\t\t\t\t</tbody>\n\t\t\t</table>\n\t\t</div>"},"identifier":[{"use":"usual","type":{"coding":[{"system":"http://terminology.hl7.org/CodeSystem/v2-0203","code":"MR"}]},"system":"urn:oid:1.2.36.146.595.217.0.1","value":"12345","period":{"start":"2001-05-06"},"assigner":{"display":"Acme Healthcare"}}],"active":true,"name":[{"use":"official","family":"Chalmers","given":["Peter","James"]},{"use":"usual","given":["Jim"]},{"use":"maiden","family":"Windsor","given":["Peter","James"],"period":{"end":"2002"}}],"telecom":[{"use":"home"},{"system":"phone","value":"(03) 5555 6473","use":"work","rank":1},{"system":"phone","value":"(03) 3410 5613","use":"mobile","rank":2},{"system":"phone","value":"(03) 5555 8834","use":"old","period":{"end":"2014"}}],"gender":"male","birthDate":"1974-12-25","_birthDate":{"extension":[{"url":"http://hl7.org/fhir/StructureDefinition/patient-birthTime","valueDateTime":"1974-12-25T14:35:45-05:00"}]},"deceasedBoolean":false,"address":[{"use":"home","type":"both","text":"534 Erewhon St PeasantVille, Rainbow, Vic  3999","line":["534 Erewhon St"],"city":"PleasantVille","district":"Rainbow","state":"Vic","postalCode":"3999","period":{"start":"1974-12-25"}}],"contact":[{"relationship":[{"coding":[{"system":"http://terminology.hl7.org/CodeSystem/v2-0131","code":"N"}]}],"name":{"family":"du Marché","_family":{"extension":[{"url":"http://hl7.org/fhir/StructureDefinition/humanname-own-prefix","valueString":"VV"}]},"given":["Bénédicte"]},"telecom":[{"system":"phone","value":"+33 (237) 998327"}],"address":{"use":"home","type":"both","line":["534 Erewhon St"],"city":"PleasantVille","district":"Rainbow","state":"Vic","postalCode":"3999","period":{"start":"1974-12-25"}},"gender":"female","period":{"start":"2012"}}],"managingOrganization":{"reference":"Organization/1"}}`
}

func getSampleLog() string {
	return `{"when":1711509483695365000,"what":"Just some fake log entry for a generated log file.","where":[{"file":"example.go","line":123}],"who":"benchmark-application","level":"INFO"}`
}
