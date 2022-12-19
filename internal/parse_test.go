package flaggy

import (
	"fmt"
	"testing"
)

type SubOption struct {
	Verbose bool     `long:"--verbose" short:"-v" help:"shows help"`
	Rate    *int     `long:"--rate" short:"-r" help:"shows help"`
	Rate2   uint     `long:"--rate2" short:"-r2" help:"shows help"`
	Value   *float32 `long:"--value" short:"-vl" help:"shows help"`
	StrPtr  *string  `long:"--str-ptr" short:"-sp" help:"shows help"`
	Str     *string  `long:"--str" short:"-s" help:"shows help"`
	Array   []int    `long:"--array" short:"-ar" help:"shows help"`
	Array2  []*int   `long:"--array2" short:"-ar2" help:"shows help"`
}

func (subOption SubOption) Run() error {
	fmt.Println("ok")
	return nil
}

type Options struct {
	Help      bool      `long:"help" short:"h" help:"shows help"`
	SubOption SubOption `long:"test" short:"t" help:"shows help"`
}

func (options Options) Run() error {
	if !Parsed() {
		fmt.Println("nothing parsed")
		return nil
	}
	if options.Help {
		PrintHelp()
	}
	return nil
}

func TestParseComplex(t *testing.T) {
	options := Options{}
	err := Parse(&options, []string{"t", "-v", "-r", "11", "-r2", "2", "--unkown", "h", "ok", "-vl", "2.25", "ok", "-sp", "okk", "what", "the", "-s", "ok", "-ar", "1", "-ar2", "200", "-ar", "2"})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if !options.SubOption.Verbose {
		t.Log("verbose is not set")
		t.Fail()
	}
	if options.SubOption.Rate == nil || *options.SubOption.Rate != 11 {
		t.Log("rate is not set")
		t.Fail()
	}
	if options.SubOption.Rate2 != 2 {
		t.Log("rate2 is not set")
		t.Fail()
	}
	if options.SubOption.Value == nil || *options.SubOption.Value != 2.25 {
		t.Log("value is not set")
		t.Fail()
	}
	if options.SubOption.Str == nil || *options.SubOption.Str != "ok" {
		t.Log("str is not set")
		t.Fail()
	}
	if options.SubOption.StrPtr == nil || *options.SubOption.StrPtr != "okk" {
		t.Log("str is not set")
		t.Fail()
	}
	if options.SubOption.Array == nil || len(options.SubOption.Array) != 2 || options.SubOption.Array[0] != 1 || options.SubOption.Array[1] != 2 {
		t.Log("array is not set")
		t.Fail()
	}
}
