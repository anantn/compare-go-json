// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
)

const (
	blocks    = " ▏▎▍▌▋▊▉█"
	darkBlock = "▓"
)

var (
	filename          = "data/patient.json"
	singleNumRootKeys = 14
	singleNumChildren = 116
	benchErr          error

	logNumRootKeys   = 6
	smallNumChildren = 126
	smallLogFile     = "data/log-small.json"
	smallSize        = 100
	smallNumRecords  = 0

	largeNumChildren = 1179
	largeLogFile     = "data/log-large.json"
	largeSize        = 5000
	largeNumRecords  = 0
)

type testfile struct {
	handle      *os.File
	numChildren int
	numRecords  int
}

func smallTestFile() testfile {
	return testfile{openSmallLogFile(), smallNumChildren, smallNumRecords}
}

func largeTestFile() testfile {
	return testfile{openLargeLogFile(), largeNumChildren, largeNumRecords}
}

type specs struct {
	os        string
	model     string
	processor string
	cores     string
	speed     string
	memory    string
}

type call struct {
	name   string
	fun    func(b *testing.B)
	res    testing.BenchmarkResult
	ns     int64 // base adjusted
	bytes  int64 // base adjusted
	allocs int64 // base adjusted
	err    error
	caveat bool
}

type pkg struct {
	name               string
	calls              map[string]*call
	marshal            func(interface{}) ([]byte, error)
	unmarshal          func([]byte, interface{}) error
	canMarshalStruct   bool
	canUnmarshalStruct bool
}

type result struct {
	pkg  string
	call *call
	ref  bool
}

type suite struct {
	title string
	fun   string // key into the pkg calls
	ref   string // reference package for the suite
}

func main() {
	testing.Init()
	flag.Parse()
	if 0 < len(flag.Args()) {
		filename = flag.Args()[0]
	}

	s := getSpecs()
	if s != nil && strings.Contains(s.os, "mac") {
		sonicPkg.calls = map[string]*call{}
		sonicValidatePkg.calls = map[string]*call{}
		simdjsonPkg.calls = map[string]*call{}
	}

	pkgs := []*pkg{
		&jsonPkg,
		&json2Pkg,
		//&ojPkg,
		&fastjsonPkg,
		&jsoniterPkg,
		&jsonparserPkg,
		&goccyPkg,
		&segmentPkg,
		&simdjsonPkg,
		&gjsonPkg,
		&gjsonValidatePkg,
		&sonicPkg,
		&sonicValidatePkg,
		&codecPkg,
		//&jinPkg,
		//&jasonPkg,
		&djsonPkg,
		&ffjsonPkg,
		&easyjsonPkg,
		&sonnetPkg,
	}

	// Run standard jsonbench tests first
	for _, s := range jsonbenchSuites(pkgs) {
		s.exec(pkgs)
	}

	// Run custom tests
	for _, s := range []*suite{
		{fun: "single-few-keys", title: "Unmarshal single record (2kb), read few keys generically", ref: "json"},
		{fun: "single-few-keys-struct", title: "Unmarshal single record (2kb), read few keys into struct", ref: "json"},
		{fun: "single-all-keys", title: "Unmarshal single record (2kb), read all keys generically", ref: "json"},
		{fun: "single-all-keys-struct", title: "Unmarshal single record (2kb), read all keys into struct", ref: "json"},
		{fun: "small-file-few-keys", title: "Unmarshal many records (2kb each) from small file (100MB), read few keys generically", ref: "json"},
		{fun: "small-file-few-keys-struct", title: "Unmarshal many records (2kb each) from small file (100MB), read few keys into struct", ref: "json"},
		{fun: "small-file-all-keys", title: "Unmarshal many records (2kb each) from small file (100MB), read all keys generically", ref: "json"},
		{fun: "small-file-all-keys-struct", title: "Unmarshal many records (2kb each) from small file (100MB), read all keys into struct", ref: "json"},
		{fun: "large-file-few-keys", title: "Unmarshal many records (25kb each) from semi-large file (5GB), read few keys generically", ref: "json"},
		{fun: "large-file-few-keys-struct", title: "Unmarshal many records (25kb each) from semi-large file (5GB), read few keys into struct", ref: "json"},
		{fun: "large-file-all-keys", title: "Unmarshal many records from (25kb each) semi-large file (5GB), read all keys generically", ref: "json"},
		{fun: "large-file-all-keys-struct", title: "Unmarshal many records from (25kb each) semi-large file (5GB), read all keys into struct", ref: "json"},
		{fun: "marshal-builder", title: "Marshal custom data through an object builder", ref: "json"},
		{fun: "validate-bytes", title: "Validate []byte", ref: "json"},
		{fun: "validate-string", title: "Validate string", ref: "json"},
	} {
		s.exec(pkgs)
	}
	// TBD read multiple json, indented small, maybe a few patients in one file
	// TBD validate io.Reader

	fmt.Println()
	fmt.Println(" Higher values (longer bars) are better in all cases. The bar graph compares the")
	fmt.Println(" parsing performance. The lighter colored bar is the reference, the go json")
	fmt.Println(" package.")
	fmt.Println()
	if s != nil {
		fmt.Println("Tests run on:")
		if 0 < len(s.model) {
			fmt.Printf(" Machine:         %s\n", s.model)
		}
		fmt.Printf(" OS:              %s\n", s.os)
		fmt.Printf(" Processor:       %s\n", s.processor)
		fmt.Printf(" Cores:           %s\n", s.cores)
		fmt.Printf(" Processor Speed: %s\n", s.speed)
		fmt.Printf(" Memory:          %s\n", s.memory)
	}
	fmt.Println()
}

