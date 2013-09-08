package examples

import (
	"strings"
	"testing"
)

func TestSingleExample(t *testing.T) {
	h := harness()

	When("we're in this one state", h,
		It("does the right thing", func(expect Expectation) {
			expect(1).ToEqual(2)
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
		It("helps to write tests", func(expect Expectation) {
			expect("foo").ToContain("bar")
		}),

		It("helps to test thoroughly", func(expect Expectation) {
			expect("baz").ToEqual("eggs")
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
		It("helps to write tests", func(expect Expectation) {
			expect(1).ToEqual(1)
		}),
	)

	if h.Failed {
		t.Fail()
	}
}

func TestItHaltsAfterFirstFailure(t *testing.T) {
	h := harness()
	keptGoing := false

	When("executing an example", h,
		It("doesn't keep going after a failure", func(expect Expectation) {
			expect(1).ToEqual(2)

			keptGoing = true
		}),
	)

	if keptGoing {
		t.Fail()
	}
}

func TestItReportsFirstAssertionToFail(t *testing.T) {
	h := harness()

	When("writing a test library", h,
		It("helps to write tests", func(expect Expectation) {
			expect(1).ToEqual(2)
			expect(1).ToEqual(3)
		}),
	)

	if !h.Failed {
		t.Fail()
	} else if true {
		t.Log(h.Messages)
	}
}
