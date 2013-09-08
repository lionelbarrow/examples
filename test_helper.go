package examples

import "fmt"

func harness() *fakeHarness {
	return &fakeHarness{false, []string{}}
}

type fakeHarness struct {
	Failed   bool
	Messages []string
}

func (f *fakeHarness) FailNow() {
	f.Failed = true
}

func (f *fakeHarness) Log(args ...interface{}) {
	for _, arg := range args {
		f.Messages = append(f.Messages, fmt.Sprintf("%v", arg))
	}
}
