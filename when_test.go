package examples

import (
	"strings"
	"testing"
)

func TestWhenWithTwoExamples(t *testing.T) {
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

func TestKeepGoingWhenAnExampleFails(t *testing.T) {
	h := harness()

	When("writing a test library", h,
		It("helps to write tests", func(expect Expectation) {
			expect(true).ToBeTrue()
		}),

		It("helps to test thoroughly", func(expect Expectation) {
			expect("baz").ToEqual("eggs")
		}),

		It("helps to be great at writing tests", func(expect Expectation) {
			expect(1).ToEqual(1)
		}),
	)

	if !h.Failed {
		t.Fail()
	} else if len(h.Messages) != 1 {
		t.Fail()
	}
}