func getLogFile() *os.File {
	// create or open benchmarks.csv
	f, err := os.OpenFile("benchmarks.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open benchmarks.csv. %s\n", err)
	}
	// write header if file is empty
	if fi, err := f.Stat(); err == nil && fi.Size() == 0 {
		_, _ = f.WriteString("package,operation,test,ns/op,b/op,allocs/op,caveat\n")
	}
	return f
}

func (s *suite) exec(pkgs []*pkg) {
	f := getLogFile()
	defer func() { _ = f.Close() }()

	fmt.Println()
	fmt.Println(s.title)
	var results []*result
	var ref *call
	for _, p := range pkgs {
		benchErr = nil
		c := p.calls[s.fun]
		r := result{pkg: p.name, call: c, ref: s.ref == p.name}
		results = append(results, &r)
		if c == nil {
			r.call = &call{ns: math.MaxInt64, err: fmt.Errorf("not supported")}
			fmt.Printf(" %12s >>> not supported <<<\n", p.name)
			continue
		}
		if r.ref {
			ref = c
		}
		c.res = testing.Benchmark(c.fun)
		if benchErr != nil {
			c.err = benchErr
			c.ns = math.MaxInt64
			fmt.Printf(" %12s.%-15s >>> %s <<<\n", p.name, c.name, benchErr)
			continue
		}
		c.ns = c.res.NsPerOp()
		c.bytes = c.res.AllocedBytesPerOp()
		c.allocs = c.res.AllocsPerOp()
		fmt.Printf(" %12s.%-15s %12d ns/op %12d B/op %12d allocs/op\n",
			p.name, c.name, c.ns, c.bytes, c.allocs)
	}
	fmt.Println()
	scale := 4 // TBD adjust to fit screen better?
	sort.Slice(results, func(i, j int) bool { return results[i].call.ns < results[j].call.ns })
	for _, r := range results {
		c := r.call
		x := 1.0
		var bar string
		if r.pkg == s.ref {
			bar = strings.Repeat(darkBlock, scale)
		} else {
			if c.err == nil {
				x = float64(ref.ns) / float64(c.ns)
				size := x * float64(scale)
				bar = strings.Repeat(string([]rune(blocks)[8:]), int(size))
				frac := int(size*8.0) - (int(size) * 8)
				bar += string([]rune(blocks)[frac : frac+1])
			} else {
				fmt.Printf(" %12s >>> %s <<<\n", r.pkg, c.err)
				continue
			}
		}
		caveat := ""
		if c.caveat {
			caveat = "* "
		}
		fmt.Printf(" %12s %s %3.2f\n", caveat+r.pkg, bar, x)
		_, _ = fmt.Fprintf(f, "%s,%s,%s,%d,%d,%d,%t\n",
			r.pkg, strings.ToLower(c.name), strings.ToLower(s.fun),
			c.ns, c.bytes, c.allocs, c.caveat)
	}
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], []byte{'\n'})
		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
}

func openSmallLogFile() *os.File {
	f, err := os.Open(smallLogFile)
	if err != nil {
		if err = createLogFile(smallLogFile, false, smallSize); err != nil {
			log.Fatalf("Failed to create %s. %s\n", smallLogFile, err)
		}
		if f, err = os.Open(smallLogFile); err != nil {
			log.Fatalf("Failed to open %s. %s\n", smallLogFile, err)
		}
	}
	if smallNumRecords == 0 {
		if smallNumRecords, err = lineCounter(f); err != nil {
			log.Fatalf("Failed to count lines in %s. %s\n", smallLogFile, err)
		}
	}
	return f
}

func openLargeLogFile() *os.File {
	f, err := os.Open(largeLogFile)
	if err != nil {
		if err = createLogFile(largeLogFile, true, largeSize); err != nil {
			log.Fatalf("Failed to create %s. %s\n", largeLogFile, err)
		}
		if f, err = os.Open(largeLogFile); err != nil {
			log.Fatalf("Failed to open %s. %s\n", largeLogFile, err)
		}
	}
	if largeNumRecords == 0 {
		if largeNumRecords, err = lineCounter(f); err != nil {
			log.Fatalf("Failed to count lines in %s. %s\n", largeLogFile, err)
		}
	}
	return f
}

