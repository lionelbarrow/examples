package examples

import (
	"strings"
	"testing"
)

func TestFailingIt(t *testing.T) {
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

func TestPassingIt(t *testing.T) {
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
	} else if len(h.Messages) != 1 {
		t.Fail()
	} else if !strings.Contains(h.Messages[0], "Expected 2 to equal 1.") {
		t.Fail()
	}
}

func TestItHandlesPanics(t *testing.T) {
	h := harness()

	When("writing a test library", h,
		It("helps to write tests", func(expect Expectation) {
			panic("Oh no!!")
		}),
	)

	if !h.Failed {
		t.Fail()
	} else if !strings.Contains(h.Messages[0], "Panic while executing it 'helps to write tests': Oh no!!") {
		t.Fail()
	}
}
