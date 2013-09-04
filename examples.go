package examples

import (
	"github.com/lionelbarrow/quiz"
	"strings"
	"flag"
)

var filter string;

func init() {
	flag.StringVar(&filter, "examples.run", "", "define a filter for example cases")
	flag.Parse()
}

type Harness interface {
	Log(...interface{})
	Fail()
}

type Example struct {
	Description string
	Failed      bool
	Skip        bool
}

func (e *Example) Fail() {
	e.Failed = true
}

func (e *Example) Log(line string) {
	e.Description = line
}

func (e *Example) Expect(target interface{}) *quiz.Expectation {
	return quiz.NewExpecation(e, target)
}

func Describe(description string, harness Harness, results ...Example) {
	exampleBlock(description, harness, results)
}

func When(description string, harness Harness, results ...Example) {
	exampleBlock(description, harness, results)
}

func exampleBlock(description string, harness Harness, results []Example) {
	for _, result := range results {
		if !result.Skip && result.Failed {
			harness.Log("When " + description + " " + result.Description)
			harness.Fail()
		}
	}
}

func It(description string, testBody func(t *Example)) Example {
	example := &Example{Description: "", Failed: false}

	if descriptionIsFiltered(description) {
		testBody(example)
		return Example{Description: "it " + description + ": \n" + example.Description, Failed: example.Failed}
	} else {
		return Example{Skip: true}
	}
}

func descriptionIsFiltered(description string) bool {
	return len(filter) == 0 || strings.Contains(description, filter)
}
