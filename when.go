package examples

import (
	"flag"
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

func Describe(description string, harness Harness, results ...example) {
	exampleBlock(description, harness, results)
}

func When(description string, harness Harness, results ...example) {
	exampleBlock(description, harness, results)
}

func exampleBlock(description string, harness Harness, results []example) {
	for _, result := range results {
		if !result.Skip && result.Failed {
			harness.Log("When " + description + " " + result.Description)
			harness.FailNow()
		}
	}
}

type example struct {
	Description string
	Failed      bool
	Skip        bool
}
