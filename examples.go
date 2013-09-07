package examples

import (
	"flag"
	"github.com/benmills/quiz"
	"strings"
)

var filter string

func init() {
	flag.StringVar(&filter, "examples.run", "", "define a filter for example cases")
	flag.Parse()
}

type Expectation func(interface{}) *quiz.Expectation

type Harness interface {
	Log(...interface{})
	Fail()
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
			harness.Fail()
		}
	}
}

func It(description string, testBody func(Expectation)) example {
	recorder := &example{Description: "", Failed: false}

	if descriptionIsFiltered(description) {
    expectation := func(target interface{}) *quiz.Expectation {
      return quiz.NewExpecation(recorder, target)
    }
    testBody(expectation)
		return example{Description: "it " + description + ": \n" + recorder.Description, Failed: recorder.Failed}
	}

	return example{Skip: true}
}

func descriptionIsFiltered(description string) bool {
	return len(filter) == 0 || strings.Contains(description, filter)
}

type example struct {
	Description string
	Failed      bool
	Skip        bool
}

func (e *example) Fail() {
	e.Failed = true
}

func (e *example) Log(line string) {
	e.Description = line
}

func (e *example) Expect(target interface{}) *quiz.Expectation {
	return quiz.NewExpecation(e, target)
}