// size is in MB.
func createLogFile(filename string, largeRecord bool, size int) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	// Repeat of patient
	var patientData interface{}
	oj.Unmarshal([]byte(getSamplePatient()), &patientData)

	// Build a log entry.
	var b oj.Builder
	_ = b.Object()
	_ = b.Value(time.Now().UnixNano(), "when")
	_ = b.Value("Just some fake log entry for a generated log file.", "what")
	_ = b.Array("where")
	_ = b.Object()
	_ = b.Value("example.go", "file")
	_ = b.Value(123, "line")
	b.Pop()
	b.Pop()
	_ = b.Value("benchmark-application", "who")
	_ = b.Value("INFO", "level")

	// Add bunch of patient data to make the record large
	_ = b.Array("patients")
	if largeRecord {
		for i := 0; i < 10; i++ {
			_ = b.Value(patientData)
		}
	} else {
		_ = b.Value(patientData)
	}
	b.PopAll()
	entry := b.Result()

	var whenX jp.Expr
	if whenX, err = jp.Parse([]byte("when")); err != nil {
		return err
	}
	j := oj.JSON(entry)
	cnt := size * 1024 * 1024 / (len(j) + 1)
	for i := 0; i < cnt; i++ {
		// Update entry.
		if err = whenX.Set(entry, time.Now().UnixNano()); err != nil {
			return err
		}
		if err = oj.Write(f, entry); err != nil {
			return err
		}
		_, _ = f.Write([]byte{'\n'})
	}
	return nil
}

func getSpecs() (s *specs) {
	// Assume MacOS and try system_profiler. If that fails assume linux and check /proc.
	out, err := exec.Command("system_profiler", "-json", "SPHardwareDataType").Output()
	if err == nil {
		var js interface{}
		if js, err = oj.Parse(out); err == nil {
			s = &specs{
				model:     alt.String(jp.C("SPHardwareDataType").N(0).C("machine_model").First(js)),
				processor: alt.String(jp.C("SPHardwareDataType").N(0).C("cpu_type").First(js)),
				cores:     alt.String(jp.C("SPHardwareDataType").N(0).C("number_processors").First(js)),
				speed:     alt.String(jp.C("SPHardwareDataType").N(0).C("current_processor_speed").First(js)),
				memory:    alt.String(jp.C("SPHardwareDataType").N(0).C("physical_memory").First(js)),
			}
			var b []byte
			if out, err = exec.Command("sw_vers", "-productName").Output(); err == nil {
				b = append(b, bytes.TrimSpace(out)...)
				b = append(b, ' ')
			}
			if out, err = exec.Command("sw_vers", "-productVersion").Output(); err == nil {
				b = append(b, bytes.TrimSpace(out)...)
			}
			s.os = string(b)
		}
		return
	}
	// Try Ubuntu next.
	if out, err = exec.Command("lsb_release", "-d").Output(); err == nil {
		s = &specs{}
		parts := strings.Split(string(out), ":")
		if 1 < len(parts) {
			s.os = string(strings.TrimSpace(parts[1]))
		}
		if out, err = os.ReadFile("/proc/cpuinfo"); err == nil {
			cnt := 0
			for _, line := range strings.Split(string(out), "\n") {
				if strings.Contains(line, "processor") {
					cnt++
				} else if strings.Contains(line, "model name") {
					parts := strings.Split(line, ":")
					if 1 < len(parts) {
						parts = strings.Split(parts[1], "@")
						s.processor = strings.TrimSpace(parts[0])
						if 1 < len(parts) {
							s.speed = strings.TrimSpace(parts[1])
						}
					}
				} else if s.speed == "" && strings.Contains(line, "cpu MHz") {
					parts := strings.Split(line, ":")
					if 1 < len(parts) {
						speed := strings.TrimSpace(parts[1])
						speedNum, err := strconv.ParseFloat(speed, 64)
						if err == nil {
							s.speed = fmt.Sprintf("%3.2fGHz", speedNum/1000)
						}
					}
				}
				s.cores = fmt.Sprintf("%d", cnt)
			}
		}
		if out, err = os.ReadFile("/proc/meminfo"); err == nil {
			for _, line := range strings.Split(string(out), "\n") {
				if strings.Contains(line, "MemTotal") {
					parts := strings.Split(line, ":")
					if 1 < len(parts) {
						s.memory = strings.TrimSpace(parts[1])
						if strings.HasSuffix(s.memory, "kB") {
							if i, err := strconv.Atoi(strings.Split(s.memory, " ")[0]); err == nil {
								s.memory = fmt.Sprintf("%d GB", i/1000000)
							}
						}
					}
				}
			}
		}
	}
	return
}
