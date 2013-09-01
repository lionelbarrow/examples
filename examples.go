package examples

import (
	"github.com/lionelbarrow/quiz"
)

type Harness interface {
	Log(...interface{})
	Fail()
}

type Example struct {
	Description string
	Failed      bool
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
	execute(description, harness, results)
}

func When(description string, harness Harness, results ...Example) {
	execute(description, harness, results)
}

func execute(description string, harness Harness, results []Example) {
	for _, result := range results {
		if result.Failed {
			harness.Log("When " + description + " " + result.Description)
			harness.Fail()
		}
	}
}

func It(description string, testBody func(t *Example)) Example {
	example := &Example{Description: "", Failed: false}
	testBody(example)
	return Example{Description: "it " + description + ": \n" + example.Description, Failed: example.Failed}
}
