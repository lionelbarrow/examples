package examples

import (
	"fmt"
	"github.com/lionelbarrow/quiz"
	"strings"
)

type Expectation func(interface{}) *quiz.Expectation

func It(description string, testBody func(Expectation)) example {
	if !descriptionIsFiltered(description) {
		return example{Skip: true}
	}

	listener, runner := newListenerRunnerPair()

	expectation := func(target interface{}) *quiz.Expectation {
		return quiz.NewExpectation(runner, target)
	}
	go func() {
		defer func() {
			if r := recover(); r != nil {
				runner.Log(fmt.Sprintf("Panic while executing test: %v", r))
				runner.FailNow()
			}
		}()
		testBody(expectation)
		runner.Pass()
	}()
	passed, failureDescription := listener.Results()
	return example{Description: "it " + description + ": \n" + failureDescription, Failed: !passed}
}

func descriptionIsFiltered(description string) bool {
	return len(filter) == 0 || strings.Contains(description, filter)
}

func newListenerRunnerPair() (*exampleRunner, *exampleRunner) {
	successChan := make(chan bool)
	messageChan := make(chan string)

	l := &exampleRunner{Finished: successChan, Message: messageChan}
	r := &exampleRunner{Finished: successChan, Message: messageChan}
	return l, r
}

type exampleRunner struct {
	Finished    chan bool
	Message     chan string
	description string
}

func (er *exampleRunner) Results() (bool, string) {
	for {
		select {
		case success := <-er.Finished:
			return success, er.description
		case newDescription := <-er.Message:
			er.description = newDescription
		}
	}
	return true, ""
}

func (er *exampleRunner) FailNow() {
	er.Finished <- false
}

func (er *exampleRunner) Pass() {
	er.Finished <- true
}

func (er *exampleRunner) Log(line string) {
	er.Message <- line
}

func (er *exampleRunner) Expect(target interface{}) *quiz.Expectation {
	return quiz.NewExpectation(er, target)
}
