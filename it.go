package examples

import (
	"fmt"
	"github.com/benmills/quiz"
	"strings"
)

type Expectation func(interface{}) *quiz.Expectation

func It(behavior string, testBody func(Expectation)) result {
	if len(filter) > 0 && !strings.Contains(behavior, filter) {
		return result{Skip: true}
	}

	listener, runner := newListenerRunnerPair()
	go executeTestBody(behavior, runner, testBody)
	passed, desc := listener.Results()

	if !passed {
		return result{Description: "it " + behavior + ": \n" + desc, Failed: true}
	}
	return result{Description: "it " + behavior, Failed: false}
}

func executeTestBody(desc string, runner *exampleRunner, body func(Expectation)) {
	expectation := func(target interface{}) *quiz.Expectation {
		return quiz.NewExpectation(runner, target)
	}
	defer func() {
		if r := recover(); r != nil {
			runner.Log(fmt.Sprintf("Panic while executing it '"+desc+"': %v", r))
			runner.FailNow()
		}
	}()
	body(expectation)
	runner.Pass()
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
