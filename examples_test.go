package examples

import (
	"fmt"
	"strings"
	"testing"
)

func harness() *fakeHarness {
	return &fakeHarness{false, []string{}}
}

type fakeHarness struct {
	Failed   bool
	Messages []string
}

func (f *fakeHarness) Fail() {
	f.Failed = true
}

func (f *fakeHarness) Log(args ...interface{}) {
	for _, arg := range args {
		f.Messages = append(f.Messages, fmt.Sprintf("%v", arg))
	}
}

func TestSingleExample(t *testing.T) {
	h := harness()

	When("we're in this one state", h,
		It("does the right thing", func(t *Example) {
			t.Expect(1).ToEqual(2)
		}),
	)

	if !h.Failed {
		t.Fail()
	} else if !strings.Contains(h.Messages[0], "When we're in this one state it does the right thing") {
		t.Fail()
	}
}

func TestTwoExamples(t *testing.T) {
	h := harness()

	When("writing a test library", h,
		It("helps to write tests", func(t *Example) {
			t.Expect("foo").ToContain("bar")
		}),

		It("helps to test thoroughly", func(t *Example) {
			t.Expect("baz").ToEqual("eggs")
		}),
	)

	if !h.Failed {
		t.Fail()
	} else if !strings.Contains(h.Messages[0], "When writing a test library it helps to write tests") {
		t.Fail()
	} else if !strings.Contains(h.Messages[1], "When writing a test library it helps to test thoroughly") {
		t.Fail()
	}
}

func TestPassing(t *testing.T) {
	h := harness()

	When("writing a test library", h,
		It("helps to write tests", func(t *Example) {
			t.Expect(1).ToEqual(1)
		}),
	)

	if h.Failed {
		t.Fail()
	}
}
