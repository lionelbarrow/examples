package examples

import (
	"flag"
	"fmt"
	"testing"
)

var filter string

func init() {
	flag.StringVar(&filter, "examples.run", "", "define a filter for example cases")
	flag.Parse()
}

type Harness interface {
	Log(...interface{})
	FailNow()
}

func Describe(description string, harness Harness, results ...result) {
	exampleBlock(description, harness, results)
}

func When(description string, harness Harness, results ...result) {
	exampleBlock(description, harness, results)
}

func exampleBlock(description string, harness Harness, results []result) {
	failed := false
	for _, result := range results {
		if !result.Skip && result.Failed {
			harness.Log("When " + description + " " + result.Description)
			failed = true
		} else if testing.Verbose() && !result.Failed && !result.Skip {
			fmt.Println("    PASS: When " + description + " " + result.Description)
		}
	}
	if failed {
		harness.FailNow()
	}
}

type result struct {
	Description string
	Failed      bool
	Skip        bool
}
