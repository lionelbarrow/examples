### Examples

[![Build Status](https://travis-ci.org/lionelbarrow/examples.png?branch=master)](https://travis-ci.org/lionelbarrow/examples)

A lightweight behavioral testing library for Go.

Here's how it works: Write regular tests using the Go standard library, making your assertions with the [quiz library](https://github.com/benmills/quiz). If you ever feel like you're writing tests that should be grouped together, use a `When` or `Describe` function to group them, like this:

```go
import (
  e "github.com/lionelbarrow/examples"
  "testing"
)

func TestMyClass(t *testing.T) {
  e.When("writing software", t,
    e.It("helps to have tests", func(expect *e.Expectation) {
      expect(1).ToEqual(1)
    }),

    e.It("helps to have correct tests", func(expect *e.Expectation) {
      expect(1).ToEqual(2)
    }),
  )
}
```

Run the tests with `go test` as usual. With the code above, this will output:

```
--- FAIL: TestMyClass (0.00 seconds)
examples.go:32: When writing software it helps to have correct tests:
                Expected 2 to equal 1.
                /Users/lionel/fake_examples/fake_test.go:15
```

As far as the `go test` command is concerned, everything within the `TestMyClass` function is the same test. The Go standard library does a number of useful things (such as automatic parallelization) to tests run through the `go test` command, so it makes sense to keep your test groups small.

### License

The MIT License (MIT)

Copyright (c) 2013 Lionel Barrow

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
